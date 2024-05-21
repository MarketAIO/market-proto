//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/tariffs/v1/models_tariffs_pallet_response.proto (package wb.tariffs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ModelsWarehousesPalletRates } from "./models_warehouses_pallet_rates_pb.js";

/**
 * @generated from message wb.tariffs.v1.ModelsTariffsPalletResponse
 */
export class ModelsTariffsPalletResponse extends Message<ModelsTariffsPalletResponse> {
  /**
   * @generated from field: wb.tariffs.v1.ModelsWarehousesPalletRates data = 1;
   */
  data?: ModelsWarehousesPalletRates;

  constructor(data?: PartialMessage<ModelsTariffsPalletResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.tariffs.v1.ModelsTariffsPalletResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: ModelsWarehousesPalletRates },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ModelsTariffsPalletResponse {
    return new ModelsTariffsPalletResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ModelsTariffsPalletResponse {
    return new ModelsTariffsPalletResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ModelsTariffsPalletResponse {
    return new ModelsTariffsPalletResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ModelsTariffsPalletResponse | PlainMessage<ModelsTariffsPalletResponse> | undefined, b: ModelsTariffsPalletResponse | PlainMessage<ModelsTariffsPalletResponse> | undefined): boolean {
    return proto3.util.equals(ModelsTariffsPalletResponse, a, b);
  }
}

