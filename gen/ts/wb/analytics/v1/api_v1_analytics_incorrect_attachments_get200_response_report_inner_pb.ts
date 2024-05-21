//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/api_v1_analytics_incorrect_attachments_get200_response_report_inner.proto (package wb.analytics.v1, syntax proto3)
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
 * @generated from message wb.analytics.v1.ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner
 */
export class ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner extends Message<ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner> {
  /**
   * Цена, ₽
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
   * Причина удержания
   *
   * @generated from field: string lostReason = 3;
   */
  lostReason = "";

  /**
   * Артикул Wildberries
   *
   * @generated from field: int32 nmID = 4;
   */
  nmID = 0;

  /**
   * Фото
   *
   * @generated from field: string photoUrl = 5;
   */
  photoUrl = "";

  /**
   * Штрихкод
   *
   * @generated from field: int32 shkID = 6;
   */
  shkID = 0;

  constructor(data?: PartialMessage<ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "amount", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 2, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "lostReason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "photoUrl", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "shkID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner {
    return new ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner {
    return new ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner {
    return new ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner> | undefined, b: ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner | PlainMessage<ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsIncorrectAttachmentsGet200ResponseReportInner, a, b);
  }
}

