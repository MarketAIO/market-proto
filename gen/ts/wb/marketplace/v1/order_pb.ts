//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/order.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { OrderAddress } from "./order_address_pb.js";
import { OrderUser } from "./order_user_pb.js";

/**
 * @generated from message wb.marketplace.v1.Order
 */
export class Order extends Message<Order> {
  /**
   * Идентификатор сборочного задания в Маркетплейсе
   *
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * Идентификатор сборочного задания в системе Wildberries
   *
   * @generated from field: string rid = 2;
   */
  rid = "";

  /**
   * Дата создания сборочного задания (RFC3339)
   *
   * @generated from field: string createdAt = 3;
   */
  createdAt = "";

  /**
   * Идентификатор склада продавца, на который поступило сборочное задание
   *
   * @generated from field: int32 warehouseId = 4;
   */
  warehouseId = 0;

  /**
   * Идентификатор поставки. Возвращается, если заказ закреплён за поставкой
   *
   * @generated from field: string supplyId = 5;
   */
  supplyId = "";

  /**
   * Список офисов, куда следует привезти товар
   *
   * @generated from field: repeated string offices = 6;
   */
  offices: string[] = [];

  /**
   * @generated from field: wb.marketplace.v1.OrderAddress address = 7;
   */
  address?: OrderAddress;

  /**
   * @generated from field: wb.marketplace.v1.OrderUser user = 8;
   */
  user?: OrderUser;

  /**
   * Массив баркодов товара
   *
   * @generated from field: repeated string skus = 9;
   */
  skus: string[] = [];

  /**
   * Цена в валюте продажи с учетом всех скидок, умноженная на 100. Код валюты продажи в поле currencyCode.
   *
   * @generated from field: int32 price = 10;
   */
  price = 0;

  /**
   * Цена в валюте страны продавца с учетом всех скидок, кроме суммы по WB Кошельку, умноженная на 100. Предоставляется в информационных целях.
   *
   * @generated from field: int32 convertedPrice = 11;
   */
  convertedPrice = 0;

  /**
   * Код валюты продажи (ISO 4217)
   *
   * @generated from field: int32 currencyCode = 12;
   */
  currencyCode = 0;

  /**
   * Код валюты страны продавца (ISO 4217)
   *
   * @generated from field: int32 convertedCurrencyCode = 13;
   */
  convertedCurrencyCode = 0;

  /**
   * Идентификатор транзакции для группировки сборочных заданий. Сборочные задания в одной корзине покупателя будут иметь одинаковый orderUID
   *
   * @generated from field: string orderUid = 14;
   */
  orderUid = "";

  /**
   * @generated from field: wb.marketplace.v1.Order.DeliveryTypeEnum deliveryType = 15;
   */
  deliveryType = Order_DeliveryTypeEnum.DeliveryTypeEnum_DBS;

  /**
   * Артикул WB
   *
   * @generated from field: int32 nmId = 16;
   */
  nmId = 0;

  /**
   * Идентификатор размера товара в системе Wildberries
   *
   * @generated from field: int32 chrtId = 17;
   */
  chrtId = 0;

  /**
   * Артикул продавца
   *
   * @generated from field: string article = 18;
   */
  article = "";

  /**
   * Код цвета (только для колеруемых товаров)
   *
   * @generated from field: string colorCode = 19;
   */
  colorCode = "";

  /**
   * @generated from field: wb.marketplace.v1.Order.CargoTypeEnum cargoType = 20;
   */
  cargoType = Order_CargoTypeEnum.CargoTypeEnum__1;

  constructor(data?: PartialMessage<Order>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.Order";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "rid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "createdAt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "warehouseId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "supplyId", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "offices", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "address", kind: "message", T: OrderAddress },
    { no: 8, name: "user", kind: "message", T: OrderUser },
    { no: 9, name: "skus", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 10, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 11, name: "convertedPrice", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 12, name: "currencyCode", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 13, name: "convertedCurrencyCode", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 14, name: "orderUid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 15, name: "deliveryType", kind: "enum", T: proto3.getEnumType(Order_DeliveryTypeEnum) },
    { no: 16, name: "nmId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 17, name: "chrtId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 18, name: "article", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 19, name: "colorCode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 20, name: "cargoType", kind: "enum", T: proto3.getEnumType(Order_CargoTypeEnum) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Order {
    return new Order().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Order {
    return new Order().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Order {
    return new Order().fromJsonString(jsonString, options);
  }

  static equals(a: Order | PlainMessage<Order> | undefined, b: Order | PlainMessage<Order> | undefined): boolean {
    return proto3.util.equals(Order, a, b);
  }
}

/**
 * <dl> <dt>Тип доставки:</dt> <dd>fbs - доставка на склад Wildberries</dd> <dd>dbs - доставка силами продавца</dd> <dd>wbgo - доставка курьером WB</dd> <dd>edbs - экспресс-доставка силами продавца</dd> </dl> 
 *
 * @generated from enum wb.marketplace.v1.Order.DeliveryTypeEnum
 */
export enum Order_DeliveryTypeEnum {
  /**
   * @generated from enum value: DeliveryTypeEnum_DBS = 0;
   */
  DeliveryTypeEnum_DBS = 0,

  /**
   * @generated from enum value: DeliveryTypeEnum_EDBS = 1;
   */
  DeliveryTypeEnum_EDBS = 1,

  /**
   * @generated from enum value: DeliveryTypeEnum_FBS = 2;
   */
  DeliveryTypeEnum_FBS = 2,

  /**
   * @generated from enum value: DeliveryTypeEnum_WBGO = 3;
   */
  DeliveryTypeEnum_WBGO = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Order_DeliveryTypeEnum)
proto3.util.setEnumType(Order_DeliveryTypeEnum, "wb.marketplace.v1.Order.DeliveryTypeEnum", [
  { no: 0, name: "DeliveryTypeEnum_DBS" },
  { no: 1, name: "DeliveryTypeEnum_EDBS" },
  { no: 2, name: "DeliveryTypeEnum_FBS" },
  { no: 3, name: "DeliveryTypeEnum_WBGO" },
]);

/**
 * <dl> <dt>Тип товара:</dt> <dd>1 - обычный</dd> <dd>2 - СГТ (Сверхгабаритный товар)</dd> <dd>3 - КГТ (Крупногабаритный товар). Не используется на данный момент.</dd> </dl> 
 *
 * @generated from enum wb.marketplace.v1.Order.CargoTypeEnum
 */
export enum Order_CargoTypeEnum {
  /**
   * @generated from enum value: CargoTypeEnum__1 = 0;
   */
  CargoTypeEnum__1 = 0,

  /**
   * @generated from enum value: CargoTypeEnum__2 = 1;
   */
  CargoTypeEnum__2 = 1,

  /**
   * @generated from enum value: CargoTypeEnum__3 = 2;
   */
  CargoTypeEnum__3 = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Order_CargoTypeEnum)
proto3.util.setEnumType(Order_CargoTypeEnum, "wb.marketplace.v1.Order.CargoTypeEnum", [
  { no: 0, name: "CargoTypeEnum__1" },
  { no: 1, name: "CargoTypeEnum__2" },
  { no: 2, name: "CargoTypeEnum__3" },
]);

