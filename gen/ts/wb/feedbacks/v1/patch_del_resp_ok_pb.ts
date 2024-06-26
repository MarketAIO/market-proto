//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/patch_del_resp_ok.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.feedbacks.v1.PatchDelRespOK
 */
export class PatchDelRespOK extends Message<PatchDelRespOK> {
  /**
   * @generated from field: optional bool data = 1;
   */
  data?: boolean;

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

  constructor(data?: PartialMessage<PatchDelRespOK>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.PatchDelRespOK";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 2, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "errorText", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "additionalErrors", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PatchDelRespOK {
    return new PatchDelRespOK().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PatchDelRespOK {
    return new PatchDelRespOK().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PatchDelRespOK {
    return new PatchDelRespOK().fromJsonString(jsonString, options);
  }

  static equals(a: PatchDelRespOK | PlainMessage<PatchDelRespOK> | undefined, b: PatchDelRespOK | PlainMessage<PatchDelRespOK> | undefined): boolean {
    return proto3.util.equals(PatchDelRespOK, a, b);
  }
}

