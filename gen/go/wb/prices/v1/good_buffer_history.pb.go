//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/prices/v1/good_buffer_history.proto

package wbPrices

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

type GoodBufferHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Артикул Wildberries
	NmID int32 `protobuf:"varint,1,opt,name=nmID,proto3" json:"nmID,omitempty"`
	// Артикул продавца
	VendorCode string `protobuf:"bytes,2,opt,name=vendorCode,proto3" json:"vendorCode,omitempty"`
	// ID размера. В методах контента это поле `chrtID`
	SizeID int32 `protobuf:"varint,3,opt,name=sizeID,proto3" json:"sizeID,omitempty"`
	// Размер
	TechSizeName string `protobuf:"bytes,4,opt,name=techSizeName,proto3" json:"techSizeName,omitempty"`
	// Цена
	Price int32 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	// Валюта, по стандарту ISO 4217
	CurrencyIsoCode4217 string `protobuf:"bytes,6,opt,name=currencyIsoCode4217,proto3" json:"currencyIsoCode4217,omitempty"`
	// Скидка, %
	Discount int32 `protobuf:"varint,7,opt,name=discount,proto3" json:"discount,omitempty"`
	// Статус товара: `1` — в обработке
	Status int32 `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	// Текст ошибки
	ErrorText string `protobuf:"bytes,9,opt,name=errorText,proto3" json:"errorText,omitempty"`
}

func (x *GoodBufferHistory) Reset() {
	*x = GoodBufferHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_prices_v1_good_buffer_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodBufferHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodBufferHistory) ProtoMessage() {}

func (x *GoodBufferHistory) ProtoReflect() protoreflect.Message {
	mi := &file_wb_prices_v1_good_buffer_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodBufferHistory.ProtoReflect.Descriptor instead.
func (*GoodBufferHistory) Descriptor() ([]byte, []int) {
	return file_wb_prices_v1_good_buffer_history_proto_rawDescGZIP(), []int{0}
}

func (x *GoodBufferHistory) GetNmID() int32 {
	if x != nil {
		return x.NmID
	}
	return 0
}

func (x *GoodBufferHistory) GetVendorCode() string {
	if x != nil {
		return x.VendorCode
	}
	return ""
}

func (x *GoodBufferHistory) GetSizeID() int32 {
	if x != nil {
		return x.SizeID
	}
	return 0
}

func (x *GoodBufferHistory) GetTechSizeName() string {
	if x != nil {
		return x.TechSizeName
	}
	return ""
}

func (x *GoodBufferHistory) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodBufferHistory) GetCurrencyIsoCode4217() string {
	if x != nil {
		return x.CurrencyIsoCode4217
	}
	return ""
}

func (x *GoodBufferHistory) GetDiscount() int32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *GoodBufferHistory) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GoodBufferHistory) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

var File_wb_prices_v1_good_buffer_history_proto protoreflect.FileDescriptor

var file_wb_prices_v1_good_buffer_history_proto_rawDesc = []byte{
	0x0a, 0x26, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x6f, 0x6f, 0x64, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x9d, 0x02, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x64, 0x42,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x63, 0x68,
	0x53, 0x69, 0x7a, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x65, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x73,
	0x6f, 0x43, 0x6f, 0x64, 0x65, 0x34, 0x32, 0x31, 0x37, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x73, 0x6f, 0x43, 0x6f, 0x64, 0x65,
	0x34, 0x32, 0x31, 0x37, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x54, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x54, 0x65, 0x78, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x62, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_prices_v1_good_buffer_history_proto_rawDescOnce sync.Once
	file_wb_prices_v1_good_buffer_history_proto_rawDescData = file_wb_prices_v1_good_buffer_history_proto_rawDesc
)

func file_wb_prices_v1_good_buffer_history_proto_rawDescGZIP() []byte {
	file_wb_prices_v1_good_buffer_history_proto_rawDescOnce.Do(func() {
		file_wb_prices_v1_good_buffer_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_prices_v1_good_buffer_history_proto_rawDescData)
	})
	return file_wb_prices_v1_good_buffer_history_proto_rawDescData
}

var file_wb_prices_v1_good_buffer_history_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_prices_v1_good_buffer_history_proto_goTypes = []interface{}{
	(*GoodBufferHistory)(nil), // 0: wb.prices.v1.GoodBufferHistory
}
var file_wb_prices_v1_good_buffer_history_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_prices_v1_good_buffer_history_proto_init() }
func file_wb_prices_v1_good_buffer_history_proto_init() {
	if File_wb_prices_v1_good_buffer_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_prices_v1_good_buffer_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodBufferHistory); i {
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
			RawDescriptor: file_wb_prices_v1_good_buffer_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_prices_v1_good_buffer_history_proto_goTypes,
		DependencyIndexes: file_wb_prices_v1_good_buffer_history_proto_depIdxs,
		MessageInfos:      file_wb_prices_v1_good_buffer_history_proto_msgTypes,
	}.Build()
	File_wb_prices_v1_good_buffer_history_proto = out.File
	file_wb_prices_v1_good_buffer_history_proto_rawDesc = nil
	file_wb_prices_v1_good_buffer_history_proto_goTypes = nil
	file_wb_prices_v1_good_buffer_history_proto_depIdxs = nil
}
