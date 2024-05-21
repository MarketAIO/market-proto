//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/api_v1_questions_patch_request_one_of.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.feedbacks.v1.ApiV1QuestionsPatchReqOneOf
 */
export class ApiV1QuestionsPatchReqOneOf extends Message<ApiV1QuestionsPatchReqOneOf> {
  /**
   * Id вопроса
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Просмотрен (`true`), не просмотрен (`false`)
   *
   * @generated from field: bool wasViewed = 2;
   */
  wasViewed = false;

  constructor(data?: PartialMessage<ApiV1QuestionsPatchReqOneOf>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.ApiV1QuestionsPatchReqOneOf";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "wasViewed", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1QuestionsPatchReqOneOf {
    return new ApiV1QuestionsPatchReqOneOf().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1QuestionsPatchReqOneOf {
    return new ApiV1QuestionsPatchReqOneOf().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1QuestionsPatchReqOneOf {
    return new ApiV1QuestionsPatchReqOneOf().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1QuestionsPatchReqOneOf | PlainMessage<ApiV1QuestionsPatchReqOneOf> | undefined, b: ApiV1QuestionsPatchReqOneOf | PlainMessage<ApiV1QuestionsPatchReqOneOf> | undefined): boolean {
    return proto3.util.equals(ApiV1QuestionsPatchReqOneOf, a, b);
  }
}

