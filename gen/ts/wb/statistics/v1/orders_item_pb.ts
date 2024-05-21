//
//Описание API Статистики
//
//С помощью этих методов можно получить отчёты.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/statistics/v1/orders_item.proto (package wb.statistics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.statistics.v1.OrdersItem
 */
export class OrdersItem extends Message<OrdersItem> {
  /**
   * Дата и время заказа. Это поле соответствует параметру `dateFrom` в запросе, если параметр `flag`=1. Если часовой пояс не указан, то берется Московское время (UTC+3).
   *
   * @generated from field: string date = 1;
   */
  date = "";

  /**
   * Дата и время обновления информации в сервисе. Это поле соответствует параметру `dateFrom` в запросе, если параметр `flag`=0 или не указан. Если часовой пояс не указан, то берется Московское время (UTC+3).
   *
   * @generated from field: string lastChangeDate = 2;
   */
  lastChangeDate = "";

  /**
   * Склад отгрузки
   *
   * @generated from field: string warehouseName = 3;
   */
  warehouseName = "";

  /**
   * Страна
   *
   * @generated from field: string countryName = 4;
   */
  countryName = "";

  /**
   * Округ
   *
   * @generated from field: string oblastOkrugName = 5;
   */
  oblastOkrugName = "";

  /**
   * Регион
   *
   * @generated from field: string regionName = 6;
   */
  regionName = "";

  /**
   * Артикул продавца
   *
   * @generated from field: string supplierArticle = 7;
   */
  supplierArticle = "";

  /**
   * Артикул WB
   *
   * @generated from field: int32 nmId = 8;
   */
  nmId = 0;

  /**
   * Баркод
   *
   * @generated from field: string barcode = 9;
   */
  barcode = "";

  /**
   * Категория
   *
   * @generated from field: string category = 10;
   */
  category = "";

  /**
   * Предмет
   *
   * @generated from field: string subject = 11;
   */
  subject = "";

  /**
   * Бренд
   *
   * @generated from field: string brand = 12;
   */
  brand = "";

  /**
   * Размер товара
   *
   * @generated from field: string techSize = 13;
   */
  techSize = "";

  /**
   * Номер поставки
   *
   * @generated from field: int32 incomeID = 14;
   */
  incomeID = 0;

  /**
   * Договор поставки
   *
   * @generated from field: bool isSupply = 15;
   */
  isSupply = false;

  /**
   * Договор реализации
   *
   * @generated from field: bool isRealization = 16;
   */
  isRealization = false;

  /**
   * Цена без скидок
   *
   * @generated from field: float totalPrice = 17;
   */
  totalPrice = 0;

  /**
   * Скидка продавца
   *
   * @generated from field: int32 discountPercent = 18;
   */
  discountPercent = 0;

  /**
   * Скидка WB
   *
   * @generated from field: float spp = 19;
   */
  spp = 0;

  /**
   * Цена с учетом всех скидок, кроме суммы по WB Кошельку
   *
   * @generated from field: float finishedPrice = 20;
   */
  finishedPrice = 0;

  /**
   * Цена со скидкой продавца (= `totalPrice` * (1 - `discountPercent`/100))
   *
   * @generated from field: float priceWithDisc = 21;
   */
  priceWithDisc = 0;

  /**
   * Отмена заказа. true - заказ отменен
   *
   * @generated from field: bool isCancel = 22;
   */
  isCancel = false;

  /**
   * Дата и время отмены заказа. Если заказ не был отменен, то \"0001-01-01T00:00:00\".Если часовой пояс не указан, то берется Московское время UTC+3.
   *
   * @generated from field: string cancelDate = 23;
   */
  cancelDate = "";

  /**
   * Тип заказа <ul> <li> `Клиентский` — заказ, поступивший от покупателя <li> `Возврат Брака` — возврат товара продавцу <li> `Принудительный возврат` — возврат товара продавцу <li> `Возврат обезлички` — возврат товара продавцу <li> `Возврат Неверного Вложения` — возврат товара продавцу <li> `Возврат Продавца` — возврат товара продавцу </ul> 
   *
   * @generated from field: string orderType = 24;
   */
  orderType = "";

  /**
   * Идентификатор стикера
   *
   * @generated from field: string sticker = 25;
   */
  sticker = "";

  /**
   * Номер заказа
   *
   * @generated from field: string gNumber = 26;
   */
  gNumber = "";

  /**
   * Уникальный идентификатор заказа.<br> Примечание для использующих API Маркетплейс: `srid` равен `rid` в ответах методов сборочных заданий. 
   *
   * @generated from field: string srid = 27;
   */
  srid = "";

  constructor(data?: PartialMessage<OrdersItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.OrdersItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "lastChangeDate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "warehouseName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "countryName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "oblastOkrugName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "regionName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "supplierArticle", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "nmId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 9, name: "barcode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "category", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "subject", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "brand", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 13, name: "techSize", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 14, name: "incomeID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 15, name: "isSupply", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 16, name: "isRealization", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 17, name: "totalPrice", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 18, name: "discountPercent", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 19, name: "spp", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 20, name: "finishedPrice", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 21, name: "priceWithDisc", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 22, name: "isCancel", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 23, name: "cancelDate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 24, name: "orderType", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 25, name: "sticker", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 26, name: "gNumber", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 27, name: "srid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OrdersItem {
    return new OrdersItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OrdersItem {
    return new OrdersItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OrdersItem {
    return new OrdersItem().fromJsonString(jsonString, options);
  }

  static equals(a: OrdersItem | PlainMessage<OrdersItem> | undefined, b: OrdersItem | PlainMessage<OrdersItem> | undefined): boolean {
    return proto3.util.equals(OrdersItem, a, b);
  }
}
