//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/tariffs/v1/commission_uae.proto (package wb.tariffs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { CommissionUAEReportInner } from "./commission_uae_report_inner_pb.js";

/**
 * @generated from message wb.tariffs.v1.CommissionUAE
 */
export class CommissionUAE extends Message<CommissionUAE> {
  /**
   * Список комиссий
   *
   * @generated from field: repeated wb.tariffs.v1.CommissionUAEReportInner report = 1;
   */
  report: CommissionUAEReportInner[] = [];

  constructor(data?: PartialMessage<CommissionUAE>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.tariffs.v1.CommissionUAE";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "report", kind: "message", T: CommissionUAEReportInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CommissionUAE {
    return new CommissionUAE().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CommissionUAE {
    return new CommissionUAE().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CommissionUAE {
    return new CommissionUAE().fromJsonString(jsonString, options);
  }

  static equals(a: CommissionUAE | PlainMessage<CommissionUAE> | undefined, b: CommissionUAE | PlainMessage<CommissionUAE> | undefined): boolean {
    return proto3.util.equals(CommissionUAE, a, b);
  }
}

