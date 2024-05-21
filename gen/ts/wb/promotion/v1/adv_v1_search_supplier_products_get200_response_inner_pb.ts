//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_search_supplier_products_get200_response_inner.proto (package wb.promotion.v1, syntax proto3)
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
import {Any, Message, proto3} from "@bufbuild/protobuf";
import {
    AdvV1SearchSupplierProductsGet200ResponseInnerSubject
} from "./adv_v1_search_supplier_products_get200_response_inner_subject_pb.js";
import {
    AdvV1SearchSupplierProductsGet200ResponseInnerKind
} from "./adv_v1_search_supplier_products_get200_response_inner_kind_pb.js";

/**
 * @generated from message wb.promotion.v1.AdvV1SearchSupplierProductsGet200ResponseInner
 */
export class AdvV1SearchSupplierProductsGet200ResponseInner extends Message<AdvV1SearchSupplierProductsGet200ResponseInner> {
  /**
   * Артикул WB.
   *
   * @generated from field: int32 nm = 1;
   */
  nm = 0;

  /**
   * Наименование товара.
   *
   * @generated from field: google.protobuf.Any name = 2;
   */
  name?: Any;

  /**
   * @generated from field: wb.promotion.v1.AdvV1SearchSupplierProductsGet200ResponseInnerSubject subject = 3;
   */
  subject?: AdvV1SearchSupplierProductsGet200ResponseInnerSubject;

  /**
   * Бренд.
   *
   * @generated from field: string brand = 4;
   */
  brand = "";

  /**
   * @generated from field: wb.promotion.v1.AdvV1SearchSupplierProductsGet200ResponseInnerKind kind = 5;
   */
  kind?: AdvV1SearchSupplierProductsGet200ResponseInnerKind;

  constructor(data?: PartialMessage<AdvV1SearchSupplierProductsGet200ResponseInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1SearchSupplierProductsGet200ResponseInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nm", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "message", T: Any },
    { no: 3, name: "subject", kind: "message", T: AdvV1SearchSupplierProductsGet200ResponseInnerSubject },
    { no: 4, name: "brand", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "kind", kind: "message", T: AdvV1SearchSupplierProductsGet200ResponseInnerKind },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1SearchSupplierProductsGet200ResponseInner {
    return new AdvV1SearchSupplierProductsGet200ResponseInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1SearchSupplierProductsGet200ResponseInner {
    return new AdvV1SearchSupplierProductsGet200ResponseInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1SearchSupplierProductsGet200ResponseInner {
    return new AdvV1SearchSupplierProductsGet200ResponseInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1SearchSupplierProductsGet200ResponseInner | PlainMessage<AdvV1SearchSupplierProductsGet200ResponseInner> | undefined, b: AdvV1SearchSupplierProductsGet200ResponseInner | PlainMessage<AdvV1SearchSupplierProductsGet200ResponseInner> | undefined): boolean {
    return proto3.util.equals(AdvV1SearchSupplierProductsGet200ResponseInner, a, b);
  }
}

