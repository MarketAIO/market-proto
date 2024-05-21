//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_get_cards_trash_post200_response_cursor.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.content.v1.ContentV2GetCardsTrashPost200ResponseCursor
 */
export class ContentV2GetCardsTrashPost200ResponseCursor extends Message<ContentV2GetCardsTrashPost200ResponseCursor> {
  /**
   * @generated from field: string trashedAt = 1;
   */
  trashedAt = "";

  /**
   * @generated from field: int32 nmID = 2;
   */
  nmID = 0;

  /**
   * @generated from field: int32 total = 3;
   */
  total = 0;

  constructor(data?: PartialMessage<ContentV2GetCardsTrashPost200ResponseCursor>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2GetCardsTrashPost200ResponseCursor";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "trashedAt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "total", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2GetCardsTrashPost200ResponseCursor {
    return new ContentV2GetCardsTrashPost200ResponseCursor().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPost200ResponseCursor {
    return new ContentV2GetCardsTrashPost200ResponseCursor().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPost200ResponseCursor {
    return new ContentV2GetCardsTrashPost200ResponseCursor().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2GetCardsTrashPost200ResponseCursor | PlainMessage<ContentV2GetCardsTrashPost200ResponseCursor> | undefined, b: ContentV2GetCardsTrashPost200ResponseCursor | PlainMessage<ContentV2GetCardsTrashPost200ResponseCursor> | undefined): boolean {
    return proto3.util.equals(ContentV2GetCardsTrashPost200ResponseCursor, a, b);
  }
}

