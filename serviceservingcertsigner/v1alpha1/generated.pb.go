// Code generated by protoc-gen-gogo.
// source: github.com/openshift/api/serviceservingcertsigner/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		github.com/openshift/api/serviceservingcertsigner/v1alpha1/generated.proto

	It has these top-level messages:
		DelegatedAuthentication
		DelegatedAuthorization
		ServiceServingCertSignerConfig
*/
package v1alpha1

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *DelegatedAuthentication) Reset()                    { *m = DelegatedAuthentication{} }
func (*DelegatedAuthentication) ProtoMessage()               {}
func (*DelegatedAuthentication) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *DelegatedAuthorization) Reset()                    { *m = DelegatedAuthorization{} }
func (*DelegatedAuthorization) ProtoMessage()               {}
func (*DelegatedAuthorization) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *ServiceServingCertSignerConfig) Reset()      { *m = ServiceServingCertSignerConfig{} }
func (*ServiceServingCertSignerConfig) ProtoMessage() {}
func (*ServiceServingCertSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{2}
}

func init() {
	proto.RegisterType((*DelegatedAuthentication)(nil), "github.com.openshift.api.serviceservingcertsigner.v1alpha1.DelegatedAuthentication")
	proto.RegisterType((*DelegatedAuthorization)(nil), "github.com.openshift.api.serviceservingcertsigner.v1alpha1.DelegatedAuthorization")
	proto.RegisterType((*ServiceServingCertSignerConfig)(nil), "github.com.openshift.api.serviceservingcertsigner.v1alpha1.ServiceServingCertSignerConfig")
}
func (m *DelegatedAuthentication) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedAuthentication) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *DelegatedAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatedAuthorization) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *ServiceServingCertSignerConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceServingCertSignerConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ServingInfo.Size()))
	n1, err := m.ServingInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Authentication.Size()))
	n2, err := m.Authentication.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Authorization.Size()))
	n3, err := m.Authorization.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Signer.Size()))
	n4, err := m.Signer.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DelegatedAuthentication) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}

func (m *DelegatedAuthorization) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}

