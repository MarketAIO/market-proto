//
//Тарифы
//
//<p>Комиссии и тарифы на услуги можно получить с любым токеном, у которого не выбрана опция Тестовый контур.</p>
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/tariffs/v1/commission_turkey.proto (package wb.tariffs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { CommissionTurkeyReportInner } from "./commission_turkey_report_inner_pb.js";

/**
 * @generated from message wb.tariffs.v1.CommissionTurkey
 */
export class CommissionTurkey extends Message<CommissionTurkey> {
  /**
   * Список комиссий
   *
   * @generated from field: repeated wb.tariffs.v1.CommissionTurkeyReportInner report = 1;
   */
  report: CommissionTurkeyReportInner[] = [];

  constructor(data?: PartialMessage<CommissionTurkey>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.tariffs.v1.CommissionTurkey";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "report", kind: "message", T: CommissionTurkeyReportInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CommissionTurkey {
    return new CommissionTurkey().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CommissionTurkey {
    return new CommissionTurkey().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CommissionTurkey {
    return new CommissionTurkey().fromJsonString(jsonString, options);
  }

  static equals(a: CommissionTurkey | PlainMessage<CommissionTurkey> | undefined, b: CommissionTurkey | PlainMessage<CommissionTurkey> | undefined): boolean {
    return proto3.util.equals(CommissionTurkey, a, b);
  }
}

