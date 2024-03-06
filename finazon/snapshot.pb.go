// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: snapshot.proto

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

type GetSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Publisher     string `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Ticker        string `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Market        string `protobuf:"bytes,3,opt,name=market,proto3" json:"market,omitempty"`
	Mic           string `protobuf:"bytes,4,opt,name=mic,proto3" json:"mic,omitempty"`
	Country       string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Cqs           string `protobuf:"bytes,6,opt,name=cqs,proto3" json:"cqs,omitempty"`
	Cik           string `protobuf:"bytes,7,opt,name=cik,proto3" json:"cik,omitempty"`
	Cusip         string `protobuf:"bytes,8,opt,name=cusip,proto3" json:"cusip,omitempty"`
	Isin          string `protobuf:"bytes,9,opt,name=isin,proto3" json:"isin,omitempty"`
	CompositeFigi string `protobuf:"bytes,10,opt,name=composite_figi,json=compositeFigi,proto3" json:"composite_figi,omitempty"`
	ShareFigi     string `protobuf:"bytes,11,opt,name=share_figi,json=shareFigi,proto3" json:"share_figi,omitempty"`
	Lei           string `protobuf:"bytes,12,opt,name=lei,proto3" json:"lei,omitempty"`
}

func (x *GetSnapshotRequest) Reset() {
	*x = GetSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnapshotRequest) ProtoMessage() {}

func (x *GetSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnapshotRequest.ProtoReflect.Descriptor instead.
func (*GetSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *GetSnapshotRequest) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *GetSnapshotRequest) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *GetSnapshotRequest) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *GetSnapshotRequest) GetMic() string {
	if x != nil {
		return x.Mic
	}
	return ""
}

func (x *GetSnapshotRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetSnapshotRequest) GetCqs() string {
	if x != nil {
		return x.Cqs
	}
	return ""
}

func (x *GetSnapshotRequest) GetCik() string {
	if x != nil {
		return x.Cik
	}
	return ""
}

func (x *GetSnapshotRequest) GetCusip() string {
	if x != nil {
		return x.Cusip
	}
	return ""
}

func (x *GetSnapshotRequest) GetIsin() string {
	if x != nil {
		return x.Isin
	}
	return ""
}

func (x *GetSnapshotRequest) GetCompositeFigi() string {
	if x != nil {
		return x.CompositeFigi
	}
	return ""
}

func (x *GetSnapshotRequest) GetShareFigi() string {
	if x != nil {
		return x.ShareFigi
	}
	return ""
}

func (x *GetSnapshotRequest) GetLei() string {
	if x != nil {
		return x.Lei
	}
	return ""
}

