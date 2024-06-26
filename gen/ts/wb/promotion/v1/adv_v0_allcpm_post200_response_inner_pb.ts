//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v0_allcpm_post200_response_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AdvV0AllcpmPost200ResponseInnerCpmInner } from "./adv_v0_allcpm_post200_response_inner_cpm_inner_pb.js";

/**
 * @generated from message wb.promotion.v1.AdvV0AllcpmPost200ResponseInner
 */
export class AdvV0AllcpmPost200ResponseInner extends Message<AdvV0AllcpmPost200ResponseInner> {
  /**
   * Значение параметра (param) запроса
   *
   * @generated from field: int32 param = 1;
   */
  param = 0;

  /**
   * Информация о ставке(-ах)
   *
   * @generated from field: repeated wb.promotion.v1.AdvV0AllcpmPost200ResponseInnerCpmInner cpm = 2;
   */
  cpm: AdvV0AllcpmPost200ResponseInnerCpmInner[] = [];

  constructor(data?: PartialMessage<AdvV0AllcpmPost200ResponseInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV0AllcpmPost200ResponseInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "param", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "cpm", kind: "message", T: AdvV0AllcpmPost200ResponseInnerCpmInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV0AllcpmPost200ResponseInner {
    return new AdvV0AllcpmPost200ResponseInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV0AllcpmPost200ResponseInner {
    return new AdvV0AllcpmPost200ResponseInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV0AllcpmPost200ResponseInner {
    return new AdvV0AllcpmPost200ResponseInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV0AllcpmPost200ResponseInner | PlainMessage<AdvV0AllcpmPost200ResponseInner> | undefined, b: AdvV0AllcpmPost200ResponseInner | PlainMessage<AdvV0AllcpmPost200ResponseInner> | undefined): boolean {
    return proto3.util.equals(AdvV0AllcpmPost200ResponseInner, a, b);
  }
}

