//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/response_info_advert_type8_auto_params_sets_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.ResponseInfoAdvertType8AutoParamsSetsInner
 */
export class ResponseInfoAdvertType8AutoParamsSetsInner extends Message<ResponseInfoAdvertType8AutoParamsSetsInner> {
  /**
   * Идентификатор set
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Название set
   *
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<ResponseInfoAdvertType8AutoParamsSetsInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.ResponseInfoAdvertType8AutoParamsSetsInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseInfoAdvertType8AutoParamsSetsInner {
    return new ResponseInfoAdvertType8AutoParamsSetsInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseInfoAdvertType8AutoParamsSetsInner {
    return new ResponseInfoAdvertType8AutoParamsSetsInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseInfoAdvertType8AutoParamsSetsInner {
    return new ResponseInfoAdvertType8AutoParamsSetsInner().fromJsonString(jsonString, options);
  }

  static equals(a: ResponseInfoAdvertType8AutoParamsSetsInner | PlainMessage<ResponseInfoAdvertType8AutoParamsSetsInner> | undefined, b: ResponseInfoAdvertType8AutoParamsSetsInner | PlainMessage<ResponseInfoAdvertType8AutoParamsSetsInner> | undefined): boolean {
    return proto3.util.equals(ResponseInfoAdvertType8AutoParamsSetsInner, a, b);
  }
}

