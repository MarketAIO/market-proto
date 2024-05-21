//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_supplies_orders_reshipment_get200_response.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ApiV3SuppliesOrdersReshipmentGet200ResponseOrdersInner } from "./api_v3_supplies_orders_reshipment_get200_response_orders_inner_pb.js";

/**
 * @generated from message wb.marketplace.v1.ApiV3SuppliesOrdersReshipmentGet200Response
 */
export class ApiV3SuppliesOrdersReshipmentGet200Response extends Message<ApiV3SuppliesOrdersReshipmentGet200Response> {
  /**
   * Список заказов
   *
   * @generated from field: repeated wb.marketplace.v1.ApiV3SuppliesOrdersReshipmentGet200ResponseOrdersInner orders = 1;
   */
  orders: ApiV3SuppliesOrdersReshipmentGet200ResponseOrdersInner[] = [];

  constructor(data?: PartialMessage<ApiV3SuppliesOrdersReshipmentGet200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3SuppliesOrdersReshipmentGet200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "orders", kind: "message", T: ApiV3SuppliesOrdersReshipmentGet200ResponseOrdersInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3SuppliesOrdersReshipmentGet200Response {
    return new ApiV3SuppliesOrdersReshipmentGet200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3SuppliesOrdersReshipmentGet200Response {
    return new ApiV3SuppliesOrdersReshipmentGet200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3SuppliesOrdersReshipmentGet200Response {
    return new ApiV3SuppliesOrdersReshipmentGet200Response().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3SuppliesOrdersReshipmentGet200Response | PlainMessage<ApiV3SuppliesOrdersReshipmentGet200Response> | undefined, b: ApiV3SuppliesOrdersReshipmentGet200Response | PlainMessage<ApiV3SuppliesOrdersReshipmentGet200Response> | undefined): boolean {
    return proto3.util.equals(ApiV3SuppliesOrdersReshipmentGet200Response, a, b);
  }
}

