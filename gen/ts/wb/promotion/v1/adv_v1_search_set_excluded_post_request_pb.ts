//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_search_set_excluded_post_request.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV1SearchSetExcludedPostReq
 */
export class AdvV1SearchSetExcludedPostReq extends Message<AdvV1SearchSetExcludedPostReq> {
  /**
   * Минус-фразы (макс. 1000 шт.)
   *
   * @generated from field: repeated string excluded = 1;
   */
  excluded: string[] = [];

  constructor(data?: PartialMessage<AdvV1SearchSetExcludedPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1SearchSetExcludedPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "excluded", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1SearchSetExcludedPostReq {
    return new AdvV1SearchSetExcludedPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1SearchSetExcludedPostReq {
    return new AdvV1SearchSetExcludedPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1SearchSetExcludedPostReq {
    return new AdvV1SearchSetExcludedPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1SearchSetExcludedPostReq | PlainMessage<AdvV1SearchSetExcludedPostReq> | undefined, b: AdvV1SearchSetExcludedPostReq | PlainMessage<AdvV1SearchSetExcludedPostReq> | undefined): boolean {
    return proto3.util.equals(AdvV1SearchSetExcludedPostReq, a, b);
  }
}

