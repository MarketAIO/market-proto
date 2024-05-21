//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/tariffs/v1/tariffs_box_response.proto (package wb.tariffs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ModelsTariffsBoxResponse } from "./models_tariffs_box_response_pb.js";

/**
 * @generated from message wb.tariffs.v1.TariffsBoxResponse
 */
export class TariffsBoxResponse extends Message<TariffsBoxResponse> {
  /**
   * @generated from field: wb.tariffs.v1.ModelsTariffsBoxResponse response = 1;
   */
  response?: ModelsTariffsBoxResponse;

  constructor(data?: PartialMessage<TariffsBoxResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.tariffs.v1.TariffsBoxResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "response", kind: "message", T: ModelsTariffsBoxResponse },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TariffsBoxResponse {
    return new TariffsBoxResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TariffsBoxResponse {
    return new TariffsBoxResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TariffsBoxResponse {
    return new TariffsBoxResponse().fromJsonString(jsonString, options);
  }

  static equals(a: TariffsBoxResponse | PlainMessage<TariffsBoxResponse> | undefined, b: TariffsBoxResponse | PlainMessage<TariffsBoxResponse> | undefined): boolean {
    return proto3.util.equals(TariffsBoxResponse, a, b);
  }
}