type SnapshotOhlcv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open   float64 `protobuf:"fixed64,1,opt,name=open,proto3" json:"open,omitempty"`
	High   float64 `protobuf:"fixed64,2,opt,name=high,proto3" json:"high,omitempty"`
	Low    float64 `protobuf:"fixed64,3,opt,name=low,proto3" json:"low,omitempty"`
	Close  float64 `protobuf:"fixed64,4,opt,name=close,proto3" json:"close,omitempty"`
	Volume float64 `protobuf:"fixed64,5,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *SnapshotOhlcv) Reset() {
	*x = SnapshotOhlcv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotOhlcv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotOhlcv) ProtoMessage() {}

func (x *SnapshotOhlcv) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotOhlcv.ProtoReflect.Descriptor instead.
func (*SnapshotOhlcv) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotOhlcv) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *SnapshotOhlcv) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *SnapshotOhlcv) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *SnapshotOhlcv) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *SnapshotOhlcv) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type SnapshotLastTrade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventAt int64   `protobuf:"varint,1,opt,name=event_at,json=eventAt,proto3" json:"event_at,omitempty"`
	Price   float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Shares  int64   `protobuf:"varint,3,opt,name=shares,proto3" json:"shares,omitempty"`
}

func (x *SnapshotLastTrade) Reset() {
	*x = SnapshotLastTrade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotLastTrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotLastTrade) ProtoMessage() {}

func (x *SnapshotLastTrade) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotLastTrade.ProtoReflect.Descriptor instead.
func (*SnapshotLastTrade) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *SnapshotLastTrade) GetEventAt() int64 {
	if x != nil {
		return x.EventAt
	}
	return 0
}

func (x *SnapshotLastTrade) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SnapshotLastTrade) GetShares() int64 {
	if x != nil {
		return x.Shares
	}
	return 0
}

type SnapshotLastFiftyTwoWeek struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	High          float64 `protobuf:"fixed64,1,opt,name=high,proto3" json:"high,omitempty"`
	HighAt        int64   `protobuf:"varint,2,opt,name=high_at,json=highAt,proto3" json:"high_at,omitempty"`
	Low           float64 `protobuf:"fixed64,3,opt,name=low,proto3" json:"low,omitempty"`
	LowAt         int64   `protobuf:"varint,4,opt,name=low_at,json=lowAt,proto3" json:"low_at,omitempty"`
	Change        float64 `protobuf:"fixed64,5,opt,name=change,proto3" json:"change,omitempty"`
	ChangePercent float64 `protobuf:"fixed64,6,opt,name=change_percent,json=changePercent,proto3" json:"change_percent,omitempty"`
	AverageVolume int64   `protobuf:"varint,7,opt,name=average_volume,json=averageVolume,proto3" json:"average_volume,omitempty"`
}

func (x *SnapshotLastFiftyTwoWeek) Reset() {
	*x = SnapshotLastFiftyTwoWeek{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotLastFiftyTwoWeek) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotLastFiftyTwoWeek) ProtoMessage() {}

func (x *SnapshotLastFiftyTwoWeek) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotLastFiftyTwoWeek.ProtoReflect.Descriptor instead.
func (*SnapshotLastFiftyTwoWeek) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *SnapshotLastFiftyTwoWeek) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *SnapshotLastFiftyTwoWeek) GetHighAt() int64 {
	if x != nil {
		return x.HighAt
	}
	return 0
}

func (x *SnapshotLastFiftyTwoWeek) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *SnapshotLastFiftyTwoWeek) GetLowAt() int64 {
	if x != nil {
		return x.LowAt
	}
	return 0
}

func (x *SnapshotLastFiftyTwoWeek) GetChange() float64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *SnapshotLastFiftyTwoWeek) GetChangePercent() float64 {
	if x != nil {
		return x.ChangePercent
	}
	return 0
}

func (x *SnapshotLastFiftyTwoWeek) GetAverageVolume() int64 {
	if x != nil {
		return x.AverageVolume
	}
	return 0
}

type SnapshotChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DayChangePercent   float64 `protobuf:"fixed64,1,opt,name=day_change_percent,json=dayChangePercent,proto3" json:"day_change_percent,omitempty"`
	WeekChangePercent  float64 `protobuf:"fixed64,2,opt,name=week_change_percent,json=weekChangePercent,proto3" json:"week_change_percent,omitempty"`
	MonthChangePercent float64 `protobuf:"fixed64,3,opt,name=month_change_percent,json=monthChangePercent,proto3" json:"month_change_percent,omitempty"`
}

func (x *SnapshotChange) Reset() {
	*x = SnapshotChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotChange) ProtoMessage() {}

func (x *SnapshotChange) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotChange.ProtoReflect.Descriptor instead.
func (*SnapshotChange) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{4}
}

func (x *SnapshotChange) GetDayChangePercent() float64 {
	if x != nil {
		return x.DayChangePercent
	}
	return 0
}

func (x *SnapshotChange) GetWeekChangePercent() float64 {
	if x != nil {
		return x.WeekChangePercent
	}
	return 0
}

func (x *SnapshotChange) GetMonthChangePercent() float64 {
	if x != nil {
		return x.MonthChangePercent
	}
	return 0
}

type GetSnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastDay          *SnapshotOhlcv            `protobuf:"bytes,1,opt,name=last_day,json=lastDay,proto3" json:"last_day,omitempty"`
	LastMonth        *SnapshotOhlcv            `protobuf:"bytes,2,opt,name=last_month,json=lastMonth,proto3" json:"last_month,omitempty"`
	LastTrade        *SnapshotLastTrade        `protobuf:"bytes,3,opt,name=last_trade,json=lastTrade,proto3" json:"last_trade,omitempty"`
	PreviousDay      *SnapshotOhlcv            `protobuf:"bytes,4,opt,name=previous_day,json=previousDay,proto3" json:"previous_day,omitempty"`
	LastFiftyTwoWeek *SnapshotLastFiftyTwoWeek `protobuf:"bytes,5,opt,name=last_fifty_two_week,json=lastFiftyTwoWeek,proto3" json:"last_fifty_two_week,omitempty"`
	Change           *SnapshotChange           `protobuf:"bytes,6,opt,name=change,proto3" json:"change,omitempty"`
}

func (x *GetSnapshotResponse) Reset() {
	*x = GetSnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnapshotResponse) ProtoMessage() {}

func (x *GetSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnapshotResponse.ProtoReflect.Descriptor instead.
func (*GetSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{5}
}

func (x *GetSnapshotResponse) GetLastDay() *SnapshotOhlcv {
	if x != nil {
		return x.LastDay
	}
	return nil
}

func (x *GetSnapshotResponse) GetLastMonth() *SnapshotOhlcv {
	if x != nil {
		return x.LastMonth
	}
	return nil
}

func (x *GetSnapshotResponse) GetLastTrade() *SnapshotLastTrade {
	if x != nil {
		return x.LastTrade
	}
	return nil
}

func (x *GetSnapshotResponse) GetPreviousDay() *SnapshotOhlcv {
	if x != nil {
		return x.PreviousDay
	}
	return nil
}

func (x *GetSnapshotResponse) GetLastFiftyTwoWeek() *SnapshotLastFiftyTwoWeek {
	if x != nil {
		return x.LastFiftyTwoWeek
	}
	return nil
}

func (x *GetSnapshotResponse) GetChange() *SnapshotChange {
	if x != nil {
		return x.Change
	}
	return nil
}

var File_snapshot_proto protoreflect.FileDescriptor

var file_snapshot_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x22, 0xb4, 0x02, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x63,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x71,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x71, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x69, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x75, 0x73, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x75, 0x73, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x73, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x67, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x46, 0x69, 0x67, 0x69, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x66, 0x69, 0x67, 0x69, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x46, 0x69, 0x67, 0x69, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x65, 0x69, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x65, 0x69,
	0x22, 0x77, 0x0a, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4f, 0x68, 0x6c, 0x63,
	0x76, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x11, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x18, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x46, 0x69, 0x66, 0x74, 0x79, 0x54, 0x77, 0x6f,
	0x57, 0x65, 0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x69, 0x67, 0x68,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x69, 0x67, 0x68, 0x41,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6c, 0x6f, 0x77, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f, 0x77, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x61, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x10, 0x64, 0x61, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x13, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11,
	0x77, 0x65, 0x65, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x12, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x22, 0xf8, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x4f, 0x68, 0x6c, 0x63, 0x76, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x79, 0x12, 0x35,
	0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x4f, 0x68, 0x6c, 0x63, 0x76, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x69, 0x6e, 0x61,
	0x7a, 0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x39, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x64, 0x61, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4f, 0x68, 0x6c, 0x63, 0x76, 0x52, 0x0b,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x44, 0x61, 0x79, 0x12, 0x50, 0x0a, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x66, 0x74, 0x79, 0x5f, 0x74, 0x77, 0x6f, 0x5f, 0x77, 0x65,
	0x65, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a,
	0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x46,
	0x69, 0x66, 0x74, 0x79, 0x54, 0x77, 0x6f, 0x57, 0x65, 0x65, 0x6b, 0x52, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x46, 0x69, 0x66, 0x74, 0x79, 0x54, 0x77, 0x6f, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x2f, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x32, 0x5d,
	0x0a, 0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x12, 0x1b, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x66, 0x69, 0x6e, 0x61, 0x7a, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_snapshot_proto_rawDescOnce sync.Once
	file_snapshot_proto_rawDescData = file_snapshot_proto_rawDesc
)

func file_snapshot_proto_rawDescGZIP() []byte {
	file_snapshot_proto_rawDescOnce.Do(func() {
		file_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_snapshot_proto_rawDescData)
	})
	return file_snapshot_proto_rawDescData
}

var file_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_snapshot_proto_goTypes = []interface{}{
	(*GetSnapshotRequest)(nil),       // 0: finazon.GetSnapshotRequest
	(*SnapshotOhlcv)(nil),            // 1: finazon.SnapshotOhlcv
	(*SnapshotLastTrade)(nil),        // 2: finazon.SnapshotLastTrade
	(*SnapshotLastFiftyTwoWeek)(nil), // 3: finazon.SnapshotLastFiftyTwoWeek
	(*SnapshotChange)(nil),           // 4: finazon.SnapshotChange
	(*GetSnapshotResponse)(nil),      // 5: finazon.GetSnapshotResponse
}
var file_snapshot_proto_depIdxs = []int32{
	1, // 0: finazon.GetSnapshotResponse.last_day:type_name -> finazon.SnapshotOhlcv
	1, // 1: finazon.GetSnapshotResponse.last_month:type_name -> finazon.SnapshotOhlcv
	2, // 2: finazon.GetSnapshotResponse.last_trade:type_name -> finazon.SnapshotLastTrade
	1, // 3: finazon.GetSnapshotResponse.previous_day:type_name -> finazon.SnapshotOhlcv
	3, // 4: finazon.GetSnapshotResponse.last_fifty_two_week:type_name -> finazon.SnapshotLastFiftyTwoWeek
	4, // 5: finazon.GetSnapshotResponse.change:type_name -> finazon.SnapshotChange
	0, // 6: finazon.SnapshotService.GetSnapshot:input_type -> finazon.GetSnapshotRequest
	5, // 7: finazon.SnapshotService.GetSnapshot:output_type -> finazon.GetSnapshotResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_snapshot_proto_init() }
func file_snapshot_proto_init() {
	if File_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnapshotRequest); i {
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
		file_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotOhlcv); i {
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
		file_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotLastTrade); i {
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
		file_snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotLastFiftyTwoWeek); i {
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
		file_snapshot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotChange); i {
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
		file_snapshot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSnapshotResponse); i {
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
			RawDescriptor: file_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_snapshot_proto_goTypes,
		DependencyIndexes: file_snapshot_proto_depIdxs,
		MessageInfos:      file_snapshot_proto_msgTypes,
	}.Build()
	File_snapshot_proto = out.File
	file_snapshot_proto_rawDesc = nil
	file_snapshot_proto_goTypes = nil
	file_snapshot_proto_depIdxs = nil
}