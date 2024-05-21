//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/order_user.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.OrderUser
 */
export class OrderUser extends Message<OrderUser> {
  /**
   * ФИО
   *
   * @generated from field: string fio = 1;
   */
  fio = "";

  /**
   * Номер телефона
   *
   * @generated from field: string phone = 2;
   */
  phone = "";

  constructor(data?: PartialMessage<OrderUser>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.OrderUser";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "fio", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "phone", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OrderUser {
    return new OrderUser().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OrderUser {
    return new OrderUser().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OrderUser {
    return new OrderUser().fromJsonString(jsonString, options);
  }

  static equals(a: OrderUser | PlainMessage<OrderUser> | undefined, b: OrderUser | PlainMessage<OrderUser> | undefined): boolean {
    return proto3.util.equals(OrderUser, a, b);
  }
}
