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
// source: wb/tariffs/v1/commission.proto

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

type Commission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Список комиссий
	Report []*CommissionReportInner `protobuf:"bytes,1,rep,name=report,proto3" json:"report,omitempty"`
}

func (x *Commission) Reset() {
	*x = Commission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1_commission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commission) ProtoMessage() {}

func (x *Commission) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1_commission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commission.ProtoReflect.Descriptor instead.
func (*Commission) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1_commission_proto_rawDescGZIP(), []int{0}
}

func (x *Commission) GetReport() []*CommissionReportInner {
	if x != nil {
		return x.Report
	}
	return nil
}

var File_wb_tariffs_v1_commission_proto protoreflect.FileDescriptor

var file_wb_tariffs_v1_commission_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x2b, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x77, 0x62, 0x2e,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x54, 0x41, 0x52, 0x49, 0x46, 0x46, 0x53, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_tariffs_v1_commission_proto_rawDescOnce sync.Once
	file_wb_tariffs_v1_commission_proto_rawDescData = file_wb_tariffs_v1_commission_proto_rawDesc
)

func file_wb_tariffs_v1_commission_proto_rawDescGZIP() []byte {
	file_wb_tariffs_v1_commission_proto_rawDescOnce.Do(func() {
		file_wb_tariffs_v1_commission_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_tariffs_v1_commission_proto_rawDescData)
	})
	return file_wb_tariffs_v1_commission_proto_rawDescData
}

var file_wb_tariffs_v1_commission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_tariffs_v1_commission_proto_goTypes = []interface{}{
	(*Commission)(nil),            // 0: wb.tariffs.v1.Commission
	(*CommissionReportInner)(nil), // 1: wb.tariffs.v1.CommissionReportInner
}
var file_wb_tariffs_v1_commission_proto_depIdxs = []int32{
	1, // 0: wb.tariffs.v1.Commission.report:type_name -> wb.tariffs.v1.CommissionReportInner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_tariffs_v1_commission_proto_init() }
func file_wb_tariffs_v1_commission_proto_init() {
	if File_wb_tariffs_v1_commission_proto != nil {
		return
	}
	file_wb_tariffs_v1_commission_report_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_tariffs_v1_commission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commission); i {
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
			RawDescriptor: file_wb_tariffs_v1_commission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_tariffs_v1_commission_proto_goTypes,
		DependencyIndexes: file_wb_tariffs_v1_commission_proto_depIdxs,
		MessageInfos:      file_wb_tariffs_v1_commission_proto_msgTypes,
	}.Build()
	File_wb_tariffs_v1_commission_proto = out.File
	file_wb_tariffs_v1_commission_proto_rawDesc = nil
	file_wb_tariffs_v1_commission_proto_goTypes = nil
	file_wb_tariffs_v1_commission_proto_depIdxs = nil
}