func (m *ServiceServingCertSignerConfig) Size() (n int) {
	var l int
	_ = l
	l = m.ServingInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Authentication.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Authorization.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Signer.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DelegatedAuthentication) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DelegatedAuthentication{`,
		`Disabled:` + fmt.Sprintf("%v", this.Disabled) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DelegatedAuthorization) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DelegatedAuthorization{`,
		`Disabled:` + fmt.Sprintf("%v", this.Disabled) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServiceServingCertSignerConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServiceServingCertSignerConfig{`,
		`ServingInfo:` + strings.Replace(strings.Replace(this.ServingInfo.String(), "HTTPServingInfo", "github_com_openshift_api_config_v1.HTTPServingInfo", 1), `&`, ``, 1) + `,`,
		`Authentication:` + strings.Replace(strings.Replace(this.Authentication.String(), "DelegatedAuthentication", "DelegatedAuthentication", 1), `&`, ``, 1) + `,`,
		`Authorization:` + strings.Replace(strings.Replace(this.Authorization.String(), "DelegatedAuthorization", "DelegatedAuthorization", 1), `&`, ``, 1) + `,`,
		`Signer:` + strings.Replace(strings.Replace(this.Signer.String(), "CertInfo", "github_com_openshift_api_config_v1.CertInfo", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DelegatedAuthentication) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: DelegatedAuthentication: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedAuthentication: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			m.Disabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *DelegatedAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: DelegatedAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatedAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
			m.Disabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ServiceServingCertSignerConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ServiceServingCertSignerConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceServingCertSignerConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServingInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServingInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Authentication.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Signer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/openshift/api/serviceservingcertsigner/v1alpha1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0x98, 0xa6, 0xca, 0xd5, 0x2a, 0x14, 0xc4, 0xa8, 0x7a, 0xf0, 0xd0, 0x4e, 0x1c,
	0x86, 0xad, 0x0e, 0x84, 0x10, 0x37, 0xb2, 0x89, 0x7f, 0x27, 0x94, 0xf4, 0xc4, 0xcd, 0x4d, 0xdf,
	0x3a, 0x66, 0x8d, 0x1d, 0x39, 0x6e, 0x25, 0x38, 0x71, 0xe6, 0x02, 0x12, 0x5f, 0xaa, 0xc7, 0x1d,
	0x77, 0x9a, 0x68, 0xf8, 0x22, 0xa8, 0x4e, 0xd6, 0x25, 0x1b, 0x11, 0x68, 0xda, 0xa9, 0xf5, 0x23,
	0x3f, 0xbf, 0xf7, 0xc9, 0xfb, 0x18, 0xbf, 0x17, 0xd2, 0x26, 0xf3, 0x31, 0x8d, 0x75, 0xca, 0x74,
	0x06, 0x2a, 0x4f, 0xe4, 0xd4, 0x32, 0x9e, 0x49, 0x96, 0x83, 0x59, 0xc8, 0x18, 0xdc, 0x8f, 0x12,
	0x31, 0x18, 0x9b, 0x4b, 0xa1, 0xc0, 0xb0, 0xc5, 0x90, 0xcf, 0xb2, 0x84, 0x0f, 0x99, 0x00, 0x05,
	0x86, 0x5b, 0x98, 0xd0, 0xcc, 0x68, 0xab, 0xfd, 0x97, 0x97, 0x2c, 0xba, 0x61, 0x51, 0x9e, 0x49,
	0xda, 0xc6, 0xa2, 0x17, 0xac, 0xc1, 0x93, 0x5a, 0x0e, 0xa1, 0x85, 0x66, 0x0e, 0x39, 0x9e, 0x4f,
	0xdd, 0xc9, 0x1d, 0xdc, 0xbf, 0x72, 0xd4, 0xe0, 0xb0, 0x35, 0x76, 0xac, 0xd5, 0x54, 0x0a, 0xb6,
	0xb8, 0x16, 0x6f, 0xf0, 0xec, 0xe4, 0x45, 0x4e, 0xa5, 0x5e, 0xdf, 0x4a, 0x79, 0x9c, 0x48, 0x05,
	0xe6, 0x33, 0xcb, 0x4e, 0xc4, 0x5a, 0xc8, 0x59, 0x0a, 0x96, 0xff, 0xcd, 0xc5, 0xda, 0x5c, 0x66,
	0xae, 0xac, 0x4c, 0xe1, 0x9a, 0xe1, 0xf9, 0xbf, 0x0c, 0x79, 0x9c, 0x40, 0xca, 0xaf, 0xfa, 0xf6,
	0xdf, 0xe0, 0x87, 0xc7, 0x30, 0x03, 0xb1, 0x96, 0x5e, 0xcd, 0x6d, 0x02, 0xca, 0xca, 0x98, 0x5b,
	0xa9, 0x95, 0x7f, 0x80, 0x3b, 0x13, 0x99, 0xf3, 0xf1, 0x0c, 0x26, 0x7d, 0xf4, 0x08, 0x3d, 0xee,
	0x04, 0xf7, 0x96, 0xe7, 0x7b, 0x5e, 0x71, 0xbe, 0xd7, 0x39, 0xae, 0xf4, 0x70, 0x73, 0x63, 0xff,
	0x35, 0xde, 0x6d, 0x80, 0xb4, 0x91, 0x5f, 0x6e, 0xc2, 0xf9, 0xb6, 0x85, 0x49, 0x54, 0x16, 0x17,
	0x95, 0xc5, 0x1d, 0x81, 0xb1, 0x91, 0x2b, 0xee, 0xc8, 0x6d, 0xd9, 0xff, 0x84, 0xbb, 0x55, 0xa7,
	0xef, 0xd4, 0x54, 0x3b, 0x66, 0xf7, 0xf0, 0x29, 0x6d, 0x7d, 0x07, 0x65, 0x39, 0x74, 0x31, 0xa4,
	0x6f, 0x47, 0xa3, 0x0f, 0xd1, 0xa5, 0x35, 0xb8, 0x5f, 0x05, 0xe9, 0xd6, 0xc4, 0xb0, 0x0e, 0xf7,
	0x7f, 0x22, 0xdc, 0xe3, 0x8d, 0xbd, 0xf4, 0xef, 0xb8, 0x79, 0x11, 0xbd, 0xf9, 0xbb, 0xa3, 0x2d,
	0x2b, 0x0f, 0x76, 0xab, 0x3c, 0xbd, 0xa6, 0x1e, 0x5e, 0x89, 0xe0, 0x7f, 0x47, 0x78, 0x87, 0xd7,
	0x97, 0xdc, 0xbf, 0xeb, 0x42, 0x85, 0xb7, 0x16, 0x6a, 0x43, 0x0e, 0x1e, 0x54, 0x99, 0x76, 0x1a,
	0x72, 0xd8, 0x9c, 0xef, 0x8f, 0xf0, 0x76, 0xc9, 0xeb, 0x6f, 0xb9, 0x24, 0x07, 0xff, 0x53, 0xc7,
	0xba, 0x59, 0xd7, 0x43, 0xaf, 0x9a, 0xb1, 0x5d, 0xf6, 0x1c, 0x56, 0xac, 0x80, 0x2e, 0x57, 0xc4,
	0x3b, 0x5d, 0x11, 0xef, 0x6c, 0x45, 0xbc, 0xaf, 0x05, 0x41, 0xcb, 0x82, 0xa0, 0xd3, 0x82, 0xa0,
	0xb3, 0x82, 0xa0, 0x5f, 0x05, 0x41, 0x3f, 0x7e, 0x13, 0xef, 0x63, 0xe7, 0xe2, 0x13, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x76, 0x2c, 0xb0, 0x8e, 0x58, 0x04, 0x00, 0x00,
}