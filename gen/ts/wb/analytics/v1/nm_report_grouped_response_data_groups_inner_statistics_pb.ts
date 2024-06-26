//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/nm_report_grouped_response_data_groups_inner_statistics.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { NmReportGroupedResponseDataGroupsInnerStatisticsSelectedPeriod } from "./nm_report_grouped_response_data_groups_inner_statistics_selected_period_pb.js";
import { NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod } from "./nm_report_grouped_response_data_groups_inner_statistics_previous_period_pb.js";
import { NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparison } from "./nm_report_grouped_response_data_groups_inner_statistics_period_comparison_pb.js";

/**
 * @generated from message wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatistics
 */
export class NmReportGroupedResponseDataGroupsInnerStatistics extends Message<NmReportGroupedResponseDataGroupsInnerStatistics> {
  /**
   * @generated from field: wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsSelectedPeriod selectedPeriod = 1;
   */
  selectedPeriod?: NmReportGroupedResponseDataGroupsInnerStatisticsSelectedPeriod;

  /**
   * @generated from field: wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod previousPeriod = 2;
   */
  previousPeriod?: NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod;

  /**
   * @generated from field: wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparison periodComparison = 3;
   */
  periodComparison?: NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparison;

  constructor(data?: PartialMessage<NmReportGroupedResponseDataGroupsInnerStatistics>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.NmReportGroupedResponseDataGroupsInnerStatistics";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "selectedPeriod", kind: "message", T: NmReportGroupedResponseDataGroupsInnerStatisticsSelectedPeriod },
    { no: 2, name: "previousPeriod", kind: "message", T: NmReportGroupedResponseDataGroupsInnerStatisticsPreviousPeriod },
    { no: 3, name: "periodComparison", kind: "message", T: NmReportGroupedResponseDataGroupsInnerStatisticsPeriodComparison },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NmReportGroupedResponseDataGroupsInnerStatistics {
    return new NmReportGroupedResponseDataGroupsInnerStatistics().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NmReportGroupedResponseDataGroupsInnerStatistics {
    return new NmReportGroupedResponseDataGroupsInnerStatistics().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NmReportGroupedResponseDataGroupsInnerStatistics {
    return new NmReportGroupedResponseDataGroupsInnerStatistics().fromJsonString(jsonString, options);
  }

  static equals(a: NmReportGroupedResponseDataGroupsInnerStatistics | PlainMessage<NmReportGroupedResponseDataGroupsInnerStatistics> | undefined, b: NmReportGroupedResponseDataGroupsInnerStatistics | PlainMessage<NmReportGroupedResponseDataGroupsInnerStatistics> | undefined): boolean {
    return proto3.util.equals(NmReportGroupedResponseDataGroupsInnerStatistics, a, b);
  }
}

