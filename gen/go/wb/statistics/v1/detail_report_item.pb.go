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
// source: wb/statistics/v1/detail_report_item.proto

package wbSTATISTICS

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DetailReportItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Номер отчёта
	RealizationreportId int32 `protobuf:"varint,1,opt,name=realizationreport_id,json=realizationreportId,proto3" json:"realizationreport_id,omitempty"`
	// Дата начала отчётного периода
	DateFrom string `protobuf:"bytes,2,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	// Дата конца отчётного периода
	DateTo string `protobuf:"bytes,3,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
	// Дата формирования отчёта
	CreateDt string `protobuf:"bytes,4,opt,name=create_dt,json=createDt,proto3" json:"create_dt,omitempty"`
	// Валюта отчёта
	CurrencyName string `protobuf:"bytes,5,opt,name=currency_name,json=currencyName,proto3" json:"currency_name,omitempty"`
	// Договор
	SuppliercontractCode *anypb.Any `protobuf:"bytes,6,opt,name=suppliercontract_code,json=suppliercontractCode,proto3" json:"suppliercontract_code,omitempty"`
	// Номер строки
	RrdId int32 `protobuf:"varint,7,opt,name=rrd_id,json=rrdId,proto3" json:"rrd_id,omitempty"`
	// Номер поставки
	GiId int32 `protobuf:"varint,8,opt,name=gi_id,json=giId,proto3" json:"gi_id,omitempty"`
	// Предмет
	SubjectName string `protobuf:"bytes,9,opt,name=subject_name,json=subjectName,proto3" json:"subject_name,omitempty"`
	// Артикул WB
	NmId int32 `protobuf:"varint,10,opt,name=nm_id,json=nmId,proto3" json:"nm_id,omitempty"`
	// Бренд
	BrandName string `protobuf:"bytes,11,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	// Артикул продавца
	SaName string `protobuf:"bytes,12,opt,name=sa_name,json=saName,proto3" json:"sa_name,omitempty"`
	// Размер
	TsName string `protobuf:"bytes,13,opt,name=ts_name,json=tsName,proto3" json:"ts_name,omitempty"`
	// Баркод
	Barcode string `protobuf:"bytes,14,opt,name=barcode,proto3" json:"barcode,omitempty"`
	// Тип документа
	DocTypeName string `protobuf:"bytes,15,opt,name=doc_type_name,json=docTypeName,proto3" json:"doc_type_name,omitempty"`
	// Количество
	Quantity int32 `protobuf:"varint,16,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// Цена розничная
	RetailPrice float32 `protobuf:"fixed32,17,opt,name=retail_price,json=retailPrice,proto3" json:"retail_price,omitempty"`
	// Сумма продаж (возвратов)
	RetailAmount float32 `protobuf:"fixed32,18,opt,name=retail_amount,json=retailAmount,proto3" json:"retail_amount,omitempty"`
	// Согласованная скидка
	SalePercent int32 `protobuf:"varint,19,opt,name=sale_percent,json=salePercent,proto3" json:"sale_percent,omitempty"`
	// Процент комиссии
	CommissionPercent float32 `protobuf:"fixed32,20,opt,name=commission_percent,json=commissionPercent,proto3" json:"commission_percent,omitempty"`
	// Склад
	OfficeName string `protobuf:"bytes,21,opt,name=office_name,json=officeName,proto3" json:"office_name,omitempty"`
	// Обоснование для оплаты
	SupplierOperName string `protobuf:"bytes,22,opt,name=supplier_oper_name,json=supplierOperName,proto3" json:"supplier_oper_name,omitempty"`
	// Дата заказа. <br>Присылается с явным указанием часового пояса
	OrderDt string `protobuf:"bytes,23,opt,name=order_dt,json=orderDt,proto3" json:"order_dt,omitempty"`
	// Дата продажи. <br>Присылается с явным указанием часового пояса
	SaleDt string `protobuf:"bytes,24,opt,name=sale_dt,json=saleDt,proto3" json:"sale_dt,omitempty"`
	// Дата операции. <br> Присылается с явным указанием часового пояса
	RrDt string `protobuf:"bytes,25,opt,name=rr_dt,json=rrDt,proto3" json:"rr_dt,omitempty"`
	// Штрих-код
	ShkId int32 `protobuf:"varint,26,opt,name=shk_id,json=shkId,proto3" json:"shk_id,omitempty"`
	// Цена розничная с учетом согласованной скидки
	RetailPriceWithdiscRub float32 `protobuf:"fixed32,27,opt,name=retail_price_withdisc_rub,json=retailPriceWithdiscRub,proto3" json:"retail_price_withdisc_rub,omitempty"`
	// Количество доставок
	DeliveryAmount int32 `protobuf:"varint,28,opt,name=delivery_amount,json=deliveryAmount,proto3" json:"delivery_amount,omitempty"`
	// Количество возвратов
	ReturnAmount int32 `protobuf:"varint,29,opt,name=return_amount,json=returnAmount,proto3" json:"return_amount,omitempty"`
	// Стоимость логистики
	DeliveryRub float32 `protobuf:"fixed32,30,opt,name=delivery_rub,json=deliveryRub,proto3" json:"delivery_rub,omitempty"`
	// Тип коробов
	GiBoxTypeName string `protobuf:"bytes,31,opt,name=gi_box_type_name,json=giBoxTypeName,proto3" json:"gi_box_type_name,omitempty"`
	// Согласованный продуктовый дисконт
	ProductDiscountForReport float32 `protobuf:"fixed32,32,opt,name=product_discount_for_report,json=productDiscountForReport,proto3" json:"product_discount_for_report,omitempty"`
	// Промокод
	SupplierPromo float32 `protobuf:"fixed32,33,opt,name=supplier_promo,json=supplierPromo,proto3" json:"supplier_promo,omitempty"`
	// Уникальный идентификатор заказа
	Rid int32 `protobuf:"varint,34,opt,name=rid,proto3" json:"rid,omitempty"`
	// Скидка постоянного покупателя
	PpvzSppPrc float32 `protobuf:"fixed32,35,opt,name=ppvz_spp_prc,json=ppvzSppPrc,proto3" json:"ppvz_spp_prc,omitempty"`
	// Размер кВВ без НДС, % базовый
	PpvzKvwPrcBase float32 `protobuf:"fixed32,36,opt,name=ppvz_kvw_prc_base,json=ppvzKvwPrcBase,proto3" json:"ppvz_kvw_prc_base,omitempty"`
	// Итоговый кВВ без НДС, %
	PpvzKvwPrc float32 `protobuf:"fixed32,37,opt,name=ppvz_kvw_prc,json=ppvzKvwPrc,proto3" json:"ppvz_kvw_prc,omitempty"`
	// Размер снижения кВВ из-за рейтинга
	SupRatingPrcUp float32 `protobuf:"fixed32,38,opt,name=sup_rating_prc_up,json=supRatingPrcUp,proto3" json:"sup_rating_prc_up,omitempty"`
	// Размер снижения кВВ из-за акции
	IsKgvpV2 float32 `protobuf:"fixed32,39,opt,name=is_kgvp_v2,json=isKgvpV2,proto3" json:"is_kgvp_v2,omitempty"`
	// Вознаграждение с продаж до вычета услуг поверенного, без НДС
	PpvzSalesCommission float32 `protobuf:"fixed32,40,opt,name=ppvz_sales_commission,json=ppvzSalesCommission,proto3" json:"ppvz_sales_commission,omitempty"`
	// К перечислению продавцу за реализованный товар
	PpvzForPay float32 `protobuf:"fixed32,41,opt,name=ppvz_for_pay,json=ppvzForPay,proto3" json:"ppvz_for_pay,omitempty"`
	// Возмещение за выдачу и возврат товаров на ПВЗ
	PpvzReward float32 `protobuf:"fixed32,42,opt,name=ppvz_reward,json=ppvzReward,proto3" json:"ppvz_reward,omitempty"`
	// Возмещение издержек по эквайрингу. <br> Издержки WB за услуги эквайринга: вычитаются из вознаграждения WB и не влияют на доход продавца.
	AcquiringFee float32 `protobuf:"fixed32,43,opt,name=acquiring_fee,json=acquiringFee,proto3" json:"acquiring_fee,omitempty"`
	// Наименование банка-эквайера
	AcquiringBank string `protobuf:"bytes,44,opt,name=acquiring_bank,json=acquiringBank,proto3" json:"acquiring_bank,omitempty"`
	// Вознаграждение WB без НДС
	PpvzVw float32 `protobuf:"fixed32,45,opt,name=ppvz_vw,json=ppvzVw,proto3" json:"ppvz_vw,omitempty"`
	// НДС с вознаграждения WB
	PpvzVwNds float32 `protobuf:"fixed32,46,opt,name=ppvz_vw_nds,json=ppvzVwNds,proto3" json:"ppvz_vw_nds,omitempty"`
	// Номер офиса
	PpvzOfficeId int32 `protobuf:"varint,47,opt,name=ppvz_office_id,json=ppvzOfficeId,proto3" json:"ppvz_office_id,omitempty"`
	// Наименование офиса доставки
	PpvzOfficeName string `protobuf:"bytes,48,opt,name=ppvz_office_name,json=ppvzOfficeName,proto3" json:"ppvz_office_name,omitempty"`
	// Номер партнера
	PpvzSupplierId int32 `protobuf:"varint,49,opt,name=ppvz_supplier_id,json=ppvzSupplierId,proto3" json:"ppvz_supplier_id,omitempty"`
	// Партнер
	PpvzSupplierName string `protobuf:"bytes,50,opt,name=ppvz_supplier_name,json=ppvzSupplierName,proto3" json:"ppvz_supplier_name,omitempty"`
	// ИНН партнера
	PpvzInn string `protobuf:"bytes,51,opt,name=ppvz_inn,json=ppvzInn,proto3" json:"ppvz_inn,omitempty"`
	// Номер таможенной декларации
	DeclarationNumber string `protobuf:"bytes,52,opt,name=declaration_number,json=declarationNumber,proto3" json:"declaration_number,omitempty"`
	// Обоснование штрафов и доплат. <br>Поле будет в ответе при наличии значения
	BonusTypeName string `protobuf:"bytes,53,opt,name=bonus_type_name,json=bonusTypeName,proto3" json:"bonus_type_name,omitempty"`
	// Цифровое значение стикера, который клеится на товар в процессе сборки заказа по схеме \"Маркетплейс\"
	StickerId string `protobuf:"bytes,54,opt,name=sticker_id,json=stickerId,proto3" json:"sticker_id,omitempty"`
	// Страна продажи
	SiteCountry string `protobuf:"bytes,55,opt,name=site_country,json=siteCountry,proto3" json:"site_country,omitempty"`
	// Штрафы
	Penalty float32 `protobuf:"fixed32,56,opt,name=penalty,proto3" json:"penalty,omitempty"`
	// Доплаты
	AdditionalPayment float32 `protobuf:"fixed32,57,opt,name=additional_payment,json=additionalPayment,proto3" json:"additional_payment,omitempty"`
	// Возмещение издержек по перевозке. Поле будет в ответе при наличии значения
	RebillLogisticCost float32 `protobuf:"fixed32,58,opt,name=rebill_logistic_cost,json=rebillLogisticCost,proto3" json:"rebill_logistic_cost,omitempty"`
	// Организатор перевозки. Поле будет в ответе при наличии значения
	RebillLogisticOrg string `protobuf:"bytes,59,opt,name=rebill_logistic_org,json=rebillLogisticOrg,proto3" json:"rebill_logistic_org,omitempty"`
	// Код маркировки. <br> Поле будет в ответе при наличии значения
	Kiz string `protobuf:"bytes,60,opt,name=kiz,proto3" json:"kiz,omitempty"`
	// Стоимость хранения
	StorageFee float32 `protobuf:"fixed32,61,opt,name=storage_fee,json=storageFee,proto3" json:"storage_fee,omitempty"`
	// Прочие удержания/выплаты
	Deduction float32 `protobuf:"fixed32,62,opt,name=deduction,proto3" json:"deduction,omitempty"`
	// Стоимость платной приёмки
	Acceptance float32 `protobuf:"fixed32,63,opt,name=acceptance,proto3" json:"acceptance,omitempty"`
	// Уникальный идентификатор заказа.  Примечание для использующих API Marketplace: `srid` равен `rid` в ответах методов сборочных заданий.
	Srid string `protobuf:"bytes,64,opt,name=srid,proto3" json:"srid,omitempty"`
	// Тип отчёта:    * `1` — стандартный   * `2` — для уведомления о выкупе
	ReportType int32 `protobuf:"varint,65,opt,name=report_type,json=reportType,proto3" json:"report_type,omitempty"`
}

