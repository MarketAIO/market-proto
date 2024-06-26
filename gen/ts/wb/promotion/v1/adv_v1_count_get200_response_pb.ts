//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_count_get200_response.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AdvV1CountGet200ResponseAdverts } from "./adv_v1_count_get200_response_adverts_pb.js";

/**
 * @generated from message wb.promotion.v1.AdvV1CountGet200Response
 */
export class AdvV1CountGet200Response extends Message<AdvV1CountGet200Response> {
  /**
   * Общее количество медиакампаний всех статусов и типов
   *
   * @generated from field: int32 all = 1;
   */
  all = 0;

  /**
   * @generated from field: wb.promotion.v1.AdvV1CountGet200ResponseAdverts adverts = 2;
   */
  adverts?: AdvV1CountGet200ResponseAdverts;

  constructor(data?: PartialMessage<AdvV1CountGet200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1CountGet200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "all", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "adverts", kind: "message", T: AdvV1CountGet200ResponseAdverts },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1CountGet200Response {
    return new AdvV1CountGet200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1CountGet200Response {
    return new AdvV1CountGet200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1CountGet200Response {
    return new AdvV1CountGet200Response().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1CountGet200Response | PlainMessage<AdvV1CountGet200Response> | undefined, b: AdvV1CountGet200Response | PlainMessage<AdvV1CountGet200Response> | undefined): boolean {
    return proto3.util.equals(AdvV1CountGet200Response, a, b);
  }
}

