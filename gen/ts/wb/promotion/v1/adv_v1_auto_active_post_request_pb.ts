//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_auto_active_post_request.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
    BinaryReadOptions,
    FieldList,
    JsonReadOptions,
    JsonValue,
    PartialMessage,
    PlainMessage
} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV1AutoActivePostReq
 */
export class AdvV1AutoActivePostReq extends Message<AdvV1AutoActivePostReq> {
  /**
   * Рекомендации на главной (`false` - отключены, `true` - включены)
   *
   * @generated from field: bool recom = 1;
   */
  recom = false;

  /**
   * Поиск/Каталог (`false` - отключены, `true` - включены)
   *
   * @generated from field: bool booster = 2;
   */
  booster = false;

  /**
   * Карточка товара (`false` - отключены, `true` - включены)
   *
   * @generated from field: bool carousel = 3;
   */
  carousel = false;

  constructor(data?: PartialMessage<AdvV1AutoActivePostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1AutoActivePostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recom", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "booster", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "carousel", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1AutoActivePostReq {
    return new AdvV1AutoActivePostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1AutoActivePostReq {
    return new AdvV1AutoActivePostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1AutoActivePostReq {
    return new AdvV1AutoActivePostReq().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1AutoActivePostReq | PlainMessage<AdvV1AutoActivePostReq> | undefined, b: AdvV1AutoActivePostReq | PlainMessage<AdvV1AutoActivePostReq> | undefined): boolean {
    return proto3.util.equals(AdvV1AutoActivePostReq, a, b);
  }
}

