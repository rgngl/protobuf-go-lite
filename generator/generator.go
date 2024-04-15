// Copyright (c) 2021 PlanetScale Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generator

import (
	"fmt"
	"runtime/debug"

	"github.com/aperturerobotics/protobuf-go-lite/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"

	generator_base "github.com/aperturerobotics/protobuf-go-lite/generator/base"
	"github.com/aperturerobotics/protobuf-go-lite/generator/pattern"
)

type ObjectSet struct {
	mp map[string]bool
}

func NewObjectSet() ObjectSet {
	return ObjectSet{
		mp: map[string]bool{},
	}
}

func (o ObjectSet) String() string {
	return fmt.Sprintf("%#v", o)
}

func (o ObjectSet) Contains(g protogen.GoIdent) bool {
	objectPath := fmt.Sprintf("%s.%s", string(g.GoImportPath), g.GoName)

	for wildcard := range o.mp {
		// Ignore malformed pattern error because pattern already checked in Set
		if ok, _ := pattern.Match(wildcard, objectPath); ok {
			return true
		}
	}

	return false
}

func (o ObjectSet) Set(s string) error {
	if !pattern.ValidatePattern(s) {
		return pattern.ErrBadPattern
	}
	o.mp[s] = true
	return nil
}

type Config struct {
	// Poolable rules determines if pool feature generate for particular message
	Poolable ObjectSet
	// PoolableExclude rules determines if pool feature disabled for particular message
	PoolableExclude ObjectSet
	WellKnownTypes  bool
	AllowEmpty      bool
	BuildTag        string
}

type Generator struct {
	plugin   *protogen.Plugin
	cfg      *Config
	features []Feature
	local    map[protoreflect.FullName]bool
}

const SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

func NewGenerator(plugin *protogen.Plugin, featureNames []string, cfg *Config) (*Generator, error) {
	plugin.SupportedFeatures = SupportedFeatures

	features, err := findFeatures(featureNames)
	if err != nil {
		return nil, err
	}

	local := make(map[protoreflect.FullName]bool)
	for _, f := range plugin.Files {
		if f.Generate {
			local[f.Desc.Package()] = true
		}
	}

	return &Generator{
		plugin:   plugin,
		cfg:      cfg,
		features: features,
		local:    local,
	}, nil
}

func (gen *Generator) Generate() {
	for _, file := range gen.plugin.Files {
		if !file.Generate {
			continue
		}

		importPath := file.GoImportPath
		gf := gen.plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".pb.go", importPath)
		p := &GeneratedFile{
			GeneratedFile: gf,
			Config:        gen.cfg,
			LocalPackages: gen.local,
		}

		// Generate header
		gen.generateHeader(p, file)

		// Generate base protobuf adapted from protobuf-go
		generator_base.GenerateFile(gen.plugin, file, gf)

		// Generate vtproto features
		gen.generateFile(p, file)
	}
}

func (gen *Generator) generateHeader(p *GeneratedFile, file *protogen.File) {
	if p.Config.BuildTag != "" {
		p.P("//go:build ", p.Config.BuildTag)
	}
	p.P("// Code generated by protoc-gen-go-lite. DO NOT EDIT.")
	if bi, ok := debug.ReadBuildInfo(); ok {
		p.P("// protoc-gen-go-lite version: ", bi.Main.Version)
	}
	if file.Proto.GetOptions().GetDeprecated() {
		p.P("// ", file.Desc.Path(), " is a deprecated file.")
	}
	p.P("// source: ", file.Desc.Path())
	p.P()
	p.P(generator_base.GenPackageKnownComment(file), "package ", file.GoPackageName)
	p.P()
}

func (gen *Generator) generateFile(p *GeneratedFile, file *protogen.File) {
	for _, feat := range gen.features {
		featGenerator := feat(p)
		_ = featGenerator.GenerateFile(file)
	}
}