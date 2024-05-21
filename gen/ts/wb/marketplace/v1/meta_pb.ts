//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/meta.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.Meta
 */
export class Meta extends Message<Meta> {
  /**
   * IMEI
   *
   * @generated from field: optional string imei = 1;
   */
  imei?: string;

  /**
   * УИН
   *
   * @generated from field: optional string uin = 2;
   */
  uin?: string;

  /**
   * GTIN
   *
   * @generated from field: optional string gtin = 3;
   */
  gtin?: string;

  /**
   * КиЗ (Маркировка честного знака)
   *
   * @generated from field: optional string sgtin = 4;
   */
  sgtin?: string;

  constructor(data?: PartialMessage<Meta>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.Meta";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "imei", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "uin", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "gtin", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "sgtin", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Meta {
    return new Meta().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Meta {
    return new Meta().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Meta {
    return new Meta().fromJsonString(jsonString, options);
  }

  static equals(a: Meta | PlainMessage<Meta> | undefined, b: Meta | PlainMessage<Meta> | undefined): boolean {
    return proto3.util.equals(Meta, a, b);
  }
}
