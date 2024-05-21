//
//Описание API Marketplace
//
//No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/marketplace/v1/warehouse.proto (package wb.marketplace.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.marketplace.v1.Warehouse
 */
export class Warehouse extends Message<Warehouse> {
  /**
   * Название склада продавца
   *
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * ID склада WB
   *
   * @generated from field: int64 officeId = 2;
   */
  officeId = protoInt64.zero;

  /**
   * ID склада продавца
   *
   * @generated from field: int64 id = 3;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: wb.marketplace.v1.Warehouse.CargoTypeEnum cargoType = 4;
   */
  cargoType = Warehouse_CargoTypeEnum.CargoTypeEnum__1;

  /**
   * @generated from field: wb.marketplace.v1.Warehouse.DeliveryTypeEnum deliveryType = 5;
   */
  deliveryType = Warehouse_DeliveryTypeEnum.DeliveryTypeEnum__1;

  constructor(data?: PartialMessage<Warehouse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.marketplace.v1.Warehouse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "officeId", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "cargoType", kind: "enum", T: proto3.getEnumType(Warehouse_CargoTypeEnum) },
    { no: 5, name: "deliveryType", kind: "enum", T: proto3.getEnumType(Warehouse_DeliveryTypeEnum) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Warehouse {
    return new Warehouse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Warehouse {
    return new Warehouse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Warehouse {
    return new Warehouse().fromJsonString(jsonString, options);
  }

  static equals(a: Warehouse | PlainMessage<Warehouse> | undefined, b: Warehouse | PlainMessage<Warehouse> | undefined): boolean {
    return proto3.util.equals(Warehouse, a, b);
  }
}

/**
 * <dl> <dt>Тип товара, который принимает склад:</dt> <dd>1 - обычный</dd> <dd>2 - СГТ (Сверхгабаритный товар)</dd> <dd>3 - КГТ (Крупногабаритный товар). Не используется на данный момент.</dd> </dl> 
 *
 * @generated from enum wb.marketplace.v1.Warehouse.CargoTypeEnum
 */
export enum Warehouse_CargoTypeEnum {
  /**
   * @generated from enum value: CargoTypeEnum__1 = 0;
   */
  CargoTypeEnum__1 = 0,

  /**
   * @generated from enum value: CargoTypeEnum__2 = 1;
   */
  CargoTypeEnum__2 = 1,

  /**
   * @generated from enum value: CargoTypeEnum__3 = 2;
   */
  CargoTypeEnum__3 = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Warehouse_CargoTypeEnum)
proto3.util.setEnumType(Warehouse_CargoTypeEnum, "wb.marketplace.v1.Warehouse.CargoTypeEnum", [
  { no: 0, name: "CargoTypeEnum__1" },
  { no: 1, name: "CargoTypeEnum__2" },
  { no: 2, name: "CargoTypeEnum__3" },
]);

/**
 * <dl> <dt>Тип доставки, который принимает склад:</dt> <dd>1 - доставка на склад Wildberries</dd> <dd>2 - доставка силами продавца</dd> <dd>3 - доставка курьером WB</dd> </dl> 
 *
 * @generated from enum wb.marketplace.v1.Warehouse.DeliveryTypeEnum
 */
export enum Warehouse_DeliveryTypeEnum {
  /**
   * @generated from enum value: DeliveryTypeEnum__1 = 0;
   */
  DeliveryTypeEnum__1 = 0,

  /**
   * @generated from enum value: DeliveryTypeEnum__2 = 1;
   */
  DeliveryTypeEnum__2 = 1,

  /**
   * @generated from enum value: DeliveryTypeEnum__3 = 2;
   */
  DeliveryTypeEnum__3 = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Warehouse_DeliveryTypeEnum)
proto3.util.setEnumType(Warehouse_DeliveryTypeEnum, "wb.marketplace.v1.Warehouse.DeliveryTypeEnum", [
  { no: 0, name: "DeliveryTypeEnum__1" },
  { no: 1, name: "DeliveryTypeEnum__2" },
  { no: 2, name: "DeliveryTypeEnum__3" },
]);

