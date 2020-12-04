// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: arith.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 算术运算请求结构
type ArithRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *ArithRequest) Reset() {
	*x = ArithRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arith_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArithRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArithRequest) ProtoMessage() {}

func (x *ArithRequest) ProtoReflect() protoreflect.Message {
	mi := &file_arith_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArithRequest.ProtoReflect.Descriptor instead.
func (*ArithRequest) Descriptor() ([]byte, []int) {
	return file_arith_proto_rawDescGZIP(), []int{0}
}

func (x *ArithRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *ArithRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

// 算术运算响应结构
type ArithResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pro int32 `protobuf:"varint,1,opt,name=pro,proto3" json:"pro,omitempty"` // 乘积
}

func (x *ArithResponse) Reset() {
	*x = ArithResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_arith_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArithResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArithResponse) ProtoMessage() {}

func (x *ArithResponse) ProtoReflect() protoreflect.Message {
	mi := &file_arith_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArithResponse.ProtoReflect.Descriptor instead.
func (*ArithResponse) Descriptor() ([]byte, []int) {
	return file_arith_proto_rawDescGZIP(), []int{1}
}

func (x *ArithResponse) GetPro() int32 {
	if x != nil {
		return x.Pro
	}
	return 0
}

var File_arith_proto protoreflect.FileDescriptor

var file_arith_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x72, 0x69, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0c, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62,
	0x22, 0x21, 0x0a, 0x0d, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x70, 0x72, 0x6f, 0x32, 0x45, 0x0a, 0x0c, 0x41, 0x72, 0x69, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x69, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x69,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_arith_proto_rawDescOnce sync.Once
	file_arith_proto_rawDescData = file_arith_proto_rawDesc
)

func file_arith_proto_rawDescGZIP() []byte {
	file_arith_proto_rawDescOnce.Do(func() {
		file_arith_proto_rawDescData = protoimpl.X.CompressGZIP(file_arith_proto_rawDescData)
	})
	return file_arith_proto_rawDescData
}

var file_arith_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_arith_proto_goTypes = []interface{}{
	(*ArithRequest)(nil),  // 0: proto.ArithRequest
	(*ArithResponse)(nil), // 1: proto.ArithResponse
}
var file_arith_proto_depIdxs = []int32{
	0, // 0: proto.ArithService.multiply:input_type -> proto.ArithRequest
	1, // 1: proto.ArithService.multiply:output_type -> proto.ArithResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_arith_proto_init() }
func file_arith_proto_init() {
	if File_arith_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_arith_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArithRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_arith_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArithResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_arith_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_arith_proto_goTypes,
		DependencyIndexes: file_arith_proto_depIdxs,
		MessageInfos:      file_arith_proto_msgTypes,
	}.Build()
	File_arith_proto = out.File
	file_arith_proto_rawDesc = nil
	file_arith_proto_goTypes = nil
	file_arith_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ArithServiceClient is the client API for ArithService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArithServiceClient interface {
	Multiply(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error)
}

type arithServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithServiceClient(cc grpc.ClientConnInterface) ArithServiceClient {
	return &arithServiceClient{cc}
}

func (c *arithServiceClient) Multiply(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error) {
	out := new(ArithResponse)
	err := c.cc.Invoke(ctx, "/proto.ArithService/multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithServiceServer is the server API for ArithService service.
type ArithServiceServer interface {
	Multiply(context.Context, *ArithRequest) (*ArithResponse, error)
}

// UnimplementedArithServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArithServiceServer struct {
}

func (*UnimplementedArithServiceServer) Multiply(context.Context, *ArithRequest) (*ArithResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}

func RegisterArithServiceServer(s *grpc.Server, srv ArithServiceServer) {
	s.RegisterService(&_ArithService_serviceDesc, srv)
}

func _ArithService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ArithService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServiceServer).Multiply(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArithService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ArithService",
	HandlerType: (*ArithServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "multiply",
			Handler:    _ArithService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arith.proto",
}