func (x *DetailReportItem) Reset() {
	*x = DetailReportItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_statistics_v1_detail_report_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailReportItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailReportItem) ProtoMessage() {}

func (x *DetailReportItem) ProtoReflect() protoreflect.Message {
	mi := &file_wb_statistics_v1_detail_report_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailReportItem.ProtoReflect.Descriptor instead.
func (*DetailReportItem) Descriptor() ([]byte, []int) {
	return file_wb_statistics_v1_detail_report_item_proto_rawDescGZIP(), []int{0}
}

func (x *DetailReportItem) GetRealizationreportId() int32 {
	if x != nil {
		return x.RealizationreportId
	}
	return 0
}

func (x *DetailReportItem) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *DetailReportItem) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *DetailReportItem) GetCreateDt() string {
	if x != nil {
		return x.CreateDt
	}
	return ""
}

func (x *DetailReportItem) GetCurrencyName() string {
	if x != nil {
		return x.CurrencyName
	}
	return ""
}

func (x *DetailReportItem) GetSuppliercontractCode() *anypb.Any {
	if x != nil {
		return x.SuppliercontractCode
	}
	return nil
}

func (x *DetailReportItem) GetRrdId() int32 {
	if x != nil {
		return x.RrdId
	}
	return 0
}

