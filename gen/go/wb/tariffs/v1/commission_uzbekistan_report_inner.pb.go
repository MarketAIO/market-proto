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
// source: wb/tariffs/v1/commission_uzbekistan_report_inner.proto

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

type CommissionUzbekistanReportInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Комиссия для продавцов из Узбекистана, %
	KgvpUzbekistan float32 `protobuf:"fixed32,1,opt,name=kgvpUzbekistan,proto3" json:"kgvpUzbekistan,omitempty"`
	// ID родительской категории
	ParentID int32 `protobuf:"varint,2,opt,name=parentID,proto3" json:"parentID,omitempty"`
	// Название родительской категории
	ParentName string `protobuf:"bytes,3,opt,name=parentName,proto3" json:"parentName,omitempty"`
	// ID предмета
	SubjectID int32 `protobuf:"varint,4,opt,name=subjectID,proto3" json:"subjectID,omitempty"`
	// Название предмета
	SubjectName string `protobuf:"bytes,5,opt,name=subjectName,proto3" json:"subjectName,omitempty"`
}

func (x *CommissionUzbekistanReportInner) Reset() {
	*x = CommissionUzbekistanReportInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommissionUzbekistanReportInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommissionUzbekistanReportInner) ProtoMessage() {}

func (x *CommissionUzbekistanReportInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommissionUzbekistanReportInner.ProtoReflect.Descriptor instead.
func (*CommissionUzbekistanReportInner) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescGZIP(), []int{0}
}

func (x *CommissionUzbekistanReportInner) GetKgvpUzbekistan() float32 {
	if x != nil {
		return x.KgvpUzbekistan
	}
	return 0
}

func (x *CommissionUzbekistanReportInner) GetParentID() int32 {
	if x != nil {
		return x.ParentID
	}
	return 0
}

func (x *CommissionUzbekistanReportInner) GetParentName() string {
	if x != nil {
		return x.ParentName
	}
	return ""
}

func (x *CommissionUzbekistanReportInner) GetSubjectID() int32 {
	if x != nil {
		return x.SubjectID
	}
	return 0
}

func (x *CommissionUzbekistanReportInner) GetSubjectName() string {
	if x != nil {
		return x.SubjectName
	}
	return ""
}

var File_wb_tariffs_v1_commission_uzbekistan_report_inner_proto protoreflect.FileDescriptor

var file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDesc = []byte{
	0x0a, 0x36, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x7a, 0x62, 0x65, 0x6b,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xc5, 0x01, 0x0a, 0x1f, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x7a, 0x62, 0x65, 0x6b, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x6b,
	0x67, 0x76, 0x70, 0x55, 0x7a, 0x62, 0x65, 0x6b, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0e, 0x6b, 0x67, 0x76, 0x70, 0x55, 0x7a, 0x62, 0x65, 0x6b, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescOnce sync.Once
	file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescData = file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDesc
)

func file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescGZIP() []byte {
	file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescOnce.Do(func() {
		file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescData)
	})
	return file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDescData
}

var file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_goTypes = []interface{}{
	(*CommissionUzbekistanReportInner)(nil), // 0: wb.tariffs.v1.CommissionUzbekistanReportInner
}
var file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_init() }
func file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_init() {
	if File_wb_tariffs_v1_commission_uzbekistan_report_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommissionUzbekistanReportInner); i {
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
			RawDescriptor: file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_goTypes,
		DependencyIndexes: file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_depIdxs,
		MessageInfos:      file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_msgTypes,
	}.Build()
	File_wb_tariffs_v1_commission_uzbekistan_report_inner_proto = out.File
	file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_rawDesc = nil
	file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_goTypes = nil
	file_wb_tariffs_v1_commission_uzbekistan_report_inner_proto_depIdxs = nil
}
