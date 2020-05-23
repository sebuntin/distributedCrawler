// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: saveItem.proto

package itemService

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

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saveItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saveItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_saveItem_proto_rawDescGZIP(), []int{0}
}

func (x *SaveRequest) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

type SaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saveItem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saveItem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_saveItem_proto_rawDescGZIP(), []int{1}
}

func (x *SaveResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_saveItem_proto protoreflect.FileDescriptor

var file_saveItem_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x61, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x21, 0x0a,
	0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0x26, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x50, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x18, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x61, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b,
	0x69, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_saveItem_proto_rawDescOnce sync.Once
	file_saveItem_proto_rawDescData = file_saveItem_proto_rawDesc
)

func file_saveItem_proto_rawDescGZIP() []byte {
	file_saveItem_proto_rawDescOnce.Do(func() {
		file_saveItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_saveItem_proto_rawDescData)
	})
	return file_saveItem_proto_rawDescData
}

var file_saveItem_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_saveItem_proto_goTypes = []interface{}{
	(*SaveRequest)(nil),  // 0: saveService.SaveRequest
	(*SaveResponse)(nil), // 1: saveService.SaveResponse
}
var file_saveItem_proto_depIdxs = []int32{
	0, // 0: saveService.SaveService.SaveItem:input_type -> saveService.SaveRequest
	1, // 1: saveService.SaveService.SaveItem:output_type -> saveService.SaveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_saveItem_proto_init() }
func file_saveItem_proto_init() {
	if File_saveItem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_saveItem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_saveItem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResponse); i {
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
			RawDescriptor: file_saveItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_saveItem_proto_goTypes,
		DependencyIndexes: file_saveItem_proto_depIdxs,
		MessageInfos:      file_saveItem_proto_msgTypes,
	}.Build()
	File_saveItem_proto = out.File
	file_saveItem_proto_rawDesc = nil
	file_saveItem_proto_goTypes = nil
	file_saveItem_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SaveServiceClient is the client API for SaveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SaveServiceClient interface {
	SaveItem(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
}

type saveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSaveServiceClient(cc grpc.ClientConnInterface) SaveServiceClient {
	return &saveServiceClient{cc}
}

func (c *saveServiceClient) SaveItem(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/saveService.SaveService/SaveItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaveServiceServer is the server API for SaveService service.
type SaveServiceServer interface {
	SaveItem(context.Context, *SaveRequest) (*SaveResponse, error)
}

// UnimplementedSaveServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSaveServiceServer struct {
}

func (*UnimplementedSaveServiceServer) SaveItem(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveItem not implemented")
}

func RegisterSaveServiceServer(s *grpc.Server, srv SaveServiceServer) {
	s.RegisterService(&_SaveService_serviceDesc, srv)
}

func _SaveService_SaveItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaveServiceServer).SaveItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saveService.SaveService/SaveItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaveServiceServer).SaveItem(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SaveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "saveService.SaveService",
	HandlerType: (*SaveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveItem",
			Handler:    _SaveService_SaveItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saveItem.proto",
}
