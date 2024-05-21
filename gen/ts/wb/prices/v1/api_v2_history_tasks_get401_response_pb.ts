//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/prices/v1/api_v2_history_tasks_get401_response.proto (package wb.prices.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.prices.v1.ApiV2HistoryTasksGet401Response
 */
export class ApiV2HistoryTasksGet401Response extends Message<ApiV2HistoryTasksGet401Response> {
  /**
   * @generated from field: google.protobuf.Any data = 1;
   */
  data?: Any;

  /**
   * @generated from field: bool error = 2;
   */
  error = false;

  /**
   * @generated from field: string errorText = 3;
   */
  errorText = "";

  constructor(data?: PartialMessage<ApiV2HistoryTasksGet401Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.prices.v1.ApiV2HistoryTasksGet401Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Any },
    { no: 2, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "errorText", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2HistoryTasksGet401Response {
    return new ApiV2HistoryTasksGet401Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2HistoryTasksGet401Response {
    return new ApiV2HistoryTasksGet401Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2HistoryTasksGet401Response {
    return new ApiV2HistoryTasksGet401Response().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2HistoryTasksGet401Response | PlainMessage<ApiV2HistoryTasksGet401Response> | undefined, b: ApiV2HistoryTasksGet401Response | PlainMessage<ApiV2HistoryTasksGet401Response> | undefined): boolean {
    return proto3.util.equals(ApiV2HistoryTasksGet401Response, a, b);
  }
}

