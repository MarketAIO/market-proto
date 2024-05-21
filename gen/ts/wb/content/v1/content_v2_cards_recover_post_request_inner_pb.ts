//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_cards_recover_post_request_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.content.v1.ContentV2CardsRecoverPostReqInner
 */
export class ContentV2CardsRecoverPostReqInner extends Message<ContentV2CardsRecoverPostReqInner> {
  /**
   * Артикул WB (max. 1000)
   *
   * @generated from field: int32 nmIDs = 1;
   */
  nmIDs = 0;

  constructor(data?: PartialMessage<ContentV2CardsRecoverPostReqInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2CardsRecoverPostReqInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmIDs", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2CardsRecoverPostReqInner {
    return new ContentV2CardsRecoverPostReqInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2CardsRecoverPostReqInner {
    return new ContentV2CardsRecoverPostReqInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2CardsRecoverPostReqInner {
    return new ContentV2CardsRecoverPostReqInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2CardsRecoverPostReqInner | PlainMessage<ContentV2CardsRecoverPostReqInner> | undefined, b: ContentV2CardsRecoverPostReqInner | PlainMessage<ContentV2CardsRecoverPostReqInner> | undefined): boolean {
    return proto3.util.equals(ContentV2CardsRecoverPostReqInner, a, b);
  }
}
