//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_balance_get200_response.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
    BinaryReadOptions,
    FieldList,
    JsonReadOptions,
    JsonValue,
    PartialMessage,
    PlainMessage
} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV1BalanceGet200Response
 */
export class AdvV1BalanceGet200Response extends Message<AdvV1BalanceGet200Response> {
  /**
   * Счёт, рублей
   *
   * @generated from field: int32 balance = 1;
   */
  balance = 0;

  /**
   * Баланс, рублей
   *
   * @generated from field: int32 net = 2;
   */
  net = 0;

  /**
   * Бонусы, рублей
   *
   * @generated from field: int32 bonus = 3;
   */
  bonus = 0;

  constructor(data?: PartialMessage<AdvV1BalanceGet200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1BalanceGet200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "balance", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "net", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "bonus", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1BalanceGet200Response {
    return new AdvV1BalanceGet200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1BalanceGet200Response {
    return new AdvV1BalanceGet200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1BalanceGet200Response {
    return new AdvV1BalanceGet200Response().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1BalanceGet200Response | PlainMessage<AdvV1BalanceGet200Response> | undefined, b: AdvV1BalanceGet200Response | PlainMessage<AdvV1BalanceGet200Response> | undefined): boolean {
    return proto3.util.equals(AdvV1BalanceGet200Response, a, b);
  }
}
