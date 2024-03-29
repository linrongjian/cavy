// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// VerifyReq 账号验证请求
type VerifyReq struct {
	CheckCode            *string  `protobuf:"bytes,1,opt,name=checkCode" json:"checkCode,omitempty"`
	Server               *string  `protobuf:"bytes,2,opt,name=server" json:"server,omitempty"`
	GameId               *int32   `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	Version              *string  `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyReq) Reset()         { *m = VerifyReq{} }
func (m *VerifyReq) String() string { return proto.CompactTextString(m) }
func (*VerifyReq) ProtoMessage()    {}
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}
func (m *VerifyReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyReq.Merge(m, src)
}
func (m *VerifyReq) XXX_Size() int {
	return m.Size()
}
func (m *VerifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyReq proto.InternalMessageInfo

func (m *VerifyReq) GetCheckCode() string {
	if m != nil && m.CheckCode != nil {
		return *m.CheckCode
	}
	return ""
}

func (m *VerifyReq) GetServer() string {
	if m != nil && m.Server != nil {
		return *m.Server
	}
	return ""
}

func (m *VerifyReq) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *VerifyReq) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

// VerifyRsp 账号验证响应
type VerifyRsp struct {
	Success              *bool    `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Account              []byte   `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRsp) Reset()         { *m = VerifyRsp{} }
func (m *VerifyRsp) String() string { return proto.CompactTextString(m) }
func (*VerifyRsp) ProtoMessage()    {}
func (*VerifyRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}
func (m *VerifyRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRsp.Merge(m, src)
}
func (m *VerifyRsp) XXX_Size() int {
	return m.Size()
}
func (m *VerifyRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRsp.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRsp proto.InternalMessageInfo

func (m *VerifyRsp) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *VerifyRsp) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*VerifyReq)(nil), "rpc.VerifyReq")
	proto.RegisterType((*VerifyRsp)(nil), "rpc.VerifyRsp")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0xf2, 0xe7, 0xe2,
	0x0c, 0x4b, 0x2d, 0xca, 0x4c, 0xab, 0x0c, 0x4a, 0x2d, 0x14, 0x12, 0xe4, 0xe2, 0x4c, 0xce, 0x48,
	0x4d, 0xce, 0x76, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe3, 0x62,
	0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x92, 0x60, 0x82, 0xf1, 0xd3, 0x13, 0x73, 0x53, 0x3d, 0x53,
	0x24, 0x98, 0x15, 0x18, 0x35, 0x58, 0x85, 0xf8, 0xb9, 0xd8, 0xcb, 0x52, 0x8b, 0x8a, 0x33, 0xf3,
	0xf3, 0x24, 0x58, 0x40, 0x0a, 0x94, 0x74, 0xe1, 0x06, 0x16, 0x17, 0x80, 0x64, 0x8b, 0x4b, 0x93,
	0x93, 0x53, 0x8b, 0x8b, 0xc1, 0xc6, 0x71, 0x80, 0x04, 0x12, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a,
	0xc0, 0xe6, 0xf1, 0x18, 0x59, 0x70, 0xb1, 0xfa, 0x80, 0xdc, 0x24, 0xa4, 0xcf, 0xc5, 0x5d, 0x06,
	0xd6, 0x07, 0xe1, 0xf2, 0xe9, 0x15, 0x15, 0x24, 0xeb, 0xc1, 0x9d, 0x26, 0x85, 0xc2, 0x2f, 0x2e,
	0x50, 0x62, 0x70, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x67, 0x3c, 0x96, 0x63, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0xad, 0x0b, 0x1e, 0x90, 0xde,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	// login verify code
	VerifyLogin(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRsp, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) VerifyLogin(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRsp, error) {
	out := new(VerifyRsp)
	err := c.cc.Invoke(ctx, "/rpc.Login/verifyLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	// login verify code
	VerifyLogin(context.Context, *VerifyReq) (*VerifyRsp, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_VerifyLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).VerifyLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Login/VerifyLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).VerifyLogin(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "verifyLogin",
			Handler:    _Login_VerifyLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}

func (m *VerifyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CheckCode != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(*m.CheckCode)))
		i += copy(dAtA[i:], *m.CheckCode)
	}
	if m.Server != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(*m.Server)))
		i += copy(dAtA[i:], *m.Server)
	}
	if m.GameId != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLogin(dAtA, i, uint64(*m.GameId))
	}
	if m.Version != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(*m.Version)))
		i += copy(dAtA[i:], *m.Version)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *VerifyRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyRsp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success != nil {
		dAtA[i] = 0x8
		i++
		if *m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Account != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLogin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VerifyReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckCode != nil {
		l = len(*m.CheckCode)
		n += 1 + l + sovLogin(uint64(l))
	}
	if m.Server != nil {
		l = len(*m.Server)
		n += 1 + l + sovLogin(uint64(l))
	}
	if m.GameId != nil {
		n += 1 + sovLogin(uint64(*m.GameId))
	}
	if m.Version != nil {
		l = len(*m.Version)
		n += 1 + l + sovLogin(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VerifyRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success != nil {
		n += 2
	}
	if m.Account != nil {
		l = len(m.Account)
		n += 1 + l + sovLogin(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLogin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLogin(x uint64) (n int) {
	return sovLogin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerifyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: VerifyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.CheckCode = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Server", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Server = &s
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.GameId = &v
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Version = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerifyRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: VerifyRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
			m.Success = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = append(m.Account[:0], dAtA[iNdEx:postIndex]...)
			if m.Account == nil {
				m.Account = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLogin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
			if length < 0 {
				return 0, ErrInvalidLengthLogin
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLogin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLogin
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
				next, err := skipLogin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLogin
				}
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
	ErrInvalidLengthLogin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogin   = fmt.Errorf("proto: integer overflow")
)
