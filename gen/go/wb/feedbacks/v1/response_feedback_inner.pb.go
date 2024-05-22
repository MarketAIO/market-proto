//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд.
//
//The version of the OpenAPI document:
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: wb/feedbacks/v1/response_feedback_inner.proto

package wbFeedbacks

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

type ResponseFeedbackInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id отзыва
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Имя автора отзыва
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// Соответствие заявленного размера реальному. <br>Возможные значения: - ` ` - для безразмерных товаров - `ок` - соответствует размеру - `smaller` - маломерит - `bigger` - большемерит
	MatchingSize string `protobuf:"bytes,3,opt,name=matchingSize,proto3" json:"matchingSize,omitempty"`
	// Текст отзыва
	Text string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	// Оценка товара от покупателя
	ProductValuation int32 `protobuf:"varint,5,opt,name=productValuation,proto3" json:"productValuation,omitempty"`
	// Дата и время создания отзыва
	CreatedDate string `protobuf:"bytes,6,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	// Статус отзыва:   - `none` - не обработан (новый)   - `wbRu` - обработан
	State          string                                         `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Answer         *ResponseFeedbackInnerAnswer                   `protobuf:"bytes,8,opt,name=answer,proto3" json:"answer,omitempty"`
	ProductDetails *ApiV1FeedbackGet200ResponseDataProductDetails `protobuf:"bytes,9,opt,name=productDetails,proto3" json:"productDetails,omitempty"`
	// Массив структур фотографий
	PhotoLinks []*ApiV1FeedbackGet200ResponseDataPhotoLinksInner `protobuf:"bytes,10,rep,name=photoLinks,proto3" json:"photoLinks,omitempty"`
	Video      *ApiV1FeedbackGet200ResponseDataVideo             `protobuf:"bytes,11,opt,name=video,proto3" json:"video,omitempty"`
	// Просмотрен ли отзыв
	WasViewed bool `protobuf:"varint,12,opt,name=wasViewed,proto3" json:"wasViewed,omitempty"`
	// Доступна ли продавцу оценка отзыва (`true` - доступна, `false` - не доступна)
	IsAbleSupplierFeedbackValuation bool `protobuf:"varint,13,opt,name=isAbleSupplierFeedbackValuation,proto3" json:"isAbleSupplierFeedbackValuation,omitempty"`
	// Оценка отзыва, оставленная продавцом
	SupplierFeedbackValuation int32 `protobuf:"varint,14,opt,name=supplierFeedbackValuation,proto3" json:"supplierFeedbackValuation,omitempty"`
	// Доступна ли продавцу оценка товара (`true` - доступна, `false` - не доступна)
	IsAbleSupplierProductValuation bool `protobuf:"varint,15,opt,name=isAbleSupplierProductValuation,proto3" json:"isAbleSupplierProductValuation,omitempty"`
	// Оценка товара, оставленная продавцом
	SupplierProductValuation int32 `protobuf:"varint,16,opt,name=supplierProductValuation,proto3" json:"supplierProductValuation,omitempty"`
	// Доступна ли товару опция возврата (`false` - нет, `true` - да)
	IsAbleReturnProductOrders bool `protobuf:"varint,17,opt,name=isAbleReturnProductOrders,proto3" json:"isAbleReturnProductOrders,omitempty"`
	// Дата и время, когда на запрос возврата был получен ответ со статус-кодом 200.
	ReturnProductOrdersDate string `protobuf:"bytes,18,opt,name=returnProductOrdersDate,proto3" json:"returnProductOrdersDate,omitempty"`
	// Список тегов покупателя
	Bables []string `protobuf:"bytes,19,rep,name=bables,proto3" json:"bables,omitempty"`
}

func (x *ResponseFeedbackInner) Reset() {
	*x = ResponseFeedbackInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_feedbacks_v1_response_feedback_inner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFeedbackInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFeedbackInner) ProtoMessage() {}

