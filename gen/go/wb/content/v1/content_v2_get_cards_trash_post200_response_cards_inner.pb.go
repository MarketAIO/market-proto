//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code>
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/content/v1/content_v2_get_cards_trash_post200_response_cards_inner.proto

package wbContent

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

type ContentV2GetCardsTrashPost200ResponseCardsInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Артикул WB
	NmID int32 `protobuf:"varint,1,opt,name=nmID,proto3" json:"nmID,omitempty"`
	// Артикул продавца
	VendorCode string `protobuf:"bytes,2,opt,name=vendorCode,proto3" json:"vendorCode,omitempty"`
	// Идентификатор предмета
	SubjectID string `protobuf:"bytes,3,opt,name=subjectID,proto3" json:"subjectID,omitempty"`
	// Название предмета
	SubjectName string `protobuf:"bytes,4,opt,name=subjectName,proto3" json:"subjectName,omitempty"`
	// Массив фото
	Photos []*ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner `protobuf:"bytes,5,rep,name=photos,proto3" json:"photos,omitempty"`
	// URL видео
	Video string `protobuf:"bytes,6,opt,name=video,proto3" json:"video,omitempty"`
	// Массив размеров
	Sizes      []*ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner `protobuf:"bytes,7,rep,name=sizes,proto3" json:"sizes,omitempty"`
	Dimensions *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions   `protobuf:"bytes,8,opt,name=dimensions,proto3" json:"dimensions,omitempty"`
	// Массив характеристик, при наличии
	Characteristics []*ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner `protobuf:"bytes,9,rep,name=characteristics,proto3" json:"characteristics,omitempty"`
	CreatedAt       string                                                                 `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	TrashedAt       string                                                                 `protobuf:"bytes,11,opt,name=trashedAt,proto3" json:"trashedAt,omitempty"`
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) Reset() {
	*x = ContentV2GetCardsTrashPost200ResponseCardsInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentV2GetCardsTrashPost200ResponseCardsInner) ProtoMessage() {}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentV2GetCardsTrashPost200ResponseCardsInner.ProtoReflect.Descriptor instead.
func (*ContentV2GetCardsTrashPost200ResponseCardsInner) Descriptor() ([]byte, []int) {
	return file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescGZIP(), []int{0}
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetNmID() int32 {
	if x != nil {
		return x.NmID
	}
	return 0
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVendorCode() string {
	if x != nil {
		return x.VendorCode
	}
	return ""
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectID() string {
	if x != nil {
		return x.SubjectID
	}
	return ""
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSubjectName() string {
	if x != nil {
		return x.SubjectName
	}
	return ""
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetPhotos() []*ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetSizes() []*ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner {
	if x != nil {
		return x.Sizes
	}
	return nil
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetDimensions() *ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCharacteristics() []*ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner {
	if x != nil {
		return x.Characteristics
	}
	return nil
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ContentV2GetCardsTrashPost200ResponseCardsInner) GetTrashedAt() string {
	if x != nil {
		return x.TrashedAt
	}
	return ""
}

var File_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto protoreflect.FileDescriptor

var file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x32,
	0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x57, 0x77, 0x62,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x61, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x73, 0x68, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x56, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x76, 0x32, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x74, 0x72, 0x61,
	0x73, 0x68, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x57, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x32, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x64, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x32,
	0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x5f, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x05, 0x0a, 0x2f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x54,
	0x72, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x60, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x48, 0x2e, 0x77, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x5e, 0x0a, 0x05, 0x73, 0x69, 0x7a,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x77, 0x62, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x56, 0x32, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x54, 0x72, 0x61, 0x73, 0x68, 0x50,
	0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x73, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x05, 0x73, 0x69, 0x7a, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x0a, 0x64, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e,
	0x77, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x54,
	0x72, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x7c, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x77,
	0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x56, 0x32, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x73, 0x54, 0x72,
	0x61, 0x73, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x42, 0x42, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescOnce sync.Once
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescData = file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDesc
)

func file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescGZIP() []byte {
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescOnce.Do(func() {
		file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescData)
	})
	return file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDescData
}

var file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_goTypes = []interface{}{
	(*ContentV2GetCardsTrashPost200ResponseCardsInner)(nil),                     // 0: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInner
	(*ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner)(nil),           // 1: wb.content.v1.ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner
	(*ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner)(nil),           // 2: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner
	(*ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions)(nil),           // 3: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions
	(*ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner)(nil), // 4: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner
}
var file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_depIdxs = []int32{
	1, // 0: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInner.photos:type_name -> wb.content.v1.ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner
	2, // 1: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInner.sizes:type_name -> wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner
	3, // 2: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInner.dimensions:type_name -> wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions
	4, // 3: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInner.characteristics:type_name -> wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_init() }
func file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_init() {
	if File_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto != nil {
		return
	}
	file_wb_content_v1_content_v2_get_cards_list_post200_response_cards_inner_photos_inner_proto_init()
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_characteristics_inner_proto_init()
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_dimensions_proto_init()
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_sizes_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentV2GetCardsTrashPost200ResponseCardsInner); i {
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
			RawDescriptor: file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_goTypes,
		DependencyIndexes: file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_depIdxs,
		MessageInfos:      file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_msgTypes,
	}.Build()
	File_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto = out.File
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_rawDesc = nil
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_goTypes = nil
	file_wb_content_v1_content_v2_get_cards_trash_post200_response_cards_inner_proto_depIdxs = nil
}
