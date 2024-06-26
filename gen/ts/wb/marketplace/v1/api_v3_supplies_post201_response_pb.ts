//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_supplies_post201_response.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.ApiV3SuppliesPost201Response
 */
export class ApiV3SuppliesPost201Response extends Message<ApiV3SuppliesPost201Response> {
  /**
   * Идентификатор поставки
   *
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<ApiV3SuppliesPost201Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3SuppliesPost201Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3SuppliesPost201Response {
    return new ApiV3SuppliesPost201Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3SuppliesPost201Response {
    return new ApiV3SuppliesPost201Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3SuppliesPost201Response {
    return new ApiV3SuppliesPost201Response().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3SuppliesPost201Response | PlainMessage<ApiV3SuppliesPost201Response> | undefined, b: ApiV3SuppliesPost201Response | PlainMessage<ApiV3SuppliesPost201Response> | undefined): boolean {
    return proto3.util.equals(ApiV3SuppliesPost201Response, a, b);
  }
}

