//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/api_v1_analytics_acceptance_report_get200_response_report_inner.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner
 */
export class ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner extends Message<ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner> {
  /**
   * Количество товаров, шт.
   *
   * @generated from field: int32 count = 1;
   */
  count = 0;

  /**
   * Дата создания поставки
   *
   * @generated from field: string giCreateDate = 2;
   */
  giCreateDate = "";

  /**
   * Номер поставки
   *
   * @generated from field: int32 incomeId = 3;
   */
  incomeId = 0;

  /**
   * Артикул Wildberries
   *
   * @generated from field: int32 nmID = 4;
   */
  nmID = 0;

  /**
   * Дата приёмки
   *
   * @generated from field: string createDate = 5;
   */
  createDate = "";

  /**
   * Предмет (подкатегория)
   *
   * @generated from field: string subjectName = 6;
   */
  subjectName = "";

  /**
   * Суммарная стоимость приёмки, ₽
   *
   * @generated from field: int32 sum = 7;
   */
  sum = 0;

  /**
   * Суммарная стоимость приёмки, ₽ с копейками
   *
   * @generated from field: float total = 8;
   */
  total = 0;

  constructor(data?: PartialMessage<ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "giCreateDate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "incomeId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "createDate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "subjectName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "sum", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "total", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner {
    return new ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner {
    return new ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner {
    return new ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner> | undefined, b: ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsAcceptanceReportGet200ResponseReportInner, a, b);
  }
}

