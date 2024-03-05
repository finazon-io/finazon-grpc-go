// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: publishers.proto

package finazon_grpc_go

import (
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

type PublisherItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PublisherItem) Reset() {
	*x = PublisherItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publishers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherItem) ProtoMessage() {}

func (x *PublisherItem) ProtoReflect() protoreflect.Message {
	mi := &file_publishers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherItem.ProtoReflect.Descriptor instead.
func (*PublisherItem) Descriptor() ([]byte, []int) {
	return file_publishers_proto_rawDescGZIP(), []int{0}
}

func (x *PublisherItem) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PublisherItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPublishersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetPublishersRequest) Reset() {
	*x = GetPublishersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publishers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublishersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublishersRequest) ProtoMessage() {}

func (x *GetPublishersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publishers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublishersRequest.ProtoReflect.Descriptor instead.
func (*GetPublishersRequest) Descriptor() ([]byte, []int) {
	return file_publishers_proto_rawDescGZIP(), []int{1}
}

func (x *GetPublishersRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetPublishersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPublishersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetPublishersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*PublisherItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetPublishersResponse) Reset() {
	*x = GetPublishersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publishers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublishersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublishersResponse) ProtoMessage() {}

func (x *GetPublishersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publishers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublishersResponse.ProtoReflect.Descriptor instead.
func (*GetPublishersResponse) Descriptor() ([]byte, []int) {
	return file_publishers_proto_rawDescGZIP(), []int{2}
}

func (x *GetPublishersResponse) GetResult() []*PublisherItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_publishers_proto protoreflect.FileDescriptor

var file_publishers_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0d, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x47, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x64, 0x0a, 0x10, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x12,
	0x1d, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publishers_proto_rawDescOnce sync.Once
	file_publishers_proto_rawDescData = file_publishers_proto_rawDesc
)

func file_publishers_proto_rawDescGZIP() []byte {
	file_publishers_proto_rawDescOnce.Do(func() {
		file_publishers_proto_rawDescData = protoimpl.X.CompressGZIP(file_publishers_proto_rawDescData)
	})
	return file_publishers_proto_rawDescData
}

var file_publishers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_publishers_proto_goTypes = []interface{}{
	(*PublisherItem)(nil),         // 0: finazon.PublisherItem
	(*GetPublishersRequest)(nil),  // 1: finazon.GetPublishersRequest
	(*GetPublishersResponse)(nil), // 2: finazon.GetPublishersResponse
}
var file_publishers_proto_depIdxs = []int32{
	0, // 0: finazon.GetPublishersResponse.result:type_name -> finazon.PublisherItem
	1, // 1: finazon.PublisherService.GetPublishers:input_type -> finazon.GetPublishersRequest
	2, // 2: finazon.PublisherService.GetPublishers:output_type -> finazon.GetPublishersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_publishers_proto_init() }
func file_publishers_proto_init() {
	if File_publishers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publishers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherItem); i {
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
		file_publishers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublishersRequest); i {
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
		file_publishers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublishersResponse); i {
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
			RawDescriptor: file_publishers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publishers_proto_goTypes,
		DependencyIndexes: file_publishers_proto_depIdxs,
		MessageInfos:      file_publishers_proto_msgTypes,
	}.Build()
	File_publishers_proto = out.File
	file_publishers_proto_rawDesc = nil
	file_publishers_proto_goTypes = nil
	file_publishers_proto_depIdxs = nil
}