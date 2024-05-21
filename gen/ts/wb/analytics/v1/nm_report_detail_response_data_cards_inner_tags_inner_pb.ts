//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/nm_report_detail_response_data_cards_inner_tags_inner.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.analytics.v1.NmReportDetailResponseDataCardsInnerTagsInner
 */
export class NmReportDetailResponseDataCardsInnerTagsInner extends Message<NmReportDetailResponseDataCardsInnerTagsInner> {
  /**
   * Идентификатор тега
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Название тега
   *
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<NmReportDetailResponseDataCardsInnerTagsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.NmReportDetailResponseDataCardsInnerTagsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NmReportDetailResponseDataCardsInnerTagsInner {
    return new NmReportDetailResponseDataCardsInnerTagsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NmReportDetailResponseDataCardsInnerTagsInner {
    return new NmReportDetailResponseDataCardsInnerTagsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NmReportDetailResponseDataCardsInnerTagsInner {
    return new NmReportDetailResponseDataCardsInnerTagsInner().fromJsonString(jsonString, options);
  }

  static equals(a: NmReportDetailResponseDataCardsInnerTagsInner | PlainMessage<NmReportDetailResponseDataCardsInnerTagsInner> | undefined, b: NmReportDetailResponseDataCardsInnerTagsInner | PlainMessage<NmReportDetailResponseDataCardsInnerTagsInner> | undefined): boolean {
    return proto3.util.equals(NmReportDetailResponseDataCardsInnerTagsInner, a, b);
  }
}

