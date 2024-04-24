// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.1.2
// source: github.com/aperturerobotics/protobuf-go-lite/testproto/disable-json/disable-json.proto

package disable_json

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
	strings "strings"
	unsafe "unsafe"

	protobuf_go_lite "github.com/aperturerobotics/protobuf-go-lite"
)

// protobuf-go-lite:disable-json
type MessageDisableJson struct {
	unknownFields []byte
	// Types that are assignable to Body:
	//
	//	*MessageDisableJson_Hello
	//	*MessageDisableJson_World
	Body isMessageDisableJson_Body `protobuf_oneof:"body"`
}

func (x *MessageDisableJson) Reset() {
	*x = MessageDisableJson{}
}

func (*MessageDisableJson) ProtoMessage() {}

func (m *MessageDisableJson) GetBody() isMessageDisableJson_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *MessageDisableJson) GetHello() bool {
	if x, ok := x.GetBody().(*MessageDisableJson_Hello); ok {
		return x.Hello
	}
	return false
}

func (x *MessageDisableJson) GetWorld() string {
	if x, ok := x.GetBody().(*MessageDisableJson_World); ok {
		return x.World
	}
	return ""
}

type isMessageDisableJson_Body interface {
	isMessageDisableJson_Body()
}

type MessageDisableJson_Hello struct {
	Hello bool `protobuf:"varint,1,opt,name=hello,proto3,oneof"`
}

type MessageDisableJson_World struct {
	World string `protobuf:"bytes,2,opt,name=world,proto3,oneof"`
}

func (*MessageDisableJson_Hello) isMessageDisableJson_Body() {}

func (*MessageDisableJson_World) isMessageDisableJson_Body() {}

func (m *MessageDisableJson) CloneVT() *MessageDisableJson {
	if m == nil {
		return (*MessageDisableJson)(nil)
	}
	r := new(MessageDisableJson)
	if m.Body != nil {
		r.Body = m.Body.(interface {
			CloneVT() isMessageDisableJson_Body
		}).CloneVT()
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *MessageDisableJson) CloneMessageVT() protobuf_go_lite.CloneMessage {
	return m.CloneVT()
}

func (m *MessageDisableJson_Hello) CloneVT() *MessageDisableJson_Hello {
	if m == nil {
		return (*MessageDisableJson_Hello)(nil)
	}
	r := new(MessageDisableJson_Hello)
	r.Hello = m.Hello
	return r
}

func (m *MessageDisableJson_World) CloneVT() *MessageDisableJson_World {
	if m == nil {
		return (*MessageDisableJson_World)(nil)
	}
	r := new(MessageDisableJson_World)
	r.World = m.World
	return r
}

func (this *MessageDisableJson) EqualVT(that *MessageDisableJson) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Body == nil && that.Body != nil {
		return false
	} else if this.Body != nil {
		if that.Body == nil {
			return false
		}
		if !this.Body.(interface {
			EqualVT(isMessageDisableJson_Body) bool
		}).EqualVT(that.Body) {
			return false
		}
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *MessageDisableJson) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*MessageDisableJson)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *MessageDisableJson_Hello) EqualVT(thatIface isMessageDisableJson_Body) bool {
	that, ok := thatIface.(*MessageDisableJson_Hello)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.Hello != that.Hello {
		return false
	}
	return true
}

func (this *MessageDisableJson_World) EqualVT(thatIface isMessageDisableJson_Body) bool {
	that, ok := thatIface.(*MessageDisableJson_World)
	if !ok {
		return false
	}
	if this == that {
		return true
	}
	if this == nil && that != nil || this != nil && that == nil {
		return false
	}
	if this.World != that.World {
		return false
	}
	return true
}

func (m *MessageDisableJson) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageDisableJson) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MessageDisableJson) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.Body.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *MessageDisableJson_Hello) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MessageDisableJson_Hello) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.Hello {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *MessageDisableJson_World) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MessageDisableJson_World) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.World)
	copy(dAtA[i:], m.World)
	i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.World)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *MessageDisableJson) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageDisableJson) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MessageDisableJson) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if msg, ok := m.Body.(*MessageDisableJson_World); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.Body.(*MessageDisableJson_Hello); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *MessageDisableJson_Hello) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MessageDisableJson_Hello) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.Hello {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *MessageDisableJson_World) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *MessageDisableJson_World) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.World)
	copy(dAtA[i:], m.World)
	i = protobuf_go_lite.EncodeVarint(dAtA, i, uint64(len(m.World)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *MessageDisableJson) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Body.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *MessageDisableJson_Hello) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *MessageDisableJson_World) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.World)
	n += 1 + l + protobuf_go_lite.SizeOfVarint(uint64(l))
	return n
}
func (x *MessageDisableJson) MarshalProtoText() string {
	var sb strings.Builder
	sb.WriteString("MessageDisableJson { ")
	switch body := x.Body.(type) {
	case *MessageDisableJson_Hello:
		if body.Hello {
			sb.WriteString(" hello: ")
			sb.WriteString(strconv.FormatBool(body.Hello))
		}
	case *MessageDisableJson_World:
		if body.World != "" {
			sb.WriteString(" world: ")
			sb.WriteString(strconv.Quote(body.World))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
func (x *MessageDisableJson) String() string {
	return x.MarshalProtoText()
}
func (m *MessageDisableJson) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageDisableJson: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageDisableJson: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hello", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Body = &MessageDisableJson_Hello{Hello: b}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field World", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = &MessageDisableJson_World{World: string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MessageDisableJson) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protobuf_go_lite.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageDisableJson: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageDisableJson: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hello", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Body = &MessageDisableJson_Hello{Hello: b}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field World", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protobuf_go_lite.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			m.Body = &MessageDisableJson_World{World: stringValue}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := protobuf_go_lite.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protobuf_go_lite.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
