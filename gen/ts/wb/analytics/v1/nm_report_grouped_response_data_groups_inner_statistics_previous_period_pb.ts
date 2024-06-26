//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/nm_report_grouped_response_data_groups_inner_statistics_previous_period.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions } from "./nm_report_grouped_response_data_groups_inner_statistics_previous_period_conversions_pb.js";

/**
 * @generated from message wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod
 */
export class NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod extends Message<NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod> {
  /**
   * Начало периода
   *
   * @generated from field: string begin = 1;
   */
  begin = "";

  /**
   * Конец периода
   *
   * @generated from field: string end = 2;
   */
  end = "";

  /**
   * Количество переходов в КТ
   *
   * @generated from field: int32 openCardCount = 3;
   */
  openCardCount = 0;

  /**
   * Положили в корзину, штук
   *
   * @generated from field: int32 addToCartCount = 4;
   */
  addToCartCount = 0;

  /**
   * Заказли товаров, штук
   *
   * @generated from field: int32 ordersCount = 5;
   */
  ordersCount = 0;

  /**
   * Заказали на сумму, руб.
   *
   * @generated from field: float ordersSumRub = 6;
   */
  ordersSumRub = 0;

  /**
   * Выкупили товаров, штук
   *
   * @generated from field: int32 buyoutsCount = 7;
   */
  buyoutsCount = 0;

  /**
   * Выкупили на сумму, руб.
   *
   * @generated from field: float buyoutsSumRub = 8;
   */
  buyoutsSumRub = 0;

  /**
   * Отменили товаров, штук
   *
   * @generated from field: int32 cancelCount = 9;
   */
  cancelCount = 0;

  /**
   * Отменили на сумму, руб
   *
   * @generated from field: int32 cancelSumRub = 10;
   */
  cancelSumRub = 0;

  /**
   * Средняя цена, руб.
   *
   * @generated from field: float avgPriceRub = 11;
   */
  avgPriceRub = 0;

  /**
   * Среднее количество заказов в день, шт.
   *
   * @generated from field: float avgOrdersCountPerDay = 12;
   */
  avgOrdersCountPerDay = 0;

  /**
   * @generated from field: wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions conversions = 13;
   */
  conversions?: NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions;

  constructor(data?: PartialMessage<NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "begin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "end", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "openCardCount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "addToCartCount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "ordersCount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "ordersSumRub", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 7, name: "buyoutsCount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "buyoutsSumRub", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 9, name: "cancelCount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "cancelSumRub", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 11, name: "avgPriceRub", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 12, name: "avgOrdersCountPerDay", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 13, name: "conversions", kind: "message", T: NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriodConversions },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod {
    return new NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod {
    return new NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod {
    return new NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod().fromJsonString(jsonString, options);
  }

  static equals(a: NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod | PlainMessage<NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod> | undefined, b: NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod | PlainMessage<NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod> | undefined): boolean {
    return proto3.util.equals(NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod, a, b);
  }
}

