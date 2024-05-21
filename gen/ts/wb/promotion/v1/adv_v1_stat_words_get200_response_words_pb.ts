//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_stat_words_get200_response_words.proto (package wb.promotion.v1, syntax proto3)
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
    AdvV1StatWordsGet200ResponseWordsKeywordsInner
} from "./adv_v1_stat_words_get200_response_words_keywords_inner_pb.js";

/**
 * @generated from message wb.promotion.v1.AdvV1StatWordsGet200ResponseWords
 */
export class AdvV1StatWordsGet200ResponseWords extends Message<AdvV1StatWordsGet200ResponseWords> {
  /**
   * Фразовое соответствие (минус фразы)
   *
   * @generated from field: repeated string phrase = 1;
   */
  phrase: string[] = [];

  /**
   * Точное соответствие (минус фразы)
   *
   * @generated from field: repeated string strong = 2;
   */
  strong: string[] = [];

  /**
   * Минус фразы из поиска
   *
   * @generated from field: repeated string excluded = 3;
   */
  excluded: string[] = [];

  /**
   * Фиксированные фразы
   *
   * @generated from field: repeated string pluse = 4;
   */
  pluse: string[] = [];

  /**
   * Блок со статистикой по ключевым фразам
   *
   * @generated from field: repeated wb.promotion.v1.AdvV1StatWordsGet200ResponseWordsKeywordsInner keywords = 5;
   */
  keywords: AdvV1StatWordsGet200ResponseWordsKeywordsInner[] = [];

  /**
   * Фиксированные ключевые фразы (`true` - включены, `false` - выключены) 
   *
   * @generated from field: bool fixed = 6;
   */
  fixed = false;

  constructor(data?: PartialMessage<AdvV1StatWordsGet200ResponseWords>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1StatWordsGet200ResponseWords";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "phrase", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "strong", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "excluded", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "pluse", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "keywords", kind: "message", T: AdvV1StatWordsGet200ResponseWordsKeywordsInner, repeated: true },
    { no: 6, name: "fixed", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1StatWordsGet200ResponseWords {
    return new AdvV1StatWordsGet200ResponseWords().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1StatWordsGet200ResponseWords {
    return new AdvV1StatWordsGet200ResponseWords().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1StatWordsGet200ResponseWords {
    return new AdvV1StatWordsGet200ResponseWords().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1StatWordsGet200ResponseWords | PlainMessage<AdvV1StatWordsGet200ResponseWords> | undefined, b: AdvV1StatWordsGet200ResponseWords | PlainMessage<AdvV1StatWordsGet200ResponseWords> | undefined): boolean {
    return proto3.util.equals(AdvV1StatWordsGet200ResponseWords, a, b);
  }
}