func (x *DetailReportItem) GetGiId() int32 {
	if x != nil {
		return x.GiId
	}
	return 0
}

func (x *DetailReportItem) GetSubjectName() string {
	if x != nil {
		return x.SubjectName
	}
	return ""
}

func (x *DetailReportItem) GetNmId() int32 {
	if x != nil {
		return x.NmId
	}
	return 0
}

func (x *DetailReportItem) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *DetailReportItem) GetSaName() string {
	if x != nil {
		return x.SaName
	}
	return ""
}

func (x *DetailReportItem) GetTsName() string {
	if x != nil {
		return x.TsName
	}
	return ""
}

func (x *DetailReportItem) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *DetailReportItem) GetDocTypeName() string {
	if x != nil {
		return x.DocTypeName
	}
	return ""
}

func (x *DetailReportItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *DetailReportItem) GetRetailPrice() float32 {
	if x != nil {
		return x.RetailPrice
	}
	return 0
}

func (x *DetailReportItem) GetRetailAmount() float32 {
	if x != nil {
		return x.RetailAmount
	}
	return 0
}

func (x *DetailReportItem) GetSalePercent() int32 {
	if x != nil {
		return x.SalePercent
	}
	return 0
}

func (x *DetailReportItem) GetCommissionPercent() float32 {
	if x != nil {
		return x.CommissionPercent
	}
	return 0
}

