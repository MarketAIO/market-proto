//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/marketplace/v1/supply_order.proto

package wbMARKETPLACE

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

// <dl> <dt>Тип товара:</dt> <dd>1 - обычный</dd> <dd>2 - СГТ (Сверхгабаритный товар)</dd> <dd>3 - КГТ (Крупногабаритный товар). Не используется на данный момент.</dd> </dl>
type SupplyOrder_CargoTypeEnum int32

const (
	SupplyOrder_CargoTypeEnum__1 SupplyOrder_CargoTypeEnum = 0
	SupplyOrder_CargoTypeEnum__2 SupplyOrder_CargoTypeEnum = 1
	SupplyOrder_CargoTypeEnum__3 SupplyOrder_CargoTypeEnum = 2
)

// Enum value maps for SupplyOrder_CargoTypeEnum.
var (
	SupplyOrder_CargoTypeEnum_name = map[int32]string{
		0: "CargoTypeEnum__1",
		1: "CargoTypeEnum__2",
		2: "CargoTypeEnum__3",
	}
	SupplyOrder_CargoTypeEnum_value = map[string]int32{
		"CargoTypeEnum__1": 0,
		"CargoTypeEnum__2": 1,
		"CargoTypeEnum__3": 2,
	}
)

func (x SupplyOrder_CargoTypeEnum) Enum() *SupplyOrder_CargoTypeEnum {
	p := new(SupplyOrder_CargoTypeEnum)
	*p = x
	return p
}

func (x SupplyOrder_CargoTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupplyOrder_CargoTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_wb_marketplace_v1_supply_order_proto_enumTypes[0].Descriptor()
}

func (SupplyOrder_CargoTypeEnum) Type() protoreflect.EnumType {
	return &file_wb_marketplace_v1_supply_order_proto_enumTypes[0]
}

func (x SupplyOrder_CargoTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupplyOrder_CargoTypeEnum.Descriptor instead.
func (SupplyOrder_CargoTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_supply_order_proto_rawDescGZIP(), []int{0, 0}
}

type SupplyOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор сборочного задания в Маркетплейсе
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Идентификатор сборочного задания в системе Wildberries
	Rid string `protobuf:"bytes,2,opt,name=rid,proto3" json:"rid,omitempty"`
	// Дата создания сборочного задания (RFC3339)
	CreatedAt string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// Идентификатор склада продавца, на который поступило сборочное задание
	WarehouseId int32 `protobuf:"varint,4,opt,name=warehouseId,proto3" json:"warehouseId,omitempty"`
	// Список офисов, куда следует привезти товар
	Offices []string   `protobuf:"bytes,5,rep,name=offices,proto3" json:"offices,omitempty"`
	User    *OrderUser `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	// Массив баркодов товара
	Skus []string `protobuf:"bytes,7,rep,name=skus,proto3" json:"skus,omitempty"`
	// Цена в валюте продажи с учетом всех скидок, умноженная на 100. Код валюты продажи в поле currencyCode.
	Price int32 `protobuf:"varint,8,opt,name=price,proto3" json:"price,omitempty"`
	// Цена в валюте страны продавца с учетом всех скидок, кроме суммы по WB Кошельку, умноженная на 100. Предоставляется в информационных целях.
	ConvertedPrice int32 `protobuf:"varint,9,opt,name=convertedPrice,proto3" json:"convertedPrice,omitempty"`
	// Код валюты продажи (ISO 4217)
	CurrencyCode int32 `protobuf:"varint,10,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	// Код валюты страны продавца (ISO 4217)
	ConvertedCurrencyCode int32 `protobuf:"varint,11,opt,name=convertedCurrencyCode,proto3" json:"convertedCurrencyCode,omitempty"`
	// Идентификатор транзакции для группировки сборочных заданий. Сборочные задания в одной корзине покупателя будут иметь одинаковый orderUID
	OrderUid string `protobuf:"bytes,12,opt,name=orderUid,proto3" json:"orderUid,omitempty"`
	// Артикул WB
	NmId int32 `protobuf:"varint,13,opt,name=nmId,proto3" json:"nmId,omitempty"`
	// Идентификатор размера товара в системе Wildberries
	ChrtId int32 `protobuf:"varint,14,opt,name=chrtId,proto3" json:"chrtId,omitempty"`
	// Артикул продавца
	Article string `protobuf:"bytes,15,opt,name=article,proto3" json:"article,omitempty"`
	// Код цвета (только для колеруемых товаров)
	ColorCode string                    `protobuf:"bytes,16,opt,name=colorCode,proto3" json:"colorCode,omitempty"`
	CargoType SupplyOrder_CargoTypeEnum `protobuf:"varint,17,opt,name=cargoType,proto3,enum=wb.marketplace.v1.SupplyOrder_CargoTypeEnum" json:"cargoType,omitempty"`
}

func (x *SupplyOrder) Reset() {
	*x = SupplyOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_marketplace_v1_supply_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyOrder) ProtoMessage() {}

