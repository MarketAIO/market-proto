//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_get_cards_trash_post200_response_cards_inner_sizes_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner
 */
export class ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner extends Message<ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner> {
  /**
   * ID размера
   *
   * @generated from field: int32 chrtID = 1;
   */
  chrtID = 0;

  /**
   * Размер товара
   *
   * @generated from field: string techSize = 2;
   */
  techSize = "";

  /**
   * Российский размер товара
   *
   * @generated from field: string wbSize = 3;
   */
  wbSize = "";

  /**
   * Массив баркодов
   *
   * @generated from field: repeated string skus = 4;
   */
  skus: string[] = [];

  constructor(data?: PartialMessage<ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chrtID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "techSize", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "wbSize", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "skus", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner {
    return new ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner {
    return new ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner {
    return new ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner | PlainMessage<ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner> | undefined, b: ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner | PlainMessage<ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner> | undefined): boolean {
    return proto3.util.equals(ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner, a, b);
  }
}

