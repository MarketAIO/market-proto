//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/tariffs/v1/models_return_tariffs_response.proto (package wb.tariffs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ModelsWarehousesReturnRates } from "./models_warehouses_return_rates_pb.js";

/**
 * @generated from message wb.tariffs.v1.ModelsReturnTariffsResponse
 */
export class ModelsReturnTariffsResponse extends Message<ModelsReturnTariffsResponse> {
  /**
   * @generated from field: wb.tariffs.v1.ModelsWarehousesReturnRates data = 1;
   */
  data?: ModelsWarehousesReturnRates;

  constructor(data?: PartialMessage<ModelsReturnTariffsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.tariffs.v1.ModelsReturnTariffsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: ModelsWarehousesReturnRates },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ModelsReturnTariffsResponse {
    return new ModelsReturnTariffsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ModelsReturnTariffsResponse {
    return new ModelsReturnTariffsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ModelsReturnTariffsResponse {
    return new ModelsReturnTariffsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ModelsReturnTariffsResponse | PlainMessage<ModelsReturnTariffsResponse> | undefined, b: ModelsReturnTariffsResponse | PlainMessage<ModelsReturnTariffsResponse> | undefined): boolean {
    return proto3.util.equals(ModelsReturnTariffsResponse, a, b);
  }
}

