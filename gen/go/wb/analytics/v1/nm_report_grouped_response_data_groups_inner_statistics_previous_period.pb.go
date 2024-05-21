//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam).
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/analytics/v1/nm_report_grouped_response_data_groups_inner_statistics_previous_period.proto

package wbANALYTICS

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

type NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Начало периода
	Begin string `protobuf:"bytes,1,opt,name=begin,proto3" json:"begin,omitempty"`
	// Конец периода
	End string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	// Количество переходов в КТ
	OpenCardCount int32 `protobuf:"varint,3,opt,name=openCardCount,proto3" json:"openCardCount,omitempty"`
	// Положили в корзину, штук
	AddToCartCount int32 `protobuf:"varint,4,opt,name=addToCartCount,proto3" json:"addToCartCount,omitempty"`
	// Заказли товаров, штук
	OrdersCount int32 `protobuf:"varint,5,opt,name=ordersCount,proto3" json:"ordersCount,omitempty"`
	// Заказали на сумму, руб.
	OrdersSumRub float32 `protobuf:"fixed32,6,opt,name=ordersSumRub,proto3" json:"ordersSumRub,omitempty"`
	// Выкупили товаров, штук
	BuyoutsCount int32 `protobuf:"varint,7,opt,name=buyoutsCount,proto3" json:"buyoutsCount,omitempty"`
	// Выкупили на сумму, руб.
	BuyoutsSumRub float32 `protobuf:"fixed32,8,opt,name=buyoutsSumRub,proto3" json:"buyoutsSumRub,omitempty"`
	// Отменили товаров, штук
	CancelCount int32 `protobuf:"varint,9,opt,name=cancelCount,proto3" json:"cancelCount,omitempty"`
	// Отменили на сумму, руб
	CancelSumRub int32 `protobuf:"varint,10,opt,name=cancelSumRub,proto3" json:"cancelSumRub,omitempty"`
	// Средняя цена, руб.
	AvgPriceRub float32 `protobuf:"fixed32,11,opt,name=avgPriceRub,proto3" json:"avgPriceRub,omitempty"`
	// Среднее количество заказов в день, шт.
	AvgOrdersCountPerDay float32                                                                    `protobuf:"fixed32,12,opt,name=avgOrdersCountPerDay,proto3" json:"avgOrdersCountPerDay,omitempty"`
	Conversions          *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions `protobuf:"bytes,13,opt,name=conversions,proto3" json:"conversions,omitempty"`
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) Reset() {
	*x = NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) ProtoMessage() {}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) ProtoReflect() protoreflect.Message {
	mi := &file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod.ProtoReflect.Descriptor instead.
func (*NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) Descriptor() ([]byte, []int) {
	return file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescGZIP(), []int{0}
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetBegin() string {
	if x != nil {
		return x.Begin
	}
	return ""
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetOpenCardCount() int32 {
	if x != nil {
		return x.OpenCardCount
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetAddToCartCount() int32 {
	if x != nil {
		return x.AddToCartCount
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetOrdersCount() int32 {
	if x != nil {
		return x.OrdersCount
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetOrdersSumRub() float32 {
	if x != nil {
		return x.OrdersSumRub
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetBuyoutsCount() int32 {
	if x != nil {
		return x.BuyoutsCount
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetBuyoutsSumRub() float32 {
	if x != nil {
		return x.BuyoutsSumRub
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetCancelCount() int32 {
	if x != nil {
		return x.CancelCount
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetCancelSumRub() int32 {
	if x != nil {
		return x.CancelSumRub
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetAvgPriceRub() float32 {
	if x != nil {
		return x.AvgPriceRub
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetAvgOrdersCountPerDay() float32 {
	if x != nil {
		return x.AvgOrdersCountPerDay
	}
	return 0
}

func (x *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod) GetConversions() *NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions {
	if x != nil {
		return x.Conversions
	}
	return nil
}

var File_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto protoreflect.FileDescriptor

var file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDesc = []byte{
	0x0a, 0x5d, 0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6d, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x77, 0x62, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x69, 0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6d, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x04, 0x0a, 0x3e,
	0x4e, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x43, 0x61,
	0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6f,
	0x70, 0x65, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x61, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x53, 0x75, 0x6d, 0x52, 0x75, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x53, 0x75, 0x6d, 0x52, 0x75, 0x62, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x75,
	0x79, 0x6f, 0x75, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x62, 0x75, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x62, 0x75, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x53, 0x75, 0x6d, 0x52, 0x75, 0x62, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x62, 0x75, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x53, 0x75,
	0x6d, 0x52, 0x75, 0x62, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x53, 0x75, 0x6d, 0x52, 0x75, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x53, 0x75, 0x6d, 0x52, 0x75, 0x62, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x76,
	0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x75, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x61, 0x76, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x75, 0x62, 0x12, 0x32, 0x0a, 0x14,
	0x61, 0x76, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65,
	0x72, 0x44, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x61, 0x76, 0x67, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x44, 0x61, 0x79,
	0x12, 0x7c, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x77, 0x62, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x49,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62,
	0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62,
	0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescOnce sync.Once
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescData = file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDesc
)

func file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescGZIP() []byte {
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescOnce.Do(func() {
		file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescData)
	})
	return file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDescData
}

var file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_goTypes = []interface{}{
	(*NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod)(nil),            // 0: wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod
	(*NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions)(nil), // 1: wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions
}
var file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_depIdxs = []int32{
	1, // 0: wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod.conversions:type_name -> wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_init()
}
func file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_init() {
	if File_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto != nil {
		return
	}
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_conversions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod); i {
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
			RawDescriptor: file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_goTypes,
		DependencyIndexes: file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_depIdxs,
		MessageInfos:      file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_msgTypes,
	}.Build()
	File_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto = out.File
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_rawDesc = nil
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_goTypes = nil
	file_wb_analytics_v1_nm_report_grouped_response_data_groups_inner_statistics_previous_period_proto_depIdxs = nil
}