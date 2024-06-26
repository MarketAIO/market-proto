//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/days_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { DaysInnerAppsInner } from "./days_inner_apps_inner_pb.js";

/**
 * @generated from message wb.promotion.v1.DaysInner
 */
export class DaysInner extends Message<DaysInner> {
  /**
   * Дата, за которую представлены данные
   *
   * @generated from field: string date = 1;
   */
  date = "";

  /**
   * Количество просмотров
   *
   * @generated from field: int32 views = 2;
   */
  views = 0;

  /**
   * Количество кликов
   *
   * @generated from field: int32 clicks = 3;
   */
  clicks = 0;

  /**
   * Показатель кликабельности, отношение числа кликов к количеству показов, % 
   *
   * @generated from field: float ctr = 4;
   */
  ctr = 0;

  /**
   * Средняя стоимость клика, ₽.
   *
   * @generated from field: float cpc = 5;
   */
  cpc = 0;

  /**
   * Затраты, ₽.
   *
   * @generated from field: float sum = 6;
   */
  sum = 0;

  /**
   * Количество добавлений товаров в корзину
   *
   * @generated from field: int32 atbs = 7;
   */
  atbs = 0;

  /**
   * Количество заказов
   *
   * @generated from field: int32 orders = 8;
   */
  orders = 0;

  /**
   * CR(conversion rate) — отношение количества заказов к общему количеству посещений кампании 
   *
   * @generated from field: int32 cr = 9;
   */
  cr = 0;

  /**
   * Количество заказанных товаров, шт.
   *
   * @generated from field: int32 shks = 10;
   */
  shks = 0;

  /**
   * Заказов на сумму, ₽
   *
   * @generated from field: float sum_price = 11;
   */
  sumPrice = 0;

  /**
   * Блок информации о платформе
   *
   * @generated from field: repeated wb.promotion.v1.DaysInnerAppsInner apps = 12;
   */
  apps: DaysInnerAppsInner[] = [];

  constructor(data?: PartialMessage<DaysInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.DaysInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "views", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "clicks", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "ctr", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "cpc", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 6, name: "sum", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 7, name: "atbs", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "orders", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 9, name: "cr", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "shks", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 11, name: "sum_price", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 12, name: "apps", kind: "message", T: DaysInnerAppsInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DaysInner {
    return new DaysInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DaysInner {
    return new DaysInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DaysInner {
    return new DaysInner().fromJsonString(jsonString, options);
  }

  static equals(a: DaysInner | PlainMessage<DaysInner> | undefined, b: DaysInner | PlainMessage<DaysInner> | undefined): boolean {
    return proto3.util.equals(DaysInner, a, b);
  }
}