func (x *DetailReportItem) GetOfficeName() string {
	if x != nil {
		return x.OfficeName
	}
	return ""
}

func (x *DetailReportItem) GetSupplierOperName() string {
	if x != nil {
		return x.SupplierOperName
	}
	return ""
}

func (x *DetailReportItem) GetOrderDt() string {
	if x != nil {
		return x.OrderDt
	}
	return ""
}

func (x *DetailReportItem) GetSaleDt() string {
	if x != nil {
		return x.SaleDt
	}
	return ""
}

func (x *DetailReportItem) GetRrDt() string {
	if x != nil {
		return x.RrDt
	}
	return ""
}

func (x *DetailReportItem) GetShkId() int32 {
	if x != nil {
		return x.ShkId
	}
	return 0
}

func (x *DetailReportItem) GetRetailPriceWithdiscRub() float32 {
	if x != nil {
		return x.RetailPriceWithdiscRub
	}
	return 0
}

func (x *DetailReportItem) GetDeliveryAmount() int32 {
	if x != nil {
		return x.DeliveryAmount
	}
	return 0
}

func (x *DetailReportItem) GetReturnAmount() int32 {
	if x != nil {
		return x.ReturnAmount
	}
	return 0
}

func (x *DetailReportItem) GetDeliveryRub() float32 {
	if x != nil {
		return x.DeliveryRub
	}
	return 0
}

