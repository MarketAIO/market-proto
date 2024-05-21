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
// source: wb/marketplace/v1/api_v3_supplies_supply_id_orders_get200_response.proto

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

type ApiV3SuppliesSupplyIdOrdersGet200Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*SupplyOrder `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ApiV3SuppliesSupplyIdOrdersGet200Response) Reset() {
	*x = ApiV3SuppliesSupplyIdOrdersGet200Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiV3SuppliesSupplyIdOrdersGet200Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiV3SuppliesSupplyIdOrdersGet200Response) ProtoMessage() {}

func (x *ApiV3SuppliesSupplyIdOrdersGet200Response) ProtoReflect() protoreflect.Message {
	mi := &file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiV3SuppliesSupplyIdOrdersGet200Response.ProtoReflect.Descriptor instead.
func (*ApiV3SuppliesSupplyIdOrdersGet200Response) Descriptor() ([]byte, []int) {
	return file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescGZIP(), []int{0}
}

func (x *ApiV3SuppliesSupplyIdOrdersGet200Response) GetOrders() []*SupplyOrder {
	if x != nil {
		return x.Orders
	}
	return nil
}

var File_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto protoreflect.FileDescriptor

var file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDesc = []byte{
	0x0a, 0x48, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x33, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x62, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x77,
	0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x29, 0x41, 0x70, 0x69, 0x56, 0x33, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x77, 0x62, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x4d, 0x41, 0x52, 0x4b,
	0x45, 0x54, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescOnce sync.Once
	file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescData = file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDesc
)

func file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescGZIP() []byte {
	file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescOnce.Do(func() {
		file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescData)
	})
	return file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDescData
}

var file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_goTypes = []interface{}{
	(*ApiV3SuppliesSupplyIdOrdersGet200Response)(nil), // 0: wb.marketplace.v1.ApiV3SuppliesSupplyIdOrdersGet200Response
	(*SupplyOrder)(nil), // 1: wb.marketplace.v1.SupplyOrder
}
var file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_depIdxs = []int32{
	1, // 0: wb.marketplace.v1.ApiV3SuppliesSupplyIdOrdersGet200Response.orders:type_name -> wb.marketplace.v1.SupplyOrder
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_init() }
func file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_init() {
	if File_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto != nil {
		return
	}
	file_wb_marketplace_v1_supply_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiV3SuppliesSupplyIdOrdersGet200Response); i {
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
			RawDescriptor: file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_goTypes,
		DependencyIndexes: file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_depIdxs,
		MessageInfos:      file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_msgTypes,
	}.Build()
	File_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto = out.File
	file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_rawDesc = nil
	file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_goTypes = nil
	file_wb_marketplace_v1_api_v3_supplies_supply_id_orders_get200_response_proto_depIdxs = nil
}
