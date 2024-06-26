//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/promotion/v1/adv_v1_advert_get200_response_extended.proto

package wbPromotion

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

type AdvV1AdvertGet200ResponseExtended struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Комментарий модератора (при наличии)
	Reason *string `protobuf:"bytes,1,opt,name=reason,proto3,oneof" json:"reason,omitempty"`
	// Затраты
	Expenses int32 `protobuf:"varint,2,opt,name=expenses,proto3" json:"expenses,omitempty"`
	// Начало показов медиакампании
	From string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	// Конец показов медиакампании
	To string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	// Время изменения медиакампании
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Стоимость размещения по дням (для типа 1)
	Price int32 `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	// Остаток бюджета (для типа 2)
	Budget int32 `protobuf:"varint,7,opt,name=budget,proto3" json:"budget,omitempty"`
	// Источник списания (1 - баланс, 2 - счет)
	Operation int32 `protobuf:"varint,8,opt,name=operation,proto3" json:"operation,omitempty"`
	// Идентификатор контракта (для продавцов на контракте)
	ContractId int32 `protobuf:"varint,9,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
}

func (x *AdvV1AdvertGet200ResponseExtended) Reset() {
	*x = AdvV1AdvertGet200ResponseExtended{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1AdvertGet200ResponseExtended) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1AdvertGet200ResponseExtended) ProtoMessage() {}

func (x *AdvV1AdvertGet200ResponseExtended) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1AdvertGet200ResponseExtended.ProtoReflect.Descriptor instead.
func (*AdvV1AdvertGet200ResponseExtended) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1AdvertGet200ResponseExtended) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseExtended) GetExpenses() int32 {
	if x != nil {
		return x.Expenses
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseExtended) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseExtended) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseExtended) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseExtended) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseExtended) GetBudget() int32 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseExtended) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseExtended) GetContractId() int32 {
	if x != nil {
		return x.ContractId
	}
	return 0
}

var File_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f,
	0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22,
	0x97, 0x02, 0x0a, 0x21, 0x41, 0x64, 0x76, 0x56, 0x31, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x47,
	0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49,
	0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescData = file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_goTypes = []interface{}{
	(*AdvV1AdvertGet200ResponseExtended)(nil), // 0: wb.promotion.v1.AdvV1AdvertGet200ResponseExtended
}
var file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_init() }
func file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_init() {
	if File_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1AdvertGet200ResponseExtended); i {
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
	file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto = out.File
	file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_advert_get200_response_extended_proto_depIdxs = nil
}