func (x *DetailReportItem) GetGiBoxTypeName() string {
	if x != nil {
		return x.GiBoxTypeName
	}
	return ""
}

func (x *DetailReportItem) GetProductDiscountForReport() float32 {
	if x != nil {
		return x.ProductDiscountForReport
	}
	return 0
}

func (x *DetailReportItem) GetSupplierPromo() float32 {
	if x != nil {
		return x.SupplierPromo
	}
	return 0
}

func (x *DetailReportItem) GetRid() int32 {
	if x != nil {
		return x.Rid
	}
	return 0
}

func (x *DetailReportItem) GetPpvzSppPrc() float32 {
	if x != nil {
		return x.PpvzSppPrc
	}
	return 0
}

func (x *DetailReportItem) GetPpvzKvwPrcBase() float32 {
	if x != nil {
		return x.PpvzKvwPrcBase
	}
	return 0
}

func (x *DetailReportItem) GetPpvzKvwPrc() float32 {
	if x != nil {
		return x.PpvzKvwPrc
	}
	return 0
}

func (x *DetailReportItem) GetSupRatingPrcUp() float32 {
	if x != nil {
		return x.SupRatingPrcUp
	}
	return 0
}

func (x *DetailReportItem) GetIsKgvpV2() float32 {
	if x != nil {
		return x.IsKgvpV2
	}
	return 0
}

func (x *DetailReportItem) GetPpvzSalesCommission() float32 {
	if x != nil {
		return x.PpvzSalesCommission
	}
	return 0
}

func (x *DetailReportItem) GetPpvzForPay() float32 {
	if x != nil {
		return x.PpvzForPay
	}
	return 0
}

func (x *DetailReportItem) GetPpvzReward() float32 {
	if x != nil {
		return x.PpvzReward
	}
	return 0
}

func (x *DetailReportItem) GetAcquiringFee() float32 {
	if x != nil {
		return x.AcquiringFee
	}
	return 0
}

func (x *DetailReportItem) GetAcquiringBank() string {
	if x != nil {
		return x.AcquiringBank
	}
	return ""
}

func (x *DetailReportItem) GetPpvzVw() float32 {
	if x != nil {
		return x.PpvzVw
	}
	return 0
}

func (x *DetailReportItem) GetPpvzVwNds() float32 {
	if x != nil {
		return x.PpvzVwNds
	}
	return 0
}

func (x *DetailReportItem) GetPpvzOfficeId() int32 {
	if x != nil {
		return x.PpvzOfficeId
	}
	return 0
}

func (x *DetailReportItem) GetPpvzOfficeName() string {
	if x != nil {
		return x.PpvzOfficeName
	}
	return ""
}

func (x *DetailReportItem) GetPpvzSupplierId() int32 {
	if x != nil {
		return x.PpvzSupplierId
	}
	return 0
}

func (x *DetailReportItem) GetPpvzSupplierName() string {
	if x != nil {
		return x.PpvzSupplierName
	}
	return ""
}

func (x *DetailReportItem) GetPpvzInn() string {
	if x != nil {
		return x.PpvzInn
	}
	return ""
}

func (x *DetailReportItem) GetDeclarationNumber() string {
	if x != nil {
		return x.DeclarationNumber
	}
	return ""
}

