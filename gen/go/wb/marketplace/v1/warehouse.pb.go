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
// source: wb/marketplace/v1/warehouse.proto

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

// <dl> <dt>Тип товара, который принимает склад:</dt> <dd>1 - обычный</dd> <dd>2 - СГТ (Сверхгабаритный товар)</dd> <dd>3 - КГТ (Крупногабаритный товар). Не используется на данный момент.</dd> </dl>
type Warehouse_CargoTypeEnum int32

const (
	Warehouse_CargoTypeEnum__1 Warehouse_CargoTypeEnum = 0
	Warehouse_CargoTypeEnum__2 Warehouse_CargoTypeEnum = 1
	Warehouse_CargoTypeEnum__3 Warehouse_CargoTypeEnum = 2
)

// Enum value maps for Warehouse_CargoTypeEnum.
var (
	Warehouse_CargoTypeEnum_name = map[int32]string{
		0: "CargoTypeEnum__1",
		1: "CargoTypeEnum__2",
		2: "CargoTypeEnum__3",
	}
	Warehouse_CargoTypeEnum_value = map[string]int32{
		"CargoTypeEnum__1": 0,
		"CargoTypeEnum__2": 1,
		"CargoTypeEnum__3": 2,
	}
)

func (x Warehouse_CargoTypeEnum) Enum() *Warehouse_CargoTypeEnum {
	p := new(Warehouse_CargoTypeEnum)
	*p = x
	return p
}

func (x Warehouse_CargoTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Warehouse_CargoTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_wb_marketplace_v1_warehouse_proto_enumTypes[0].Descriptor()
}

func (Warehouse_CargoTypeEnum) Type() protoreflect.EnumType {
	return &file_wb_marketplace_v1_warehouse_proto_enumTypes[0]
}

