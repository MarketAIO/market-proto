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
// source: wb/marketplace/v1/supply.proto

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

// <dl> <dt>Тип поставки:</dt> <dd>0 - признак отсутствует</dd> <dd>1 - обычная</dd> <dd>2 - СГТ (Содержит сверхгабаритные товары)</dd> <dd>3 - КГТ (Содержит крупногабаритные товары). Не используется на данный момент.</dd> </dl>
type Supply_CargoTypeEnum int32

const (
	Supply_CargoTypeEnum__0 Supply_CargoTypeEnum = 0
	Supply_CargoTypeEnum__1 Supply_CargoTypeEnum = 1
	Supply_CargoTypeEnum__2 Supply_CargoTypeEnum = 2
	Supply_CargoTypeEnum__3 Supply_CargoTypeEnum = 3
)

// Enum value maps for Supply_CargoTypeEnum.
var (
	Supply_CargoTypeEnum_name = map[int32]string{
		0: "CargoTypeEnum__0",
		1: "CargoTypeEnum__1",
		2: "CargoTypeEnum__2",
		3: "CargoTypeEnum__3",
	}
	Supply_CargoTypeEnum_value = map[string]int32{
		"CargoTypeEnum__0": 0,
		"CargoTypeEnum__1": 1,
		"CargoTypeEnum__2": 2,
		"CargoTypeEnum__3": 3,
	}
)

func (x Supply_CargoTypeEnum) Enum() *Supply_CargoTypeEnum {
	p := new(Supply_CargoTypeEnum)
	*p = x
	return p
}

func (x Supply_CargoTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Supply_CargoTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_wb_marketplace_v1_supply_proto_enumTypes[0].Descriptor()
}

func (Supply_CargoTypeEnum) Type() protoreflect.EnumType {
	return &file_wb_marketplace_v1_supply_proto_enumTypes[0]
}

func (x Supply_CargoTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Supply_CargoTypeEnum.Descriptor instead.
func (Supply_CargoTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_supply_proto_rawDescGZIP(), []int{0, 0}
}

type Supply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор поставки
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Флаг закрытия поставки
	Done bool `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	// Дата создания поставки (RFC3339)
	CreatedAt string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// Дата закрытия поставки (RFC3339)
	ClosedAt string `protobuf:"bytes,4,opt,name=closedAt,proto3" json:"closedAt,omitempty"`
	// Дата скана поставки (RFC3339)
	ScanDt string `protobuf:"bytes,5,opt,name=scanDt,proto3" json:"scanDt,omitempty"`
	// Наименование поставки
	Name      string               `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	CargoType Supply_CargoTypeEnum `protobuf:"varint,7,opt,name=cargoType,proto3,enum=wb.marketplace.v1.Supply_CargoTypeEnum" json:"cargoType,omitempty"`
}

func (x *Supply) Reset() {
	*x = Supply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_marketplace_v1_supply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supply) ProtoMessage() {}

func (x *Supply) ProtoReflect() protoreflect.Message {
	mi := &file_wb_marketplace_v1_supply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supply.ProtoReflect.Descriptor instead.
func (*Supply) Descriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_supply_proto_rawDescGZIP(), []int{0}
}

func (x *Supply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Supply) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Supply) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Supply) GetClosedAt() string {
	if x != nil {
		return x.ClosedAt
	}
	return ""
}

func (x *Supply) GetScanDt() string {
	if x != nil {
		return x.ScanDt
	}
	return ""
}

func (x *Supply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Supply) GetCargoType() Supply_CargoTypeEnum {
	if x != nil {
		return x.CargoType
	}
	return Supply_CargoTypeEnum__0
}

var File_wb_marketplace_v1_supply_proto protoreflect.FileDescriptor

var file_wb_marketplace_v1_supply_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0xc2, 0x02, 0x0a, 0x06, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f,
	0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x61, 0x6e, 0x44, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x61, 0x6e, 0x44, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x67,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x77, 0x62,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x67, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x5f, 0x5f, 0x30, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x32,
	0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x33, 0x10, 0x03, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_marketplace_v1_supply_proto_rawDescOnce sync.Once
	file_wb_marketplace_v1_supply_proto_rawDescData = file_wb_marketplace_v1_supply_proto_rawDesc
)

func file_wb_marketplace_v1_supply_proto_rawDescGZIP() []byte {
	file_wb_marketplace_v1_supply_proto_rawDescOnce.Do(func() {
		file_wb_marketplace_v1_supply_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_marketplace_v1_supply_proto_rawDescData)
	})
	return file_wb_marketplace_v1_supply_proto_rawDescData
}

var file_wb_marketplace_v1_supply_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wb_marketplace_v1_supply_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_marketplace_v1_supply_proto_goTypes = []interface{}{
	(Supply_CargoTypeEnum)(0), // 0: wb.marketplace.v1.Supply.CargoTypeEnum
	(*Supply)(nil),            // 1: wb.marketplace.v1.Supply
}
var file_wb_marketplace_v1_supply_proto_depIdxs = []int32{
	0, // 0: wb.marketplace.v1.Supply.cargoType:type_name -> wb.marketplace.v1.Supply.CargoTypeEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_marketplace_v1_supply_proto_init() }
func file_wb_marketplace_v1_supply_proto_init() {
	if File_wb_marketplace_v1_supply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_marketplace_v1_supply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supply); i {
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
			RawDescriptor: file_wb_marketplace_v1_supply_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_marketplace_v1_supply_proto_goTypes,
		DependencyIndexes: file_wb_marketplace_v1_supply_proto_depIdxs,
		EnumInfos:         file_wb_marketplace_v1_supply_proto_enumTypes,
		MessageInfos:      file_wb_marketplace_v1_supply_proto_msgTypes,
	}.Build()
	File_wb_marketplace_v1_supply_proto = out.File
	file_wb_marketplace_v1_supply_proto_rawDesc = nil
	file_wb_marketplace_v1_supply_proto_goTypes = nil
	file_wb_marketplace_v1_supply_proto_depIdxs = nil
}
