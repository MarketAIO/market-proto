//
//Описание API Вопросов и Отзывов
//
//`Важно.` Допускается 1 запрос в секунду на методы вопросов и отзывов в целом. При превышении лимита до 3 запросов в секунду последует блокировка на 60 секунд. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/feedbacks/v1/product_rating.proto (package wb.feedbacks.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.feedbacks.v1.ProductRating
 */
export class ProductRating extends Message<ProductRating> {
  /**
   * Артикул WB
   *
   * @generated from field: int32 nmId = 1;
   */
  nmId = 0;

  /**
   * Идентификатор карточки товара
   *
   * @generated from field: int32 imtId = 2;
   */
  imtId = 0;

  /**
   * Сумма оценок
   *
   * @generated from field: int32 valuationsSum = 3;
   */
  valuationsSum = 0;

  /**
   * Количество отзывов
   *
   * @generated from field: int32 feedbacksCount = 4;
   */
  feedbacksCount = 0;

  /**
   * Средняя оценка
   *
   * @generated from field: float valuation = 5;
   */
  valuation = 0;

  /**
   * Название товара
   *
   * @generated from field: optional string productName = 6;
   */
  productName?: string;

  /**
   * Артикул продавца
   *
   * @generated from field: string supplierArticle = 7;
   */
  supplierArticle = "";

  /**
   * Бренд товара
   *
   * @generated from field: string brandName = 8;
   */
  brandName = "";

  constructor(data?: PartialMessage<ProductRating>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.feedbacks.v1.ProductRating";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "imtId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "valuationsSum", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "feedbacksCount", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "valuation", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 6, name: "productName", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "supplierArticle", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "brandName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProductRating {
    return new ProductRating().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProductRating {
    return new ProductRating().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProductRating {
    return new ProductRating().fromJsonString(jsonString, options);
  }

  static equals(a: ProductRating | PlainMessage<ProductRating> | undefined, b: ProductRating | PlainMessage<ProductRating> | undefined): boolean {
    return proto3.util.equals(ProductRating, a, b);
  }
}