func (x *ResponseFeedbackInner) ProtoReflect() protoreflect.Message {
	mi := &file_wb_feedbacks_v1_response_feedback_inner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFeedbackInner.ProtoReflect.Descriptor instead.
func (*ResponseFeedbackInner) Descriptor() ([]byte, []int) {
	return file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseFeedbackInner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseFeedbackInner) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ResponseFeedbackInner) GetMatchingSize() string {
	if x != nil {
		return x.MatchingSize
	}
	return ""
}

func (x *ResponseFeedbackInner) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ResponseFeedbackInner) GetProductValuation() int32 {
	if x != nil {
		return x.ProductValuation
	}
	return 0
}

func (x *ResponseFeedbackInner) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *ResponseFeedbackInner) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ResponseFeedbackInner) GetAnswer() *ResponseFeedbackInnerAnswer {
	if x != nil {
		return x.Answer
	}
	return nil
}

func (x *ResponseFeedbackInner) GetProductDetails() *ApiV1FeedbackGet200ResponseDataProductDetails {
	if x != nil {
		return x.ProductDetails
	}
	return nil
}

func (x *ResponseFeedbackInner) GetPhotoLinks() []*ApiV1FeedbackGet200ResponseDataPhotoLinksInner {
	if x != nil {
		return x.PhotoLinks
	}
	return nil
}

func (x *ResponseFeedbackInner) GetVideo() *ApiV1FeedbackGet200ResponseDataVideo {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *ResponseFeedbackInner) GetWasViewed() bool {
	if x != nil {
		return x.WasViewed
	}
	return false
}

func (x *ResponseFeedbackInner) GetIsAbleSupplierFeedbackValuation() bool {
	if x != nil {
		return x.IsAbleSupplierFeedbackValuation
	}
	return false
}

func (x *ResponseFeedbackInner) GetSupplierFeedbackValuation() int32 {
	if x != nil {
		return x.SupplierFeedbackValuation
	}
	return 0
}

func (x *ResponseFeedbackInner) GetIsAbleSupplierProductValuation() bool {
	if x != nil {
		return x.IsAbleSupplierProductValuation
	}
	return false
}

func (x *ResponseFeedbackInner) GetSupplierProductValuation() int32 {
	if x != nil {
		return x.SupplierProductValuation
	}
	return 0
}

func (x *ResponseFeedbackInner) GetIsAbleReturnProductOrders() bool {
	if x != nil {
		return x.IsAbleReturnProductOrders
	}
	return false
}

func (x *ResponseFeedbackInner) GetReturnProductOrdersDate() string {
	if x != nil {
		return x.ReturnProductOrdersDate
	}
	return ""
}

func (x *ResponseFeedbackInner) GetBables() []string {
	if x != nil {
		return x.Bables
	}
	return nil
}

var File_wb_feedbacks_v1_response_feedback_inner_proto protoreflect.FileDescriptor

var file_wb_feedbacks_v1_response_feedback_inner_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x77, 0x62, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x77, 0x62, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x4c, 0x77, 0x62, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x5f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a,
	0x77, 0x62, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x67, 0x65, 0x74, 0x32, 0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x77, 0x62, 0x2f, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x31, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x32,
	0x30, 0x30, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x77, 0x62,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf5, 0x07, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x77, 0x62, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x77, 0x62, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x31, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x5f, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x77, 0x62, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x62, 0x61, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x31, 0x46, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x4b, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x77, 0x62, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x31, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x47, 0x65, 0x74, 0x32, 0x30, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x77, 0x61, 0x73, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x77, 0x61, 0x73, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x12, 0x48,
	0x0a, 0x1f, 0x69, 0x73, 0x41, 0x62, 0x6c, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x69, 0x73, 0x41, 0x62, 0x6c, 0x65, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x19, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x56, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x1e, 0x69, 0x73, 0x41, 0x62, 0x6c, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e,
	0x69, 0x73, 0x41, 0x62, 0x6c, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x18, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x18, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x19, 0x69, 0x73,
	0x41, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x69,
	0x73, 0x41, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x13, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41,
	0x49, 0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x62, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63,
	0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescOnce sync.Once
	file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescData = file_wb_feedbacks_v1_response_feedback_inner_proto_rawDesc
)

