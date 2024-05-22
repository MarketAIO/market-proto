//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v2_auto_stat_words_get200_response_clusters_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV2AutoStatWordsGet200ResponseClustersInner
 */
export class AdvV2AutoStatWordsGet200ResponseClustersInner extends Message<AdvV2AutoStatWordsGet200ResponseClustersInner> {
  /**
   * Кластер — набор похожих ключевых фраз
   *
   * @generated from field: string cluster = 1;
   */
  cluster = "";

  /**
   * Сколько раз товары показывались по всем фразам из кластера
   *
   * @generated from field: int32 count = 2;
   */
  count = 0;

  /**
   * Ключевые фразы из кластера, по которым товары показывались хотя бы один раз
   *
   * @generated from field: repeated string keywords = 3;
   */
  keywords: string[] = [];

  constructor(data?: PartialMessage<AdvV2AutoStatWordsGet200ResponseClustersInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV2AutoStatWordsGet200ResponseClustersInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cluster", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "keywords", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV2AutoStatWordsGet200ResponseClustersInner {
    return new AdvV2AutoStatWordsGet200ResponseClustersInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV2AutoStatWordsGet200ResponseClustersInner {
    return new AdvV2AutoStatWordsGet200ResponseClustersInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV2AutoStatWordsGet200ResponseClustersInner {
    return new AdvV2AutoStatWordsGet200ResponseClustersInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV2AutoStatWordsGet200ResponseClustersInner | PlainMessage<AdvV2AutoStatWordsGet200ResponseClustersInner> | undefined, b: AdvV2AutoStatWordsGet200ResponseClustersInner | PlainMessage<AdvV2AutoStatWordsGet200ResponseClustersInner> | undefined): boolean {
    return proto3.util.equals(AdvV2AutoStatWordsGet200ResponseClustersInner, a, b);
  }
}

