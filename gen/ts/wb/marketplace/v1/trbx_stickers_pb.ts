//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/trbx_stickers.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.TrbxStickers
 */
export class TrbxStickers extends Message<TrbxStickers> {
  /**
   * Закодированное значение этикетки.
   *
   * @generated from field: string barcode = 1;
   */
  barcode = "";

  /**
   * Полное представление этикетки в заданном формате. (кодировка base64)
   *
   * @generated from field: bytes file = 2;
   */
  file = new Uint8Array(0);

  constructor(data?: PartialMessage<TrbxStickers>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.TrbxStickers";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "barcode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "file", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TrbxStickers {
    return new TrbxStickers().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TrbxStickers {
    return new TrbxStickers().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TrbxStickers {
    return new TrbxStickers().fromJsonString(jsonString, options);
  }

  static equals(a: TrbxStickers | PlainMessage<TrbxStickers> | undefined, b: TrbxStickers | PlainMessage<TrbxStickers> | undefined): boolean {
    return proto3.util.equals(TrbxStickers, a, b);
  }
}