func (x *DetailReportItem) GetBonusTypeName() string {
	if x != nil {
		return x.BonusTypeName
	}
	return ""
}

func (x *DetailReportItem) GetStickerId() string {
	if x != nil {
		return x.StickerId
	}
	return ""
}

func (x *DetailReportItem) GetSiteCountry() string {
	if x != nil {
		return x.SiteCountry
	}
	return ""
}

func (x *DetailReportItem) GetPenalty() float32 {
	if x != nil {
		return x.Penalty
	}
	return 0
}

func (x *DetailReportItem) GetAdditionalPayment() float32 {
	if x != nil {
		return x.AdditionalPayment
	}
	return 0
}

func (x *DetailReportItem) GetRebillLogisticCost() float32 {
	if x != nil {
		return x.RebillLogisticCost
	}
	return 0
}

func (x *DetailReportItem) GetRebillLogisticOrg() string {
	if x != nil {
		return x.RebillLogisticOrg
	}
	return ""
}

func (x *DetailReportItem) GetKiz() string {
	if x != nil {
		return x.Kiz
	}
	return ""
}

func (x *DetailReportItem) GetStorageFee() float32 {
	if x != nil {
		return x.StorageFee
	}
	return 0
}

func (x *DetailReportItem) GetDeduction() float32 {
	if x != nil {
		return x.Deduction
	}
	return 0
}

func (x *DetailReportItem) GetAcceptance() float32 {
	if x != nil {
		return x.Acceptance
	}
	return 0
}

func (x *DetailReportItem) GetSrid() string {
	if x != nil {
		return x.Srid
	}
	return ""
}

func (x *DetailReportItem) GetReportType() int32 {
	if x != nil {
		return x.ReportType
	}
	return 0
}

var File_wb_statistics_v1_detail_report_item_proto protoreflect.FileDescriptor

