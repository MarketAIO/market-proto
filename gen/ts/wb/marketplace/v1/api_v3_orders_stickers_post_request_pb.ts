//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_orders_stickers_post_request.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.ApiV3OrdersStickersPostReq
 */
export class ApiV3OrdersStickersPostReq extends Message<ApiV3OrdersStickersPostReq> {
  /**
   * Массив идентификаторов сборочных заданий
   *
   * @generated from field: repeated int64 orders = 1;
   */
  orders: bigint[] = [];

  constructor(data?: PartialMessage<ApiV3OrdersStickersPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3OrdersStickersPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "orders", kind: "scalar", T: 3 /* ScalarType.INT64 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3OrdersStickersPostReq {
    return new ApiV3OrdersStickersPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3OrdersStickersPostReq {
    return new ApiV3OrdersStickersPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3OrdersStickersPostReq {
    return new ApiV3OrdersStickersPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3OrdersStickersPostReq | PlainMessage<ApiV3OrdersStickersPostReq> | undefined, b: ApiV3OrdersStickersPostReq | PlainMessage<ApiV3OrdersStickersPostReq> | undefined): boolean {
    return proto3.util.equals(ApiV3OrdersStickersPostReq, a, b);
  }
}

