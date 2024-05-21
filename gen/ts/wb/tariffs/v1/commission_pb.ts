//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/tariffs/v1/commission.proto (package wb.tariffs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { CommissionReportInner } from "./commission_report_inner_pb.js";

/**
 * @generated from message wb.tariffs.v1.Commission
 */
export class Commission extends Message<Commission> {
  /**
   * Список комиссий
   *
   * @generated from field: repeated wb.tariffs.v1.CommissionReportInner report = 1;
   */
  report: CommissionReportInner[] = [];

  constructor(data?: PartialMessage<Commission>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.tariffs.v1.Commission";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "report", kind: "message", T: CommissionReportInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Commission {
    return new Commission().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Commission {
    return new Commission().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Commission {
    return new Commission().fromJsonString(jsonString, options);
  }

  static equals(a: Commission | PlainMessage<Commission> | undefined, b: Commission | PlainMessage<Commission> | undefined): boolean {
    return proto3.util.equals(Commission, a, b);
  }
}

