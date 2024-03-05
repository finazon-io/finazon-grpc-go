// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: crypto.proto

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

type CryptoTimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Open      float64 `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	Close     float64 `protobuf:"fixed64,3,opt,name=close,proto3" json:"close,omitempty"`
	High      float64 `protobuf:"fixed64,4,opt,name=high,proto3" json:"high,omitempty"`
	Low       float64 `protobuf:"fixed64,5,opt,name=low,proto3" json:"low,omitempty"`
	Volume    float64 `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *CryptoTimeSeries) Reset() {
	*x = CryptoTimeSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoTimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoTimeSeries) ProtoMessage() {}

func (x *CryptoTimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoTimeSeries.ProtoReflect.Descriptor instead.
func (*CryptoTimeSeries) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoTimeSeries) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CryptoTimeSeries) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *CryptoTimeSeries) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *CryptoTimeSeries) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *CryptoTimeSeries) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *CryptoTimeSeries) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type GetCryptoTimeSeriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker   string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Timezone string `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Interval string `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
	StartAt  int64  `protobuf:"varint,5,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt    int64  `protobuf:"varint,6,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	Page     int32  `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Order    string `protobuf:"bytes,9,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetCryptoTimeSeriesRequest) Reset() {
	*x = GetCryptoTimeSeriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCryptoTimeSeriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCryptoTimeSeriesRequest) ProtoMessage() {}

func (x *GetCryptoTimeSeriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCryptoTimeSeriesRequest.ProtoReflect.Descriptor instead.
func (*GetCryptoTimeSeriesRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *GetCryptoTimeSeriesRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *GetCryptoTimeSeriesRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *GetCryptoTimeSeriesRequest) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *GetCryptoTimeSeriesRequest) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *GetCryptoTimeSeriesRequest) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *GetCryptoTimeSeriesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetCryptoTimeSeriesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetCryptoTimeSeriesRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetCryptoTimeSeriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*CryptoTimeSeries `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetCryptoTimeSeriesResponse) Reset() {
	*x = GetCryptoTimeSeriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCryptoTimeSeriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCryptoTimeSeriesResponse) ProtoMessage() {}

func (x *GetCryptoTimeSeriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCryptoTimeSeriesResponse.ProtoReflect.Descriptor instead.
func (*GetCryptoTimeSeriesResponse) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *GetCryptoTimeSeriesResponse) GetResult() []*CryptoTimeSeries {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_crypto_proto protoreflect.FileDescriptor

var file_crypto_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x50, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x6d, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_proto_rawDescOnce sync.Once
	file_crypto_proto_rawDescData = file_crypto_proto_rawDesc
)

func file_crypto_proto_rawDescGZIP() []byte {
	file_crypto_proto_rawDescOnce.Do(func() {
		file_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_proto_rawDescData)
	})
	return file_crypto_proto_rawDescData
}

var file_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_crypto_proto_goTypes = []interface{}{
	(*CryptoTimeSeries)(nil),            // 0: finazon.CryptoTimeSeries
	(*GetCryptoTimeSeriesRequest)(nil),  // 1: finazon.GetCryptoTimeSeriesRequest
	(*GetCryptoTimeSeriesResponse)(nil), // 2: finazon.GetCryptoTimeSeriesResponse
}
var file_crypto_proto_depIdxs = []int32{
	0, // 0: finazon.GetCryptoTimeSeriesResponse.result:type_name -> finazon.CryptoTimeSeries
	1, // 1: finazon.CryptoService.GetTimeSeries:input_type -> finazon.GetCryptoTimeSeriesRequest
	2, // 2: finazon.CryptoService.GetTimeSeries:output_type -> finazon.GetCryptoTimeSeriesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_crypto_proto_init() }
func file_crypto_proto_init() {
	if File_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoTimeSeries); i {
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
		file_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCryptoTimeSeriesRequest); i {
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
		file_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCryptoTimeSeriesResponse); i {
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
			RawDescriptor: file_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_proto_goTypes,
		DependencyIndexes: file_crypto_proto_depIdxs,
		MessageInfos:      file_crypto_proto_msgTypes,
	}.Build()
	File_crypto_proto = out.File
	file_crypto_proto_rawDesc = nil
	file_crypto_proto_goTypes = nil
	file_crypto_proto_depIdxs = nil
}
