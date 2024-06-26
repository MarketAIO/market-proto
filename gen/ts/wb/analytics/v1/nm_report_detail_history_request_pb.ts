//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/nm_report_detail_history_request.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { NmReportDetailHistoryReqPeriod } from "./nm_report_detail_history_request_period_pb.js";

/**
 * @generated from message wb.analytics.v1.NmReportDetailHistoryReq
 */
export class NmReportDetailHistoryReq extends Message<NmReportDetailHistoryReq> {
  /**
   * Артикул Wildberries (максимум 20)
   *
   * @generated from field: repeated int32 nmIDs = 1;
   */
  nmIDs: number[] = [];

  /**
   * @generated from field: wb.analytics.v1.NmReportDetailHistoryReqPeriod period = 2;
   */
  period?: NmReportDetailHistoryReqPeriod;

  /**
   * Временная зона.<br> Если не указано, то по умолчанию используется Europe/Moscow. 
   *
   * @generated from field: string timezone = 3;
   */
  timezone = "";

  /**
   * Тип агрегации. Если не указано, то по умолчанию используется агрегация по дням. <br> Доступные уровни агрегации `day`, `week` 
   *
   * @generated from field: string aggregationLevel = 4;
   */
  aggregationLevel = "";

  constructor(data?: PartialMessage<NmReportDetailHistoryReq>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.NmReportDetailHistoryReq";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmIDs", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 2, name: "period", kind: "message", T: NmReportDetailHistoryReqPeriod },
    { no: 3, name: "timezone", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "aggregationLevel", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NmReportDetailHistoryReq {
    return new NmReportDetailHistoryReq().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NmReportDetailHistoryReq {
    return new NmReportDetailHistoryReq().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NmReportDetailHistoryReq {
    return new NmReportDetailHistoryReq().fromJsonString(jsonString, options);
  }

  static equals(a: NmReportDetailHistoryReq | PlainMessage<NmReportDetailHistoryReq> | undefined, b: NmReportDetailHistoryReq | PlainMessage<NmReportDetailHistoryReq> | undefined): boolean {
    return proto3.util.equals(NmReportDetailHistoryReq, a, b);
  }
}

