// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v0.0.0-dev
// - protoc             v3.9.1
// source: enums.proto

package test

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// CustomEnum_customname contains custom string values that override CustomEnum_name.
var CustomEnum_customname = map[int32]string{
	1: "1.0",
	2: "1.0.1",
}

// MarshalProtoJSON marshals the CustomEnum to JSON.
func (x CustomEnum) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), CustomEnum_customname, CustomEnum_name)
}

// CustomEnum_customvalue contains custom string values that extend CustomEnum_value.
var CustomEnum_customvalue = map[string]int32{
	"UNKNOWN": 0,
	"V1_0":    1,
	"1.0":     1,
	"1.0.0":   1,
	"V1_0_1":  2,
	"1.0.1":   2,
}

// UnmarshalProtoJSON unmarshals the CustomEnum from JSON.
func (x *CustomEnum) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(CustomEnum_value, CustomEnum_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read CustomEnum enum: %v", err)
		return
	}
	*x = CustomEnum(v)
}

// MarshalProtoJSON marshals the CustomEnumValue message to JSON.
func (x *CustomEnumValue) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	x.Value.MarshalProtoJSON(s)
	return
}

// UnmarshalProtoJSON unmarshals the CustomEnumValue message from JSON.
func (x *CustomEnumValue) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	x.Value.UnmarshalProtoJSON(s)
	return
}

// MarshalProtoJSON marshals the MessageWithEnums message to JSON.
func (x *MessageWithEnums) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Regular != 0 || s.HasField("regular") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("regular")
		s.WriteEnum(int32(x.Regular), RegularEnum_name)
	}
	if len(x.Regulars) > 0 || s.HasField("regulars") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("regulars")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Regulars {
			s.WriteMoreIf(&wroteElement)
			s.WriteEnum(int32(element), RegularEnum_name)
		}
		s.WriteArrayEnd()
	}
	if x.Custom != 0 || s.HasField("custom") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("custom")
		x.Custom.MarshalProtoJSON(s)
	}
	if len(x.Customs) > 0 || s.HasField("customs") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("customs")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Customs {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.WrappedCustom != nil || s.HasField("wrapped_custom") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("wrapped_custom")
		x.WrappedCustom.MarshalProtoJSON(s.WithField("wrapped_custom"))
	}
	if len(x.WrappedCustoms) > 0 || s.HasField("wrapped_customs") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("wrapped_customs")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.WrappedCustoms {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("wrapped_customs"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the MessageWithEnums message from JSON.
func (x *MessageWithEnums) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "regular":
			s.AddField("regular")
			x.Regular = RegularEnum(s.ReadEnum(RegularEnum_value))
		case "regulars":
			s.AddField("regulars")
			s.ReadArray(func() {
				x.Regulars = append(x.Regulars, RegularEnum(s.ReadEnum(RegularEnum_value)))
			})
		case "custom":
			s.AddField("custom")
			x.Custom.UnmarshalProtoJSON(s)
		case "customs":
			s.AddField("customs")
			s.ReadArray(func() {
				var v CustomEnum
				v.UnmarshalProtoJSON(s)
				x.Customs = append(x.Customs, v)
			})
		case "wrapped_custom", "wrappedCustom":
			s.AddField("wrapped_custom")
			if !s.ReadNil() {
				x.WrappedCustom = &CustomEnumValue{}
				x.WrappedCustom.UnmarshalProtoJSON(s.WithField("wrapped_custom", false))
			}
		case "wrapped_customs", "wrappedCustoms":
			s.AddField("wrapped_customs")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.WrappedCustoms = append(x.WrappedCustoms, nil)
					return
				}
				v := &CustomEnumValue{}
				v.UnmarshalProtoJSON(s.WithField("wrapped_customs", false))
				if s.Err() != nil {
					return
				}
				x.WrappedCustoms = append(x.WrappedCustoms, v)
			})
		}
	})
}

// MarshalProtoJSON marshals the MessageWithOneofEnums message to JSON.
func (x *MessageWithOneofEnums) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Value != nil {
		switch ov := x.Value.(type) {
		case *MessageWithOneofEnums_Regular:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("regular")
			s.WriteEnum(int32(ov.Regular), RegularEnum_name)
		case *MessageWithOneofEnums_Custom:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("custom")
			ov.Custom.MarshalProtoJSON(s)
		case *MessageWithOneofEnums_WrappedCustom:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("wrapped_custom")
			ov.WrappedCustom.MarshalProtoJSON(s.WithField("wrapped_custom"))
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the MessageWithOneofEnums message from JSON.
func (x *MessageWithOneofEnums) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "regular":
			s.AddField("regular")
			ov := &MessageWithOneofEnums_Regular{}
			ov.Regular = RegularEnum(s.ReadEnum(RegularEnum_value))
			x.Value = ov
		case "custom":
			s.AddField("custom")
			ov := &MessageWithOneofEnums_Custom{}
			ov.Custom.UnmarshalProtoJSON(s)
			x.Value = ov
		case "wrapped_custom", "wrappedCustom":
			s.AddField("wrapped_custom")
			ov := &MessageWithOneofEnums_WrappedCustom{}
			if !s.ReadNil() {
				ov.WrappedCustom = &CustomEnumValue{}
				ov.WrappedCustom.UnmarshalProtoJSON(s.WithField("wrapped_custom", false))
			}
			x.Value = ov
		}
	})
}
