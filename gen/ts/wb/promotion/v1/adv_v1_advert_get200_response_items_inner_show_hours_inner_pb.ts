//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_advert_get200_response_items_inner_show_hours_inner.proto (package wb.promotion.v1, syntax proto3)
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
 * @generated from message wb.promotion.v1.AdvV1AdvertGet200ResponseItemsInnerShowHoursInner
 */
export class AdvV1AdvertGet200ResponseItemsInnerShowHoursInner extends Message<AdvV1AdvertGet200ResponseItemsInnerShowHoursInner> {
  /**
   * Начало показа
   *
   * @generated from field: int32 From = 1;
   */
  From = 0;

  /**
   * Конец показа
   *
   * @generated from field: int32 To = 2;
   */
  To = 0;

  constructor(data?: PartialMessage<AdvV1AdvertGet200ResponseItemsInnerShowHoursInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1AdvertGet200ResponseItemsInnerShowHoursInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "From", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "To", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1AdvertGet200ResponseItemsInnerShowHoursInner {
    return new AdvV1AdvertGet200ResponseItemsInnerShowHoursInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1AdvertGet200ResponseItemsInnerShowHoursInner {
    return new AdvV1AdvertGet200ResponseItemsInnerShowHoursInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1AdvertGet200ResponseItemsInnerShowHoursInner {
    return new AdvV1AdvertGet200ResponseItemsInnerShowHoursInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1AdvertGet200ResponseItemsInnerShowHoursInner | PlainMessage<AdvV1AdvertGet200ResponseItemsInnerShowHoursInner> | undefined, b: AdvV1AdvertGet200ResponseItemsInnerShowHoursInner | PlainMessage<AdvV1AdvertGet200ResponseItemsInnerShowHoursInner> | undefined): boolean {
    return proto3.util.equals(AdvV1AdvertGet200ResponseItemsInnerShowHoursInner, a, b);
  }
}

