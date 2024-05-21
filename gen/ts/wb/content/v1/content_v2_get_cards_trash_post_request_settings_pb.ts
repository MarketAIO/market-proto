//
//Описание API Контента
//
//<dl> <dt>Словарь сокращений:</dt> <dd>КТ — карточка товара</dd> <dd>НМ — номенклатура</dd> </dl> Ограничения по количеству запросов: <dd>Допускается максимум 100 запросов в минуту на методы контента в целом.</dd>  <br> Публичное API Контента создано для синхронизации данных между серверами Wildberries и серверами продавцов. <br> Вы загружаете данные на свои носители, работаете с ними на своих мощностях и синхронизируетесь с нашими серверами по мере необходимости. <br> <code>Не допускается использование API Контента в качестве внешней базы данных. При превышении лимитов на запросы доступ к API будет ограничен.</code> 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/content/v1/content_v2_get_cards_trash_post_request_settings.proto (package wb.content.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ContentV2GetCardsTrashPostReqSettingsSort } from "./content_v2_get_cards_trash_post_request_settings_sort_pb.js";
import { ContentV2GetCardsTrashPostReqSettingsCursor } from "./content_v2_get_cards_trash_post_request_settings_cursor_pb.js";
import { ContentV2GetCardsTrashPostReqSettingsFilter } from "./content_v2_get_cards_trash_post_request_settings_filter_pb.js";

/**
 * @generated from message wb.content.v1.ContentV2GetCardsTrashPostReqSettings
 */
export class ContentV2GetCardsTrashPostReqSettings extends Message<ContentV2GetCardsTrashPostReqSettings> {
  /**
   * @generated from field: wb.content.v1.ContentV2GetCardsTrashPostReqSettingsSort sort = 1;
   */
  sort?: ContentV2GetCardsTrashPostReqSettingsSort;

  /**
   * @generated from field: wb.content.v1.ContentV2GetCardsTrashPostReqSettingsCursor cursor = 2;
   */
  cursor?: ContentV2GetCardsTrashPostReqSettingsCursor;

  /**
   * @generated from field: wb.content.v1.ContentV2GetCardsTrashPostReqSettingsFilter filter = 3;
   */
  filter?: ContentV2GetCardsTrashPostReqSettingsFilter;

  constructor(data?: PartialMessage<ContentV2GetCardsTrashPostReqSettings>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.content.v1.ContentV2GetCardsTrashPostReqSettings";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sort", kind: "message", T: ContentV2GetCardsTrashPostReqSettingsSort },
    { no: 2, name: "cursor", kind: "message", T: ContentV2GetCardsTrashPostReqSettingsCursor },
    { no: 3, name: "filter", kind: "message", T: ContentV2GetCardsTrashPostReqSettingsFilter },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ContentV2GetCardsTrashPostReqSettings {
    return new ContentV2GetCardsTrashPostReqSettings().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPostReqSettings {
    return new ContentV2GetCardsTrashPostReqSettings().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ContentV2GetCardsTrashPostReqSettings {
    return new ContentV2GetCardsTrashPostReqSettings().fromJsonString(jsonString, options);
  }

  static equals(a: ContentV2GetCardsTrashPostReqSettings | PlainMessage<ContentV2GetCardsTrashPostReqSettings> | undefined, b: ContentV2GetCardsTrashPostReqSettings | PlainMessage<ContentV2GetCardsTrashPostReqSettings> | undefined): boolean {
    return proto3.util.equals(ContentV2GetCardsTrashPostReqSettings, a, b);
  }
}

