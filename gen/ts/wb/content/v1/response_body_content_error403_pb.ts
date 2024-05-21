//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/response_body_content_error403.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.content.v1.ResponseBodyContentError403
 */
export class ResponseBodyContentError403 extends Message<ResponseBodyContentError403> {
  /**
   * @generated from field: google.protobuf.Any data = 1;
   */
  data?: Any;

  /**
   * Флаг ошибки
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
   * @generated from field: string additionalErrors = 4;
   */
  additionalErrors = "";

  constructor(data?: PartialMessage<ResponseBodyContentError403>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ResponseBodyContentError403";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Any },
    { no: 2, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "errorText", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "additionalErrors", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseBodyContentError403 {
    return new ResponseBodyContentError403().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseBodyContentError403 {
    return new ResponseBodyContentError403().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseBodyContentError403 {
    return new ResponseBodyContentError403().fromJsonString(jsonString, options);
  }

  static equals(a: ResponseBodyContentError403 | PlainMessage<ResponseBodyContentError403> | undefined, b: ResponseBodyContentError403 | PlainMessage<ResponseBodyContentError403> | undefined): boolean {
    return proto3.util.equals(ResponseBodyContentError403, a, b);
  }
}

