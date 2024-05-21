//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_supplies_get200_response.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Supply } from "./supply_pb.js";

/**
 * @generated from message wb.marketplace.v1.ApiV3SuppliesGet200Response
 */
export class ApiV3SuppliesGet200Response extends Message<ApiV3SuppliesGet200Response> {
  /**
   * Параметр пагинации. Содержит значение, которое необходимо указать в запросе для получения следующего пакета данных
   *
   * @generated from field: int64 next = 1;
   */
  next = protoInt64.zero;

  /**
   * Список поставок
   *
   * @generated from field: repeated wb.marketplace.v1.Supply supplies = 2;
   */
  supplies: Supply[] = [];

  constructor(data?: PartialMessage<ApiV3SuppliesGet200Response>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3SuppliesGet200Response";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "supplies", kind: "message", T: Supply, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3SuppliesGet200Response {
    return new ApiV3SuppliesGet200Response().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3SuppliesGet200Response {
    return new ApiV3SuppliesGet200Response().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3SuppliesGet200Response {
    return new ApiV3SuppliesGet200Response().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3SuppliesGet200Response | PlainMessage<ApiV3SuppliesGet200Response> | undefined, b: ApiV3SuppliesGet200Response | PlainMessage<ApiV3SuppliesGet200Response> | undefined): boolean {
    return proto3.util.equals(ApiV3SuppliesGet200Response, a, b);
  }
}
