//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/response_info_advert_type9_united_params_inner_menus_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInnerMenusInner
 */
export class ResponseInfoAdvertType9UnitedParamsInnerMenusInner extends Message<ResponseInfoAdvertType9UnitedParamsInnerMenusInner> {
  /**
   * ID меню, где размещается кампания
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Название меню
   *
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<ResponseInfoAdvertType9UnitedParamsInnerMenusInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.ResponseInfoAdvertType9UnitedParamsInnerMenusInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResponseInfoAdvertType9UnitedParamsInnerMenusInner {
    return new ResponseInfoAdvertType9UnitedParamsInnerMenusInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResponseInfoAdvertType9UnitedParamsInnerMenusInner {
    return new ResponseInfoAdvertType9UnitedParamsInnerMenusInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResponseInfoAdvertType9UnitedParamsInnerMenusInner {
    return new ResponseInfoAdvertType9UnitedParamsInnerMenusInner().fromJsonString(jsonString, options);
  }

  static equals(a: ResponseInfoAdvertType9UnitedParamsInnerMenusInner | PlainMessage<ResponseInfoAdvertType9UnitedParamsInnerMenusInner> | undefined, b: ResponseInfoAdvertType9UnitedParamsInnerMenusInner | PlainMessage<ResponseInfoAdvertType9UnitedParamsInnerMenusInner> | undefined): boolean {
    return proto3.util.equals(ResponseInfoAdvertType9UnitedParamsInnerMenusInner, a, b);
  }
}

