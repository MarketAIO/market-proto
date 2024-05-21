//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/nm_report_grouped_response.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { NmReportGroupedResponseData } from "./nm_report_grouped_response_data_pb.js";
import { ResponseErrorAdditionalErrorsInner } from "./response_error_additional_errors_inner_pb.js";

/**
 * @generated from message wb.analytics.v1.NmReportGroupedResponse
 */
export class NmReportGroupedResponse extends Message<NmReportGroupedResponse> {
  /**
   * @generated from field: wb.analytics.v1.NmReportGroupedResponseData data = 1;
   */
  data?: NmReportGroupedResponseData;

  /**
   * Флаг ошибки
   *
   * @generated from field: bool error = 2;
   */
  error = false;

  /**
   * Описание ошибки
   *
   * @generated from field: string errorText = 3;
   */
  errorText = "";

  /**
   * Дополнительные ошибки
   *
   * @generated from field: repeated wb.analytics.v1.ResponseErrorAdditionalErrorsInner additionalErrors = 4;
   */
  additionalErrors: ResponseErrorAdditionalErrorsInner[] = [];

  constructor(data?: PartialMessage<NmReportGroupedResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.NmReportGroupedResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: NmReportGroupedResponseData },
    { no: 2, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "errorText", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "additionalErrors", kind: "message", T: ResponseErrorAdditionalErrorsInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NmReportGroupedResponse {
    return new NmReportGroupedResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NmReportGroupedResponse {
    return new NmReportGroupedResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NmReportGroupedResponse {
    return new NmReportGroupedResponse().fromJsonString(jsonString, options);
  }

  static equals(a: NmReportGroupedResponse | PlainMessage<NmReportGroupedResponse> | undefined, b: NmReportGroupedResponse | PlainMessage<NmReportGroupedResponse> | undefined): boolean {
    return proto3.util.equals(NmReportGroupedResponse, a, b);
  }
}

