//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_stocks_warehouse_id_post200_response_stocks_inner.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.ApiV3StocksWarehouseIdPost200ResponseStocksInner
 */
export class ApiV3StocksWarehouseIdPost200ResponseStocksInner extends Message<ApiV3StocksWarehouseIdPost200ResponseStocksInner> {
  /**
   * Баркод
   *
   * @generated from field: string sku = 1;
   */
  sku = "";

  /**
   * Остаток
   *
   * @generated from field: int32 amount = 2;
   */
  amount = 0;

  constructor(data?: PartialMessage<ApiV3StocksWarehouseIdPost200ResponseStocksInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3StocksWarehouseIdPost200ResponseStocksInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sku", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "amount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3StocksWarehouseIdPost200ResponseStocksInner {
    return new ApiV3StocksWarehouseIdPost200ResponseStocksInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3StocksWarehouseIdPost200ResponseStocksInner {
    return new ApiV3StocksWarehouseIdPost200ResponseStocksInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3StocksWarehouseIdPost200ResponseStocksInner {
    return new ApiV3StocksWarehouseIdPost200ResponseStocksInner().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3StocksWarehouseIdPost200ResponseStocksInner | PlainMessage<ApiV3StocksWarehouseIdPost200ResponseStocksInner> | undefined, b: ApiV3StocksWarehouseIdPost200ResponseStocksInner | PlainMessage<ApiV3StocksWarehouseIdPost200ResponseStocksInner> | undefined): boolean {
    return proto3.util.equals(ApiV3StocksWarehouseIdPost200ResponseStocksInner, a, b);
  }
}

