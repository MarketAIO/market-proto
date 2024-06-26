//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/adv_v1_adverts_get200_response_inner.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message wb.promotion.v1.AdvV1AdvertsGet200ResponseInner
 */
export class AdvV1AdvertsGet200ResponseInner extends Message<AdvV1AdvertsGet200ResponseInner> {
  /**
   * Идентификатор медиакампании
   *
   * @generated from field: int32 advertId = 1;
   */
  advertId = 0;

  /**
   * Название медиакампании
   *
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * Название бренда
   *
   * @generated from field: string brand = 3;
   */
  brand = "";

  /**
   * <dl> <dt>Тип медиакампании:</dt> <dd><code>1</code> - размещение по дням</dd> <dd><code>2</code> - размещение по просмотрам</dd> </dl> 
   *
   * @generated from field: int32 type = 4;
   */
  type = 0;

  /**
   * <dl> <dt>Статус медиакампании:</dt>   <dd><code>1</code> - черновик</dd>   <dd><code>2</code> - модерация   <dd><code>3</code> - отклонено (с возможностью вернуть на модерацию)</dd>   <dd><code>4</code> - одобрено</dd>   <dd><code>5</code> - запланировано</dd>   <dd><code>6</code> - на показах</dd>   <dd><code>7</code> - завершено</dd>   <dd><code>8</code> - отказался</dd>   <dd><code>9</code> - приостановлена продавцом</dd>   <dd><code>10</code> - пауза по дневному лимиту</dd>   <dd><code>11</code> - пауза по расходу бюджета</dd> </dl> 
   *
   * @generated from field: int32 status = 5;
   */
  status = 0;

  /**
   * Время создания медиакампании
   *
   * @generated from field: string createTime = 6;
   */
  createTime = "";

  /**
   * Время завершения медиакампании (при наличии)
   *
   * @generated from field: string endTime = 7;
   */
  endTime = "";

  constructor(data?: PartialMessage<AdvV1AdvertsGet200ResponseInner>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.promotion.v1.AdvV1AdvertsGet200ResponseInner";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "advertId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "brand", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "type", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "status", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "createTime", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "endTime", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvV1AdvertsGet200ResponseInner {
    return new AdvV1AdvertsGet200ResponseInner().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvV1AdvertsGet200ResponseInner {
    return new AdvV1AdvertsGet200ResponseInner().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvV1AdvertsGet200ResponseInner {
    return new AdvV1AdvertsGet200ResponseInner().fromJsonString(jsonString, options);
  }

  static equals(a: AdvV1AdvertsGet200ResponseInner | PlainMessage<AdvV1AdvertsGet200ResponseInner> | undefined, b: AdvV1AdvertsGet200ResponseInner | PlainMessage<AdvV1AdvertsGet200ResponseInner> | undefined): boolean {
    return proto3.util.equals(AdvV1AdvertsGet200ResponseInner, a, b);
  }
}

