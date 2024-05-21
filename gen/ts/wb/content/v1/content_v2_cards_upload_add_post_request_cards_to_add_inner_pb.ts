//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_cards_upload_add_post_request_cards_to_add_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ContentV2CardsUploadPostReqInnerVariantsInnerDimensions } from "./content_v2_cards_upload_post_request_inner_variants_inner_dimensions_pb.js";
import { ContentV2CardsUploadAddPostReqCardsToAddInnerCharacteristicsInner } from "./content_v2_cards_upload_add_post_request_cards_to_add_inner_characteristics_inner_pb.js";
import { ContentV2CardsUploadAddPostReqCardsToAddInnerSizesInner } from "./content_v2_cards_upload_add_post_request_cards_to_add_inner_sizes_inner_pb.js";

/**
 * @generated from message wb.content.v1.ContentV2CardsUploadAddPostReqCardsToAddInner
 */
export class ContentV2CardsUploadAddPostReqCardsToAddInner extends Message<ContentV2CardsUploadAddPostReqCardsToAddInner> {
  /**
   * Бренд
   *
   * @generated from field: string brand = 1;
   */
  brand = "";

  /**
   * Артикул продавца
   *
   * @generated from field: string vendorCode = 2;
   */
  vendorCode = "";

  /**
   * Наименование товара
   *
   * @generated from field: string title = 3;
   */
  title = "";

  /**
   * Описание товара. Максимальное количество символов зависит от категории товара. Стандарт — 2000, минимум — 1000, максимум — 5000.<br> Подробно о правилах описания в **Правилах заполнения карточки товара** в разделе [Инструкции](https://seller.wildberries.ru/training) на портале продавцов. 
   *
   * @generated from field: string description = 4;
   */
  description = "";

  /**
   * @generated from field: wb.content.v1.ContentV2CardsUploadPostReqInnerVariantsInnerDimensions dimensions = 5;
   */
  dimensions?: ContentV2CardsUploadPostReqInnerVariantsInnerDimensions;

  /**
   * Характеристики товара
   *
   * @generated from field: repeated wb.content.v1.ContentV2CardsUploadAddPostReqCardsToAddInnerCharacteristicsInner characteristics = 6;
   */
  characteristics: ContentV2CardsUploadAddPostReqCardsToAddInnerCharacteristicsInner[] = [];

  /**
   * Массив с размерами. <br>  Если для размерного товара (обувь, одежда и др.) не указать этот массив, то системой в карточке он будет сгенерирован автоматически с `techSize` = \"A\" и `wbSize` = \"1\" и баркодом. 
   *
   * @generated from field: repeated wb.content.v1.ContentV2CardsUploadAddPostReqCardsToAddInnerSizesInner sizes = 7;
   */
  sizes: ContentV2CardsUploadAddPostReqCardsToAddInnerSizesInner[] = [];

  constructor(data?: PartialMessage<ContentV2CardsUploadAddPostReqCardsToAddInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2CardsUploadAddPostReqCardsToAddInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "brand", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "vendorCode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "dimensions", kind: "message", T: ContentV2CardsUploadPostReqInnerVariantsInnerDimensions },
    { no: 6, name: "characteristics", kind: "message", T: ContentV2CardsUploadAddPostReqCardsToAddInnerCharacteristicsInner, repeated: true },
    { no: 7, name: "sizes", kind: "message", T: ContentV2CardsUploadAddPostReqCardsToAddInnerSizesInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2CardsUploadAddPostReqCardsToAddInner {
    return new ContentV2CardsUploadAddPostReqCardsToAddInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2CardsUploadAddPostReqCardsToAddInner {
    return new ContentV2CardsUploadAddPostReqCardsToAddInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2CardsUploadAddPostReqCardsToAddInner {
    return new ContentV2CardsUploadAddPostReqCardsToAddInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2CardsUploadAddPostReqCardsToAddInner | PlainMessage<ContentV2CardsUploadAddPostReqCardsToAddInner> | undefined, b: ContentV2CardsUploadAddPostReqCardsToAddInner | PlainMessage<ContentV2CardsUploadAddPostReqCardsToAddInner> | undefined): boolean {
    return proto3.util.equals(ContentV2CardsUploadAddPostReqCardsToAddInner, a, b);
  }
}
