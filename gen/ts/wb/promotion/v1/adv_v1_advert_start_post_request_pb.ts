//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_advert_start_post_request.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV1AdvertStartPostReq
 */
export class AdvV1AdvertStartPostReq extends Message<AdvV1AdvertStartPostReq> {
  /**
   * ID медиакампании
   *
   * @generated from field: int32 advert_id = 1;
   */
  advertId = 0;

  /**
   * Описание причины запуска
   *
   * @generated from field: string reason = 2;
   */
  reason = "";

  constructor(data?: PartialMessage<AdvV1AdvertStartPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1AdvertStartPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "advert_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1AdvertStartPostReq {
    return new AdvV1AdvertStartPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1AdvertStartPostReq {
    return new AdvV1AdvertStartPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1AdvertStartPostReq {
    return new AdvV1AdvertStartPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1AdvertStartPostReq | PlainMessage<AdvV1AdvertStartPostReq> | undefined, b: AdvV1AdvertStartPostReq | PlainMessage<AdvV1AdvertStartPostReq> | undefined): boolean {
    return proto3.util.equals(AdvV1AdvertStartPostReq, a, b);
  }
}

