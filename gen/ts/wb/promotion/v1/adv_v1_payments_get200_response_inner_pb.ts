//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_payments_get200_response_inner.proto (package wb.promotion.v1, syntax proto3)
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
 * @generated from message wb.promotion.v1.AdvV1PaymentsGet200ResponseInner
 */
export class AdvV1PaymentsGet200ResponseInner extends Message<AdvV1PaymentsGet200ResponseInner> {
  /**
   * Идентификатор платежа
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Дата платежа
   *
   * @generated from field: string date = 2;
   */
  date = "";

  /**
   * Сумма платежа
   *
   * @generated from field: int32 sum = 3;
   */
  sum = 0;

  /**
   * <dl> <dt>Тип источника списания:</dt> <dd><code>0</code> - Счёт</dd> <dd><code>1</code> - Баланс</dd> <dd><code>3</code> - Картой</dd> </dl> 
   *
   * @generated from field: int32 type = 4;
   */
  type = 0;

  /**
   * <dl> <dt>Статус:</dt> <dd><code>0</code> - ошибка</dd> <dd><code>1</code> - обработано</dd> </dl> 
   *
   * @generated from field: int32 statusId = 5;
   */
  statusId = 0;

  /**
   * <dl> <dt>Статус операции(при оплате картой):</dt> <dd><b>success</b> - успех</dd> <dd><b>fail</b> - неуспех</dd> <dd><b>pending</b> - в ожидании ответа</dd> <dd><b>unknown</b> - неизвестно</dd> </dl> 
   *
   * @generated from field: string cardStatus = 6;
   */
  cardStatus = "";

  constructor(data?: PartialMessage<AdvV1PaymentsGet200ResponseInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1PaymentsGet200ResponseInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "sum", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "type", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "statusId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "cardStatus", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1PaymentsGet200ResponseInner {
    return new AdvV1PaymentsGet200ResponseInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1PaymentsGet200ResponseInner {
    return new AdvV1PaymentsGet200ResponseInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1PaymentsGet200ResponseInner {
    return new AdvV1PaymentsGet200ResponseInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1PaymentsGet200ResponseInner | PlainMessage<AdvV1PaymentsGet200ResponseInner> | undefined, b: AdvV1PaymentsGet200ResponseInner | PlainMessage<AdvV1PaymentsGet200ResponseInner> | undefined): boolean {
    return proto3.util.equals(AdvV1PaymentsGet200ResponseInner, a, b);
  }
}
