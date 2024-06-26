//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_cards_update_post_request_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ContentV2CardsUpdatePostReqInnerDimensions } from "./content_v2_cards_update_post_request_inner_dimensions_pb.js";
import { ContentV2CardsUpdatePostReqInnerCharacteristicsInner } from "./content_v2_cards_update_post_request_inner_characteristics_inner_pb.js";
import { ContentV2CardsUpdatePostReqInnerSizesInner } from "./content_v2_cards_update_post_request_inner_sizes_inner_pb.js";

/**
 * @generated from message wb.content.v1.ContentV2CardsUpdatePostReqInner
 */
export class ContentV2CardsUpdatePostReqInner extends Message<ContentV2CardsUpdatePostReqInner> {
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
   * Бренд
   *
   * @generated from field: string brand = 3;
   */
  brand = "";

  /**
   * Наименование товара
   *
   * @generated from field: string title = 4;
   */
  title = "";

  /**
   * Описание товара. Максимальное количество символов зависит от категории товара. Стандарт — 2000, минимум — 1000, максимум — 5000.<br> Подробно о правилах описания в **Правилах заполнения карточки товара** в разделе [Инструкции](https://seller.wildberries.ru/training) на портале продавцов. 
   *
   * @generated from field: string description = 5;
   */
  description = "";

  /**
   * @generated from field: wb.content.v1.ContentV2CardsUpdatePostReqInnerDimensions dimensions = 6;
   */
  dimensions?: ContentV2CardsUpdatePostReqInnerDimensions;

  /**
   * Характеристики товара
   *
   * @generated from field: repeated wb.content.v1.ContentV2CardsUpdatePostReqInnerCharacteristicsInner characteristics = 7;
   */
  characteristics: ContentV2CardsUpdatePostReqInnerCharacteristicsInner[] = [];

  /**
   * Массив размеров артикула. <br> Для безразмерного товара все равно нужно передавать данный массив без параметров (wbSize и techSize), но с баркодом.                           
   *
   * @generated from field: repeated wb.content.v1.ContentV2CardsUpdatePostReqInnerSizesInner sizes = 8;
   */
  sizes: ContentV2CardsUpdatePostReqInnerSizesInner[] = [];

  constructor(data?: PartialMessage<ContentV2CardsUpdatePostReqInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2CardsUpdatePostReqInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "vendorCode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "brand", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "dimensions", kind: "message", T: ContentV2CardsUpdatePostReqInnerDimensions },
    { no: 7, name: "characteristics", kind: "message", T: ContentV2CardsUpdatePostReqInnerCharacteristicsInner, repeated: true },
    { no: 8, name: "sizes", kind: "message", T: ContentV2CardsUpdatePostReqInnerSizesInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2CardsUpdatePostReqInner {
    return new ContentV2CardsUpdatePostReqInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2CardsUpdatePostReqInner {
    return new ContentV2CardsUpdatePostReqInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2CardsUpdatePostReqInner {
    return new ContentV2CardsUpdatePostReqInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2CardsUpdatePostReqInner | PlainMessage<ContentV2CardsUpdatePostReqInner> | undefined, b: ContentV2CardsUpdatePostReqInner | PlainMessage<ContentV2CardsUpdatePostReqInner> | undefined): boolean {
    return proto3.util.equals(ContentV2CardsUpdatePostReqInner, a, b);
  }
}

