//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/nm_report_grouped_response_data_groups_inner_statistics_period_comparison_conversions.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions
 */
export class NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions extends Message<NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions> {
  /**
   * Конверсия в корзину, % (Какой процент посетителей, открывших карточку товара, добавили товар в корзину)
   *
   * @generated from field: int32 addToCartPercent = 1;
   */
  addToCartPercent = 0;

  /**
   * Конверсия в заказ, % (Какой процент посетителей, добавивших товар в корзину, сделали заказ)
   *
   * @generated from field: int32 cartToOrderPercent = 2;
   */
  cartToOrderPercent = 0;

  /**
   * Процент выкупа, % (Какой процент посетителей, заказавших товар, его выкупили. Без учёта товаров, которые еще доставляются покупателю.)
   *
   * @generated from field: int32 buyoutsPercent = 3;
   */
  buyoutsPercent = 0;

  constructor(data?: PartialMessage<NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "addToCartPercent", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "cartToOrderPercent", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "buyoutsPercent", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions {
    return new NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions {
    return new NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions {
    return new NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions().fromJsonString(jsonString, options);
  }

  static equals(a: NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions | PlainMessage<NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions> | undefined, b: NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions | PlainMessage<NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions> | undefined): boolean {
    return proto3.util.equals(NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparisonConversions, a, b);
  }
}

