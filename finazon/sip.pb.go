// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: sip.proto

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

type SipTradesItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TradeDate       int64   `protobuf:"varint,1,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"`
	MarketCenter    string  `protobuf:"bytes,2,opt,name=market_center,json=marketCenter,proto3" json:"market_center,omitempty"`
	Topic           string  `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Price           float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Quantity        int64   `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SalesConditions string  `protobuf:"bytes,6,opt,name=sales_conditions,json=salesConditions,proto3" json:"sales_conditions,omitempty"`
}

func (x *SipTradesItem) Reset() {
	*x = SipTradesItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipTradesItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipTradesItem) ProtoMessage() {}

func (x *SipTradesItem) ProtoReflect() protoreflect.Message {
	mi := &file_sip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipTradesItem.ProtoReflect.Descriptor instead.
func (*SipTradesItem) Descriptor() ([]byte, []int) {
	return file_sip_proto_rawDescGZIP(), []int{0}
}

func (x *SipTradesItem) GetTradeDate() int64 {
	if x != nil {
		return x.TradeDate
	}
	return 0
}

func (x *SipTradesItem) GetMarketCenter() string {
	if x != nil {
		return x.MarketCenter
	}
	return ""
}

func (x *SipTradesItem) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SipTradesItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SipTradesItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *SipTradesItem) GetSalesConditions() string {
	if x != nil {
		return x.SalesConditions
	}
	return ""
}

type GetSipTradesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker        string `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Mic           string `protobuf:"bytes,2,opt,name=mic,proto3" json:"mic,omitempty"`
	Country       string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Market        string `protobuf:"bytes,4,opt,name=market,proto3" json:"market,omitempty"`
	StartAt       int64  `protobuf:"varint,5,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt         int64  `protobuf:"varint,6,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	Tape          string `protobuf:"bytes,7,opt,name=tape,proto3" json:"tape,omitempty"`
	Page          int32  `protobuf:"varint,8,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32  `protobuf:"varint,9,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Order         string `protobuf:"bytes,10,opt,name=order,proto3" json:"order,omitempty"`
	Cqs           string `protobuf:"bytes,11,opt,name=cqs,proto3" json:"cqs,omitempty"`
	Cik           string `protobuf:"bytes,12,opt,name=cik,proto3" json:"cik,omitempty"`
	Cusip         string `protobuf:"bytes,13,opt,name=cusip,proto3" json:"cusip,omitempty"`
	Isin          string `protobuf:"bytes,14,opt,name=isin,proto3" json:"isin,omitempty"`
	CompositeFigi string `protobuf:"bytes,15,opt,name=composite_figi,json=compositeFigi,proto3" json:"composite_figi,omitempty"`
	ShareFigi     string `protobuf:"bytes,16,opt,name=share_figi,json=shareFigi,proto3" json:"share_figi,omitempty"`
	Lei           string `protobuf:"bytes,17,opt,name=lei,proto3" json:"lei,omitempty"`
}

func (x *GetSipTradesRequest) Reset() {
	*x = GetSipTradesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSipTradesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSipTradesRequest) ProtoMessage() {}

func (x *GetSipTradesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSipTradesRequest.ProtoReflect.Descriptor instead.
func (*GetSipTradesRequest) Descriptor() ([]byte, []int) {
	return file_sip_proto_rawDescGZIP(), []int{1}
}

func (x *GetSipTradesRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *GetSipTradesRequest) GetMic() string {
	if x != nil {
		return x.Mic
	}
	return ""
}

func (x *GetSipTradesRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetSipTradesRequest) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *GetSipTradesRequest) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *GetSipTradesRequest) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *GetSipTradesRequest) GetTape() string {
	if x != nil {
		return x.Tape
	}
	return ""
}

func (x *GetSipTradesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetSipTradesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetSipTradesRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *GetSipTradesRequest) GetCqs() string {
	if x != nil {
		return x.Cqs
	}
	return ""
}

func (x *GetSipTradesRequest) GetCik() string {
	if x != nil {
		return x.Cik
	}
	return ""
}

func (x *GetSipTradesRequest) GetCusip() string {
	if x != nil {
		return x.Cusip
	}
	return ""
}

func (x *GetSipTradesRequest) GetIsin() string {
	if x != nil {
		return x.Isin
	}
	return ""
}

func (x *GetSipTradesRequest) GetCompositeFigi() string {
	if x != nil {
		return x.CompositeFigi
	}
	return ""
}

func (x *GetSipTradesRequest) GetShareFigi() string {
	if x != nil {
		return x.ShareFigi
	}
	return ""
}

func (x *GetSipTradesRequest) GetLei() string {
	if x != nil {
		return x.Lei
	}
	return ""
}

type GetSipTradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*SipTradesItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetSipTradesResponse) Reset() {
	*x = GetSipTradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSipTradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSipTradesResponse) ProtoMessage() {}

