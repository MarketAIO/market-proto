//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/daily_stats2_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { DailyStats2InnerAppTypeStatsInner } from "./daily_stats2_inner_app_type_stats_inner_pb.js";

/**
 * @generated from message wb.promotion.v1.DailyStats2Inner
 */
export class DailyStats2Inner extends Message<DailyStats2Inner> {
  /**
   * Дата
   *
   * @generated from field: string date = 1;
   */
  date = "";

  /**
   * Статистика по платформам
   *
   * @generated from field: repeated wb.promotion.v1.DailyStats2InnerAppTypeStatsInner app_type_stats = 2;
   */
  appTypeStats: DailyStats2InnerAppTypeStatsInner[] = [];

  constructor(data?: PartialMessage<DailyStats2Inner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.DailyStats2Inner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "app_type_stats", kind: "message", T: DailyStats2InnerAppTypeStatsInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DailyStats2Inner {
    return new DailyStats2Inner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DailyStats2Inner {
    return new DailyStats2Inner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DailyStats2Inner {
    return new DailyStats2Inner().fromJsonString(jsonString, options);
  }

  static equals(a: DailyStats2Inner | PlainMessage<DailyStats2Inner> | undefined, b: DailyStats2Inner | PlainMessage<DailyStats2Inner> | undefined): boolean {
    return proto3.util.equals(DailyStats2Inner, a, b);
  }
}
