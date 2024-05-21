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
// source: wb/tariffs/v1/models_warehouses_box_rates.proto

package wbTARIFFS

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

type ModelsWarehousesBoxRates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Дата начала тарифа
	DtFromMin string `protobuf:"bytes,1,opt,name=dtFromMin,proto3" json:"dtFromMin,omitempty"`
	// Дата начала следующего тарифа
	DtNextBox string `protobuf:"bytes,2,opt,name=dtNextBox,proto3" json:"dtNextBox,omitempty"`
	// Дата окончания тарифа
	DtTillMax string `protobuf:"bytes,3,opt,name=dtTillMax,proto3" json:"dtTillMax,omitempty"`
	// Тарифы для коробов, сгруппированные по складам
	WarehouseList []*ModelsWarehouseBoxRates `protobuf:"bytes,4,rep,name=warehouseList,proto3" json:"warehouseList,omitempty"`
}

func (x *ModelsWarehousesBoxRates) Reset() {
	*x = ModelsWarehousesBoxRates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1_models_warehouses_box_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelsWarehousesBoxRates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelsWarehousesBoxRates) ProtoMessage() {}

func (x *ModelsWarehousesBoxRates) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1_models_warehouses_box_rates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelsWarehousesBoxRates.ProtoReflect.Descriptor instead.
func (*ModelsWarehousesBoxRates) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescGZIP(), []int{0}
}

func (x *ModelsWarehousesBoxRates) GetDtFromMin() string {
	if x != nil {
		return x.DtFromMin
	}
	return ""
}

func (x *ModelsWarehousesBoxRates) GetDtNextBox() string {
	if x != nil {
		return x.DtNextBox
	}
	return ""
}

func (x *ModelsWarehousesBoxRates) GetDtTillMax() string {
	if x != nil {
		return x.DtTillMax
	}
	return ""
}

func (x *ModelsWarehousesBoxRates) GetWarehouseList() []*ModelsWarehouseBoxRates {
	if x != nil {
		return x.WarehouseList
	}
	return nil
}

var File_wb_tariffs_v1_models_warehouses_box_rates_proto protoreflect.FileDescriptor

var file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x73, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x2e, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc2, 0x01, 0x0a, 0x18, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x73, 0x42, 0x6f, 0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x74, 0x4e, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x74, 0x4e, 0x65, 0x78, 0x74, 0x42, 0x6f, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x74, 0x54,
	0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x74,
	0x54, 0x69, 0x6c, 0x6c, 0x4d, 0x61, 0x78, 0x12, 0x4c, 0x0a, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x6f,
	0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x77, 0x62, 0x54, 0x41, 0x52, 0x49, 0x46, 0x46, 0x53, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescOnce sync.Once
	file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescData = file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDesc
)

func file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescGZIP() []byte {
	file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescOnce.Do(func() {
		file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescData)
	})
	return file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDescData
}

var file_wb_tariffs_v1_models_warehouses_box_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_tariffs_v1_models_warehouses_box_rates_proto_goTypes = []interface{}{
	(*ModelsWarehousesBoxRates)(nil), // 0: wb.tariffs.v1.ModelsWarehousesBoxRates
	(*ModelsWarehouseBoxRates)(nil),  // 1: wb.tariffs.v1.ModelsWarehouseBoxRates
}
var file_wb_tariffs_v1_models_warehouses_box_rates_proto_depIdxs = []int32{
	1, // 0: wb.tariffs.v1.ModelsWarehousesBoxRates.warehouseList:type_name -> wb.tariffs.v1.ModelsWarehouseBoxRates
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_tariffs_v1_models_warehouses_box_rates_proto_init() }
func file_wb_tariffs_v1_models_warehouses_box_rates_proto_init() {
	if File_wb_tariffs_v1_models_warehouses_box_rates_proto != nil {
		return
	}
	file_wb_tariffs_v1_models_warehouse_box_rates_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_tariffs_v1_models_warehouses_box_rates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelsWarehousesBoxRates); i {
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
			RawDescriptor: file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_tariffs_v1_models_warehouses_box_rates_proto_goTypes,
		DependencyIndexes: file_wb_tariffs_v1_models_warehouses_box_rates_proto_depIdxs,
		MessageInfos:      file_wb_tariffs_v1_models_warehouses_box_rates_proto_msgTypes,
	}.Build()
	File_wb_tariffs_v1_models_warehouses_box_rates_proto = out.File
	file_wb_tariffs_v1_models_warehouses_box_rates_proto_rawDesc = nil
	file_wb_tariffs_v1_models_warehouses_box_rates_proto_goTypes = nil
	file_wb_tariffs_v1_models_warehouses_box_rates_proto_depIdxs = nil
}
