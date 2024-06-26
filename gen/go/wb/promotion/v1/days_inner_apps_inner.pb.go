//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/promotion/v1/days_inner_apps_inner.proto

package wbPromotion

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

type DaysInnerAppsInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Количество просмотров
	Views int32 `protobuf:"varint,1,opt,name=views,proto3" json:"views,omitempty"`
	// Количество кликов
	Clicks int32 `protobuf:"varint,2,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// Показатель кликабельности, отношение числа кликов к количеству показов, %
	Ctr float32 `protobuf:"fixed32,3,opt,name=ctr,proto3" json:"ctr,omitempty"`
	// Средняя стоимость клика, ₽.
	Cpc float32 `protobuf:"fixed32,4,opt,name=cpc,proto3" json:"cpc,omitempty"`
	// Затраты, ₽.
	Sum float32 `protobuf:"fixed32,5,opt,name=sum,proto3" json:"sum,omitempty"`
	// Количество добавлений товаров в корзину
	Atbs int32 `protobuf:"varint,6,opt,name=atbs,proto3" json:"atbs,omitempty"`
	// Количество заказов
	Orders int32 `protobuf:"varint,7,opt,name=orders,proto3" json:"orders,omitempty"`
	// CR(conversion rate) — это отношение количества заказов к общему количеству посещений кампании
	Cr int32 `protobuf:"varint,8,opt,name=cr,proto3" json:"cr,omitempty"`
	// Количество заказанных товаров, шт.
	Shks int32 `protobuf:"varint,9,opt,name=shks,proto3" json:"shks,omitempty"`
	// Заказов на сумму, ₽
	SumPrice float32 `protobuf:"fixed32,10,opt,name=sum_price,json=sumPrice,proto3" json:"sum_price,omitempty"`
	// Блок статистики по артикулам WB
	Nm []*DaysInnerAppsInnerNmInner `protobuf:"bytes,11,rep,name=nm,proto3" json:"nm,omitempty"`
	// Тип платформы (`1` - сайт, `32` - Android, `64` - IOS)
	AppType int32 `protobuf:"varint,12,opt,name=appType,proto3" json:"appType,omitempty"`
}

func (x *DaysInnerAppsInner) Reset() {
	*x = DaysInnerAppsInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_days_inner_apps_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DaysInnerAppsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DaysInnerAppsInner) ProtoMessage() {}

func (x *DaysInnerAppsInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_days_inner_apps_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DaysInnerAppsInner.ProtoReflect.Descriptor instead.
func (*DaysInnerAppsInner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescGZIP(), []int{0}
}

func (x *DaysInnerAppsInner) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *DaysInnerAppsInner) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *DaysInnerAppsInner) GetCtr() float32 {
	if x != nil {
		return x.Ctr
	}
	return 0
}

func (x *DaysInnerAppsInner) GetCpc() float32 {
	if x != nil {
		return x.Cpc
	}
	return 0
}

func (x *DaysInnerAppsInner) GetSum() float32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *DaysInnerAppsInner) GetAtbs() int32 {
	if x != nil {
		return x.Atbs
	}
	return 0
}

func (x *DaysInnerAppsInner) GetOrders() int32 {
	if x != nil {
		return x.Orders
	}
	return 0
}

func (x *DaysInnerAppsInner) GetCr() int32 {
	if x != nil {
		return x.Cr
	}
	return 0
}

func (x *DaysInnerAppsInner) GetShks() int32 {
	if x != nil {
		return x.Shks
	}
	return 0
}

func (x *DaysInnerAppsInner) GetSumPrice() float32 {
	if x != nil {
		return x.SumPrice
	}
	return 0
}

func (x *DaysInnerAppsInner) GetNm() []*DaysInnerAppsInnerNmInner {
	if x != nil {
		return x.Nm
	}
	return nil
}

func (x *DaysInnerAppsInner) GetAppType() int32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

var File_wb_promotion_v1_days_inner_apps_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_days_inner_apps_inner_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x70,
	0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x34,
	0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x79, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x5f,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x6d, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x12, 0x44, 0x61, 0x79, 0x73, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x70, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x70, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x74, 0x62, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61,
	0x74, 0x62, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x63,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x63, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x68, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x68, 0x6b, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x73, 0x75, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x02,
	0x6e, 0x6d, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x79, 0x73, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x6d, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x02, 0x6e, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77,
	0x62, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescData = file_wb_promotion_v1_days_inner_apps_inner_proto_rawDesc
)

func file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_days_inner_apps_inner_proto_rawDescData
}

var file_wb_promotion_v1_days_inner_apps_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_days_inner_apps_inner_proto_goTypes = []interface{}{
	(*DaysInnerAppsInner)(nil),        // 0: wb.promotion.v1.DaysInnerAppsInner
	(*DaysInnerAppsInnerNmInner)(nil), // 1: wb.promotion.v1.DaysInnerAppsInnerNmInner
}
var file_wb_promotion_v1_days_inner_apps_inner_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.DaysInnerAppsInner.nm:type_name -> wb.promotion.v1.DaysInnerAppsInnerNmInner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_days_inner_apps_inner_proto_init() }
func file_wb_promotion_v1_days_inner_apps_inner_proto_init() {
	if File_wb_promotion_v1_days_inner_apps_inner_proto != nil {
		return
	}
	file_wb_promotion_v1_days_inner_apps_inner_nm_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_days_inner_apps_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DaysInnerAppsInner); i {
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
			RawDescriptor: file_wb_promotion_v1_days_inner_apps_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_days_inner_apps_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_days_inner_apps_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_days_inner_apps_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_days_inner_apps_inner_proto = out.File
	file_wb_promotion_v1_days_inner_apps_inner_proto_rawDesc = nil
	file_wb_promotion_v1_days_inner_apps_inner_proto_goTypes = nil
	file_wb_promotion_v1_days_inner_apps_inner_proto_depIdxs = nil
}
