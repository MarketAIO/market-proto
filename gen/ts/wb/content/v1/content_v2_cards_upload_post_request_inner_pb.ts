//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_cards_upload_post_request_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ContentV2CardsUploadPostReqInnerVariantsInner } from "./content_v2_cards_upload_post_request_inner_variants_inner_pb.js";

/**
 * @generated from message wb.content.v1.ContentV2CardsUploadPostReqInner
 */
export class ContentV2CardsUploadPostReqInner extends Message<ContentV2CardsUploadPostReqInner> {
  /**
   * ID предмета
   *
   * @generated from field: int32 subjectID = 1;
   */
  subjectID = 0;

  /**
   * Массив вариантов товара. В каждой КТ может быть не более 30 вариантов (НМ)
   *
   * @generated from field: repeated wb.content.v1.ContentV2CardsUploadPostReqInnerVariantsInner variants = 2;
   */
  variants: ContentV2CardsUploadPostReqInnerVariantsInner[] = [];

  constructor(data?: PartialMessage<ContentV2CardsUploadPostReqInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2CardsUploadPostReqInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "subjectID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "variants", kind: "message", T: ContentV2CardsUploadPostReqInnerVariantsInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2CardsUploadPostReqInner {
    return new ContentV2CardsUploadPostReqInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2CardsUploadPostReqInner {
    return new ContentV2CardsUploadPostReqInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2CardsUploadPostReqInner {
    return new ContentV2CardsUploadPostReqInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2CardsUploadPostReqInner | PlainMessage<ContentV2CardsUploadPostReqInner> | undefined, b: ContentV2CardsUploadPostReqInner | PlainMessage<ContentV2CardsUploadPostReqInner> | undefined): boolean {
    return proto3.util.equals(ContentV2CardsUploadPostReqInner, a, b);
  }
}