func file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescGZIP() []byte {
	file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescOnce.Do(func() {
		file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescData)
	})
	return file_wb_feedbacks_v1_response_feedback_inner_proto_rawDescData
}

var file_wb_feedbacks_v1_response_feedback_inner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_feedbacks_v1_response_feedback_inner_proto_goTypes = []interface{}{
	(*ResponseFeedbackInner)(nil),                          // 0: wb.feedbacks.v1.ResponseFeedbackInner
	(*ResponseFeedbackInnerAnswer)(nil),                    // 1: wb.feedbacks.v1.ResponseFeedbackInnerAnswer
	(*ApiV1FeedbackGet200ResponseDataProductDetails)(nil),  // 2: wb.feedbacks.v1.ApiV1FeedbackGet200ResponseDataProductDetails
	(*ApiV1FeedbackGet200ResponseDataPhotoLinksInner)(nil), // 3: wb.feedbacks.v1.ApiV1FeedbackGet200ResponseDataPhotoLinksInner
	(*ApiV1FeedbackGet200ResponseDataVideo)(nil),           // 4: wb.feedbacks.v1.ApiV1FeedbackGet200ResponseDataVideo
}
var file_wb_feedbacks_v1_response_feedback_inner_proto_depIdxs = []int32{
	1, // 0: wb.feedbacks.v1.ResponseFeedbackInner.answer:type_name -> wb.feedbacks.v1.ResponseFeedbackInnerAnswer
	2, // 1: wb.feedbacks.v1.ResponseFeedbackInner.productDetails:type_name -> wb.feedbacks.v1.ApiV1FeedbackGet200ResponseDataProductDetails
	3, // 2: wb.feedbacks.v1.ResponseFeedbackInner.photoLinks:type_name -> wb.feedbacks.v1.ApiV1FeedbackGet200ResponseDataPhotoLinksInner
	4, // 3: wb.feedbacks.v1.ResponseFeedbackInner.video:type_name -> wb.feedbacks.v1.ApiV1FeedbackGet200ResponseDataVideo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wb_feedbacks_v1_response_feedback_inner_proto_init() }
func file_wb_feedbacks_v1_response_feedback_inner_proto_init() {
	if File_wb_feedbacks_v1_response_feedback_inner_proto != nil {
		return
	}
	file_wb_feedbacks_v1_api_v1_feedback_get200_response_data_photo_links_inner_proto_init()
	file_wb_feedbacks_v1_api_v1_feedback_get200_response_data_product_details_proto_init()
	file_wb_feedbacks_v1_api_v1_feedback_get200_response_data_video_proto_init()
	file_wb_feedbacks_v1_response_feedback_inner_answer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_feedbacks_v1_response_feedback_inner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFeedbackInner); i {
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
			RawDescriptor: file_wb_feedbacks_v1_response_feedback_inner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_feedbacks_v1_response_feedback_inner_proto_goTypes,
		DependencyIndexes: file_wb_feedbacks_v1_response_feedback_inner_proto_depIdxs,
		MessageInfos:      file_wb_feedbacks_v1_response_feedback_inner_proto_msgTypes,
	}.Build()
	File_wb_feedbacks_v1_response_feedback_inner_proto = out.File
	file_wb_feedbacks_v1_response_feedback_inner_proto_rawDesc = nil
	file_wb_feedbacks_v1_response_feedback_inner_proto_goTypes = nil
	file_wb_feedbacks_v1_response_feedback_inner_proto_depIdxs = nil
}
