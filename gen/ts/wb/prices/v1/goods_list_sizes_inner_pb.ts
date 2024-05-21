//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/prices/v1/goods_list_sizes_inner.proto (package wb.prices.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.prices.v1.GoodsListSizesInner
 */
export class GoodsListSizesInner extends Message<GoodsListSizesInner> {
  /**
   * ID размера. В методах контента это поле `chrtID`
   *
   * @generated from field: int64 sizeID = 1;
   */
  sizeID = protoInt64.zero;

  /**
   * Цена
   *
   * @generated from field: int32 price = 2;
   */
  price = 0;

  /**
   * Цена со скидкой
   *
   * @generated from field: float discountedPrice = 3;
   */
  discountedPrice = 0;

  /**
   * Размер товара
   *
   * @generated from field: string techSizeName = 4;
   */
  techSizeName = "";

  constructor(data?: PartialMessage<GoodsListSizesInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.prices.v1.GoodsListSizesInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sizeID", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "discountedPrice", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "techSizeName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GoodsListSizesInner {
    return new GoodsListSizesInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GoodsListSizesInner {
    return new GoodsListSizesInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GoodsListSizesInner {
    return new GoodsListSizesInner().fromJsonString(jsonString, options);
  }

  static equals(a: GoodsListSizesInner | PlainMessage<GoodsListSizesInner> | undefined, b: GoodsListSizesInner | PlainMessage<GoodsListSizesInner> | undefined): boolean {
    return proto3.util.equals(GoodsListSizesInner, a, b);
  }
}
