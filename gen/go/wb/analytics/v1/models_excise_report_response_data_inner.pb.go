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
// source: wb/analytics/v1/models_excise_report_response_data_inner.proto

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

type ModelsExciseReportResponseDataInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Страна покупателя
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Цена товара, с НДС
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	// Валюта
	CurrencyNameShort string `protobuf:"bytes,3,opt,name=currency_name_short,json=currencyNameShort,proto3" json:"currency_name_short,omitempty"`
	// Код маркировки
	ExciseShort string `protobuf:"bytes,4,opt,name=excise_short,json=exciseShort,proto3" json:"excise_short,omitempty"`
	// Баркод
	Barcode string `protobuf:"bytes,5,opt,name=barcode,proto3" json:"barcode,omitempty"`
	// Артикул Wildberries
	NmId int32 `protobuf:"varint,6,opt,name=nm_id,json=nmId,proto3" json:"nm_id,omitempty"`
	// Тип операции, если есть:    * `1` — вывод из оборота   * `2` — возврат в оборот
	OperationTypeId int32 `protobuf:"varint,7,opt,name=operation_type_id,json=operationTypeId,proto3" json:"operation_type_id,omitempty"`
	// Номер фискального документа (чека полного расчёта), если есть
	FiscalDocNumber int32 `protobuf:"varint,8,opt,name=fiscal_doc_number,json=fiscalDocNumber,proto3" json:"fiscal_doc_number,omitempty"`
	// Дата фискализации (дата в чеке), если есть, `ГГГГ-ММ-ДД`
	FiscalDt string `protobuf:"bytes,9,opt,name=fiscal_dt,json=fiscalDt,proto3" json:"fiscal_dt,omitempty"`
	// Номер фискального накопителя, если есть
	FiscalDriveNumber string `protobuf:"bytes,10,opt,name=fiscal_drive_number,json=fiscalDriveNumber,proto3" json:"fiscal_drive_number,omitempty"`
	// `Rid`
	Rid int32 `protobuf:"varint,11,opt,name=rid,proto3" json:"rid,omitempty"`
	// `Srid`
	Srid string `protobuf:"bytes,12,opt,name=srid,proto3" json:"srid,omitempty"`
}

func (x *ModelsExciseReportResponseDataInner) Reset() {
	*x = ModelsExciseReportResponseDataInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_analytics_v1_models_excise_report_response_data_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelsExciseReportResponseDataInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelsExciseReportResponseDataInner) ProtoMessage() {}

func (x *ModelsExciseReportResponseDataInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_analytics_v1_models_excise_report_response_data_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelsExciseReportResponseDataInner.ProtoReflect.Descriptor instead.
func (*ModelsExciseReportResponseDataInner) Descriptor() ([]byte, []int) {
	return file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescGZIP(), []int{0}
}

func (x *ModelsExciseReportResponseDataInner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelsExciseReportResponseDataInner) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ModelsExciseReportResponseDataInner) GetCurrencyNameShort() string {
	if x != nil {
		return x.CurrencyNameShort
	}
	return ""
}

func (x *ModelsExciseReportResponseDataInner) GetExciseShort() string {
	if x != nil {
		return x.ExciseShort
	}
	return ""
}

func (x *ModelsExciseReportResponseDataInner) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *ModelsExciseReportResponseDataInner) GetNmId() int32 {
	if x != nil {
		return x.NmId
	}
	return 0
}

func (x *ModelsExciseReportResponseDataInner) GetOperationTypeId() int32 {
	if x != nil {
		return x.OperationTypeId
	}
	return 0
}

func (x *ModelsExciseReportResponseDataInner) GetFiscalDocNumber() int32 {
	if x != nil {
		return x.FiscalDocNumber
	}
	return 0
}

func (x *ModelsExciseReportResponseDataInner) GetFiscalDt() string {
	if x != nil {
		return x.FiscalDt
	}
	return ""
}

func (x *ModelsExciseReportResponseDataInner) GetFiscalDriveNumber() string {
	if x != nil {
		return x.FiscalDriveNumber
	}
	return ""
}

func (x *ModelsExciseReportResponseDataInner) GetRid() int32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *ModelsExciseReportResponseDataInner) GetSrid() string {
	if x != nil {
		return x.Srid
	}
	return ""
}

var File_wb_analytics_v1_models_excise_report_response_data_inner_proto protoreflect.FileDescriptor

var file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5f, 0x65, 0x78, 0x63, 0x69, 0x73, 0x65, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x77, 0x62, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x31, 0x22, 0x9c, 0x03, 0x0a, 0x23, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x45, 0x78, 0x63, 0x69,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x69, 0x73, 0x65, 0x5f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x69, 0x73,
	0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x13, 0x0a, 0x05, 0x6e, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6e, 0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x66, 0x69,
	0x73, 0x63, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x44, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69,
	0x73, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x72, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x72, 0x69, 0x64,
	0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x77, 0x62, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x62, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescOnce sync.Once
	file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescData = file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDesc
)

func file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescGZIP() []byte {
	file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescOnce.Do(func() {
		file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescData)
	})
	return file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDescData
}

var file_wb_analytics_v1_models_excise_report_response_data_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_analytics_v1_models_excise_report_response_data_inner_proto_goTypes = []interface{}{
	(*ModelsExciseReportResponseDataInner)(nil), // 0: wb.analytics.v1.ModelsExciseReportResponseDataInner
}
var file_wb_analytics_v1_models_excise_report_response_data_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_analytics_v1_models_excise_report_response_data_inner_proto_init() }
func file_wb_analytics_v1_models_excise_report_response_data_inner_proto_init() {
	if File_wb_analytics_v1_models_excise_report_response_data_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_analytics_v1_models_excise_report_response_data_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelsExciseReportResponseDataInner); i {
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
			RawDescriptor: file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_analytics_v1_models_excise_report_response_data_inner_proto_goTypes,
		DependencyIndexes: file_wb_analytics_v1_models_excise_report_response_data_inner_proto_depIdxs,
		MessageInfos:      file_wb_analytics_v1_models_excise_report_response_data_inner_proto_msgTypes,
	}.Build()
	File_wb_analytics_v1_models_excise_report_response_data_inner_proto = out.File
	file_wb_analytics_v1_models_excise_report_response_data_inner_proto_rawDesc = nil
	file_wb_analytics_v1_models_excise_report_response_data_inner_proto_goTypes = nil
	file_wb_analytics_v1_models_excise_report_response_data_inner_proto_depIdxs = nil
}