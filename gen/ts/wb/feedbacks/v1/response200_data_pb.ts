//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/response200_data.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Response200DataTemplatesInner } from "./response200_data_templates_inner_pb.js";

/**
 * @generated from message wb.feedbacks.v1.Response200Data
 */
export class Response200Data extends Message<Response200Data> {
  /**
   * @generated from field: repeated wb.feedbacks.v1.Response200DataTemplatesInner templates = 1;
   */
  templates: Response200DataTemplatesInner[] = [];

  constructor(data?: PartialMessage<Response200Data>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.Response200Data";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "templates", kind: "message", T: Response200DataTemplatesInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Response200Data {
    return new Response200Data().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Response200Data {
    return new Response200Data().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Response200Data {
    return new Response200Data().fromJsonString(jsonString, options);
  }

  static equals(a: Response200Data | PlainMessage<Response200Data> | undefined, b: Response200Data | PlainMessage<Response200Data> | undefined): boolean {
    return proto3.util.equals(Response200Data, a, b);
  }
}

