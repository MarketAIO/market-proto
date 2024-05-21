//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_orders_status_post_request.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.ApiV3OrdersStatusPostReq
 */
export class ApiV3OrdersStatusPostReq extends Message<ApiV3OrdersStatusPostReq> {
  /**
   * Список идентификаторов сборочных заданий
   *
   * @generated from field: repeated int64 orders = 1;
   */
  orders: bigint[] = [];

  constructor(data?: PartialMessage<ApiV3OrdersStatusPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3OrdersStatusPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "orders", kind: "scalar", T: 3 /* ScalarType.INT64 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3OrdersStatusPostReq {
    return new ApiV3OrdersStatusPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3OrdersStatusPostReq {
    return new ApiV3OrdersStatusPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3OrdersStatusPostReq {
    return new ApiV3OrdersStatusPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3OrdersStatusPostReq | PlainMessage<ApiV3OrdersStatusPostReq> | undefined, b: ApiV3OrdersStatusPostReq | PlainMessage<ApiV3OrdersStatusPostReq> | undefined): boolean {
    return proto3.util.equals(ApiV3OrdersStatusPostReq, a, b);
  }
}

