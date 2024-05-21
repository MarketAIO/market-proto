//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_promotion_count_get200_response.proto (package wb.promotion.v1, syntax proto3)
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
import {
    AdvV1PromotionCountGet200ResponseAdvertsInner
} from "./adv_v1_promotion_count_get200_response_adverts_inner_pb.js";

/**
 * @generated from message wb.promotion.v1.AdvV1PromotionCountGet200Response
 */
export class AdvV1PromotionCountGet200Response extends Message<AdvV1PromotionCountGet200Response> {
  /**
   * Данные по кампаниям
   *
   * @generated from field: repeated wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInner adverts = 1;
   */
  adverts: AdvV1PromotionCountGet200ResponseAdvertsInner[] = [];

  constructor(data?: PartialMessage<AdvV1PromotionCountGet200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1PromotionCountGet200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "adverts", kind: "message", T: AdvV1PromotionCountGet200ResponseAdvertsInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1PromotionCountGet200Response {
    return new AdvV1PromotionCountGet200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1PromotionCountGet200Response {
    return new AdvV1PromotionCountGet200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1PromotionCountGet200Response {
    return new AdvV1PromotionCountGet200Response().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1PromotionCountGet200Response | PlainMessage<AdvV1PromotionCountGet200Response> | undefined, b: AdvV1PromotionCountGet200Response | PlainMessage<AdvV1PromotionCountGet200Response> | undefined): boolean {
    return proto3.util.equals(AdvV1PromotionCountGet200Response, a, b);
  }
}

