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
// source: wb/marketplace/v1/order.proto

package wbMarketplace

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

// <dl> <dt>Тип доставки:</dt> <dd>fbs - доставка на склад Wildberries</dd> <dd>dbs - доставка силами продавца</dd> <dd>wbgo - доставка курьером WB</dd> <dd>edbs - экспресс-доставка силами продавца</dd> </dl>
type Order_DeliveryTypeEnum int32

const (
	Order_DeliveryTypeEnum_DBS  Order_DeliveryTypeEnum = 0
	Order_DeliveryTypeEnum_EDBS Order_DeliveryTypeEnum = 1
	Order_DeliveryTypeEnum_FBS  Order_DeliveryTypeEnum = 2
	Order_DeliveryTypeEnum_WBGO Order_DeliveryTypeEnum = 3
)

// Enum value maps for Order_DeliveryTypeEnum.
var (
	Order_DeliveryTypeEnum_name = map[int32]string{
		0: "DeliveryTypeEnum_DBS",
		1: "DeliveryTypeEnum_EDBS",
		2: "DeliveryTypeEnum_FBS",
		3: "DeliveryTypeEnum_WBGO",
	}
	Order_DeliveryTypeEnum_value = map[string]int32{
		"DeliveryTypeEnum_DBS":  0,
		"DeliveryTypeEnum_EDBS": 1,
		"DeliveryTypeEnum_FBS":  2,
		"DeliveryTypeEnum_WBGO": 3,
	}
)

func (x Order_DeliveryTypeEnum) Enum() *Order_DeliveryTypeEnum {
	p := new(Order_DeliveryTypeEnum)
	*p = x
	return p
}

func (x Order_DeliveryTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_DeliveryTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_wb_marketplace_v1_order_proto_enumTypes[0].Descriptor()
}

func (Order_DeliveryTypeEnum) Type() protoreflect.EnumType {
	return &file_wb_marketplace_v1_order_proto_enumTypes[0]
}

