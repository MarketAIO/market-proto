//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_upd_get200_response_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV1UpdGet200ResponseInner
 */
export class AdvV1UpdGet200ResponseInner extends Message<AdvV1UpdGet200ResponseInner> {
  /**
   * Номер выставленного документа (при наличии)
   *
   * @generated from field: int32 updNum = 1;
   */
  updNum = 0;

  /**
   * Время списания
   *
   * @generated from field: string updTime = 2;
   */
  updTime = "";

  /**
   * Выставленная сумма
   *
   * @generated from field: int32 updSum = 3;
   */
  updSum = 0;

  /**
   * Идентификатор кампании
   *
   * @generated from field: int32 advertId = 4;
   */
  advertId = 0;

  /**
   * Название кампании
   *
   * @generated from field: string campName = 5;
   */
  campName = "";

  /**
   * Тип кампании
   *
   * @generated from field: int32 advertType = 6;
   */
  advertType = 0;

  /**
   * <dl> <dt>Источник списания:</dt> <dd>Баланс</dd> <dd>Бонусы</dd> <dd>Счет</dd> </dl> 
   *
   * @generated from field: string paymentType = 7;
   */
  paymentType = "";

  /**
   * <dl>   <dt>Статус кампании:</dt>   <dd><code>4</code> - готова к запуску </dd>   <dd><code>7</code> - завершена</dd>   <dd><code>8</code> - отказался</dd>   <dd><code>9</code> - активна</dd>   <dd><code>11</code> - приостановлена</dd> </dl> 
   *
   * @generated from field: int32 advertStatus = 8;
   */
  advertStatus = 0;

  constructor(data?: PartialMessage<AdvV1UpdGet200ResponseInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1UpdGet200ResponseInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "updNum", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "updTime", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "updSum", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "advertId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "campName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "advertType", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "paymentType", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "advertStatus", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1UpdGet200ResponseInner {
    return new AdvV1UpdGet200ResponseInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1UpdGet200ResponseInner {
    return new AdvV1UpdGet200ResponseInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1UpdGet200ResponseInner {
    return new AdvV1UpdGet200ResponseInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1UpdGet200ResponseInner | PlainMessage<AdvV1UpdGet200ResponseInner> | undefined, b: AdvV1UpdGet200ResponseInner | PlainMessage<AdvV1UpdGet200ResponseInner> | undefined): boolean {
    return proto3.util.equals(AdvV1UpdGet200ResponseInner, a, b);
  }
}

