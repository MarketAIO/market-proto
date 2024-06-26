//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/days_inner_apps_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { DaysInnerAppsInnerNmInner } from "./days_inner_apps_inner_nm_inner_pb.js";

/**
 * @generated from message wb.promotion.v1.DaysInnerAppsInner
 */
export class DaysInnerAppsInner extends Message<DaysInnerAppsInner> {
  /**
   * Количество просмотров
   *
   * @generated from field: int32 views = 1;
   */
  views = 0;

  /**
   * Количество кликов
   *
   * @generated from field: int32 clicks = 2;
   */
  clicks = 0;

  /**
   * Показатель кликабельности, отношение числа кликов к количеству показов, % 
   *
   * @generated from field: float ctr = 3;
   */
  ctr = 0;

  /**
   * Средняя стоимость клика, ₽.
   *
   * @generated from field: float cpc = 4;
   */
  cpc = 0;

  /**
   * Затраты, ₽.
   *
   * @generated from field: float sum = 5;
   */
  sum = 0;

  /**
   * Количество добавлений товаров в корзину
   *
   * @generated from field: int32 atbs = 6;
   */
  atbs = 0;

  /**
   * Количество заказов
   *
   * @generated from field: int32 orders = 7;
   */
  orders = 0;

  /**
   * CR(conversion rate) — это отношение количества заказов к общему количеству посещений кампании 
   *
   * @generated from field: int32 cr = 8;
   */
  cr = 0;

  /**
   * Количество заказанных товаров, шт.
   *
   * @generated from field: int32 shks = 9;
   */
  shks = 0;

  /**
   * Заказов на сумму, ₽
   *
   * @generated from field: float sum_price = 10;
   */
  sumPrice = 0;

  /**
   * Блок статистики по артикулам WB
   *
   * @generated from field: repeated wb.promotion.v1.DaysInnerAppsInnerNmInner nm = 11;
   */
  nm: DaysInnerAppsInnerNmInner[] = [];

  /**
   * Тип платформы (`1` - сайт, `32` - Android, `64` - IOS)
   *
   * @generated from field: int32 appType = 12;
   */
  appType = 0;

  constructor(data?: PartialMessage<DaysInnerAppsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.DaysInnerAppsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "views", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "clicks", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "ctr", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "cpc", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "sum", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 6, name: "atbs", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "orders", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "cr", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 9, name: "shks", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "sum_price", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 11, name: "nm", kind: "message", T: DaysInnerAppsInnerNmInner, repeated: true },
    { no: 12, name: "appType", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DaysInnerAppsInner {
    return new DaysInnerAppsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DaysInnerAppsInner {
    return new DaysInnerAppsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DaysInnerAppsInner {
    return new DaysInnerAppsInner().fromJsonString(jsonString, options);
  }

  static equals(a: DaysInnerAppsInner | PlainMessage<DaysInnerAppsInner> | undefined, b: DaysInnerAppsInner | PlainMessage<DaysInnerAppsInner> | undefined): boolean {
    return proto3.util.equals(DaysInnerAppsInner, a, b);
  }
}

