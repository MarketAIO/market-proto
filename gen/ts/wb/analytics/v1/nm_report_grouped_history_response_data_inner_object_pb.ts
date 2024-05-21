//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/nm_report_grouped_history_response_data_inner_object.proto (package wb.analytics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.analytics.v1.NmReportGroupedHistoryResponseDataInnerObject
 */
export class NmReportGroupedHistoryResponseDataInnerObject extends Message<NmReportGroupedHistoryResponseDataInnerObject> {
  /**
   * Идентификатор предмета
   *
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * Название предмета
   *
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<NmReportGroupedHistoryResponseDataInnerObject>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.NmReportGroupedHistoryResponseDataInnerObject";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NmReportGroupedHistoryResponseDataInnerObject {
    return new NmReportGroupedHistoryResponseDataInnerObject().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NmReportGroupedHistoryResponseDataInnerObject {
    return new NmReportGroupedHistoryResponseDataInnerObject().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NmReportGroupedHistoryResponseDataInnerObject {
    return new NmReportGroupedHistoryResponseDataInnerObject().fromJsonString(jsonString, options);
  }

  static equals(a: NmReportGroupedHistoryResponseDataInnerObject | PlainMessage<NmReportGroupedHistoryResponseDataInnerObject> | undefined, b: NmReportGroupedHistoryResponseDataInnerObject | PlainMessage<NmReportGroupedHistoryResponseDataInnerObject> | undefined): boolean {
    return proto3.util.equals(NmReportGroupedHistoryResponseDataInnerObject, a, b);
  }
}
