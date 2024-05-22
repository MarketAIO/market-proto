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
// source: wb/promotion/v1/adv_v1_adverts_get200_response_inner.proto

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

type AdvV1AdvertsGet200ResponseInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор медиакампании
	AdvertId int32 `protobuf:"varint,1,opt,name=advertId,proto3" json:"advertId,omitempty"`
	// Название медиакампании
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Название бренда
	Brand string `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	// <dl> <dt>Тип медиакампании:</dt> <dd><code>1</code> - размещение по дням</dd> <dd><code>2</code> - размещение по просмотрам</dd> </dl>
	Type int32 `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	// <dl> <dt>Статус медиакампании:</dt>   <dd><code>1</code> - черновик</dd>   <dd><code>2</code> - модерация   <dd><code>3</code> - отклонено (с возможностью вернуть на модерацию)</dd>   <dd><code>4</code> - одобрено</dd>   <dd><code>5</code> - запланировано</dd>   <dd><code>6</code> - на показах</dd>   <dd><code>7</code> - завершено</dd>   <dd><code>8</code> - отказался</dd>   <dd><code>9</code> - приостановлена продавцом</dd>   <dd><code>10</code> - пауза по дневному лимиту</dd>   <dd><code>11</code> - пауза по расходу бюджета</dd> </dl>
	Status int32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	// Время создания медиакампании
	CreateTime string `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// Время завершения медиакампании (при наличии)
	EndTime string `protobuf:"bytes,7,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *AdvV1AdvertsGet200ResponseInner) Reset() {
	*x = AdvV1AdvertsGet200ResponseInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1AdvertsGet200ResponseInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1AdvertsGet200ResponseInner) ProtoMessage() {}

func (x *AdvV1AdvertsGet200ResponseInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1AdvertsGet200ResponseInner.ProtoReflect.Descriptor instead.
func (*AdvV1AdvertsGet200ResponseInner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1AdvertsGet200ResponseInner) GetAdvertId() int32 {
	if x != nil {
		return x.AdvertId
	}
	return 0
}

func (x *AdvV1AdvertsGet200ResponseInner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdvV1AdvertsGet200ResponseInner) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *AdvV1AdvertsGet200ResponseInner) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AdvV1AdvertsGet200ResponseInner) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdvV1AdvertsGet200ResponseInner) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AdvV1AdvertsGet200ResponseInner) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

var File_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73,
	0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xcd, 0x01,
	0x0a, 0x1f, 0x41, 0x64, 0x76, 0x56, 0x31, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x73, 0x47, 0x65,
	0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x46, 0x5a,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescData = file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_goTypes = []interface{}{
	(*AdvV1AdvertsGet200ResponseInner)(nil), // 0: wb.promotion.v1.AdvV1AdvertsGet200ResponseInner
}
var file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_init() }
func file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_init() {
	if File_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1AdvertsGet200ResponseInner); i {
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
			RawDescriptor: file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto = out.File
	file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_adverts_get200_response_inner_proto_depIdxs = nil
}
