//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/create_task_response.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { CreateTaskResponseData } from "./create_task_response_data_pb.js";

/**
 * @generated from message wb.analytics.v1.CreateTaskResponse
 */
export class CreateTaskResponse extends Message<CreateTaskResponse> {
  /**
   * @generated from field: wb.analytics.v1.CreateTaskResponseData data = 1;
   */
  data?: CreateTaskResponseData;

  constructor(data?: PartialMessage<CreateTaskResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.CreateTaskResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: CreateTaskResponseData },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTaskResponse {
    return new CreateTaskResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTaskResponse | PlainMessage<CreateTaskResponse> | undefined, b: CreateTaskResponse | PlainMessage<CreateTaskResponse> | undefined): boolean {
    return proto3.util.equals(CreateTaskResponse, a, b);
  }
}
