//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/excise_report_response.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ModelsExciseReportResponse } from "./models_excise_report_response_pb.js";

/**
 * @generated from message wb.analytics.v1.ExciseReportResponse
 */
export class ExciseReportResponse extends Message<ExciseReportResponse> {
  /**
   * @generated from field: wb.analytics.v1.ModelsExciseReportResponse response = 1;
   */
  response?: ModelsExciseReportResponse;

  constructor(data?: PartialMessage<ExciseReportResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ExciseReportResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "response", kind: "message", T: ModelsExciseReportResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExciseReportResponse {
    return new ExciseReportResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExciseReportResponse {
    return new ExciseReportResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExciseReportResponse {
    return new ExciseReportResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ExciseReportResponse | PlainMessage<ExciseReportResponse> | undefined, b: ExciseReportResponse | PlainMessage<ExciseReportResponse> | undefined): boolean {
    return proto3.util.equals(ExciseReportResponse, a, b);
  }
}

