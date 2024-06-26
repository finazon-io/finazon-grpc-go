// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: datasets.proto

package finazon

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

type DatasetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DatasetItem) Reset() {
	*x = DatasetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetItem) ProtoMessage() {}

func (x *DatasetItem) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetItem.ProtoReflect.Descriptor instead.
func (*DatasetItem) Descriptor() ([]byte, []int) {
	return file_datasets_proto_rawDescGZIP(), []int{0}
}

func (x *DatasetItem) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DatasetItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDatasetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetDatasetsRequest) Reset() {
	*x = GetDatasetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatasetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatasetsRequest) ProtoMessage() {}

func (x *GetDatasetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatasetsRequest.ProtoReflect.Descriptor instead.
func (*GetDatasetsRequest) Descriptor() ([]byte, []int) {
	return file_datasets_proto_rawDescGZIP(), []int{1}
}

func (x *GetDatasetsRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetDatasetsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetDatasetsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetDatasetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*DatasetItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetDatasetsResponse) Reset() {
	*x = GetDatasetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatasetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatasetsResponse) ProtoMessage() {}

func (x *GetDatasetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatasetsResponse.ProtoReflect.Descriptor instead.
func (*GetDatasetsResponse) Descriptor() ([]byte, []int) {
	return file_datasets_proto_rawDescGZIP(), []int{2}
}

func (x *GetDatasetsResponse) GetResult() []*DatasetItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_datasets_proto protoreflect.FileDescriptor

var file_datasets_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x0b, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x59, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x5d, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x12, 0x1b, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datasets_proto_rawDescOnce sync.Once
	file_datasets_proto_rawDescData = file_datasets_proto_rawDesc
)

func file_datasets_proto_rawDescGZIP() []byte {
	file_datasets_proto_rawDescOnce.Do(func() {
		file_datasets_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasets_proto_rawDescData)
	})
	return file_datasets_proto_rawDescData
}

var file_datasets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datasets_proto_goTypes = []interface{}{
	(*DatasetItem)(nil),         // 0: finazon.DatasetItem
	(*GetDatasetsRequest)(nil),  // 1: finazon.GetDatasetsRequest
	(*GetDatasetsResponse)(nil), // 2: finazon.GetDatasetsResponse
}
var file_datasets_proto_depIdxs = []int32{
	0, // 0: finazon.GetDatasetsResponse.result:type_name -> finazon.DatasetItem
	1, // 1: finazon.DatasetsService.GetDatasets:input_type -> finazon.GetDatasetsRequest
	2, // 2: finazon.DatasetsService.GetDatasets:output_type -> finazon.GetDatasetsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datasets_proto_init() }
func file_datasets_proto_init() {
	if File_datasets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatasetItem); i {
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
		file_datasets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatasetsRequest); i {
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
		file_datasets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatasetsResponse); i {
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
			RawDescriptor: file_datasets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datasets_proto_goTypes,
		DependencyIndexes: file_datasets_proto_depIdxs,
		MessageInfos:      file_datasets_proto_msgTypes,
	}.Build()
	File_datasets_proto = out.File
	file_datasets_proto_rawDesc = nil
	file_datasets_proto_goTypes = nil
	file_datasets_proto_depIdxs = nil
}