func (x Order_DeliveryTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_DeliveryTypeEnum.Descriptor instead.
func (Order_DeliveryTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_order_proto_rawDescGZIP(), []int{0, 0}
}

// <dl> <dt>Тип товара:</dt> <dd>1 - обычный</dd> <dd>2 - СГТ (Сверхгабаритный товар)</dd> <dd>3 - КГТ (Крупногабаритный товар). Не используется на данный момент.</dd> </dl>
type Order_CargoTypeEnum int32

const (
	Order_CargoTypeEnum__1 Order_CargoTypeEnum = 0
	Order_CargoTypeEnum__2 Order_CargoTypeEnum = 1
	Order_CargoTypeEnum__3 Order_CargoTypeEnum = 2
)

// Enum value maps for Order_CargoTypeEnum.
var (
	Order_CargoTypeEnum_name = map[int32]string{
		0: "CargoTypeEnum__1",
		1: "CargoTypeEnum__2",
		2: "CargoTypeEnum__3",
	}
	Order_CargoTypeEnum_value = map[string]int32{
		"CargoTypeEnum__1": 0,
		"CargoTypeEnum__2": 1,
		"CargoTypeEnum__3": 2,
	}
)

func (x Order_CargoTypeEnum) Enum() *Order_CargoTypeEnum {
	p := new(Order_CargoTypeEnum)
	*p = x
	return p
}

func (x Order_CargoTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_CargoTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_wb_marketplace_v1_order_proto_enumTypes[1].Descriptor()
}

func (Order_CargoTypeEnum) Type() protoreflect.EnumType {
	return &file_wb_marketplace_v1_order_proto_enumTypes[1]
}

func (x Order_CargoTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_CargoTypeEnum.Descriptor instead.
func (Order_CargoTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_order_proto_rawDescGZIP(), []int{0, 1}
}

type Order struct {
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
	// Идентификатор поставки. Возвращается, если заказ закреплён за поставкой
	SupplyId string `protobuf:"bytes,5,opt,name=supplyId,proto3" json:"supplyId,omitempty"`
	// Список офисов, куда следует привезти товар
	Offices []string      `protobuf:"bytes,6,rep,name=offices,proto3" json:"offices,omitempty"`
	Address *OrderAddress `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	User    *OrderUser    `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	// Массив баркодов товара
	Skus []string `protobuf:"bytes,9,rep,name=skus,proto3" json:"skus,omitempty"`
	// Цена в валюте продажи с учетом всех скидок, умноженная на 100. Код валюты продажи в поле currencyCode.
	Price int32 `protobuf:"varint,10,opt,name=price,proto3" json:"price,omitempty"`
	// Цена в валюте страны продавца с учетом всех скидок, кроме суммы по WB Кошельку, умноженная на 100. Предоставляется в информационных целях.
	ConvertedPrice int32 `protobuf:"varint,11,opt,name=convertedPrice,proto3" json:"convertedPrice,omitempty"`
	// Код валюты продажи (ISO 4217)
	CurrencyCode int32 `protobuf:"varint,12,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	// Код валюты страны продавца (ISO 4217)
	ConvertedCurrencyCode int32 `protobuf:"varint,13,opt,name=convertedCurrencyCode,proto3" json:"convertedCurrencyCode,omitempty"`
	// Идентификатор транзакции для группировки сборочных заданий. Сборочные задания в одной корзине покупателя будут иметь одинаковый orderUID
	OrderUid     string                 `protobuf:"bytes,14,opt,name=orderUid,proto3" json:"orderUid,omitempty"`
	DeliveryType Order_DeliveryTypeEnum `protobuf:"varint,15,opt,name=deliveryType,proto3,enum=wb.marketplace.v1.Order_DeliveryTypeEnum" json:"deliveryType,omitempty"`
	// Артикул WB
	NmId int32 `protobuf:"varint,16,opt,name=nmId,proto3" json:"nmId,omitempty"`
	// Идентификатор размера товара в системе Wildberries
	ChrtId int32 `protobuf:"varint,17,opt,name=chrtId,proto3" json:"chrtId,omitempty"`
	// Артикул продавца
	Article string `protobuf:"bytes,18,opt,name=article,proto3" json:"article,omitempty"`
	// Код цвета (только для колеруемых товаров)
	ColorCode string              `protobuf:"bytes,19,opt,name=colorCode,proto3" json:"colorCode,omitempty"`
	CargoType Order_CargoTypeEnum `protobuf:"varint,20,opt,name=cargoType,proto3,enum=wb.marketplace.v1.Order_CargoTypeEnum" json:"cargoType,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_marketplace_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_wb_marketplace_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Order) GetWarehouseId() int32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *Order) GetSupplyId() string {
	if x != nil {
		return x.SupplyId
	}
	return ""
}

func (x *Order) GetOffices() []string {
	if x != nil {
		return x.Offices
	}
	return nil
}

func (x *Order) GetAddress() *OrderAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Order) GetUser() *OrderUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Order) GetSkus() []string {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *Order) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetConvertedPrice() int32 {
	if x != nil {
		return x.ConvertedPrice
	}
	return 0
}

func (x *Order) GetCurrencyCode() int32 {
	if x != nil {
		return x.CurrencyCode
	}
	return 0
}

func (x *Order) GetConvertedCurrencyCode() int32 {
	if x != nil {
		return x.ConvertedCurrencyCode
	}
	return 0
}

func (x *Order) GetOrderUid() string {
	if x != nil {
		return x.OrderUid
	}
	return ""
}

func (x *Order) GetDeliveryType() Order_DeliveryTypeEnum {
	if x != nil {
		return x.DeliveryType
	}
	return Order_DeliveryTypeEnum_DBS
}

func (x *Order) GetNmId() int32 {
	if x != nil {
		return x.NmId
	}
	return 0
}

