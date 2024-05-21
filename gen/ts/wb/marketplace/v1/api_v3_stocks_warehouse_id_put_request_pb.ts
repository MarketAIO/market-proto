//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_stocks_warehouse_id_put_request.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ApiV3StocksWarehouseIdPutReqStocksInner } from "./api_v3_stocks_warehouse_id_put_request_stocks_inner_pb.js";

/**
 * @generated from message wb.marketplace.v1.ApiV3StocksWarehouseIdPutReq
 */
export class ApiV3StocksWarehouseIdPutReq extends Message<ApiV3StocksWarehouseIdPutReq> {
  /**
   * Массив баркодов товаров и их остатков
   *
   * @generated from field: repeated wb.marketplace.v1.ApiV3StocksWarehouseIdPutReqStocksInner stocks = 1;
   */
  stocks: ApiV3StocksWarehouseIdPutReqStocksInner[] = [];

  constructor(data?: PartialMessage<ApiV3StocksWarehouseIdPutReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3StocksWarehouseIdPutReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "stocks", kind: "message", T: ApiV3StocksWarehouseIdPutReqStocksInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3StocksWarehouseIdPutReq {
    return new ApiV3StocksWarehouseIdPutReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3StocksWarehouseIdPutReq {
    return new ApiV3StocksWarehouseIdPutReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3StocksWarehouseIdPutReq {
    return new ApiV3StocksWarehouseIdPutReq().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3StocksWarehouseIdPutReq | PlainMessage<ApiV3StocksWarehouseIdPutReq> | undefined, b: ApiV3StocksWarehouseIdPutReq | PlainMessage<ApiV3StocksWarehouseIdPutReq> | undefined): boolean {
    return proto3.util.equals(ApiV3StocksWarehouseIdPutReq, a, b);
  }
}