func (x *GetSipTradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSipTradesResponse.ProtoReflect.Descriptor instead.
func (*GetSipTradesResponse) Descriptor() ([]byte, []int) {
	return file_sip_proto_rawDescGZIP(), []int{2}
}

func (x *GetSipTradesResponse) GetResult() []*SipTradesItem {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetSipMarketCenterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSipMarketCenterRequest) Reset() {
	*x = GetSipMarketCenterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSipMarketCenterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSipMarketCenterRequest) ProtoMessage() {}

func (x *GetSipMarketCenterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSipMarketCenterRequest.ProtoReflect.Descriptor instead.
func (*GetSipMarketCenterRequest) Descriptor() ([]byte, []int) {
	return file_sip_proto_rawDescGZIP(), []int{3}
}

type SipMarketCenterItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SipMarketCenterItem) Reset() {
	*x = SipMarketCenterItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SipMarketCenterItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SipMarketCenterItem) ProtoMessage() {}

func (x *SipMarketCenterItem) ProtoReflect() protoreflect.Message {
	mi := &file_sip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SipMarketCenterItem.ProtoReflect.Descriptor instead.
func (*SipMarketCenterItem) Descriptor() ([]byte, []int) {
	return file_sip_proto_rawDescGZIP(), []int{4}
}

func (x *SipMarketCenterItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SipMarketCenterItem) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetSipMarketCenterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*SipMarketCenterItem `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetSipMarketCenterResponse) Reset() {
	*x = GetSipMarketCenterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSipMarketCenterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSipMarketCenterResponse) ProtoMessage() {}

func (x *GetSipMarketCenterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSipMarketCenterResponse.ProtoReflect.Descriptor instead.
func (*GetSipMarketCenterResponse) Descriptor() ([]byte, []int) {
	return file_sip_proto_rawDescGZIP(), []int{5}
}

func (x *GetSipMarketCenterResponse) GetResult() []*SipMarketCenterItem {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_sip_proto protoreflect.FileDescriptor

var file_sip_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x69, 0x6e,
	0x61, 0x7a, 0x6f, 0x6e, 0x22, 0xc6, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x70, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x61,
	0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa4, 0x03,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x69, 0x70, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x63, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x71, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x71, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x69, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x75, 0x73, 0x69, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x75, 0x73, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x69, 0x6e,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x67, 0x69, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x46,
	0x69, 0x67, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x66, 0x69, 0x67,
	0x69, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x46, 0x69,
	0x67, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x69, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6c, 0x65, 0x69, 0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x69, 0x70, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66,
	0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x70, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1b, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x13, 0x53, 0x69, 0x70,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53,
	0x69, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e,
	0x2e, 0x53, 0x69, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xb6, 0x01, 0x0a,
	0x0a, 0x53, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x70, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x70, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x70, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x70, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sip_proto_rawDescOnce sync.Once
	file_sip_proto_rawDescData = file_sip_proto_rawDesc
)

func file_sip_proto_rawDescGZIP() []byte {
	file_sip_proto_rawDescOnce.Do(func() {
		file_sip_proto_rawDescData = protoimpl.X.CompressGZIP(file_sip_proto_rawDescData)
	})
	return file_sip_proto_rawDescData
}

var file_sip_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sip_proto_goTypes = []interface{}{
	(*SipTradesItem)(nil),              // 0: finazon.SipTradesItem
	(*GetSipTradesRequest)(nil),        // 1: finazon.GetSipTradesRequest
	(*GetSipTradesResponse)(nil),       // 2: finazon.GetSipTradesResponse
	(*GetSipMarketCenterRequest)(nil),  // 3: finazon.GetSipMarketCenterRequest
	(*SipMarketCenterItem)(nil),        // 4: finazon.SipMarketCenterItem
	(*GetSipMarketCenterResponse)(nil), // 5: finazon.GetSipMarketCenterResponse
}
var file_sip_proto_depIdxs = []int32{
	0, // 0: finazon.GetSipTradesResponse.result:type_name -> finazon.SipTradesItem
	4, // 1: finazon.GetSipMarketCenterResponse.result:type_name -> finazon.SipMarketCenterItem
	1, // 2: finazon.SipService.GetTrades:input_type -> finazon.GetSipTradesRequest
	3, // 3: finazon.SipService.GetMarketCenter:input_type -> finazon.GetSipMarketCenterRequest
	2, // 4: finazon.SipService.GetTrades:output_type -> finazon.GetSipTradesResponse
	5, // 5: finazon.SipService.GetMarketCenter:output_type -> finazon.GetSipMarketCenterResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sip_proto_init() }
func file_sip_proto_init() {
	if File_sip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipTradesItem); i {
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
		file_sip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSipTradesRequest); i {
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
		file_sip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSipTradesResponse); i {
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
		file_sip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSipMarketCenterRequest); i {
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
		file_sip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SipMarketCenterItem); i {
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
		file_sip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSipMarketCenterResponse); i {
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
			RawDescriptor: file_sip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sip_proto_goTypes,
		DependencyIndexes: file_sip_proto_depIdxs,
		MessageInfos:      file_sip_proto_msgTypes,
	}.Build()
	File_sip_proto = out.File
	file_sip_proto_rawDesc = nil
	file_sip_proto_goTypes = nil
	file_sip_proto_depIdxs = nil
}
