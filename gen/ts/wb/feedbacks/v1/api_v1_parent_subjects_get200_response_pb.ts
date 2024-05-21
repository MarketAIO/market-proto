//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/api_v1_parent_subjects_get200_response.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ApiV1ParentSubjectsGet200ResponseDataInner } from "./api_v1_parent_subjects_get200_response_data_inner_pb.js";

/**
 * @generated from message wb.feedbacks.v1.ApiV1ParentSubjectsGet200Response
 */
export class ApiV1ParentSubjectsGet200Response extends Message<ApiV1ParentSubjectsGet200Response> {
  /**
   * @generated from field: repeated wb.feedbacks.v1.ApiV1ParentSubjectsGet200ResponseDataInner data = 1;
   */
  data: ApiV1ParentSubjectsGet200ResponseDataInner[] = [];

  /**
   * Есть ли ошибка
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
   * @generated from field: repeated string additionalErrors = 4;
   */
  additionalErrors: string[] = [];

  constructor(data?: PartialMessage<ApiV1ParentSubjectsGet200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.ApiV1ParentSubjectsGet200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: ApiV1ParentSubjectsGet200ResponseDataInner, repeated: true },
    { no: 2, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "errorText", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "additionalErrors", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1ParentSubjectsGet200Response {
    return new ApiV1ParentSubjectsGet200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1ParentSubjectsGet200Response {
    return new ApiV1ParentSubjectsGet200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1ParentSubjectsGet200Response {
    return new ApiV1ParentSubjectsGet200Response().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1ParentSubjectsGet200Response | PlainMessage<ApiV1ParentSubjectsGet200Response> | undefined, b: ApiV1ParentSubjectsGet200Response | PlainMessage<ApiV1ParentSubjectsGet200Response> | undefined): boolean {
    return proto3.util.equals(ApiV1ParentSubjectsGet200Response, a, b);
  }
}