func (x *SupplyOrder) ProtoReflect() protoreflect.Message {
	mi := &file_wb_marketplace_v1_supply_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyOrder.ProtoReflect.Descriptor instead.
func (*SupplyOrder) Descriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_supply_order_proto_rawDescGZIP(), []int{0}
}

func (x *SupplyOrder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SupplyOrder) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *SupplyOrder) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SupplyOrder) GetWarehouseId() int32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *SupplyOrder) GetOffices() []string {
	if x != nil {
		return x.Offices
	}
	return nil
}

func (x *SupplyOrder) GetUser() *OrderUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SupplyOrder) GetSkus() []string {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *SupplyOrder) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SupplyOrder) GetConvertedPrice() int32 {
	if x != nil {
		return x.ConvertedPrice
	}
	return 0
}

func (x *SupplyOrder) GetCurrencyCode() int32 {
	if x != nil {
		return x.CurrencyCode
	}
	return 0
}

func (x *SupplyOrder) GetConvertedCurrencyCode() int32 {
	if x != nil {
		return x.ConvertedCurrencyCode
	}
	return 0
}

func (x *SupplyOrder) GetOrderUid() string {
	if x != nil {
		return x.OrderUid
	}
	return ""
}

func (x *SupplyOrder) GetNmId() int32 {
	if x != nil {
		return x.NmId
	}
	return 0
}

func (x *SupplyOrder) GetChrtId() int32 {
	if x != nil {
		return x.ChrtId
	}
	return 0
}

func (x *SupplyOrder) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *SupplyOrder) GetColorCode() string {
	if x != nil {
		return x.ColorCode
	}
	return ""
}

func (x *SupplyOrder) GetCargoType() SupplyOrder_CargoTypeEnum {
	if x != nil {
		return x.CargoType
	}
	return SupplyOrder_CargoTypeEnum__1
}

var File_wb_marketplace_v1_supply_order_proto protoreflect.FileDescriptor

var file_wb_marketplace_v1_supply_order_proto_rawDesc = []byte{
	0x0a, 0x24, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x77, 0x62, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x05,
	0x0a, 0x0b, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6b, 0x75, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x15, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x72, 0x74, 0x49,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x72, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x67, 0x6f,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x77, 0x62, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x67, 0x6f,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x31, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61,
	0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x32, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x5f, 0x5f, 0x33, 0x10, 0x02, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54,
	0x50, 0x4c, 0x41, 0x43, 0x45, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_marketplace_v1_supply_order_proto_rawDescOnce sync.Once
	file_wb_marketplace_v1_supply_order_proto_rawDescData = file_wb_marketplace_v1_supply_order_proto_rawDesc
)

func file_wb_marketplace_v1_supply_order_proto_rawDescGZIP() []byte {
	file_wb_marketplace_v1_supply_order_proto_rawDescOnce.Do(func() {
		file_wb_marketplace_v1_supply_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_marketplace_v1_supply_order_proto_rawDescData)
	})
	return file_wb_marketplace_v1_supply_order_proto_rawDescData
}

var file_wb_marketplace_v1_supply_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wb_marketplace_v1_supply_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_marketplace_v1_supply_order_proto_goTypes = []interface{}{
	(SupplyOrder_CargoTypeEnum)(0), // 0: wb.marketplace.v1.SupplyOrder.CargoTypeEnum
	(*SupplyOrder)(nil),            // 1: wb.marketplace.v1.SupplyOrder
	(*OrderUser)(nil),              // 2: wb.marketplace.v1.OrderUser
}
var file_wb_marketplace_v1_supply_order_proto_depIdxs = []int32{
	2, // 0: wb.marketplace.v1.SupplyOrder.user:type_name -> wb.marketplace.v1.OrderUser
	0, // 1: wb.marketplace.v1.SupplyOrder.cargoType:type_name -> wb.marketplace.v1.SupplyOrder.CargoTypeEnum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wb_marketplace_v1_supply_order_proto_init() }
func file_wb_marketplace_v1_supply_order_proto_init() {
	if File_wb_marketplace_v1_supply_order_proto != nil {
		return
	}
	file_wb_marketplace_v1_order_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_marketplace_v1_supply_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyOrder); i {
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
			RawDescriptor: file_wb_marketplace_v1_supply_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_marketplace_v1_supply_order_proto_goTypes,
		DependencyIndexes: file_wb_marketplace_v1_supply_order_proto_depIdxs,
		EnumInfos:         file_wb_marketplace_v1_supply_order_proto_enumTypes,
		MessageInfos:      file_wb_marketplace_v1_supply_order_proto_msgTypes,
	}.Build()
	File_wb_marketplace_v1_supply_order_proto = out.File
	file_wb_marketplace_v1_supply_order_proto_rawDesc = nil
	file_wb_marketplace_v1_supply_order_proto_goTypes = nil
	file_wb_marketplace_v1_supply_order_proto_depIdxs = nil
}
