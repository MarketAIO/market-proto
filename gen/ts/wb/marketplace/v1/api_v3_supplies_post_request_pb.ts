//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_supplies_post_request.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.ApiV3SuppliesPostReq
 */
export class ApiV3SuppliesPostReq extends Message<ApiV3SuppliesPostReq> {
  /**
   * Наименование поставки
   *
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<ApiV3SuppliesPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3SuppliesPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3SuppliesPostReq {
    return new ApiV3SuppliesPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3SuppliesPostReq {
    return new ApiV3SuppliesPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3SuppliesPostReq {
    return new ApiV3SuppliesPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3SuppliesPostReq | PlainMessage<ApiV3SuppliesPostReq> | undefined, b: ApiV3SuppliesPostReq | PlainMessage<ApiV3SuppliesPostReq> | undefined): boolean {
    return proto3.util.equals(ApiV3SuppliesPostReq, a, b);
  }
}
