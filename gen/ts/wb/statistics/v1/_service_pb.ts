//
//Описание API Статистики
//
//С помощью этих методов можно получить отчёты.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/statistics/v1/_service.proto (package wb.statistics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
    BinaryReadOptions,
    FieldList,
    JsonReadOptions,
    JsonValue,
    PartialMessage,
    PlainMessage
} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";
import {IncomesItem} from "./incomes_item_pb.js";
import {OrdersItem} from "./orders_item_pb.js";
import {SalesItem} from "./sales_item_pb.js";
import {StocksItem} from "./stocks_item_pb.js";
import {DetailReportItem} from "./detail_report_item_pb.js";

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierIncomesGetRequest
 */
export class ApiV1SupplierIncomesGetRequest extends Message<ApiV1SupplierIncomesGetRequest> {
  /**
   * Дата и время последнего изменения по поставке. <br> Дата в формате RFC3339. Можно передать дату или дату со временем.  Время можно указывать с точностью до секунд или миллисекунд. <br> Время передаётся в часовом поясе Мск (UTC+3). <br>Примеры: <ul> <li> `2019-06-20` <li> `2019-06-20T23:59:59` <li> `2019-06-20T00:00:00.12345` <li> `2017-03-25T00:00:00` </ul> 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  constructor(data?: PartialMessage<ApiV1SupplierIncomesGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierIncomesGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierIncomesGetRequest {
    return new ApiV1SupplierIncomesGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierIncomesGetRequest {
    return new ApiV1SupplierIncomesGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierIncomesGetRequest {
    return new ApiV1SupplierIncomesGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierIncomesGetRequest | PlainMessage<ApiV1SupplierIncomesGetRequest> | undefined, b: ApiV1SupplierIncomesGetRequest | PlainMessage<ApiV1SupplierIncomesGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierIncomesGetRequest, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierIncomesGetResponse
 */
export class ApiV1SupplierIncomesGetResponse extends Message<ApiV1SupplierIncomesGetResponse> {
  /**
   * @generated from field: repeated wb.statistics.v1.IncomesItem data = 1;
   */
  data: IncomesItem[] = [];