var file_wb_statistics_v1_detail_report_item_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x62, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x31, 0x0a,
	0x14, 0x72, 0x65, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x72, 0x65, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x64, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x15, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x14, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x72, 0x64, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x67, 0x69,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x67, 0x69, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6e, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x61,
	0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x64,
	0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x44, 0x74, 0x12,
	0x13, 0x0a, 0x05, 0x72, 0x72, 0x5f, 0x64, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x72, 0x44, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x68, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x68, 0x6b, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x69, 0x73, 0x63, 0x5f, 0x72, 0x75, 0x62, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x16,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x69, 0x73, 0x63, 0x52, 0x75, 0x62, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x72, 0x75, 0x62, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x75, 0x62, 0x12, 0x27, 0x0a, 0x10, 0x67, 0x69, 0x5f, 0x62, 0x6f,
	0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x67, 0x69, 0x42, 0x6f, 0x78, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x02, 0x52, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x18, 0x21, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x70, 0x76, 0x7a,
	0x5f, 0x73, 0x70, 0x70, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x23, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x70, 0x70, 0x76, 0x7a, 0x53, 0x70, 0x70, 0x50, 0x72, 0x63, 0x12, 0x29, 0x0a, 0x11, 0x70, 0x70,
	0x76, 0x7a, 0x5f, 0x6b, 0x76, 0x77, 0x5f, 0x70, 0x72, 0x63, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x24, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x70, 0x70, 0x76, 0x7a, 0x4b, 0x76, 0x77, 0x50, 0x72,
	0x63, 0x42, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x6b, 0x76,
	0x77, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x25, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x70, 0x76,
	0x7a, 0x4b, 0x76, 0x77, 0x50, 0x72, 0x63, 0x12, 0x29, 0x0a, 0x11, 0x73, 0x75, 0x70, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x63, 0x5f, 0x75, 0x70, 0x18, 0x26, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0e, 0x73, 0x75, 0x70, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x63,
	0x55, 0x70, 0x12, 0x1c, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6b, 0x67, 0x76, 0x70, 0x5f, 0x76, 0x32,
	0x18, 0x27, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x69, 0x73, 0x4b, 0x67, 0x76, 0x70, 0x56, 0x32,
	0x12, 0x32, 0x0a, 0x15, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x13, 0x70, 0x70, 0x76, 0x7a, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x66, 0x6f, 0x72,
	0x5f, 0x70, 0x61, 0x79, 0x18, 0x29, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x70, 0x76, 0x7a,
	0x46, 0x6f, 0x72, 0x50, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x70, 0x76,
	0x7a, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c,
	0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x2c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x42,
	0x61, 0x6e, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x76, 0x77, 0x18, 0x2d,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x70, 0x70, 0x76, 0x7a, 0x56, 0x77, 0x12, 0x1e, 0x0a, 0x0b,
	0x70, 0x70, 0x76, 0x7a, 0x5f, 0x76, 0x77, 0x5f, 0x6e, 0x64, 0x73, 0x18, 0x2e, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x70, 0x70, 0x76, 0x7a, 0x56, 0x77, 0x4e, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x70, 0x76, 0x7a, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x2f,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x70, 0x76, 0x7a, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x70,
	0x76, 0x7a, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x70, 0x70, 0x76, 0x7a, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x31, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x70, 0x76, 0x7a, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x70, 0x76, 0x7a, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x70, 0x76, 0x7a, 0x5f, 0x69, 0x6e, 0x6e,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x70, 0x76, 0x7a, 0x49, 0x6e, 0x6e, 0x12,
	0x2d, 0x0a, 0x12, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x34, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x63,
	0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x36, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x37, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x61,
	0x6c, 0x74, 0x79, 0x18, 0x38, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x61, 0x6c,
	0x74, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x39, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x62, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x12, 0x72, 0x65, 0x62, 0x69, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43,
	0x6f, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x62, 0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x6f, 0x72, 0x67, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x62, 0x69, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x4f, 0x72, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x69, 0x7a, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x69, 0x7a, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x66, 0x65, 0x65, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x3e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x64, 0x65, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x3f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x72, 0x69, 0x64, 0x18, 0x40, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x72, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x41, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x41, 0x49,
	0x4f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x62, 0x53, 0x54, 0x41, 0x54,
	0x49, 0x53, 0x54, 0x49, 0x43, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_statistics_v1_detail_report_item_proto_rawDescOnce sync.Once
	file_wb_statistics_v1_detail_report_item_proto_rawDescData = file_wb_statistics_v1_detail_report_item_proto_rawDesc
)

func file_wb_statistics_v1_detail_report_item_proto_rawDescGZIP() []byte {
	file_wb_statistics_v1_detail_report_item_proto_rawDescOnce.Do(func() {
		file_wb_statistics_v1_detail_report_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_statistics_v1_detail_report_item_proto_rawDescData)
	})
	return file_wb_statistics_v1_detail_report_item_proto_rawDescData
}

var file_wb_statistics_v1_detail_report_item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_statistics_v1_detail_report_item_proto_goTypes = []interface{}{
	(*DetailReportItem)(nil), // 0: wb.statistics.v1.DetailReportItem
	(*anypb.Any)(nil),        // 1: google.protobuf.Any
}
var file_wb_statistics_v1_detail_report_item_proto_depIdxs = []int32{
	1, // 0: wb.statistics.v1.DetailReportItem.suppliercontract_code:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wb_statistics_v1_detail_report_item_proto_init() }
func file_wb_statistics_v1_detail_report_item_proto_init() {
	if File_wb_statistics_v1_detail_report_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wb_statistics_v1_detail_report_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailReportItem); i {
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
			RawDescriptor: file_wb_statistics_v1_detail_report_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wb_statistics_v1_detail_report_item_proto_goTypes,
		DependencyIndexes: file_wb_statistics_v1_detail_report_item_proto_depIdxs,
		MessageInfos:      file_wb_statistics_v1_detail_report_item_proto_msgTypes,
	}.Build()
	File_wb_statistics_v1_detail_report_item_proto = out.File
	file_wb_statistics_v1_detail_report_item_proto_rawDesc = nil
	file_wb_statistics_v1_detail_report_item_proto_goTypes = nil
	file_wb_statistics_v1_detail_report_item_proto_depIdxs = nil
}