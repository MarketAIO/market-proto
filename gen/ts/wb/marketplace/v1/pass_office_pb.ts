//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/pass_office.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.PassOffice
 */
export class PassOffice extends Message<PassOffice> {
  /**
   * Название
   *
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * Адрес
   *
   * @generated from field: string address = 2;
   */
  address = "";

  /**
   * ID
   *
   * @generated from field: int64 id = 3;
   */
  id = protoInt64.zero;

  constructor(data?: PartialMessage<PassOffice>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.PassOffice";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PassOffice {
    return new PassOffice().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PassOffice {
    return new PassOffice().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PassOffice {
    return new PassOffice().fromJsonString(jsonString, options);
  }

  static equals(a: PassOffice | PlainMessage<PassOffice> | undefined, b: PassOffice | PlainMessage<PassOffice> | undefined): boolean {
    return proto3.util.equals(PassOffice, a, b);
  }
}
