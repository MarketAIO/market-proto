//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam).
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/analytics/v1/api_v1_analytics_storage_coefficient_get200_response_report_inner.proto

package wbANALYTICS

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

type ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Фактическая высота, см
	ActualHeight float32 `protobuf:"fixed32,1,opt,name=actualHeight,proto3" json:"actualHeight,omitempty"`
	// Фактическая длина, см
	ActualLength float32 `protobuf:"fixed32,2,opt,name=actualLength,proto3" json:"actualLength,omitempty"`
	// Фактический объём, л
	ActualVolume float32 `protobuf:"fixed32,3,opt,name=actualVolume,proto3" json:"actualVolume,omitempty"`
	// Фактическая ширина, см
	ActualWidth float32 `protobuf:"fixed32,4,opt,name=actualWidth,proto3" json:"actualWidth,omitempty"`
	// Дата измерения в формате RFC 3339. Для расчёта используются измерения за 30 дней, до начала отчётной недели
	Date string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	// Разница в габаритах, %
	DimensionDifference float32 `protobuf:"fixed32,6,opt,name=dimensionDifference,proto3" json:"dimensionDifference,omitempty"`
	// Высота из карточки, см
	Height float32 `protobuf:"fixed32,7,opt,name=height,proto3" json:"height,omitempty"`
	// Длина из карточки, см
	Length float32 `protobuf:"fixed32,8,opt,name=length,proto3" json:"length,omitempty"`
	// Коэффициент логистики и хранения для товара
	LogWarehouseCoef float32 `protobuf:"fixed32,9,opt,name=logWarehouseCoef,proto3" json:"logWarehouseCoef,omitempty"`
	// Артикул Wildberries
	NmID int32 `protobuf:"varint,10,opt,name=nmID,proto3" json:"nmID,omitempty"`
	// Фото измерений
	PhotoUrls string `protobuf:"bytes,11,opt,name=photoUrls,proto3" json:"photoUrls,omitempty"`
	// Наименование товара
	Title string `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	// Объём из карточки, л
	Volume float32 `protobuf:"fixed32,13,opt,name=volume,proto3" json:"volume,omitempty"`
	// Ширина из карточки, см
	Width float32 `protobuf:"fixed32,14,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) Reset() {
	*x = ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) ProtoMessage() {}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner.ProtoReflect.Descriptor instead.
func (*ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) Descriptor() ([]byte, []int) {
	return file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescGZIP(), []int{0}
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetActualHeight() float32 {
	if x != nil {
		return x.ActualHeight
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetActualLength() float32 {
	if x != nil {
		return x.ActualLength
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetActualVolume() float32 {
	if x != nil {
		return x.ActualVolume
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetActualWidth() float32 {
	if x != nil {
		return x.ActualWidth
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetDimensionDifference() float32 {
	if x != nil {
		return x.DimensionDifference
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetLength() float32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetLogWarehouseCoef() float32 {
	if x != nil {
		return x.LogWarehouseCoef
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetNmID() int32 {
	if x != nil {
		return x.NmID
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetPhotoUrls() string {
	if x != nil {
		return x.PhotoUrls
	}
	return ""
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetVolume() float32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

var File_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto protoreflect.FileDescriptor

var file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDesc = []byte{
	0x0a, 0x57, 0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x65, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xe1, 0x03, 0x0a, 0x39, 0x41,
	0x70, 0x69, 0x56, 0x31, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x47,
	0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x75,
	0x61, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x57, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x75, 0x61,
	0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x10,
	0x6c, 0x6f, 0x67, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x43, 0x6f, 0x65, 0x66,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x6c, 0x6f, 0x67, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x43, 0x6f, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6d, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x42, 0x49,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62,
	0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62,
	0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescOnce sync.Once
	file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescData = file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDesc
)

func file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescGZIP() []byte {
	file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescOnce.Do(func() {
		file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescData)
	})
	return file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDescData
}

var file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_goTypes = []interface{}{
	(*ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner)(nil), // 0: wb.analytics.v1.ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner
}
var file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_init()
}
func file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_init() {
	if File_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner); i {
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
			RawDescriptor: file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_goTypes,
		DependencyIndexes: file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_depIdxs,
		MessageInfos:      file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_msgTypes,
	}.Build()
	File_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto = out.File
	file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_rawDesc = nil
	file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_goTypes = nil
	file_wb_analytics_v1_api_v1_analytics_storage_coefficient_get200_response_report_inner_proto_depIdxs = nil
}
