//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/response_content_error1_additional_errors.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.content.v1.ResponseContentError1AdditionalErrors
 */
export class ResponseContentError1AdditionalErrors extends Message<ResponseContentError1AdditionalErrors> {
  /**
   * @generated from field: string error = 1;
   */
  error = "";

  constructor(data?: PartialMessage<ResponseContentError1AdditionalErrors>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ResponseContentError1AdditionalErrors";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseContentError1AdditionalErrors {
    return new ResponseContentError1AdditionalErrors().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseContentError1AdditionalErrors {
    return new ResponseContentError1AdditionalErrors().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseContentError1AdditionalErrors {
    return new ResponseContentError1AdditionalErrors().fromJsonString(jsonString, options);
  }

  static equals(a: ResponseContentError1AdditionalErrors | PlainMessage<ResponseContentError1AdditionalErrors> | undefined, b: ResponseContentError1AdditionalErrors | PlainMessage<ResponseContentError1AdditionalErrors> | undefined): boolean {
    return proto3.util.equals(ResponseContentError1AdditionalErrors, a, b);
  }
}
