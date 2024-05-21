//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_get_cards_list_post200_response_cards_inner_characteristics_inner.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.content.v1.ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner
 */
export class ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner extends Message<ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner> {
  /**
   * Идентификатор характеристики
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Название характеристики
   *
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * Значение характеристики. Тип значения зависит от типа характеристики
   *
   * @generated from field: google.protobuf.Any value = 3;
   */
  value?: Any;

  constructor(data?: PartialMessage<ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "value", kind: "message", T: Any },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner {
    return new ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner {
    return new ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner {
    return new ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner | PlainMessage<ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner> | undefined, b: ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner | PlainMessage<ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner> | undefined): boolean {
    return proto3.util.equals(ContentV2GetCardsListPost200ResponseCardsInnerCharacteristicsInner, a, b);
  }
}
