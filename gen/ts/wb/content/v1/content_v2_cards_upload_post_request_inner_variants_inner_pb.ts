//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_cards_upload_post_request_inner_variants_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ContentV2CardsUploadPostReqInnerVariantsInnerDimensions } from "./content_v2_cards_upload_post_request_inner_variants_inner_dimensions_pb.js";
import { ContentV2CardsUploadPostReqInnerVariantsInnerSizesInner } from "./content_v2_cards_upload_post_request_inner_variants_inner_sizes_inner_pb.js";
import { ContentV2CardsUploadPostReqInnerVariantsInnerCharacteristicsInner } from "./content_v2_cards_upload_post_request_inner_variants_inner_characteristics_inner_pb.js";

/**
 * @generated from message wb.content.v1.ContentV2CardsUploadPostReqInnerVariantsInner
 */
export class ContentV2CardsUploadPostReqInnerVariantsInner extends Message<ContentV2CardsUploadPostReqInnerVariantsInner> {
  /**
   * Бренд
   *
   * @generated from field: string brand = 1;
   */
  brand = "";

  /**
   * Наименование товара
   *
   * @generated from field: string title = 2;
   */
  title = "";

  /**
   * Описание товара. Максимальное количество символов зависит от категории товара. Стандарт — 2000, минимум — 1000, максимум — 5000.<br> Подробно о правилах описания в **Правилах заполнения карточки товара** в разделе [Инструкции](https://seller.wildberries.ru/training) на портале продавцов. 
   *
   * @generated from field: string description = 3;
   */
  description = "";

  /**
   * Артикул продавца
   *
   * @generated from field: string vendorCode = 4;
   */
  vendorCode = "";

  /**
   * @generated from field: wb.content.v1.ContentV2CardsUploadPostReqInnerVariantsInnerDimensions dimensions = 5;
   */
  dimensions?: ContentV2CardsUploadPostReqInnerVariantsInnerDimensions;

  /**
   * Массив с размерами. <br>  Если для размерного товара (обувь, одежда и др.) не указать этот массив, то системой в карточке он будет сгенерирован автоматически с `techSize` = \"A\" и `wbSize` = \"1\" и баркодом. 
   *
   * @generated from field: repeated wb.content.v1.ContentV2CardsUploadPostReqInnerVariantsInnerSizesInner sizes = 6;
   */
  sizes: ContentV2CardsUploadPostReqInnerVariantsInnerSizesInner[] = [];

  /**
   * Массив характеристик товара
   *
   * @generated from field: repeated wb.content.v1.ContentV2CardsUploadPostReqInnerVariantsInnerCharacteristicsInner characteristics = 7;
   */
  characteristics: ContentV2CardsUploadPostReqInnerVariantsInnerCharacteristicsInner[] = [];

  constructor(data?: PartialMessage<ContentV2CardsUploadPostReqInnerVariantsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2CardsUploadPostReqInnerVariantsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "brand", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "vendorCode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "dimensions", kind: "message", T: ContentV2CardsUploadPostReqInnerVariantsInnerDimensions },
    { no: 6, name: "sizes", kind: "message", T: ContentV2CardsUploadPostReqInnerVariantsInnerSizesInner, repeated: true },
    { no: 7, name: "characteristics", kind: "message", T: ContentV2CardsUploadPostReqInnerVariantsInnerCharacteristicsInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2CardsUploadPostReqInnerVariantsInner {
    return new ContentV2CardsUploadPostReqInnerVariantsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2CardsUploadPostReqInnerVariantsInner {
    return new ContentV2CardsUploadPostReqInnerVariantsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2CardsUploadPostReqInnerVariantsInner {
    return new ContentV2CardsUploadPostReqInnerVariantsInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2CardsUploadPostReqInnerVariantsInner | PlainMessage<ContentV2CardsUploadPostReqInnerVariantsInner> | undefined, b: ContentV2CardsUploadPostReqInnerVariantsInner | PlainMessage<ContentV2CardsUploadPostReqInnerVariantsInner> | undefined): boolean {
    return proto3.util.equals(ContentV2CardsUploadPostReqInnerVariantsInner, a, b);
  }
}
