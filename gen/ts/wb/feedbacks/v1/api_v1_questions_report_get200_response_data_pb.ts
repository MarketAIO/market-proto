//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/api_v1_questions_report_get200_response_data.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.feedbacks.v1.ApiV1QuestionsReportGet200ResponseData
 */
export class ApiV1QuestionsReportGet200ResponseData extends Message<ApiV1QuestionsReportGet200ResponseData> {
  /**
   * Имя файла
   *
   * @generated from field: string fileName = 1;
   */
  fileName = "";

  /**
   * Файл
   *
   * @generated from field: string file = 2;
   */
  file = "";

  /**
   * Тип контента
   *
   * @generated from field: string contentType = 3;
   */
  contentType = "";

  constructor(data?: PartialMessage<ApiV1QuestionsReportGet200ResponseData>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.ApiV1QuestionsReportGet200ResponseData";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "fileName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "file", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "contentType", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1QuestionsReportGet200ResponseData {
    return new ApiV1QuestionsReportGet200ResponseData().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1QuestionsReportGet200ResponseData {
    return new ApiV1QuestionsReportGet200ResponseData().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1QuestionsReportGet200ResponseData {
    return new ApiV1QuestionsReportGet200ResponseData().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1QuestionsReportGet200ResponseData | PlainMessage<ApiV1QuestionsReportGet200ResponseData> | undefined, b: ApiV1QuestionsReportGet200ResponseData | PlainMessage<ApiV1QuestionsReportGet200ResponseData> | undefined): boolean {
    return proto3.util.equals(ApiV1QuestionsReportGet200ResponseData, a, b);
  }
}

