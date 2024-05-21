//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/tariffs/v1/models_warehouses_box_rates.proto (package wb.tariffs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ModelsWarehouseBoxRates } from "./models_warehouse_box_rates_pb.js";

/**
 * @generated from message wb.tariffs.v1.ModelsWarehousesBoxRates
 */
export class ModelsWarehousesBoxRates extends Message<ModelsWarehousesBoxRates> {
  /**
   * Дата начала тарифа
   *
   * @generated from field: string dtFromMin = 1;
   */
  dtFromMin = "";

  /**
   * Дата начала следующего тарифа
   *
   * @generated from field: string dtNextBox = 2;
   */
  dtNextBox = "";

  /**
   * Дата окончания тарифа
   *
   * @generated from field: string dtTillMax = 3;
   */
  dtTillMax = "";

  /**
   * Тарифы для коробов, сгруппированные по складам
   *
   * @generated from field: repeated wb.tariffs.v1.ModelsWarehouseBoxRates warehouseList = 4;
   */
  warehouseList: ModelsWarehouseBoxRates[] = [];

  constructor(data?: PartialMessage<ModelsWarehousesBoxRates>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.tariffs.v1.ModelsWarehousesBoxRates";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dtFromMin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dtNextBox", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "dtTillMax", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "warehouseList", kind: "message", T: ModelsWarehouseBoxRates, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ModelsWarehousesBoxRates {
    return new ModelsWarehousesBoxRates().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ModelsWarehousesBoxRates {
    return new ModelsWarehousesBoxRates().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ModelsWarehousesBoxRates {
    return new ModelsWarehousesBoxRates().fromJsonString(jsonString, options);
  }

  static equals(a: ModelsWarehousesBoxRates | PlainMessage<ModelsWarehousesBoxRates> | undefined, b: ModelsWarehousesBoxRates | PlainMessage<ModelsWarehousesBoxRates> | undefined): boolean {
    return proto3.util.equals(ModelsWarehousesBoxRates, a, b);
  }
}
