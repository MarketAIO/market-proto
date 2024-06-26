//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_orders_status_post200_response_orders_inner.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner
 */
export class ApiV3OrdersStatusPost200ResponseOrdersInner extends Message<ApiV3OrdersStatusPost200ResponseOrdersInner> {
  /**
   * Идентификатор сборочного задания
   *
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner.SupplierStatusEnum supplierStatus = 2;
   */
  supplierStatus = ApiV3OrdersStatusPost200ResponseOrdersInner_SupplierStatusEnum.SupplierStatusEnum_NEW;

  /**
   * @generated from field: wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner.WbStatusEnum wbStatus = 3;
   */
  wbStatus = ApiV3OrdersStatusPost200ResponseOrdersInner_WbStatusEnum.WbStatusEnum_WAITING;

  constructor(data?: PartialMessage<ApiV3OrdersStatusPost200ResponseOrdersInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "supplierStatus", kind: "enum", T: proto3.getEnumType(ApiV3OrdersStatusPost200ResponseOrdersInner_SupplierStatusEnum) },
    { no: 3, name: "wbStatus", kind: "enum", T: proto3.getEnumType(ApiV3OrdersStatusPost200ResponseOrdersInner_WbStatusEnum) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3OrdersStatusPost200ResponseOrdersInner {
    return new ApiV3OrdersStatusPost200ResponseOrdersInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3OrdersStatusPost200ResponseOrdersInner {
    return new ApiV3OrdersStatusPost200ResponseOrdersInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3OrdersStatusPost200ResponseOrdersInner {
    return new ApiV3OrdersStatusPost200ResponseOrdersInner().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3OrdersStatusPost200ResponseOrdersInner | PlainMessage<ApiV3OrdersStatusPost200ResponseOrdersInner> | undefined, b: ApiV3OrdersStatusPost200ResponseOrdersInner | PlainMessage<ApiV3OrdersStatusPost200ResponseOrdersInner> | undefined): boolean {
    return proto3.util.equals(ApiV3OrdersStatusPost200ResponseOrdersInner, a, b);
  }
}

/**
 * Статус сборочного задания продавца (устанавливается продавцом)
 *
 * @generated from enum wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner.SupplierStatusEnum
 */
export enum ApiV3OrdersStatusPost200ResponseOrdersInner_SupplierStatusEnum {
  /**
   * @generated from enum value: SupplierStatusEnum_NEW = 0;
   */
  SupplierStatusEnum_NEW = 0,

  /**
   * @generated from enum value: SupplierStatusEnum_CONFIRM = 1;
   */
  SupplierStatusEnum_CONFIRM = 1,

  /**
   * @generated from enum value: SupplierStatusEnum_COMPLETE = 2;
   */
  SupplierStatusEnum_COMPLETE = 2,

  /**
   * @generated from enum value: SupplierStatusEnum_CANCEL = 3;
   */
  SupplierStatusEnum_CANCEL = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(ApiV3OrdersStatusPost200ResponseOrdersInner_SupplierStatusEnum)
proto3.util.setEnumType(ApiV3OrdersStatusPost200ResponseOrdersInner_SupplierStatusEnum, "wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner.SupplierStatusEnum", [
  { no: 0, name: "SupplierStatusEnum_NEW" },
  { no: 1, name: "SupplierStatusEnum_CONFIRM" },
  { no: 2, name: "SupplierStatusEnum_COMPLETE" },
  { no: 3, name: "SupplierStatusEnum_CANCEL" },
]);

/**
 * Статус сборочного задания в системе Wildberries
 *
 * @generated from enum wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner.WbStatusEnum
 */
export enum ApiV3OrdersStatusPost200ResponseOrdersInner_WbStatusEnum {
  /**
   * @generated from enum value: WbStatusEnum_WAITING = 0;
   */
  WbStatusEnum_WAITING = 0,

  /**
   * @generated from enum value: WbStatusEnum_SORTED = 1;
   */
  WbStatusEnum_SORTED = 1,

  /**
   * @generated from enum value: WbStatusEnum_SOLD = 2;
   */
  WbStatusEnum_SOLD = 2,

  /**
   * @generated from enum value: WbStatusEnum_CANCELED = 3;
   */
  WbStatusEnum_CANCELED = 3,

  /**
   * @generated from enum value: WbStatusEnum_CANCELED_BY_CLIENT = 4;
   */
  WbStatusEnum_CANCELED_BY_CLIENT = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(ApiV3OrdersStatusPost200ResponseOrdersInner_WbStatusEnum)
proto3.util.setEnumType(ApiV3OrdersStatusPost200ResponseOrdersInner_WbStatusEnum, "wb.marketplace.v1.ApiV3OrdersStatusPost200ResponseOrdersInner.WbStatusEnum", [
  { no: 0, name: "WbStatusEnum_WAITING" },
  { no: 1, name: "WbStatusEnum_SORTED" },
  { no: 2, name: "WbStatusEnum_SOLD" },
  { no: 3, name: "WbStatusEnum_CANCELED" },
  { no: 4, name: "WbStatusEnum_CANCELED_BY_CLIENT" },
]);

