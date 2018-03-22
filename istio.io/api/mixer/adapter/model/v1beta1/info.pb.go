// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/adapter/model/v1beta1/info.proto

package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Info describes an adapter or a backend that wants to provide telemetry and policy functionality to Mixer as an
// out of process adapter.
type Info struct {
	// Name of the adapter. It must be RFC 1035 compatible DNS label.
	// Regex: "^[a-z]([-a-z0-9]*[a-z0-9])?$"
	// Name is used in Istio configuration, therefore it should be descriptive but short.
	// example: denier
	// Vendor adapters should use a vendor prefix.
	// example: mycompany-denier
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User-friendly description of the adapter.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Base64 encoded proto descriptor of all the templates the adapter wants to serve.
	Templates []string `protobuf:"bytes,3,rep,name=templates" json:"templates,omitempty"`
	// Base64 encoded proto descriptor of the adapter configuration.
	Config string `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptorInfo, []int{0} }

func (m *Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Info) GetTemplates() []string {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *Info) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func init() {
	proto.RegisterType((*Info)(nil), "istio.mixer.adapter.model.v1beta1.Info")
}
func (this *Info) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Info)
	if !ok {
		that2, ok := that.(Info)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Templates) != len(that1.Templates) {
		return false
	}
	for i := range this.Templates {
		if this.Templates[i] != that1.Templates[i] {
			return false
		}
	}
	if this.Config != that1.Config {
		return false
	}
	return true
}
func (this *Info) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&v1beta1.Info{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "Templates: "+fmt.Sprintf("%#v", this.Templates)+",\n")
	s = append(s, "Config: "+fmt.Sprintf("%#v", this.Config)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringInfo(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Templates) > 0 {
		for _, s := range m.Templates {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Config) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Config)))
		i += copy(dAtA[i:], m.Config)
	}
	return i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Info) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if len(m.Templates) > 0 {
		for _, s := range m.Templates {
			l = len(s)
			n += 1 + l + sovInfo(uint64(l))
		}
	}
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	return n
}

func sovInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Info) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Info{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`Templates:` + fmt.Sprintf("%v", this.Templates) + `,`,
		`Config:` + fmt.Sprintf("%v", this.Config) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringInfo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Templates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Templates = append(m.Templates, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInfo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipInfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/model/v1beta1/info.proto", fileDescriptorInfo) }

var fileDescriptorInfo = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x73, 0x34, 0xaa, 0x14, 0xb3, 0x79, 0x40, 0x19, 0xd0, 0x29, 0x30, 0xa0, 0x4c, 0xb6,
	0x2a, 0xfe, 0x80, 0x8d, 0xb5, 0x03, 0x03, 0x9b, 0xdb, 0x38, 0xe8, 0xa4, 0xc6, 0x67, 0x39, 0x16,
	0x62, 0xe4, 0x13, 0xf8, 0x0c, 0x3e, 0x85, 0xb1, 0x23, 0x23, 0x31, 0x0b, 0x63, 0x3e, 0x01, 0xc9,
	0x44, 0x82, 0x89, 0xed, 0xee, 0xdd, 0xbb, 0xe1, 0x89, 0xab, 0x81, 0x9e, 0x6c, 0xd0, 0xa6, 0x33,
	0x3e, 0xda, 0xa0, 0x07, 0xee, 0xec, 0x41, 0x3f, 0x6e, 0x76, 0x36, 0x9a, 0x8d, 0x26, 0xd7, 0xb3,
	0xf2, 0x81, 0x23, 0xcb, 0x0b, 0x1a, 0x23, 0xb1, 0xca, 0xb6, 0x5a, 0x6c, 0x95, 0x6d, 0xb5, 0xd8,
	0x97, 0x41, 0x94, 0xb7, 0xae, 0x67, 0x29, 0x45, 0xe9, 0xcc, 0x60, 0x6b, 0x68, 0xa0, 0xad, 0xb6,
	0x79, 0x96, 0x8d, 0x38, 0xed, 0xec, 0xb8, 0x0f, 0xe4, 0x23, 0xb1, 0xab, 0x4f, 0xf2, 0xe9, 0x2f,
	0x92, 0xe7, 0xa2, 0x8a, 0x76, 0xf0, 0x07, 0x13, 0xed, 0x58, 0xaf, 0x9a, 0x55, 0x5b, 0x6d, 0x7f,
	0x81, 0x3c, 0x13, 0xeb, 0x3d, 0xbb, 0x9e, 0x1e, 0xea, 0x32, 0xbf, 0x2e, 0xdb, 0xcd, 0xdd, 0x71,
	0xc2, 0xe2, 0x7d, 0xc2, 0x62, 0x9e, 0x10, 0x9e, 0x13, 0xc2, 0x6b, 0x42, 0x78, 0x4b, 0x08, 0xc7,
	0x84, 0xf0, 0x91, 0x10, 0xbe, 0x12, 0x16, 0x73, 0x42, 0x78, 0xf9, 0xc4, 0xe2, 0xbe, 0xfd, 0x89,
	0x20, 0xd6, 0xc6, 0x93, 0xfe, 0xa7, 0x7c, 0xb7, 0xce, 0xd5, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x04, 0x13, 0x80, 0x48, 0x1f, 0x01, 0x00, 0x00,
}
