//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/api_v1_feedbacks_count_unanswered_get200_response_data.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.feedbacks.v1.ApiV1FeedbacksCountUnansweredGet200ResponseData
 */
export class ApiV1FeedbacksCountUnansweredGet200ResponseData extends Message<ApiV1FeedbacksCountUnansweredGet200ResponseData> {
  /**
   * Количество необработанных отзывов
   *
   * @generated from field: int32 countUnanswered = 1;
   */
  countUnanswered = 0;

  /**
   * Количество необработанных отзывов за сегодня
   *
   * @generated from field: int32 countUnansweredToday = 2;
   */
  countUnansweredToday = 0;

  /**
   * Средняя оценка всех отзывов
   *
   * @generated from field: string valuation = 3;
   */
  valuation = "";

  constructor(data?: PartialMessage<ApiV1FeedbacksCountUnansweredGet200ResponseData>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.ApiV1FeedbacksCountUnansweredGet200ResponseData";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "countUnanswered", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "countUnansweredToday", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "valuation", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1FeedbacksCountUnansweredGet200ResponseData {
    return new ApiV1FeedbacksCountUnansweredGet200ResponseData().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1FeedbacksCountUnansweredGet200ResponseData {
    return new ApiV1FeedbacksCountUnansweredGet200ResponseData().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1FeedbacksCountUnansweredGet200ResponseData {
    return new ApiV1FeedbacksCountUnansweredGet200ResponseData().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1FeedbacksCountUnansweredGet200ResponseData | PlainMessage<ApiV1FeedbacksCountUnansweredGet200ResponseData> | undefined, b: ApiV1FeedbacksCountUnansweredGet200ResponseData | PlainMessage<ApiV1FeedbacksCountUnansweredGet200ResponseData> | undefined): boolean {
    return proto3.util.equals(ApiV1FeedbacksCountUnansweredGet200ResponseData, a, b);
  }
}
