//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/api_v1_analytics_antifraud_details_get200_response.proto (package wb.analytics.v1, syntax proto3)
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
    ApiV1AnalyticsAntifraudDetailsGet200ResponseDetailsInner
} from "./api_v1_analytics_antifraud_details_get200_response_details_inner_pb.js";

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsAntifraudDetailsGet200Response
 */
export class ApiV1AnalyticsAntifraudDetailsGet200Response extends Message<ApiV1AnalyticsAntifraudDetailsGet200Response> {
  /**
   * @generated from field: repeated wb.analytics.v1.ApiV1AnalyticsAntifraudDetailsGet200ResponseDetailsInner details = 1;
   */
  details: ApiV1AnalyticsAntifraudDetailsGet200ResponseDetailsInner[] = [];

  constructor(data?: PartialMessage<ApiV1AnalyticsAntifraudDetailsGet200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsAntifraudDetailsGet200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "details", kind: "message", T: ApiV1AnalyticsAntifraudDetailsGet200ResponseDetailsInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsAntifraudDetailsGet200Response {
    return new ApiV1AnalyticsAntifraudDetailsGet200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAntifraudDetailsGet200Response {
    return new ApiV1AnalyticsAntifraudDetailsGet200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAntifraudDetailsGet200Response {
    return new ApiV1AnalyticsAntifraudDetailsGet200Response().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsAntifraudDetailsGet200Response | PlainMessage<ApiV1AnalyticsAntifraudDetailsGet200Response> | undefined, b: ApiV1AnalyticsAntifraudDetailsGet200Response | PlainMessage<ApiV1AnalyticsAntifraudDetailsGet200Response> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsAntifraudDetailsGet200Response, a, b);
  }
}
