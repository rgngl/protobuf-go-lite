//go:build deps_only
// +build deps_only

package hack

import (
	// _ imports golangci-lint
	_ "github.com/golangci/golangci-lint/pkg/golinters"
	// _ imports golangci-lint commands
	_ "github.com/golangci/golangci-lint/pkg/commands"
	// _ imports goimports
	_ "golang.org/x/tools/cmd/goimports"
	// _ imports gofumpt
	_ "mvdan.cc/gofumpt"
	// _ imports protoc-gen-go-lite
	_ "github.com/aperturerobotics/protobuf-go-lite/cmd/protoc-gen-go-lite"
	// _ imports protoc-gen-go-lite-vtproto
	_ "github.com/aperturerobotics/vtprotobuf-lite/cmd/protoc-gen-go-lite-vtproto"
	// _ imports protowrap
	_ "github.com/aperturerobotics/goprotowrap/cmd/protowrap"
)
