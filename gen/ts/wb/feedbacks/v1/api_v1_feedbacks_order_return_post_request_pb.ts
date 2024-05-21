//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/api_v1_feedbacks_order_return_post_request.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.feedbacks.v1.ApiV1FeedbacksOrderReturnPostReq
 */
export class ApiV1FeedbacksOrderReturnPostReq extends Message<ApiV1FeedbacksOrderReturnPostReq> {
  /**
   * Идентификатор отзыва
   *
   * @generated from field: string feedbackId = 1;
   */
  feedbackId = "";

  constructor(data?: PartialMessage<ApiV1FeedbacksOrderReturnPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.ApiV1FeedbacksOrderReturnPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "feedbackId", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1FeedbacksOrderReturnPostReq {
    return new ApiV1FeedbacksOrderReturnPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1FeedbacksOrderReturnPostReq {
    return new ApiV1FeedbacksOrderReturnPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1FeedbacksOrderReturnPostReq {
    return new ApiV1FeedbacksOrderReturnPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1FeedbacksOrderReturnPostReq | PlainMessage<ApiV1FeedbacksOrderReturnPostReq> | undefined, b: ApiV1FeedbacksOrderReturnPostReq | PlainMessage<ApiV1FeedbacksOrderReturnPostReq> | undefined): boolean {
    return proto3.util.equals(ApiV1FeedbacksOrderReturnPostReq, a, b);
  }
}
