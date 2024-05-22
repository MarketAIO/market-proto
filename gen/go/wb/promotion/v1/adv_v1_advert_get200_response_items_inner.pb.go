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
// source: wb/promotion/v1/adv_v1_advert_get200_response_items_inner.proto

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

type AdvV1AdvertGet200ResponseItemsInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID баннера
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Бренд
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Статус (такой же как у медиакампании)
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// Позиция на странице размещения
	Place int32 `protobuf:"varint,4,opt,name=place,proto3" json:"place,omitempty"`
	// Бюджет
	Budget int32 `protobuf:"varint,5,opt,name=budget,proto3" json:"budget,omitempty"`
	// Дневной лимит (для баннеров по показам)
	DailyLimit int32 `protobuf:"varint,6,opt,name=daily_limit,json=dailyLimit,proto3" json:"daily_limit,omitempty"`
	// Название категории размещения
	CategoryName string `protobuf:"bytes,7,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	// Ставка
	Cpm int32 `protobuf:"varint,8,opt,name=cpm,proto3" json:"cpm,omitempty"`
	// URL страницы, на которую попадает пользователь при клике по баннеру
	Url string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	// <dl> <dt>Тип продвижения:</dt> <dd><code>1</code> - баннер</dd> <dd><code>2</code> - всплывающее меню</dd> <dd><code>3</code> - почтовая рассылка</dd> <dd><code>4</code> - социальные сети</dd> <dd><code>5</code> - push-уведомления в мобильном приложении</dd> </dl>
	AdvertType int32 `protobuf:"varint,10,opt,name=advert_type,json=advertType,proto3" json:"advert_type,omitempty"`
	// Дата создания баннера
	CreatedAt string `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Дата обновления баннера
	UpdatedAt string `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Дата начала работы баннера
	DateFrom string `protobuf:"bytes,13,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	// Дата завершения работы баннера
	DateTo string `protobuf:"bytes,14,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
	// Подборка артикулов WB
	Nms []int32 `protobuf:"varint,15,rep,packed,name=nms,proto3" json:"nms,omitempty"`
	// Текст под плашкой баннера
	BottomText1 string `protobuf:"bytes,16,opt,name=bottomText1,proto3" json:"bottomText1,omitempty"`
	// 2-я строка с текстом под плашкой баннера
	BottomText2 string `protobuf:"bytes,17,opt,name=bottomText2,proto3" json:"bottomText2,omitempty"`
	// Текст push-уведомления или рассылки
	Message string `protobuf:"bytes,18,opt,name=message,proto3" json:"message,omitempty"`
	// Дополнительные настройки. <dl> <dt>Формат почтовой рассылки:</dt> <dd><code>1</code> - общий</dd> <dd><code>2</code> - частичный</dd> <dd><code>3</code> - уникальный</dd> </dl> <dl> <dt>Социальная сеть:</dt> <dd><code>1</code> - VKontakte</dd> <dd><code>2</code> - OK (Одноклассники)</dd> </dl>
	AdditionalSettings int32 `protobuf:"varint,19,opt,name=additionalSettings,proto3" json:"additionalSettings,omitempty"`
	// Кол-во получателей push-уведомлений
	ReceiversCount int32 `protobuf:"varint,20,opt,name=receiversCount,proto3" json:"receiversCount,omitempty"`
	// ID родительской категории товара
	SubjectId int32 `protobuf:"varint,21,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// Название родительской категории товара
	SubjectName string `protobuf:"bytes,22,opt,name=subject_name,json=subjectName,proto3" json:"subject_name,omitempty"`
	// Название акции
	ActionName string `protobuf:"bytes,23,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	// Часы показа
	ShowHours []*AdvV1AdvertGet200ResponseItemsInnerShowHoursInner `protobuf:"bytes,24,rep,name=show_hours,json=showHours,proto3" json:"show_hours,omitempty"`
	// Уникальный идентификатор медиакампании для работы с ОРД
	Erid string `protobuf:"bytes,25,opt,name=Erid,proto3" json:"Erid,omitempty"`
}

func (x *AdvV1AdvertGet200ResponseItemsInner) Reset() {
	*x = AdvV1AdvertGet200ResponseItemsInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvV1AdvertGet200ResponseItemsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvV1AdvertGet200ResponseItemsInner) ProtoMessage() {}

func (x *AdvV1AdvertGet200ResponseItemsInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvV1AdvertGet200ResponseItemsInner.ProtoReflect.Descriptor instead.
func (*AdvV1AdvertGet200ResponseItemsInner) Descriptor() ([]byte, []int) {
	return file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescGZIP(), []int{0}
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetPlace() int32 {
	if x != nil {
		return x.Place
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetBudget() int32 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetDailyLimit() int32 {
	if x != nil {
		return x.DailyLimit
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetCpm() int32 {
	if x != nil {
		return x.Cpm
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetAdvertType() int32 {
	if x != nil {
		return x.AdvertType
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetNms() []int32 {
	if x != nil {
		return x.Nms
	}
	return nil
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetBottomText1() string {
	if x != nil {
		return x.BottomText1
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetBottomText2() string {
	if x != nil {
		return x.BottomText2
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetAdditionalSettings() int32 {
	if x != nil {
		return x.AdditionalSettings
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetReceiversCount() int32 {
	if x != nil {
		return x.ReceiversCount
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetSubjectId() int32 {
	if x != nil {
		return x.SubjectId
	}
	return 0
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetSubjectName() string {
	if x != nil {
		return x.SubjectName
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetShowHours() []*AdvV1AdvertGet200ResponseItemsInnerShowHoursInner {
	if x != nil {
		return x.ShowHours
	}
	return nil
}

func (x *AdvV1AdvertGet200ResponseItemsInner) GetErid() string {
	if x != nil {
		return x.Erid
	}
	return ""
}

var File_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto protoreflect.FileDescriptor

var file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f,
	0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x50, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x76, 0x5f, 0x76, 0x31, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x68,
	0x6f, 0x77, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x06, 0x0a, 0x23, 0x41, 0x64, 0x76, 0x56, 0x31, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x70, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x70, 0x6d, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6d, 0x73, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x03, 0x6e, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f,
	0x6d, 0x54, 0x65, 0x78, 0x74, 0x31, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f,
	0x74, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x31, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6f, 0x74,
	0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x32, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x61, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x18,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x76, 0x56, 0x31, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x77, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x45, 0x72, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x45, 0x72, 0x69, 0x64, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49, 0x4f, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescOnce sync.Once
	file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescData = file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDesc
)

func file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescGZIP() []byte {
	file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescOnce.Do(func() {
		file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescData)
	})
	return file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDescData
}

var file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_goTypes = []interface{}{
	(*AdvV1AdvertGet200ResponseItemsInner)(nil),               // 0: wb.promotion.v1.AdvV1AdvertGet200ResponseItemsInner
	(*AdvV1AdvertGet200ResponseItemsInnerShowHoursInner)(nil), // 1: wb.promotion.v1.AdvV1AdvertGet200ResponseItemsInnerShowHoursInner
}
var file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_depIdxs = []int32{
	1, // 0: wb.promotion.v1.AdvV1AdvertGet200ResponseItemsInner.show_hours:type_name -> wb.promotion.v1.AdvV1AdvertGet200ResponseItemsInnerShowHoursInner
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_init() }
func file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_init() {
	if File_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto != nil {
		return
	}
	file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_show_hours_inner_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvV1AdvertGet200ResponseItemsInner); i {
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
			RawDescriptor: file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_goTypes,
		DependencyIndexes: file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_depIdxs,
		MessageInfos:      file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_msgTypes,
	}.Build()
	File_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto = out.File
	file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_rawDesc = nil
	file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_goTypes = nil
	file_wb_promotion_v1_adv_v1_advert_get200_response_items_inner_proto_depIdxs = nil
}
