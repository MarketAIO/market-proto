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
// source: wb/promotion/v1/response_info_advert.proto

package wbPROMOTION

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

type ResponseInfoAdvert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Идентификатор кампании
	AdvertId int32 `protobuf:"varint,1,opt,name=advertId,proto3" json:"advertId,omitempty"`
	// <dl>   <dt>Тип кампании:</dt>     <dd><code>4</code> - кампания  в каталоге</dd>     <dd><code>5</code> - кампания в карточке товара</dd>     <dd><code>6</code> - кампания в поиске</dd>     <dd><code>7</code> - кампания в рекомендациях на главной странице</dd>   </dl>
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	// <dl> <dt>Статус кампании:</dt> <dd><code>-1</code> - кампания в процессе удаления </dd> <dd><code>4</code> - готова к запуску </dd> <dd><code>7</code> - Кампания завершена</dd> <dd><code>8</code> - отказался</dd> <dd><code>9</code> - идут показы</dd> <dd><code>11</code> - Кампания на паузе</dd> </dl> Кампания в процессе удаления. Статус означает, что кампания была удалена, и через 3-10 минут она исчезнет из ответа метода.
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// Дневной бюджет, если не установлен, то 0
	DailyBudget int32 `protobuf:"varint,4,opt,name=dailyBudget,proto3" json:"dailyBudget,omitempty"`
	// Время создания кампании
	CreateTime string `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// Время последнего изменения кампании
	ChangeTime string `protobuf:"bytes,6,opt,name=changeTime,proto3" json:"changeTime,omitempty"`
	// Дата последнего запуска кампании
	StartTime string `protobuf:"bytes,7,opt,name=startTime,proto3" json:"startTime,omitempty"`
	// Дата завершения кампании
	EndTime string `protobuf:"bytes,8,opt,name=endTime,proto3" json:"endTime,omitempty"`
	// Название кампании
	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	// Параметры кампании
	Params []*ResponseInfoAdvertParamsInner `protobuf:"bytes,10,rep,name=params,proto3" json:"params,omitempty"`
	// Активность фиксированных фраз (Для кампаний в поиске)  <br> (`false` - отключены, `true` - включены)
	SearchPluseState bool `protobuf:"varint,11,opt,name=searchPluseState,proto3" json:"searchPluseState,omitempty"`
}

func (x *ResponseInfoAdvert) Reset() {
	*x = ResponseInfoAdvert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_response_info_advert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseInfoAdvert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseInfoAdvert) ProtoMessage() {}

func (x *ResponseInfoAdvert) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_response_info_advert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseInfoAdvert.ProtoReflect.Descriptor instead.
func (*ResponseInfoAdvert) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_response_info_advert_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseInfoAdvert) GetAdvertId() int32 {
	if x != nil {
		return x.AdvertId
	}
	return 0
}

func (x *ResponseInfoAdvert) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ResponseInfoAdvert) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ResponseInfoAdvert) GetDailyBudget() int32 {
	if x != nil {
		return x.DailyBudget
	}
	return 0
}

func (x *ResponseInfoAdvert) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ResponseInfoAdvert) GetChangeTime() string {
	if x != nil {
		return x.ChangeTime
	}
	return ""
}

func (x *ResponseInfoAdvert) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ResponseInfoAdvert) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *ResponseInfoAdvert) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseInfoAdvert) GetParams() []*ResponseInfoAdvertParamsInner {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ResponseInfoAdvert) GetSearchPluseState() bool {
	if x != nil {
		return x.SearchPluseState
	}
	return false
}

var File_wb_promotion_v1_response_info_advert_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_response_info_advert_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x37, 0x77,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x69, 0x6c,
	0x79, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x75, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x75,
	0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x49,
	0x4f, 0x4e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_response_info_advert_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_response_info_advert_proto_rawDescData = file_wb_promotion_v1_response_info_advert_proto_rawDesc
)

func file_wb_promotion_v1_response_info_advert_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_response_info_advert_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_response_info_advert_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_response_info_advert_proto_rawDescData)
	})
	return file_wb_promotion_v1_response_info_advert_proto_rawDescData
}

var file_wb_promotion_v1_response_info_advert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_response_info_advert_proto_goTypes = []interface{}{
	(*ResponseInfoAdvert)(nil),            // 0: wb.promotion.v1.ResponseInfoAdvert
	(*ResponseInfoAdvertParamsInner)(nil), // 1: wb.promotion.v1.ResponseInfoAdvertParamsInner
}
var file_wb_promotion_v1_response_info_advert_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.ResponseInfoAdvert.params:type_name -> wb.promotion.v1.ResponseInfoAdvertParamsInner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_response_info_advert_proto_init() }
func file_wb_promotion_v1_response_info_advert_proto_init() {
	if File_wb_promotion_v1_response_info_advert_proto != nil {
		return
	}
	file_wb_promotion_v1_response_info_advert_params_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_response_info_advert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseInfoAdvert); i {
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
			RawDescriptor: file_wb_promotion_v1_response_info_advert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_response_info_advert_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_response_info_advert_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_response_info_advert_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_response_info_advert_proto = out.File
	file_wb_promotion_v1_response_info_advert_proto_rawDesc = nil
	file_wb_promotion_v1_response_info_advert_proto_goTypes = nil
	file_wb_promotion_v1_response_info_advert_proto_depIdxs = nil
}