func (x *Order) GetChrtId() int32 {
	if x != nil {
		return x.ChrtId
	}
	return 0
}

func (x *Order) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *Order) GetColorCode() string {
	if x != nil {
		return x.ColorCode
	}
	return ""
}

func (x *Order) GetCargoType() Order_CargoTypeEnum {
	if x != nil {
		return x.CargoType
	}
	return Order_CargoTypeEnum__1
}

var File_wb_marketplace_v1_order_proto protoreflect.FileDescriptor

var file_wb_marketplace_v1_order_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x25, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x77, 0x62, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x07,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x39, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x62, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6b, 0x75, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6b, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x77, 0x62, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x72, 0x74, 0x49,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x72, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x67, 0x6f,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x77, 0x62, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x22, 0x7c, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x44, 0x42, 0x53, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f,
	0x45, 0x44, 0x42, 0x53, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x46, 0x42, 0x53, 0x10, 0x02,
	0x12, 0x19, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x57, 0x42, 0x47, 0x4f, 0x10, 0x03, 0x22, 0x51, 0x0a, 0x0d, 0x43,
	0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x31,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x32, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x33, 0x10, 0x02, 0x42, 0x4a,
	0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wb_marketplace_v1_order_proto_rawDescOnce sync.Once
	file_wb_marketplace_v1_order_proto_rawDescData = file_wb_marketplace_v1_order_proto_rawDesc
)

func file_wb_marketplace_v1_order_proto_rawDescGZIP() []byte {
	file_wb_marketplace_v1_order_proto_rawDescOnce.Do(func() {
		file_wb_marketplace_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_marketplace_v1_order_proto_rawDescData)
	})
	return file_wb_marketplace_v1_order_proto_rawDescData
}

var file_wb_marketplace_v1_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_wb_marketplace_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_marketplace_v1_order_proto_goTypes = []interface{}{
	(Order_DeliveryTypeEnum)(0), // 0: wb.marketplace.v1.Order.DeliveryTypeEnum
	(Order_CargoTypeEnum)(0),    // 1: wb.marketplace.v1.Order.CargoTypeEnum
	(*Order)(nil),               // 2: wb.marketplace.v1.Order
	(*OrderAddress)(nil),        // 3: wb.marketplace.v1.OrderAddress
	(*OrderUser)(nil),           // 4: wb.marketplace.v1.OrderUser
}
var file_wb_marketplace_v1_order_proto_depIdxs = []int32{
	3, // 0: wb.marketplace.v1.Order.address:type_name -> wb.marketplace.v1.OrderAddress
	4, // 1: wb.marketplace.v1.Order.user:type_name -> wb.marketplace.v1.OrderUser
	0, // 2: wb.marketplace.v1.Order.deliveryType:type_name -> wb.marketplace.v1.Order.DeliveryTypeEnum
	1, // 3: wb.marketplace.v1.Order.cargoType:type_name -> wb.marketplace.v1.Order.CargoTypeEnum
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wb_marketplace_v1_order_proto_init() }
func file_wb_marketplace_v1_order_proto_init() {
	if File_wb_marketplace_v1_order_proto != nil {
		return
	}
	file_wb_marketplace_v1_order_address_proto_init()
	file_wb_marketplace_v1_order_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_marketplace_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_wb_marketplace_v1_order_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_marketplace_v1_order_proto_goTypes,
		DependencyIndexes: file_wb_marketplace_v1_order_proto_depIdxs,
		EnumInfos:         file_wb_marketplace_v1_order_proto_enumTypes,
		MessageInfos:      file_wb_marketplace_v1_order_proto_msgTypes,
	}.Build()
	File_wb_marketplace_v1_order_proto = out.File
	file_wb_marketplace_v1_order_proto_rawDesc = nil
	file_wb_marketplace_v1_order_proto_goTypes = nil
	file_wb_marketplace_v1_order_proto_depIdxs = nil
}
