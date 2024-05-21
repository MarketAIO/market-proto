//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v2_fullstats_post200_response.proto (package wb.promotion.v1, syntax proto3)
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
 * @generated from message wb.promotion.v1.AdvV2FullstatsPost200Response
 */
export class AdvV2FullstatsPost200Response extends Message<AdvV2FullstatsPost200Response> {
  constructor(data?: PartialMessage<AdvV2FullstatsPost200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV2FullstatsPost200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV2FullstatsPost200Response {
    return new AdvV2FullstatsPost200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV2FullstatsPost200Response {
    return new AdvV2FullstatsPost200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV2FullstatsPost200Response {
    return new AdvV2FullstatsPost200Response().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV2FullstatsPost200Response | PlainMessage<AdvV2FullstatsPost200Response> | undefined, b: AdvV2FullstatsPost200Response | PlainMessage<AdvV2FullstatsPost200Response> | undefined): boolean {
    return proto3.util.equals(AdvV2FullstatsPost200Response, a, b);
  }
}
