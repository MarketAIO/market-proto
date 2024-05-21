//
//API цен и скидок
//
//С помощью этих методов можно устанавливать цены и скидки. Максимум — 10 запросов за 6 секунд суммарно для всех методов. 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/prices/v1/api_v2_upload_task_post_request.proto (package wb.prices.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Good } from "./good_pb.js";

/**
 * @generated from message wb.prices.v1.ApiV2UploadTaskPostReq
 */
export class ApiV2UploadTaskPostReq extends Message<ApiV2UploadTaskPostReq> {
  /**
   * Товары, цены и скидки для них. Максимум 1 000 товаров. Цена и скидка не могут быть пустыми одновременно. <br><br> Если новая цена со скидкой будет хотя бы в 3 раза меньше старой, она попадёт [в карантин](https://seller.wildberries.ru/discount-and-prices/quarantine) и товар будет продаваться по старой цене. Ошибка об этом будет в ответах методов [Состояния загрузок](./#tag/Sostoyaniya-zagruzok) <br><br> Вы можете изменить цену или скидку с помощью API либо вывести товар из карантина [в личном кабинете](https://seller.wildberries.ru/discount-and-prices/quarantine) 
   *
   * @generated from field: repeated wb.prices.v1.Good data = 1;
   */
  data: Good[] = [];

  constructor(data?: PartialMessage<ApiV2UploadTaskPostReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.prices.v1.ApiV2UploadTaskPostReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: Good, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2UploadTaskPostReq {
    return new ApiV2UploadTaskPostReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2UploadTaskPostReq {
    return new ApiV2UploadTaskPostReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2UploadTaskPostReq {
    return new ApiV2UploadTaskPostReq().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2UploadTaskPostReq | PlainMessage<ApiV2UploadTaskPostReq> | undefined, b: ApiV2UploadTaskPostReq | PlainMessage<ApiV2UploadTaskPostReq> | undefined): boolean {
    return proto3.util.equals(ApiV2UploadTaskPostReq, a, b);
  }
}