  constructor(data?: PartialMessage<ApiV1SupplierIncomesGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierIncomesGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: IncomesItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierIncomesGetResponse {
    return new ApiV1SupplierIncomesGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierIncomesGetResponse {
    return new ApiV1SupplierIncomesGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierIncomesGetResponse {
    return new ApiV1SupplierIncomesGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierIncomesGetResponse | PlainMessage<ApiV1SupplierIncomesGetResponse> | undefined, b: ApiV1SupplierIncomesGetResponse | PlainMessage<ApiV1SupplierIncomesGetResponse> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierIncomesGetResponse, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierOrdersGetRequest
 */
export class ApiV1SupplierOrdersGetRequest extends Message<ApiV1SupplierOrdersGetRequest> {
  /**
   * Дата и время последнего изменения по заказу. <br> Дата в формате RFC3339. Можно передать дату или дату со временем.  Время можно указывать с точностью до секунд или миллисекунд. <br> Время передаётся в часовом поясе Мск (UTC+3). <br>Примеры: <ul> <li> `2019-06-20` <li> `2019-06-20T23:59:59` <li> `2019-06-20T00:00:00.12345` <li> `2017-03-25T00:00:00` </ul> 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Если параметр `flag=0` (или не указан в строке запроса), при вызове API возвращаются данные,  у которых значение поля `lastChangeDate` (дата время обновления информации в сервисе) больше или равно переданному  значению параметра `dateFrom`.  При этом количество возвращенных строк данных варьируется в интервале от 0 до примерно 100 000. <br> Если параметр `flag=1`, то будет выгружена информация обо всех заказах или продажах с датой,  равной переданному параметру `dateFrom` (в данном случае время в дате значения не имеет).  При этом количество возвращенных строк данных будет равно количеству всех заказов или продаж,  сделанных в указанную дату, переданную в параметре `dateFrom`. 
   *
   * @generated from field: int32 flag = 2;
   */
  flag = 0;

  constructor(data?: PartialMessage<ApiV1SupplierOrdersGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierOrdersGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "flag", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierOrdersGetRequest {
    return new ApiV1SupplierOrdersGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierOrdersGetRequest {
    return new ApiV1SupplierOrdersGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierOrdersGetRequest {
    return new ApiV1SupplierOrdersGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierOrdersGetRequest | PlainMessage<ApiV1SupplierOrdersGetRequest> | undefined, b: ApiV1SupplierOrdersGetRequest | PlainMessage<ApiV1SupplierOrdersGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierOrdersGetRequest, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierOrdersGetResponse
 */
export class ApiV1SupplierOrdersGetResponse extends Message<ApiV1SupplierOrdersGetResponse> {
  /**
   * @generated from field: repeated wb.statistics.v1.OrdersItem data = 1;
   */
  data: OrdersItem[] = [];

  constructor(data?: PartialMessage<ApiV1SupplierOrdersGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierOrdersGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: OrdersItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierOrdersGetResponse {
    return new ApiV1SupplierOrdersGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierOrdersGetResponse {
    return new ApiV1SupplierOrdersGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierOrdersGetResponse {
    return new ApiV1SupplierOrdersGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierOrdersGetResponse | PlainMessage<ApiV1SupplierOrdersGetResponse> | undefined, b: ApiV1SupplierOrdersGetResponse | PlainMessage<ApiV1SupplierOrdersGetResponse> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierOrdersGetResponse, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierSalesGetRequest
 */
export class ApiV1SupplierSalesGetRequest extends Message<ApiV1SupplierSalesGetRequest> {
  /**
   * Дата и время последнего изменения по продаже/возврату. <br> Дата в формате RFC3339. Можно передать дату или дату со временем.  Время можно указывать с точностью до секунд или миллисекунд. <br> Время передаётся в часовом поясе Мск (UTC+3). <br>Примеры: <ul> <li> `2019-06-20` <li> `2019-06-20T23:59:59` <li> `2019-06-20T00:00:00.12345` <li> `2017-03-25T00:00:00` </ul> 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Если параметр `flag=0` (или не указан в строке запроса), при вызове API возвращаются данные,  у которых значение поля `lastChangeDate` (дата время обновления информации в сервисе) больше или равно переданному  значению параметра `dateFrom`.  При этом количество возвращенных строк данных варьируется в интервале от 0 до примерно 100 000. <br> Если параметр `flag=1`, то будет выгружена информация обо всех заказах или продажах с датой,  равной переданному параметру `dateFrom` (в данном случае время в дате значения не имеет).  При этом количество возвращенных строк данных будет равно количеству всех заказов или продаж,  сделанных в указанную дату, переданную в параметре `dateFrom`. 
   *
   * @generated from field: int32 flag = 2;
   */
  flag = 0;

  constructor(data?: PartialMessage<ApiV1SupplierSalesGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierSalesGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "flag", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierSalesGetRequest {
    return new ApiV1SupplierSalesGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierSalesGetRequest {
    return new ApiV1SupplierSalesGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierSalesGetRequest {
    return new ApiV1SupplierSalesGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierSalesGetRequest | PlainMessage<ApiV1SupplierSalesGetRequest> | undefined, b: ApiV1SupplierSalesGetRequest | PlainMessage<ApiV1SupplierSalesGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierSalesGetRequest, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierSalesGetResponse
 */
export class ApiV1SupplierSalesGetResponse extends Message<ApiV1SupplierSalesGetResponse> {
  /**
   * @generated from field: repeated wb.statistics.v1.SalesItem data = 1;
   */
  data: SalesItem[] = [];

  constructor(data?: PartialMessage<ApiV1SupplierSalesGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierSalesGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: SalesItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierSalesGetResponse {
    return new ApiV1SupplierSalesGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierSalesGetResponse {
    return new ApiV1SupplierSalesGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierSalesGetResponse {
    return new ApiV1SupplierSalesGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierSalesGetResponse | PlainMessage<ApiV1SupplierSalesGetResponse> | undefined, b: ApiV1SupplierSalesGetResponse | PlainMessage<ApiV1SupplierSalesGetResponse> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierSalesGetResponse, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierStocksGetRequest
 */
export class ApiV1SupplierStocksGetRequest extends Message<ApiV1SupplierStocksGetRequest> {
  /**
   * Дата и время последнего изменения по товару. <br> Для получения полного остатка следует указывать максимально раннее значение. <br> Например, `2019-06-20` <br> Дата в формате RFC3339. Можно передать дату или дату со временем.  Время можно указывать с точностью до секунд или миллисекунд. <br> Время передаётся в часовом поясе Мск (UTC+3). <br>Примеры: <ul> <li> `2019-06-20` <li> `2019-06-20T23:59:59` <li> `2019-06-20T00:00:00.12345` <li> `2017-03-25T00:00:00` </ul> 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  constructor(data?: PartialMessage<ApiV1SupplierStocksGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierStocksGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierStocksGetRequest {
    return new ApiV1SupplierStocksGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierStocksGetRequest {
    return new ApiV1SupplierStocksGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierStocksGetRequest {
    return new ApiV1SupplierStocksGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierStocksGetRequest | PlainMessage<ApiV1SupplierStocksGetRequest> | undefined, b: ApiV1SupplierStocksGetRequest | PlainMessage<ApiV1SupplierStocksGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierStocksGetRequest, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV1SupplierStocksGetResponse
 */
export class ApiV1SupplierStocksGetResponse extends Message<ApiV1SupplierStocksGetResponse> {
  /**
   * @generated from field: repeated wb.statistics.v1.StocksItem data = 1;
   */
  data: StocksItem[] = [];

  constructor(data?: PartialMessage<ApiV1SupplierStocksGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV1SupplierStocksGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: StocksItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1SupplierStocksGetResponse {
    return new ApiV1SupplierStocksGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1SupplierStocksGetResponse {
    return new ApiV1SupplierStocksGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1SupplierStocksGetResponse {
    return new ApiV1SupplierStocksGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1SupplierStocksGetResponse | PlainMessage<ApiV1SupplierStocksGetResponse> | undefined, b: ApiV1SupplierStocksGetResponse | PlainMessage<ApiV1SupplierStocksGetResponse> | undefined): boolean {
    return proto3.util.equals(ApiV1SupplierStocksGetResponse, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV5SupplierReportDetailByPeriodGetRequest
 */
export class ApiV5SupplierReportDetailByPeriodGetRequest extends Message<ApiV5SupplierReportDetailByPeriodGetRequest> {
  /**
   * Начальная дата отчета.<br> Дата в формате RFC3339. Можно передать дату или дату со временем.  Время можно указывать с точностью до секунд или миллисекунд. <br> Время передаётся в часовом поясе Мск (UTC+3). <br>Примеры: <ul> <li> `2019-06-20` <li> `2019-06-20T23:59:59` <li> `2019-06-20T00:00:00.12345` <li> `2017-03-25T00:00:00` </ul> 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Конечная дата отчета
   *
   * @generated from field: string dateTo = 2;
   */
  dateTo = "";

  /**
   * Максимальное количество строк отчета, возвращаемых методом. Не может быть более 100000.
   *
   * @generated from field: int32 limit = 3;
   */
  limit = 0;

  /**
   * Уникальный идентификатор строки отчета. Необходим для получения отчета частями.  <br> Загрузку отчета нужно начинать с `rrdid = 0` и при последующих вызовах API передавать в запросе значение `rrd_id` из последней строки, полученной в результате предыдущего вызова.  <br> Таким образом для загрузки одного отчета может понадобиться вызывать API до тех пор, пока количество возвращаемых строк не станет равным нулю. 
   *
   * @generated from field: int32 rrdid = 4;
   */
  rrdid = 0;

  constructor(data?: PartialMessage<ApiV5SupplierReportDetailByPeriodGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV5SupplierReportDetailByPeriodGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dateTo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "rrdid", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV5SupplierReportDetailByPeriodGetRequest {
    return new ApiV5SupplierReportDetailByPeriodGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV5SupplierReportDetailByPeriodGetRequest {
    return new ApiV5SupplierReportDetailByPeriodGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV5SupplierReportDetailByPeriodGetRequest {
    return new ApiV5SupplierReportDetailByPeriodGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV5SupplierReportDetailByPeriodGetRequest | PlainMessage<ApiV5SupplierReportDetailByPeriodGetRequest> | undefined, b: ApiV5SupplierReportDetailByPeriodGetRequest | PlainMessage<ApiV5SupplierReportDetailByPeriodGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV5SupplierReportDetailByPeriodGetRequest, a, b);
  }
}

/**
 * @generated from message wb.statistics.v1.ApiV5SupplierReportDetailByPeriodGetResponse
 */
export class ApiV5SupplierReportDetailByPeriodGetResponse extends Message<ApiV5SupplierReportDetailByPeriodGetResponse> {
  /**
   * @generated from field: repeated wb.statistics.v1.DetailReportItem data = 1;
   */
  data: DetailReportItem[] = [];

  constructor(data?: PartialMessage<ApiV5SupplierReportDetailByPeriodGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.statistics.v1.ApiV5SupplierReportDetailByPeriodGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: DetailReportItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV5SupplierReportDetailByPeriodGetResponse {
    return new ApiV5SupplierReportDetailByPeriodGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV5SupplierReportDetailByPeriodGetResponse {
    return new ApiV5SupplierReportDetailByPeriodGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV5SupplierReportDetailByPeriodGetResponse {
    return new ApiV5SupplierReportDetailByPeriodGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV5SupplierReportDetailByPeriodGetResponse | PlainMessage<ApiV5SupplierReportDetailByPeriodGetResponse> | undefined, b: ApiV5SupplierReportDetailByPeriodGetResponse | PlainMessage<ApiV5SupplierReportDetailByPeriodGetResponse> | undefined): boolean {
    return proto3.util.equals(ApiV5SupplierReportDetailByPeriodGetResponse, a, b);
  }
}
