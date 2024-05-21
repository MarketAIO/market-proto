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
// source: wb/tariffs/v1/_service.proto

package wbTARIFFS

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiV1TariffsBoxGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Дата в формате ГГГГ-ММ-ДД
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ApiV1TariffsBoxGetRequest) Reset() {
	*x = ApiV1TariffsBoxGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1__service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiV1TariffsBoxGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiV1TariffsBoxGetRequest) ProtoMessage() {}

func (x *ApiV1TariffsBoxGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1__service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiV1TariffsBoxGetRequest.ProtoReflect.Descriptor instead.
func (*ApiV1TariffsBoxGetRequest) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1__service_proto_rawDescGZIP(), []int{0}
}

func (x *ApiV1TariffsBoxGetRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ApiV1TariffsCommissionGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Язык полей ответа `parentName` и `subjectName`:     - `ru` — русский   - `en` — английский   - `zh` — китайский
	Locale string `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
}

func (x *ApiV1TariffsCommissionGetRequest) Reset() {
	*x = ApiV1TariffsCommissionGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1__service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiV1TariffsCommissionGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiV1TariffsCommissionGetRequest) ProtoMessage() {}

func (x *ApiV1TariffsCommissionGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1__service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiV1TariffsCommissionGetRequest.ProtoReflect.Descriptor instead.
func (*ApiV1TariffsCommissionGetRequest) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1__service_proto_rawDescGZIP(), []int{1}
}

func (x *ApiV1TariffsCommissionGetRequest) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

type ApiV1TariffsPalletGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Дата в формате ГГГГ-ММ-ДД
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ApiV1TariffsPalletGetRequest) Reset() {
	*x = ApiV1TariffsPalletGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1__service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiV1TariffsPalletGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiV1TariffsPalletGetRequest) ProtoMessage() {}

func (x *ApiV1TariffsPalletGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1__service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiV1TariffsPalletGetRequest.ProtoReflect.Descriptor instead.
func (*ApiV1TariffsPalletGetRequest) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1__service_proto_rawDescGZIP(), []int{2}
}

func (x *ApiV1TariffsPalletGetRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ApiV1TariffsReturnGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Дата в формате ГГГГ-ММ-ДД
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ApiV1TariffsReturnGetRequest) Reset() {
	*x = ApiV1TariffsReturnGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_tariffs_v1__service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiV1TariffsReturnGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiV1TariffsReturnGetRequest) ProtoMessage() {}

func (x *ApiV1TariffsReturnGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wb_tariffs_v1__service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiV1TariffsReturnGetRequest.ProtoReflect.Descriptor instead.
func (*ApiV1TariffsReturnGetRequest) Descriptor() ([]byte, []int) {
	return file_wb_tariffs_v1__service_proto_rawDescGZIP(), []int{3}
}

func (x *ApiV1TariffsReturnGetRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_wb_tariffs_v1__service_proto protoreflect.FileDescriptor

var file_wb_tariffs_v1__service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x77, 0x62, 0x2f, 0x74,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31,
	0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x77, 0x62, 0x2f, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x5f, 0x62, 0x6f,
	0x78, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x5f, 0x70, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x19, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x42, 0x6f, 0x78,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3a,
	0x0a, 0x20, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x32, 0x0a, 0x1c, 0x41, 0x70,
	0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x50, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x32,
	0x0a, 0x1c, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x32, 0xcf, 0x03, 0x0a, 0x0e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x73, 0x42, 0x6f, 0x78, 0x47, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x77, 0x62,
	0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x56,
	0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x42, 0x6f, 0x78, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x42, 0x6f, 0x78,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x19, 0x41, 0x70, 0x69,
	0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x12, 0x2f, 0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x15,
	0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x50, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x73, 0x50, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x50, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x15, 0x41, 0x70, 0x69, 0x56,
	0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x47, 0x65,
	0x74, 0x12, 0x2b, 0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x31, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x77, 0x62, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x77, 0x62, 0x54, 0x41, 0x52, 0x49, 0x46, 0x46, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wb_tariffs_v1__service_proto_rawDescOnce sync.Once
	file_wb_tariffs_v1__service_proto_rawDescData = file_wb_tariffs_v1__service_proto_rawDesc
)

func file_wb_tariffs_v1__service_proto_rawDescGZIP() []byte {
	file_wb_tariffs_v1__service_proto_rawDescOnce.Do(func() {
		file_wb_tariffs_v1__service_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_tariffs_v1__service_proto_rawDescData)
	})
	return file_wb_tariffs_v1__service_proto_rawDescData
}

var file_wb_tariffs_v1__service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wb_tariffs_v1__service_proto_goTypes = []interface{}{
	(*ApiV1TariffsBoxGetRequest)(nil),            // 0: wb.tariffs.v1.ApiV1TariffsBoxGetRequest
	(*ApiV1TariffsCommissionGetRequest)(nil),     // 1: wb.tariffs.v1.ApiV1TariffsCommissionGetRequest
	(*ApiV1TariffsPalletGetRequest)(nil),         // 2: wb.tariffs.v1.ApiV1TariffsPalletGetRequest
	(*ApiV1TariffsReturnGetRequest)(nil),         // 3: wb.tariffs.v1.ApiV1TariffsReturnGetRequest
	(*TariffsBoxResponse)(nil),                   // 4: wb.tariffs.v1.TariffsBoxResponse
	(*ApiV1TariffsCommissionGet200Response)(nil), // 5: wb.tariffs.v1.ApiV1TariffsCommissionGet200Response
	(*TariffsPalletResponse)(nil),                // 6: wb.tariffs.v1.TariffsPalletResponse
	(*ReturnTariffsResponse)(nil),                // 7: wb.tariffs.v1.ReturnTariffsResponse
}
var file_wb_tariffs_v1__service_proto_depIdxs = []int32{
	0, // 0: wb.tariffs.v1.DefaultService.ApiV1TariffsBoxGet:input_type -> wb.tariffs.v1.ApiV1TariffsBoxGetRequest
	1, // 1: wb.tariffs.v1.DefaultService.ApiV1TariffsCommissionGet:input_type -> wb.tariffs.v1.ApiV1TariffsCommissionGetRequest
	2, // 2: wb.tariffs.v1.DefaultService.ApiV1TariffsPalletGet:input_type -> wb.tariffs.v1.ApiV1TariffsPalletGetRequest
	3, // 3: wb.tariffs.v1.DefaultService.ApiV1TariffsReturnGet:input_type -> wb.tariffs.v1.ApiV1TariffsReturnGetRequest
	4, // 4: wb.tariffs.v1.DefaultService.ApiV1TariffsBoxGet:output_type -> wb.tariffs.v1.TariffsBoxResponse
	5, // 5: wb.tariffs.v1.DefaultService.ApiV1TariffsCommissionGet:output_type -> wb.tariffs.v1.ApiV1TariffsCommissionGet200Response
	6, // 6: wb.tariffs.v1.DefaultService.ApiV1TariffsPalletGet:output_type -> wb.tariffs.v1.TariffsPalletResponse
	7, // 7: wb.tariffs.v1.DefaultService.ApiV1TariffsReturnGet:output_type -> wb.tariffs.v1.ReturnTariffsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_tariffs_v1__service_proto_init() }
func file_wb_tariffs_v1__service_proto_init() {
	if File_wb_tariffs_v1__service_proto != nil {
		return
	}
	file_wb_tariffs_v1_api_v1_tariffs_commission_get200_response_proto_init()
	file_wb_tariffs_v1_return_tariffs_response_proto_init()
	file_wb_tariffs_v1_tariffs_box_response_proto_init()
	file_wb_tariffs_v1_tariffs_pallet_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_tariffs_v1__service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiV1TariffsBoxGetRequest); i {
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
		file_wb_tariffs_v1__service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiV1TariffsCommissionGetRequest); i {
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
		file_wb_tariffs_v1__service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiV1TariffsPalletGetRequest); i {
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
		file_wb_tariffs_v1__service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiV1TariffsReturnGetRequest); i {
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
			RawDescriptor: file_wb_tariffs_v1__service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wb_tariffs_v1__service_proto_goTypes,
		DependencyIndexes: file_wb_tariffs_v1__service_proto_depIdxs,
		MessageInfos:      file_wb_tariffs_v1__service_proto_msgTypes,
	}.Build()
	File_wb_tariffs_v1__service_proto = out.File
	file_wb_tariffs_v1__service_proto_rawDesc = nil
	file_wb_tariffs_v1__service_proto_goTypes = nil
	file_wb_tariffs_v1__service_proto_depIdxs = nil
}
