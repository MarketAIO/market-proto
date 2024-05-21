//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v0_allcpm_post200_response_inner_cpm_inner.proto (package wb.promotion.v1, syntax proto3)
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
 * @generated from message wb.promotion.v1.AdvV0AllcpmPost200ResponseInnerCpmInner
 */
export class AdvV0AllcpmPost200ResponseInnerCpmInner extends Message<AdvV0AllcpmPost200ResponseInnerCpmInner> {
  /**
   * Размер ставки
   *
   * @generated from field: int32 Cpm = 1;
   */
  Cpm = 0;

  /**
   * Количество ставок
   *
   * @generated from field: int32 Count = 2;
   */
  Count = 0;

  constructor(data?: PartialMessage<AdvV0AllcpmPost200ResponseInnerCpmInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV0AllcpmPost200ResponseInnerCpmInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "Cpm", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "Count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV0AllcpmPost200ResponseInnerCpmInner {
    return new AdvV0AllcpmPost200ResponseInnerCpmInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV0AllcpmPost200ResponseInnerCpmInner {
    return new AdvV0AllcpmPost200ResponseInnerCpmInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV0AllcpmPost200ResponseInnerCpmInner {
    return new AdvV0AllcpmPost200ResponseInnerCpmInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV0AllcpmPost200ResponseInnerCpmInner | PlainMessage<AdvV0AllcpmPost200ResponseInnerCpmInner> | undefined, b: AdvV0AllcpmPost200ResponseInnerCpmInner | PlainMessage<AdvV0AllcpmPost200ResponseInnerCpmInner> | undefined): boolean {
    return proto3.util.equals(AdvV0AllcpmPost200ResponseInnerCpmInner, a, b);
  }
}

