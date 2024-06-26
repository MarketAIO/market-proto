//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_get_cards_trash_post200_response_cards_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner } from "./content_v2_get_cards_list_post200_response_cards_inner_photos_inner_pb.js";
import { ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner } from "./content_v2_get_cards_trash_post200_response_cards_inner_sizes_inner_pb.js";
import { ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions } from "./content_v2_get_cards_trash_post200_response_cards_inner_dimensions_pb.js";
import { ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner } from "./content_v2_get_cards_trash_post200_response_cards_inner_characteristics_inner_pb.js";

/**
 * @generated from message wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInner
 */
export class ContentV2GetCardsTrashPost200ResponseCardsInner extends Message<ContentV2GetCardsTrashPost200ResponseCardsInner> {
  /**
   * Артикул WB
   *
   * @generated from field: int32 nmID = 1;
   */
  nmID = 0;

  /**
   * Артикул продавца
   *
   * @generated from field: string vendorCode = 2;
   */
  vendorCode = "";

  /**
   * Идентификатор предмета
   *
   * @generated from field: string subjectID = 3;
   */
  subjectID = "";

  /**
   * Название предмета
   *
   * @generated from field: string subjectName = 4;
   */
  subjectName = "";

  /**
   * Массив фото
   *
   * @generated from field: repeated wb.content.v1.ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner photos = 5;
   */
  photos: ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner[] = [];

  /**
   * URL видео
   *
   * @generated from field: string video = 6;
   */
  video = "";

  /**
   * Массив размеров
   *
   * @generated from field: repeated wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner sizes = 7;
   */
  sizes: ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner[] = [];

  /**
   * @generated from field: wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions dimensions = 8;
   */
  dimensions?: ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions;

  /**
   * Массив характеристик, при наличии
   *
   * @generated from field: repeated wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner characteristics = 9;
   */
  characteristics: ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner[] = [];

  /**
   * @generated from field: string createdAt = 10;
   */
  createdAt = "";

  /**
   * @generated from field: string trashedAt = 11;
   */
  trashedAt = "";

  constructor(data?: PartialMessage<ContentV2GetCardsTrashPost200ResponseCardsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2GetCardsTrashPost200ResponseCardsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "vendorCode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "subjectID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "subjectName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "photos", kind: "message", T: ContentV2GetCardsListPost200ResponseCardsInnerPhotosInner, repeated: true },
    { no: 6, name: "video", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "sizes", kind: "message", T: ContentV2GetCardsTrashPost200ResponseCardsInnerSizesInner, repeated: true },
    { no: 8, name: "dimensions", kind: "message", T: ContentV2GetCardsTrashPost200ResponseCardsInnerDimensions },
    { no: 9, name: "characteristics", kind: "message", T: ContentV2GetCardsTrashPost200ResponseCardsInnerCharacteristicsInner, repeated: true },
    { no: 10, name: "createdAt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "trashedAt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2GetCardsTrashPost200ResponseCardsInner {
    return new ContentV2GetCardsTrashPost200ResponseCardsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPost200ResponseCardsInner {
    return new ContentV2GetCardsTrashPost200ResponseCardsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPost200ResponseCardsInner {
    return new ContentV2GetCardsTrashPost200ResponseCardsInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2GetCardsTrashPost200ResponseCardsInner | PlainMessage<ContentV2GetCardsTrashPost200ResponseCardsInner> | undefined, b: ContentV2GetCardsTrashPost200ResponseCardsInner | PlainMessage<ContentV2GetCardsTrashPost200ResponseCardsInner> | undefined): boolean {
    return proto3.util.equals(ContentV2GetCardsTrashPost200ResponseCardsInner, a, b);
  }
}

