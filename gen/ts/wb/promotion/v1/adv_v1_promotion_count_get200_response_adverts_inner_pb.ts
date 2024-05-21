//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_promotion_count_get200_response_adverts_inner.proto (package wb.promotion.v1, syntax proto3)
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
    AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner
} from "./adv_v1_promotion_count_get200_response_adverts_inner_advert_list_inner_pb.js";

/**
 * @generated from message wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInner
 */
export class AdvV1PromotionCountGet200ResponseAdvertsInner extends Message<AdvV1PromotionCountGet200ResponseAdvertsInner> {
  /**
   * Тип кампании
   *
   * @generated from field: int32 type = 1;
   */
  type = 0;

  /**
   * Статус кампании
   *
   * @generated from field: int32 status = 2;
   */
  status = 0;

  /**
   * Количество кампаний
   *
   * @generated from field: int32 count = 3;
   */
  count = 0;

  /**
   * Список кампаний
   *
   * @generated from field: repeated wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner advert_list = 4;
   */
  advertList: AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner[] = [];

  constructor(data?: PartialMessage<AdvV1PromotionCountGet200ResponseAdvertsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1PromotionCountGet200ResponseAdvertsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "type", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "status", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "advert_list", kind: "message", T: AdvV1PromotionCountGet200ResponseAdvertsInnerAdvertListInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1PromotionCountGet200ResponseAdvertsInner {
    return new AdvV1PromotionCountGet200ResponseAdvertsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1PromotionCountGet200ResponseAdvertsInner {
    return new AdvV1PromotionCountGet200ResponseAdvertsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1PromotionCountGet200ResponseAdvertsInner {
    return new AdvV1PromotionCountGet200ResponseAdvertsInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1PromotionCountGet200ResponseAdvertsInner | PlainMessage<AdvV1PromotionCountGet200ResponseAdvertsInner> | undefined, b: AdvV1PromotionCountGet200ResponseAdvertsInner | PlainMessage<AdvV1PromotionCountGet200ResponseAdvertsInner> | undefined): boolean {
    return proto3.util.equals(AdvV1PromotionCountGet200ResponseAdvertsInner, a, b);
  }
}

