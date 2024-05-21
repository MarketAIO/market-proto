//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_supplier_subjects_get200_response_inner.proto (package wb.promotion.v1, syntax proto3)
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
 * @generated from message wb.promotion.v1.AdvV1SupplierSubjectsGet200ResponseInner
 */
export class AdvV1SupplierSubjectsGet200ResponseInner extends Message<AdvV1SupplierSubjectsGet200ResponseInner> {
  /**
   * ID предмета
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Предмет
   *
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * Количество Артикулов Wildberries (`nmId`) с таким предметом.
   *
   * @generated from field: int32 count = 3;
   */
  count = 0;

  constructor(data?: PartialMessage<AdvV1SupplierSubjectsGet200ResponseInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1SupplierSubjectsGet200ResponseInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1SupplierSubjectsGet200ResponseInner {
    return new AdvV1SupplierSubjectsGet200ResponseInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1SupplierSubjectsGet200ResponseInner {
    return new AdvV1SupplierSubjectsGet200ResponseInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1SupplierSubjectsGet200ResponseInner {
    return new AdvV1SupplierSubjectsGet200ResponseInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1SupplierSubjectsGet200ResponseInner | PlainMessage<AdvV1SupplierSubjectsGet200ResponseInner> | undefined, b: AdvV1SupplierSubjectsGet200ResponseInner | PlainMessage<AdvV1SupplierSubjectsGet200ResponseInner> | undefined): boolean {
    return proto3.util.equals(AdvV1SupplierSubjectsGet200ResponseInner, a, b);
  }
}

