//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/api_v1_analytics_storage_coefficient_get200_response_report_inner.proto (package wb.analytics.v1, syntax proto3)
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
 * @generated from message wb.analytics.v1.ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner
 */
export class ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner extends Message<ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner> {
  /**
   * Фактическая высота, см
   *
   * @generated from field: float actualHeight = 1;
   */
  actualHeight = 0;

  /**
   * Фактическая длина, см
   *
   * @generated from field: float actualLength = 2;
   */
  actualLength = 0;

  /**
   * Фактический объём, л
   *
   * @generated from field: float actualVolume = 3;
   */
  actualVolume = 0;

  /**
   * Фактическая ширина, см
   *
   * @generated from field: float actualWidth = 4;
   */
  actualWidth = 0;

  /**
   * Дата измерения в формате RFC 3339. Для расчёта используются измерения за 30 дней, до начала отчётной недели
   *
   * @generated from field: string date = 5;
   */
  date = "";

  /**
   * Разница в габаритах, %
   *
   * @generated from field: float dimensionDifference = 6;
   */
  dimensionDifference = 0;

  /**
   * Высота из карточки, см
   *
   * @generated from field: float height = 7;
   */
  height = 0;

  /**
   * Длина из карточки, см
   *
   * @generated from field: float length = 8;
   */
  length = 0;

  /**
   * Коэффициент логистики и хранения для товара
   *
   * @generated from field: float logWarehouseCoef = 9;
   */
  logWarehouseCoef = 0;

  /**
   * Артикул Wildberries
   *
   * @generated from field: int32 nmID = 10;
   */
  nmID = 0;

  /**
   * Фото измерений
   *
   * @generated from field: string photoUrls = 11;
   */
  photoUrls = "";

  /**
   * Наименование товара
   *
   * @generated from field: string title = 12;
   */
  title = "";

  /**
   * Объём из карточки, л
   *
   * @generated from field: float volume = 13;
   */
  volume = 0;

  /**
   * Ширина из карточки, см
   *
   * @generated from field: float width = 14;
   */
  width = 0;

  constructor(data?: PartialMessage<ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "actualHeight", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "actualLength", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 3, name: "actualVolume", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "actualWidth", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "dimensionDifference", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 7, name: "height", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 8, name: "length", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 9, name: "logWarehouseCoef", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 10, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 11, name: "photoUrls", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 13, name: "volume", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 14, name: "width", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner {
    return new ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner {
    return new ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner {
    return new ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner> | undefined, b: ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsStorageCoefficientGet200ResponseReportInner, a, b);
  }
}
