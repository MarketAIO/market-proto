//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/api_v1_analytics_goods_labeling_get200_response_report_inner.proto (package wb.analytics.v1, syntax proto3)
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

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner
 */
export class ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner extends Message<ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner> {
  /**
   * Сумма штрафа, руб
   *
   * @generated from field: float amount = 1;
   */
  amount = 0;

  /**
   * Дата
   *
   * @generated from field: string date = 2;
   */
  date = "";

  /**
   * Номер поставки
   *
   * @generated from field: int32 incomeId = 3;
   */
  incomeId = 0;

  /**
   * Артикул WB
   *
   * @generated from field: int32 nmID = 4;
   */
  nmID = 0;

  /**
   * URL фото товара
   *
   * @generated from field: repeated string photoUrls = 5;
   */
  photoUrls: string[] = [];

  /**
   * Штрихкод товара в Wildberries
   *
   * @generated from field: int32 shkID = 6;
   */
  shkID = 0;

  /**
   * Баркод из карточки товара
   *
   * @generated from field: string sku = 7;
   */
  sku = "";

  constructor(data?: PartialMessage<ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "amount", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "incomeId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "photoUrls", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "shkID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "sku", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner {
    return new ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner {
    return new ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner {
    return new ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner> | undefined, b: ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsGoodsLabelingGet200ResponseReportInner, a, b);
  }
}

