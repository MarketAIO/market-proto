//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/tariffs/v1/models_warehouse_box_rates.proto

package wbTariffs

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

type ModelsWarehouseBoxRates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Коэффициент, %. На него умножается стоимость доставки и хранения. Во всех тарифах этот коэффициент уже учтён
	BoxDeliveryAndStorageExpr string `protobuf:"bytes,1,opt,name=boxDeliveryAndStorageExpr,proto3" json:"boxDeliveryAndStorageExpr,omitempty"`
	// Доставка 1 литра, ₽
	BoxDeliveryBase string `protobuf:"bytes,2,opt,name=boxDeliveryBase,proto3" json:"boxDeliveryBase,omitempty"`
	// Доставка каждого дополнительного литра, ₽
	BoxDeliveryLiter string `protobuf:"bytes,3,opt,name=boxDeliveryLiter,proto3" json:"boxDeliveryLiter,omitempty"`
	// Хранение 1 литра, ₽
	BoxStorageBase string `protobuf:"bytes,4,opt,name=boxStorageBase,proto3" json:"boxStorageBase,omitempty"`
	// Хранение каждого дополнительного литра, ₽
	BoxStorageLiter string `protobuf:"bytes,5,opt,name=boxStorageLiter,proto3" json:"boxStorageLiter,omitempty"`
	// Название склада
	WarehouseName string `protobuf:"bytes,6,opt,name=warehouseName,proto3" json:"warehouseName,omitempty"`
}

func (x *ModelsWarehouseBoxRates) Reset() {
	*x = ModelsWarehouseBoxRates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1_models_warehouse_box_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelsWarehouseBoxRates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelsWarehouseBoxRates) ProtoMessage() {}

func (x *ModelsWarehouseBoxRates) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1_models_warehouse_box_rates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelsWarehouseBoxRates.ProtoReflect.Descriptor instead.
func (*ModelsWarehouseBoxRates) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescGZIP(), []int{0}
}

func (x *ModelsWarehouseBoxRates) GetBoxDeliveryAndStorageExpr() string {
	if x != nil {
		return x.BoxDeliveryAndStorageExpr
	}
	return ""
}

func (x *ModelsWarehouseBoxRates) GetBoxDeliveryBase() string {
	if x != nil {
		return x.BoxDeliveryBase
	}
	return ""
}

func (x *ModelsWarehouseBoxRates) GetBoxDeliveryLiter() string {
	if x != nil {
		return x.BoxDeliveryLiter
	}
	return ""
}

func (x *ModelsWarehouseBoxRates) GetBoxStorageBase() string {
	if x != nil {
		return x.BoxStorageBase
	}
	return ""
}

func (x *ModelsWarehouseBoxRates) GetBoxStorageLiter() string {
	if x != nil {
		return x.BoxStorageLiter
	}
	return ""
}

func (x *ModelsWarehouseBoxRates) GetWarehouseName() string {
	if x != nil {
		return x.WarehouseName
	}
	return ""
}

var File_wb_tariffs_v1_models_warehouse_box_rates_proto protoreflect.FileDescriptor

var file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x22,
	0xa5, 0x02, 0x0a, 0x17, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x19, 0x62,
	0x6f, 0x78, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x45, 0x78, 0x70, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19,
	0x62, 0x6f, 0x78, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x64, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x78, 0x70, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x6f, 0x78,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x62, 0x6f, 0x78, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x6f, 0x78, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62,
	0x6f, 0x78, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x0e, 0x62, 0x6f, 0x78, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6f, 0x78, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x6f, 0x78, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x62, 0x6f, 0x78, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x74, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x77, 0x62, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescOnce sync.Once
	file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescData = file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDesc
)

func file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescGZIP() []byte {
	file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescOnce.Do(func() {
		file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescData)
	})
	return file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDescData
}

var file_wb_tariffs_v1_models_warehouse_box_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_tariffs_v1_models_warehouse_box_rates_proto_goTypes = []interface{}{
	(*ModelsWarehouseBoxRates)(nil), // 0: wb.tariffs.v1.ModelsWarehouseBoxRates
}
var file_wb_tariffs_v1_models_warehouse_box_rates_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_tariffs_v1_models_warehouse_box_rates_proto_init() }
func file_wb_tariffs_v1_models_warehouse_box_rates_proto_init() {
	if File_wb_tariffs_v1_models_warehouse_box_rates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_tariffs_v1_models_warehouse_box_rates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelsWarehouseBoxRates); i {
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
			RawDescriptor: file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_tariffs_v1_models_warehouse_box_rates_proto_goTypes,
		DependencyIndexes: file_wb_tariffs_v1_models_warehouse_box_rates_proto_depIdxs,
		MessageInfos:      file_wb_tariffs_v1_models_warehouse_box_rates_proto_msgTypes,
	}.Build()
	File_wb_tariffs_v1_models_warehouse_box_rates_proto = out.File
	file_wb_tariffs_v1_models_warehouse_box_rates_proto_rawDesc = nil
	file_wb_tariffs_v1_models_warehouse_box_rates_proto_goTypes = nil
	file_wb_tariffs_v1_models_warehouse_box_rates_proto_depIdxs = nil
}
