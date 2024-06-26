//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_tags_get200_response_data.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.content.v1.ContentV2TagsGet200ResponseData
 */
export class ContentV2TagsGet200ResponseData extends Message<ContentV2TagsGet200ResponseData> {
  /**
   * Числовой идентификатор тега
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Цвет тега
   *
   * @generated from field: string color = 2;
   */
  color = "";

  /**
   * Имя тега
   *
   * @generated from field: string name = 3;
   */
  name = "";

  constructor(data?: PartialMessage<ContentV2TagsGet200ResponseData>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2TagsGet200ResponseData";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "color", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2TagsGet200ResponseData {
    return new ContentV2TagsGet200ResponseData().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2TagsGet200ResponseData {
    return new ContentV2TagsGet200ResponseData().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2TagsGet200ResponseData {
    return new ContentV2TagsGet200ResponseData().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2TagsGet200ResponseData | PlainMessage<ContentV2TagsGet200ResponseData> | undefined, b: ContentV2TagsGet200ResponseData | PlainMessage<ContentV2TagsGet200ResponseData> | undefined): boolean {
    return proto3.util.equals(ContentV2TagsGet200ResponseData, a, b);
  }
}

