//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_auto_updatenm_post_request.proto (package wb.promotion.v1, syntax proto3)
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
 * @generated from message wb.promotion.v1.AdvV1AutoUpdatenmPostReq
 */
export class AdvV1AutoUpdatenmPostReq extends Message<AdvV1AutoUpdatenmPostReq> {
  /**
   * Номенклатуры, которые необходимо добавить.
   *
   * @generated from field: repeated int32 add = 1;
   */
  add: number[] = [];

  /**
   * Номенклатуры, которые необходимо удалить.
   *
   * @generated from field: repeated int32 delete = 2;
   */
  delete: number[] = [];

  constructor(data?: PartialMessage<AdvV1AutoUpdatenmPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1AutoUpdatenmPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "add", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "delete", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1AutoUpdatenmPostReq {
    return new AdvV1AutoUpdatenmPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1AutoUpdatenmPostReq {
    return new AdvV1AutoUpdatenmPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1AutoUpdatenmPostReq {
    return new AdvV1AutoUpdatenmPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1AutoUpdatenmPostReq | PlainMessage<AdvV1AutoUpdatenmPostReq> | undefined, b: AdvV1AutoUpdatenmPostReq | PlainMessage<AdvV1AutoUpdatenmPostReq> | undefined): boolean {
    return proto3.util.equals(AdvV1AutoUpdatenmPostReq, a, b);
  }
}
