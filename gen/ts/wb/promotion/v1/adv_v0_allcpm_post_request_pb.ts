//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v0_allcpm_post_request.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV0AllcpmPostReq
 */
export class AdvV0AllcpmPostReq extends Message<AdvV0AllcpmPostReq> {
  /**
   * Массив параметров запроса, по которым будет получен список ставок активных кампаний: должен быть значением `menuId` (для кампании в каталоге), `subjectId` (для кампании в поиске и рекомендациях) или `setId` (для кампании в карточке товара). 
   *
   * @generated from field: repeated int32 param = 1;
   */
  param: number[] = [];

  constructor(data?: PartialMessage<AdvV0AllcpmPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV0AllcpmPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "param", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV0AllcpmPostReq {
    return new AdvV0AllcpmPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV0AllcpmPostReq {
    return new AdvV0AllcpmPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV0AllcpmPostReq {
    return new AdvV0AllcpmPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV0AllcpmPostReq | PlainMessage<AdvV0AllcpmPostReq> | undefined, b: AdvV0AllcpmPostReq | PlainMessage<AdvV0AllcpmPostReq> | undefined): boolean {
    return proto3.util.equals(AdvV0AllcpmPostReq, a, b);
  }
}

