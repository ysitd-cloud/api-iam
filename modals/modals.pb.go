// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modals.proto

/*
	Package modals is a generated protocol buffer package.

	It is generated from these files:
		modals.proto

	It has these top-level messages:
		Realm
*/
package modals

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import strings "strings"
import reflect "reflect"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Realm struct {
	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Enabled   bool      `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	NotBefore time.Time `protobuf:"bytes,4,opt,name=notBefore,stdtime" json:"notBefore"`
}

func (m *Realm) Reset()                    { *m = Realm{} }
func (*Realm) ProtoMessage()               {}
func (*Realm) Descriptor() ([]byte, []int) { return fileDescriptorModals, []int{0} }

func (m *Realm) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Realm) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Realm) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Realm) GetNotBefore() time.Time {
	if m != nil {
		return m.NotBefore
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Realm)(nil), "modals.Realm")
}
func (this *Realm) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Realm)
	if !ok {
		that2, ok := that.(Realm)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if !this.NotBefore.Equal(that1.NotBefore) {
		return false
	}
	return true
}
func (this *Realm) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&modals.Realm{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Enabled: "+fmt.Sprintf("%#v", this.Enabled)+",\n")
	s = append(s, "NotBefore: "+fmt.Sprintf("%#v", this.NotBefore)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringModals(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Realm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Realm) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintModals(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintModals(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Enabled {
		dAtA[i] = 0x18
		i++
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintModals(dAtA, i, uint64(types.SizeOfStdTime(m.NotBefore)))
	n1, err := types.StdTimeMarshalTo(m.NotBefore, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeVarintModals(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedRealm(r randyModals, easy bool) *Realm {
	this := &Realm{}
	this.Id = string(randStringModals(r))
	this.Name = string(randStringModals(r))
	this.Enabled = bool(bool(r.Intn(2) == 0))
	v1 := types.NewPopulatedStdTime(r, easy)
	this.NotBefore = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyModals interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneModals(r randyModals) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringModals(r randyModals) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneModals(r)
	}
	return string(tmps)
}
func randUnrecognizedModals(r randyModals, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldModals(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldModals(dAtA []byte, r randyModals, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateModals(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateModals(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateModals(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateModals(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateModals(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateModals(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateModals(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Realm) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovModals(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovModals(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	l = types.SizeOfStdTime(m.NotBefore)
	n += 1 + l + sovModals(uint64(l))
	return n
}

func sovModals(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozModals(x uint64) (n int) {
	return sovModals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Realm) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Realm{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Enabled:` + fmt.Sprintf("%v", this.Enabled) + `,`,
		`NotBefore:` + strings.Replace(strings.Replace(this.NotBefore.String(), "Timestamp", "google_protobuf.Timestamp", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringModals(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Realm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModals
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
			return fmt.Errorf("proto: Realm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Realm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModals
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
				return ErrInvalidLengthModals
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModals
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
				return ErrInvalidLengthModals
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotBefore", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthModals
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.NotBefore, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModals
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
func skipModals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModals
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
					return 0, ErrIntOverflowModals
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
					return 0, ErrIntOverflowModals
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
				return 0, ErrInvalidLengthModals
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowModals
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
				next, err := skipModals(dAtA[start:])
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
	ErrInvalidLengthModals = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModals   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("modals.proto", fileDescriptorModals) }

var fileDescriptorModals = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xf3, 0x97, 0x52, 0x5a, 0x83, 0x18, 0x3c, 0xa0, 0xa8, 0xc3, 0xdf, 0x8a, 0x29, 0x0b,
	0x8e, 0x04, 0x37, 0xc8, 0x8e, 0x54, 0x45, 0x4c, 0x6c, 0x4e, 0xed, 0x06, 0x4b, 0x71, 0x5c, 0x35,
	0xce, 0xc0, 0xc6, 0xca, 0xd6, 0x63, 0x70, 0x04, 0x8e, 0xd0, 0xb1, 0x23, 0x13, 0x10, 0xb3, 0x30,
	0x76, 0x64, 0x44, 0x71, 0x14, 0x75, 0xf2, 0xfb, 0x9e, 0x9e, 0xe5, 0x4f, 0x26, 0x17, 0xda, 0x08,
	0x5e, 0x54, 0x6c, 0xbd, 0x31, 0xd6, 0xd0, 0x51, 0x47, 0xd3, 0x59, 0x6e, 0x4c, 0x5e, 0xc8, 0xd8,
	0xb7, 0x59, 0xbd, 0x8a, 0xad, 0xd2, 0xb2, 0xb2, 0x5c, 0xaf, 0xbb, 0xe1, 0xf4, 0x26, 0x57, 0xf6,
	0xa9, 0xce, 0xd8, 0xd2, 0xe8, 0x38, 0x37, 0xb9, 0x39, 0x2e, 0x5b, 0xf2, 0xe0, 0x53, 0x37, 0xbf,
	0x7e, 0x05, 0x72, 0x9a, 0x4a, 0x5e, 0x68, 0x7a, 0x49, 0x06, 0x4a, 0x84, 0x30, 0x87, 0x68, 0x92,
	0x0e, 0x94, 0xa0, 0x94, 0x0c, 0x4b, 0xae, 0x65, 0x38, 0xf0, 0x8d, 0xcf, 0x34, 0x24, 0x67, 0xb2,
	0xe4, 0x59, 0x21, 0x45, 0x78, 0x32, 0x87, 0x68, 0x9c, 0xf6, 0x48, 0x13, 0x32, 0x29, 0x8d, 0x4d,
	0xe4, 0xca, 0x6c, 0x64, 0x38, 0x9c, 0x43, 0x74, 0x7e, 0x3b, 0x65, 0x9d, 0x2b, 0xeb, 0x0d, 0xd8,
	0x43, 0xef, 0x9a, 0x8c, 0x77, 0x9f, 0xb3, 0x60, 0xfb, 0x35, 0x83, 0xf4, 0x78, 0x2d, 0x79, 0xdc,
	0x37, 0x18, 0x7c, 0x34, 0x18, 0x1c, 0x1a, 0x84, 0xbf, 0x06, 0xe1, 0xc5, 0x21, 0xbc, 0x39, 0x84,
	0x77, 0x87, 0xb0, 0x73, 0x08, 0x7b, 0x87, 0xf0, 0xed, 0x10, 0x7e, 0x1d, 0x06, 0x07, 0x87, 0xb0,
	0xfd, 0xc1, 0x80, 0x5c, 0x2d, 0x0b, 0x53, 0x0b, 0xf6, 0x5c, 0x29, 0x2b, 0x98, 0xe2, 0x9a, 0x75,
	0x3f, 0x95, 0x90, 0xfb, 0xf6, 0x5c, 0xb4, 0x6f, 0x2f, 0x20, 0x1b, 0x79, 0x89, 0xbb, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8a, 0x45, 0xe5, 0x7e, 0x56, 0x01, 0x00, 0x00,
}
