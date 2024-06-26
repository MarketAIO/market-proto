//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/prices/v1/size_good_req.proto (package wb.prices.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.prices.v1.SizeGoodReq
 */
export class SizeGoodReq extends Message<SizeGoodReq> {
  /**
   * Артикул Wildberries
   *
   * @generated from field: int32 nmID = 1;
   */
  nmID = 0;

  /**
   * ID размера. Можно получить с помощью метода [Получение списка товаров по артикулам](./#tag/Spiski-tovarov/paths/~1api~1v2~1list~1goods~1filter/get), поле `sizeID`. В методах контента это поле `chrtID`
   *
   * @generated from field: int32 sizeID = 2;
   */
  sizeID = 0;

  /**
   * Цена. Валюту можно получить с помощью метода [Получение списка товаров по артикулам](./#tag/Spiski-tovarov/paths/~1api~1v2~1list~1goods~1filter/get), поле `currencyIsoCode4217`
   *
   * @generated from field: int32 price = 3;
   */
  price = 0;

  constructor(data?: PartialMessage<SizeGoodReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.prices.v1.SizeGoodReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "sizeID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SizeGoodReq {
    return new SizeGoodReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SizeGoodReq {
    return new SizeGoodReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SizeGoodReq {
    return new SizeGoodReq().fromJsonString(jsonString, options);
  }

  static equals(a: SizeGoodReq | PlainMessage<SizeGoodReq> | undefined, b: SizeGoodReq | PlainMessage<SizeGoodReq> | undefined): boolean {
    return proto3.util.equals(SizeGoodReq, a, b);
  }
}

