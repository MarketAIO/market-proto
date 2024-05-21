//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/api_v3_warehouses_warehouse_id_put_request.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.ApiV3WarehousesWarehouseIdPutReq
 */
export class ApiV3WarehousesWarehouseIdPutReq extends Message<ApiV3WarehousesWarehouseIdPutReq> {
  /**
   * Имя склада продавца
   *
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * Идентификатор склада WB
   *
   * @generated from field: int32 officeId = 2;
   */
  officeId = 0;

  constructor(data?: PartialMessage<ApiV3WarehousesWarehouseIdPutReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.ApiV3WarehousesWarehouseIdPutReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "officeId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV3WarehousesWarehouseIdPutReq {
    return new ApiV3WarehousesWarehouseIdPutReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV3WarehousesWarehouseIdPutReq {
    return new ApiV3WarehousesWarehouseIdPutReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV3WarehousesWarehouseIdPutReq {
    return new ApiV3WarehousesWarehouseIdPutReq().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV3WarehousesWarehouseIdPutReq | PlainMessage<ApiV3WarehousesWarehouseIdPutReq> | undefined, b: ApiV3WarehousesWarehouseIdPutReq | PlainMessage<ApiV3WarehousesWarehouseIdPutReq> | undefined): boolean {
    return proto3.util.equals(ApiV3WarehousesWarehouseIdPutReq, a, b);
  }
}

