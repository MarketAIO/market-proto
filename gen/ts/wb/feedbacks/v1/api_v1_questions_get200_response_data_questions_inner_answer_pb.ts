//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/api_v1_questions_get200_response_data_questions_inner_answer.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.feedbacks.v1.ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer
 */
export class ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer extends Message<ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer> {
  /**
   * Текст ответа
   *
   * @generated from field: string text = 1;
   */
  text = "";

  /**
   * Можно ли отредактировать ответ (`false` - нельзя, `true` - можно)
   *
   * @generated from field: bool editable = 2;
   */
  editable = false;

  /**
   * Дата и время создания ответа
   *
   * @generated from field: string createDate = 3;
   */
  createDate = "";

  constructor(data?: PartialMessage<ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "editable", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "createDate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer {
    return new ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer {
    return new ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer {
    return new ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer | PlainMessage<ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer> | undefined, b: ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer | PlainMessage<ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer> | undefined): boolean {
    return proto3.util.equals(ApiV1QuestionsGet200ResponseDataQuestionsInnerAnswer, a, b);
  }
}