func (x Warehouse_CargoTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Warehouse_CargoTypeEnum.Descriptor instead.
func (Warehouse_CargoTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_warehouse_proto_rawDescGZIP(), []int{0, 0}
}

// <dl> <dt>Тип доставки, который принимает склад:</dt> <dd>1 - доставка на склад Wildberries</dd> <dd>2 - доставка силами продавца</dd> <dd>3 - доставка курьером WB</dd> </dl>
type Warehouse_DeliveryTypeEnum int32

const (
	Warehouse_DeliveryTypeEnum__1 Warehouse_DeliveryTypeEnum = 0
	Warehouse_DeliveryTypeEnum__2 Warehouse_DeliveryTypeEnum = 1
	Warehouse_DeliveryTypeEnum__3 Warehouse_DeliveryTypeEnum = 2
)

// Enum value maps for Warehouse_DeliveryTypeEnum.
var (
	Warehouse_DeliveryTypeEnum_name = map[int32]string{
		0: "DeliveryTypeEnum__1",
		1: "DeliveryTypeEnum__2",
		2: "DeliveryTypeEnum__3",
	}
	Warehouse_DeliveryTypeEnum_value = map[string]int32{
		"DeliveryTypeEnum__1": 0,
		"DeliveryTypeEnum__2": 1,
		"DeliveryTypeEnum__3": 2,
	}
)

func (x Warehouse_DeliveryTypeEnum) Enum() *Warehouse_DeliveryTypeEnum {
	p := new(Warehouse_DeliveryTypeEnum)
	*p = x
	return p
}

func (x Warehouse_DeliveryTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Warehouse_DeliveryTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_wb_marketplace_v1_warehouse_proto_enumTypes[1].Descriptor()
}

func (Warehouse_DeliveryTypeEnum) Type() protoreflect.EnumType {
	return &file_wb_marketplace_v1_warehouse_proto_enumTypes[1]
}

func (x Warehouse_DeliveryTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Warehouse_DeliveryTypeEnum.Descriptor instead.
func (Warehouse_DeliveryTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_warehouse_proto_rawDescGZIP(), []int{0, 1}
}

type Warehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Название склада продавца
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID склада WB
	OfficeId int64 `protobuf:"varint,2,opt,name=officeId,proto3" json:"officeId,omitempty"`
	// ID склада продавца
	Id           int64                      `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	CargoType    Warehouse_CargoTypeEnum    `protobuf:"varint,4,opt,name=cargoType,proto3,enum=wb.marketplace.v1.Warehouse_CargoTypeEnum" json:"cargoType,omitempty"`
	DeliveryType Warehouse_DeliveryTypeEnum `protobuf:"varint,5,opt,name=deliveryType,proto3,enum=wb.marketplace.v1.Warehouse_DeliveryTypeEnum" json:"deliveryType,omitempty"`
}

func (x *Warehouse) Reset() {
	*x = Warehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_marketplace_v1_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouse) ProtoMessage() {}

func (x *Warehouse) ProtoReflect() protoreflect.Message {
	mi := &file_wb_marketplace_v1_warehouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warehouse.ProtoReflect.Descriptor instead.
func (*Warehouse) Descriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *Warehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Warehouse) GetOfficeId() int64 {
	if x != nil {
		return x.OfficeId
	}
	return 0
}

func (x *Warehouse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Warehouse) GetCargoType() Warehouse_CargoTypeEnum {
	if x != nil {
		return x.CargoType
	}
	return Warehouse_CargoTypeEnum__1
}

func (x *Warehouse) GetDeliveryType() Warehouse_DeliveryTypeEnum {
	if x != nil {
		return x.DeliveryType
	}
	return Warehouse_DeliveryTypeEnum__1
}

var File_wb_marketplace_v1_warehouse_proto protoreflect.FileDescriptor

var file_wb_marketplace_v1_warehouse_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x9a, 0x03, 0x0a, 0x09, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x48, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x51,
	0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x51, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x31, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x32, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f,
	0x5f, 0x33, 0x10, 0x02, 0x22, 0x5d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x31, 0x10,
	0x00, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x32, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f,
	0x33, 0x10, 0x02, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_marketplace_v1_warehouse_proto_rawDescOnce sync.Once
	file_wb_marketplace_v1_warehouse_proto_rawDescData = file_wb_marketplace_v1_warehouse_proto_rawDesc
)

func file_wb_marketplace_v1_warehouse_proto_rawDescGZIP() []byte {
	file_wb_marketplace_v1_warehouse_proto_rawDescOnce.Do(func() {
		file_wb_marketplace_v1_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_marketplace_v1_warehouse_proto_rawDescData)
	})
	return file_wb_marketplace_v1_warehouse_proto_rawDescData
}

var file_wb_marketplace_v1_warehouse_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_wb_marketplace_v1_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_marketplace_v1_warehouse_proto_goTypes = []interface{}{
	(Warehouse_CargoTypeEnum)(0),    // 0: wb.marketplace.v1.Warehouse.CargoTypeEnum
	(Warehouse_DeliveryTypeEnum)(0), // 1: wb.marketplace.v1.Warehouse.DeliveryTypeEnum
	(*Warehouse)(nil),               // 2: wb.marketplace.v1.Warehouse
}
var file_wb_marketplace_v1_warehouse_proto_depIdxs = []int32{
	0, // 0: wb.marketplace.v1.Warehouse.cargoType:type_name -> wb.marketplace.v1.Warehouse.CargoTypeEnum
	1, // 1: wb.marketplace.v1.Warehouse.deliveryType:type_name -> wb.marketplace.v1.Warehouse.DeliveryTypeEnum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wb_marketplace_v1_warehouse_proto_init() }
func file_wb_marketplace_v1_warehouse_proto_init() {
	if File_wb_marketplace_v1_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_marketplace_v1_warehouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warehouse); i {
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
			RawDescriptor: file_wb_marketplace_v1_warehouse_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_marketplace_v1_warehouse_proto_goTypes,
		DependencyIndexes: file_wb_marketplace_v1_warehouse_proto_depIdxs,
		EnumInfos:         file_wb_marketplace_v1_warehouse_proto_enumTypes,
		MessageInfos:      file_wb_marketplace_v1_warehouse_proto_msgTypes,
	}.Build()
	File_wb_marketplace_v1_warehouse_proto = out.File
	file_wb_marketplace_v1_warehouse_proto_rawDesc = nil
	file_wb_marketplace_v1_warehouse_proto_goTypes = nil
	file_wb_marketplace_v1_warehouse_proto_depIdxs = nil
}