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
// source: wb/promotion/v1/stats_blok2.proto

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

type StatsBlok2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID баннера
	ItemId int32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	// Название бренда
	ItemName string `protobuf:"bytes,2,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	// Название категории
	CategoryName string `protobuf:"bytes,3,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	// <dl> <dt>Тип медиакампании:</dt> <dd><code>1</code> - размещение по дням</dd> <dd><code>2</code> - размещение по просмотрам</dd> </dl>
	AdvertType int32 `protobuf:"varint,4,opt,name=advert_type,json=advertType,proto3" json:"advert_type,omitempty"`
	// Место на странице
	Place int32 `protobuf:"varint,5,opt,name=place,proto3" json:"place,omitempty"`
	// Количество просмотров
	Views int32 `protobuf:"varint,6,opt,name=views,proto3" json:"views,omitempty"`
	// Количество кликов
	Clicks int32 `protobuf:"varint,7,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// CR(conversion rate) — это отношение количества заказов к общему количеству посещений медиакампании
	Cr float32 `protobuf:"fixed32,8,opt,name=cr,proto3" json:"cr,omitempty"`
	// CTR (click-through rate) — показатель кликабельности, отношение числа кликов к количеству показов в рамках медиакампании
	Ctr float32 `protobuf:"fixed32,9,opt,name=ctr,proto3" json:"ctr,omitempty"`
	// Время начала размещения
	DateFrom string `protobuf:"bytes,10,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	// Время завершения размещения
	DateTo string `protobuf:"bytes,11,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
	// Родительская категория предмета
	SubjectName string `protobuf:"bytes,12,opt,name=subject_name,json=subjectName,proto3" json:"subject_name,omitempty"`
	// Количество добавлений товаров в корзину
	Atbs int32 `protobuf:"varint,13,opt,name=atbs,proto3" json:"atbs,omitempty"`
	// Количество заказов
	Orders int32 `protobuf:"varint,14,opt,name=orders,proto3" json:"orders,omitempty"`
	// Стоимость размещения
	Price int32 `protobuf:"varint,15,opt,name=price,proto3" json:"price,omitempty"`
	// (cost per click) - цена клика по продвигаемому товару
	Cpc float32 `protobuf:"fixed32,16,opt,name=cpc,proto3" json:"cpc,omitempty"`
	// Статус медиакампании
	Status     int32               `protobuf:"varint,17,opt,name=status,proto3" json:"status,omitempty"`
	DailyStats []*DailyStats2Inner `protobuf:"bytes,18,rep,name=daily_stats,json=dailyStats,proto3" json:"daily_stats,omitempty"`
	// Стоимость размещения баннера
	Expenses int32 `protobuf:"varint,19,opt,name=expenses,proto3" json:"expenses,omitempty"`
	// Отношение количества добавлений в корзину к количеству кликов
	Cr1 float32 `protobuf:"fixed32,20,opt,name=cr1,proto3" json:"cr1,omitempty"`
	// Отношение количества заказов к количеству добавлений в корзину
	Cr2 int32 `protobuf:"varint,21,opt,name=cr2,proto3" json:"cr2,omitempty"`
}

func (x *StatsBlok2) Reset() {
	*x = StatsBlok2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_stats_blok2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsBlok2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsBlok2) ProtoMessage() {}

func (x *StatsBlok2) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_stats_blok2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsBlok2.ProtoReflect.Descriptor instead.
func (*StatsBlok2) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_stats_blok2_proto_rawDescGZIP(), []int{0}
}

func (x *StatsBlok2) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *StatsBlok2) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *StatsBlok2) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *StatsBlok2) GetAdvertType() int32 {
	if x != nil {
		return x.AdvertType
	}
	return 0
}

func (x *StatsBlok2) GetPlace() int32 {
	if x != nil {
		return x.Place
	}
	return 0
}

func (x *StatsBlok2) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *StatsBlok2) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *StatsBlok2) GetCr() float32 {
	if x != nil {
		return x.Cr
	}
	return 0
}

func (x *StatsBlok2) GetCtr() float32 {
	if x != nil {
		return x.Ctr
	}
	return 0
}

func (x *StatsBlok2) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *StatsBlok2) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *StatsBlok2) GetSubjectName() string {
	if x != nil {
		return x.SubjectName
	}
	return ""
}

func (x *StatsBlok2) GetAtbs() int32 {
	if x != nil {
		return x.Atbs
	}
	return 0
}

func (x *StatsBlok2) GetOrders() int32 {
	if x != nil {
		return x.Orders
	}
	return 0
}

func (x *StatsBlok2) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *StatsBlok2) GetCpc() float32 {
	if x != nil {
		return x.Cpc
	}
	return 0
}

func (x *StatsBlok2) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StatsBlok2) GetDailyStats() []*DailyStats2Inner {
	if x != nil {
		return x.DailyStats
	}
	return nil
}

func (x *StatsBlok2) GetExpenses() int32 {
	if x != nil {
		return x.Expenses
	}
	return 0
}

func (x *StatsBlok2) GetCr1() float32 {
	if x != nil {
		return x.Cr1
	}
	return 0
}

func (x *StatsBlok2) GetCr2() int32 {
	if x != nil {
		return x.Cr2
	}
	return 0
}

var File_wb_promotion_v1_stats_blok2_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_stats_blok2_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x6b, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x32, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7,
	0x04, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x6c, 0x6f, 0x6b, 0x32, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x63, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x63, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x74, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x74, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x74, 0x62, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x74, 0x62, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70,
	0x63, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x70, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x32, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x72, 0x31, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x63, 0x72, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x72, 0x32, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x72, 0x32, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_stats_blok2_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_stats_blok2_proto_rawDescData = file_wb_promotion_v1_stats_blok2_proto_rawDesc
)

func file_wb_promotion_v1_stats_blok2_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_stats_blok2_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_stats_blok2_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_stats_blok2_proto_rawDescData)
	})
	return file_wb_promotion_v1_stats_blok2_proto_rawDescData
}

var file_wb_promotion_v1_stats_blok2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_stats_blok2_proto_goTypes = []interface{}{
	(*StatsBlok2)(nil),       // 0: wb.promotion.v1.StatsBlok2
	(*DailyStats2Inner)(nil), // 1: wb.promotion.v1.DailyStats2Inner
}
var file_wb_promotion_v1_stats_blok2_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.StatsBlok2.daily_stats:type_name -> wb.promotion.v1.DailyStats2Inner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_stats_blok2_proto_init() }
func file_wb_promotion_v1_stats_blok2_proto_init() {
	if File_wb_promotion_v1_stats_blok2_proto != nil {
		return
	}
	file_wb_promotion_v1_daily_stats2_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_stats_blok2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsBlok2); i {
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
			RawDescriptor: file_wb_promotion_v1_stats_blok2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_stats_blok2_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_stats_blok2_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_stats_blok2_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_stats_blok2_proto = out.File
	file_wb_promotion_v1_stats_blok2_proto_rawDesc = nil
	file_wb_promotion_v1_stats_blok2_proto_goTypes = nil
	file_wb_promotion_v1_stats_blok2_proto_depIdxs = nil
}
