//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_stat_words_get200_response_words_keywords_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV1StatWordsGet200ResponseWordsKeywordsInner
 */
export class AdvV1StatWordsGet200ResponseWordsKeywordsInner extends Message<AdvV1StatWordsGet200ResponseWordsKeywordsInner> {
  /**
   * Ключевая фраза
   *
   * @generated from field: string keyword = 1;
   */
  keyword = "";

  /**
   * Количество просмотров по ключевой фразе
   *
   * @generated from field: int32 count = 2;
   */
  count = 0;

  constructor(data?: PartialMessage<AdvV1StatWordsGet200ResponseWordsKeywordsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1StatWordsGet200ResponseWordsKeywordsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "keyword", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1StatWordsGet200ResponseWordsKeywordsInner {
    return new AdvV1StatWordsGet200ResponseWordsKeywordsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1StatWordsGet200ResponseWordsKeywordsInner {
    return new AdvV1StatWordsGet200ResponseWordsKeywordsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1StatWordsGet200ResponseWordsKeywordsInner {
    return new AdvV1StatWordsGet200ResponseWordsKeywordsInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1StatWordsGet200ResponseWordsKeywordsInner | PlainMessage<AdvV1StatWordsGet200ResponseWordsKeywordsInner> | undefined, b: AdvV1StatWordsGet200ResponseWordsKeywordsInner | PlainMessage<AdvV1StatWordsGet200ResponseWordsKeywordsInner> | undefined): boolean {
    return proto3.util.equals(AdvV1StatWordsGet200ResponseWordsKeywordsInner, a, b);
  }
}

