//
//Описание API Статистики
//
//С помощью этих методов можно получить отчёты.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/statistics/v1/incomes_item.proto

package wbSTATISTICS

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

// Текущий статус поставки
type IncomesItem_StatusEnum int32

const (
	IncomesItem_StatusEnum__ IncomesItem_StatusEnum = 0
)

// Enum value maps for IncomesItem_StatusEnum.
var (
	IncomesItem_StatusEnum_name = map[int32]string{
		0: "StatusEnum__",
	}
	IncomesItem_StatusEnum_value = map[string]int32{
		"StatusEnum__": 0,
	}
)

func (x IncomesItem_StatusEnum) Enum() *IncomesItem_StatusEnum {
	p := new(IncomesItem_StatusEnum)
	*p = x
	return p
}

func (x IncomesItem_StatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IncomesItem_StatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_wb_statistics_v1_incomes_item_proto_enumTypes[0].Descriptor()
}

func (IncomesItem_StatusEnum) Type() protoreflect.EnumType {
	return &file_wb_statistics_v1_incomes_item_proto_enumTypes[0]
}

func (x IncomesItem_StatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IncomesItem_StatusEnum.Descriptor instead.
func (IncomesItem_StatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_wb_statistics_v1_incomes_item_proto_rawDescGZIP(), []int{0, 0}
}

type IncomesItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Номер поставки
	IncomeId int32 `protobuf:"varint,1,opt,name=incomeId,proto3" json:"incomeId,omitempty"`
	// Номер УПД
	Number string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	// Дата поступления. Если часовой пояс не указан, то берется Московское время UTC+3.
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	// Дата и время обновления информации в сервисе. Это поле соответствует параметру `dateFrom` в запросе. Если часовой пояс не указан, то берется Московское время UTC+3.
	LastChangeDate string `protobuf:"bytes,4,opt,name=lastChangeDate,proto3" json:"lastChangeDate,omitempty"`
	// Артикул продавца
	SupplierArticle string `protobuf:"bytes,5,opt,name=supplierArticle,proto3" json:"supplierArticle,omitempty"`
	// Размер товара (пример S, M, L, XL, 42, 42-43)
	TechSize string `protobuf:"bytes,6,opt,name=techSize,proto3" json:"techSize,omitempty"`
	// Бар-код
	Barcode string `protobuf:"bytes,7,opt,name=barcode,proto3" json:"barcode,omitempty"`
	// Количество
	Quantity int32 `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// Цена из УПД
	TotalPrice float32 `protobuf:"fixed32,9,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	// Дата принятия (закрытия) в WB. Если часовой пояс не указан, то берется Московское время UTC+3
	DateClose string `protobuf:"bytes,10,opt,name=dateClose,proto3" json:"dateClose,omitempty"`
	// Название склада
	WarehouseName string `protobuf:"bytes,11,opt,name=warehouseName,proto3" json:"warehouseName,omitempty"`
	// Артикул WB
	NmId   int32                  `protobuf:"varint,12,opt,name=nmId,proto3" json:"nmId,omitempty"`
	Status IncomesItem_StatusEnum `protobuf:"varint,13,opt,name=status,proto3,enum=wb.statistics.v1.IncomesItem_StatusEnum" json:"status,omitempty"`
}

func (x *IncomesItem) Reset() {
	*x = IncomesItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_statistics_v1_incomes_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomesItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomesItem) ProtoMessage() {}

func (x *IncomesItem) ProtoReflect() protoreflect.Message {
	mi := &file_wb_statistics_v1_incomes_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomesItem.ProtoReflect.Descriptor instead.
func (*IncomesItem) Descriptor() ([]byte, []int) {
	return file_wb_statistics_v1_incomes_item_proto_rawDescGZIP(), []int{0}
}

func (x *IncomesItem) GetIncomeId() int32 {
	if x != nil {
		return x.IncomeId
	}
	return 0
}

func (x *IncomesItem) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *IncomesItem) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *IncomesItem) GetLastChangeDate() string {
	if x != nil {
		return x.LastChangeDate
	}
	return ""
}

func (x *IncomesItem) GetSupplierArticle() string {
	if x != nil {
		return x.SupplierArticle
	}
	return ""
}

func (x *IncomesItem) GetTechSize() string {
	if x != nil {
		return x.TechSize
	}
	return ""
}

func (x *IncomesItem) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *IncomesItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *IncomesItem) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *IncomesItem) GetDateClose() string {
	if x != nil {
		return x.DateClose
	}
	return ""
}

func (x *IncomesItem) GetWarehouseName() string {
	if x != nil {
		return x.WarehouseName
	}
	return ""
}

func (x *IncomesItem) GetNmId() int32 {
	if x != nil {
		return x.NmId
	}
	return 0
}

func (x *IncomesItem) GetStatus() IncomesItem_StatusEnum {
	if x != nil {
		return x.Status
	}
	return IncomesItem_StatusEnum__
}

var File_wb_statistics_v1_incomes_item_proto protoreflect.FileDescriptor

var file_wb_statistics_v1_incomes_item_proto_rawDesc = []byte{
	0x0a, 0x23, 0x77, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x62, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xd3, 0x03, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x63, 0x6f, 0x6d,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x63, 0x6f, 0x6d,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x77, 0x62,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1e, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x5f, 0x5f, 0x10, 0x00, 0x42, 0x4b, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62,
	0x53, 0x54, 0x41, 0x54, 0x49, 0x53, 0x54, 0x49, 0x43, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wb_statistics_v1_incomes_item_proto_rawDescOnce sync.Once
	file_wb_statistics_v1_incomes_item_proto_rawDescData = file_wb_statistics_v1_incomes_item_proto_rawDesc
)

func file_wb_statistics_v1_incomes_item_proto_rawDescGZIP() []byte {
	file_wb_statistics_v1_incomes_item_proto_rawDescOnce.Do(func() {
		file_wb_statistics_v1_incomes_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_statistics_v1_incomes_item_proto_rawDescData)
	})
	return file_wb_statistics_v1_incomes_item_proto_rawDescData
}

var file_wb_statistics_v1_incomes_item_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wb_statistics_v1_incomes_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_statistics_v1_incomes_item_proto_goTypes = []interface{}{
	(IncomesItem_StatusEnum)(0), // 0: wb.statistics.v1.IncomesItem.StatusEnum
	(*IncomesItem)(nil),         // 1: wb.statistics.v1.IncomesItem
}
var file_wb_statistics_v1_incomes_item_proto_depIdxs = []int32{
	0, // 0: wb.statistics.v1.IncomesItem.status:type_name -> wb.statistics.v1.IncomesItem.StatusEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_statistics_v1_incomes_item_proto_init() }
func file_wb_statistics_v1_incomes_item_proto_init() {
	if File_wb_statistics_v1_incomes_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_statistics_v1_incomes_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomesItem); i {
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
			RawDescriptor: file_wb_statistics_v1_incomes_item_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_statistics_v1_incomes_item_proto_goTypes,
		DependencyIndexes: file_wb_statistics_v1_incomes_item_proto_depIdxs,
		EnumInfos:         file_wb_statistics_v1_incomes_item_proto_enumTypes,
		MessageInfos:      file_wb_statistics_v1_incomes_item_proto_msgTypes,
	}.Build()
	File_wb_statistics_v1_incomes_item_proto = out.File
	file_wb_statistics_v1_incomes_item_proto_rawDesc = nil
	file_wb_statistics_v1_incomes_item_proto_goTypes = nil
	file_wb_statistics_v1_incomes_item_proto_depIdxs = nil
}