//
//Описание API Аналитика
//
//Сервис предоставляет публичный API для получения аналитических данных. С помощью этих методов вы можете получать аналитические отчёты. [Часть методов](./#tag/Voronka-prodazh-(Dzhem)/) доступна только с [подпиской на расширенную аналитику Джем](https://seller.wildberries.ru/dynamic-product-categories/jam). 
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file wb/analytics/v1/_service.proto (package wb.analytics.v1, syntax proto3)
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
import {ExciseReportReq} from "./excise_report_request_pb.js";
import {ResponsePaidStorageInner} from "./response_paid_storage_inner_pb.js";
import {NmReportDetailHistoryReq} from "./nm_report_detail_history_request_pb.js";
import {NmReportDetailReq} from "./nm_report_detail_request_pb.js";
import {ApiV2NmReportDownloadsPostReq} from "./api_v2_nm_report_downloads_post_request_pb.js";
import {NmReportRetryReportReq} from "./nm_report_retry_report_request_pb.js";
import {NmReportGroupedHistoryReq} from "./nm_report_grouped_history_request_pb.js";
import {NmReportGroupedReq} from "./nm_report_grouped_request_pb.js";

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsAcceptanceReportGetRequest
 */
export class ApiV1AnalyticsAcceptanceReportGetRequest extends Message<ApiV1AnalyticsAcceptanceReportGetRequest> {
  /**
   * Начало отчётного периода, `ГГГГ-ММ-ДД` 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Конец отчётного периода, `ГГГГ-ММ-ДД` 
   *
   * @generated from field: string dateTo = 2;
   */
  dateTo = "";

  constructor(data?: PartialMessage<ApiV1AnalyticsAcceptanceReportGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsAcceptanceReportGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dateTo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsAcceptanceReportGetRequest {
    return new ApiV1AnalyticsAcceptanceReportGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAcceptanceReportGetRequest {
    return new ApiV1AnalyticsAcceptanceReportGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAcceptanceReportGetRequest {
    return new ApiV1AnalyticsAcceptanceReportGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsAcceptanceReportGetRequest | PlainMessage<ApiV1AnalyticsAcceptanceReportGetRequest> | undefined, b: ApiV1AnalyticsAcceptanceReportGetRequest | PlainMessage<ApiV1AnalyticsAcceptanceReportGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsAcceptanceReportGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsAntifraudDetailsGetRequest
 */
export class ApiV1AnalyticsAntifraudDetailsGetRequest extends Message<ApiV1AnalyticsAntifraudDetailsGetRequest> {
  /**
   * Дата, которая входит в отчётный период, `ГГГГ-ММ-ДД`.  <br/>          Чтобы получить данные за всё время с августа 2023,  не указывайте этот параметр 
   *
   * @generated from field: string date = 1;
   */
  date = "";

  constructor(data?: PartialMessage<ApiV1AnalyticsAntifraudDetailsGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsAntifraudDetailsGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsAntifraudDetailsGetRequest {
    return new ApiV1AnalyticsAntifraudDetailsGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAntifraudDetailsGetRequest {
    return new ApiV1AnalyticsAntifraudDetailsGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsAntifraudDetailsGetRequest {
    return new ApiV1AnalyticsAntifraudDetailsGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsAntifraudDetailsGetRequest | PlainMessage<ApiV1AnalyticsAntifraudDetailsGetRequest> | undefined, b: ApiV1AnalyticsAntifraudDetailsGetRequest | PlainMessage<ApiV1AnalyticsAntifraudDetailsGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsAntifraudDetailsGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsExciseReportPostRequest
 */
export class ApiV1AnalyticsExciseReportPostRequest extends Message<ApiV1AnalyticsExciseReportPostRequest> {
  /**
   * Начало отчётного периода в формате RFC3339. Можно передать дату  или дату со временем. Примеры:    * `2023-12-01`   * `2023-12-01T23:59:59`   * `2023-12-01T00:00:00.12345`   * `2023-12-01T00:00:00` 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Конец отчётного периода в формате RFC3339. Можно передать дату  или дату со временем. Примеры:    * `2023-12-01`   * `2023-12-01T23:59:59`   * `2023-12-01T00:00:00.12345`   * `2023-12-01T00:00:00` 
   *
   * @generated from field: string dateTo = 2;
   */
  dateTo = "";

  /**
   * @generated from field: wb.analytics.v1.ExciseReportReq exciseReportReq = 3;
   */
  exciseReportReq?: ExciseReportReq;

  constructor(data?: PartialMessage<ApiV1AnalyticsExciseReportPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsExciseReportPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dateTo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "exciseReportReq", kind: "message", T: ExciseReportReq },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsExciseReportPostRequest {
    return new ApiV1AnalyticsExciseReportPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsExciseReportPostRequest {
    return new ApiV1AnalyticsExciseReportPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsExciseReportPostRequest {
    return new ApiV1AnalyticsExciseReportPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsExciseReportPostRequest | PlainMessage<ApiV1AnalyticsExciseReportPostRequest> | undefined, b: ApiV1AnalyticsExciseReportPostRequest | PlainMessage<ApiV1AnalyticsExciseReportPostRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsExciseReportPostRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsGoodsLabelingGetRequest
 */
export class ApiV1AnalyticsGoodsLabelingGetRequest extends Message<ApiV1AnalyticsGoodsLabelingGetRequest> {
  /**
   * Начало отчётного периода, `ГГГГ-ММ-ДД` 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Конец отчётного периода, `ГГГГ-ММ-ДД` 
   *
   * @generated from field: string dateTo = 2;
   */
  dateTo = "";

  constructor(data?: PartialMessage<ApiV1AnalyticsGoodsLabelingGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsGoodsLabelingGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dateTo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsGoodsLabelingGetRequest {
    return new ApiV1AnalyticsGoodsLabelingGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsGoodsLabelingGetRequest {
    return new ApiV1AnalyticsGoodsLabelingGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsGoodsLabelingGetRequest {
    return new ApiV1AnalyticsGoodsLabelingGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsGoodsLabelingGetRequest | PlainMessage<ApiV1AnalyticsGoodsLabelingGetRequest> | undefined, b: ApiV1AnalyticsGoodsLabelingGetRequest | PlainMessage<ApiV1AnalyticsGoodsLabelingGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsGoodsLabelingGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsIncorrectAttachmentsGetRequest
 */
export class ApiV1AnalyticsIncorrectAttachmentsGetRequest extends Message<ApiV1AnalyticsIncorrectAttachmentsGetRequest> {
  /**
   * Начало отчётного периода, `ГГГГ-ММ-ДД` 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Конец отчётного периода, `ГГГГ-ММ-ДД` 
   *
   * @generated from field: string dateTo = 2;
   */
  dateTo = "";

  constructor(data?: PartialMessage<ApiV1AnalyticsIncorrectAttachmentsGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsIncorrectAttachmentsGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dateTo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsIncorrectAttachmentsGetRequest {
    return new ApiV1AnalyticsIncorrectAttachmentsGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsIncorrectAttachmentsGetRequest {
    return new ApiV1AnalyticsIncorrectAttachmentsGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsIncorrectAttachmentsGetRequest {
    return new ApiV1AnalyticsIncorrectAttachmentsGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsIncorrectAttachmentsGetRequest | PlainMessage<ApiV1AnalyticsIncorrectAttachmentsGetRequest> | undefined, b: ApiV1AnalyticsIncorrectAttachmentsGetRequest | PlainMessage<ApiV1AnalyticsIncorrectAttachmentsGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsIncorrectAttachmentsGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1AnalyticsStorageCoefficientGetRequest
 */
export class ApiV1AnalyticsStorageCoefficientGetRequest extends Message<ApiV1AnalyticsStorageCoefficientGetRequest> {
  /**
   * Дата, которая входит в отчётный период, `ГГГГ-ММ-ДД` 
   *
   * @generated from field: string date = 1;
   */
  date = "";

  constructor(data?: PartialMessage<ApiV1AnalyticsStorageCoefficientGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1AnalyticsStorageCoefficientGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1AnalyticsStorageCoefficientGetRequest {
    return new ApiV1AnalyticsStorageCoefficientGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1AnalyticsStorageCoefficientGetRequest {
    return new ApiV1AnalyticsStorageCoefficientGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1AnalyticsStorageCoefficientGetRequest {
    return new ApiV1AnalyticsStorageCoefficientGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1AnalyticsStorageCoefficientGetRequest | PlainMessage<ApiV1AnalyticsStorageCoefficientGetRequest> | undefined, b: ApiV1AnalyticsStorageCoefficientGetRequest | PlainMessage<ApiV1AnalyticsStorageCoefficientGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1AnalyticsStorageCoefficientGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1PaidStorageGetRequest
 */
export class ApiV1PaidStorageGetRequest extends Message<ApiV1PaidStorageGetRequest> {
  /**
   * Начало отчётного периода в формате RFC3339. Можно передать дату или дату со временем. Примеры:    * `2019-06-20`   * `2019-06-20T23:59:59`   * `2019-06-20T00:00:00.12345`   * `2017-03-25T00:00:00` </ul> 
   *
   * @generated from field: string dateFrom = 1;
   */
  dateFrom = "";

  /**
   * Конец отчётного периода в формате RFC3339. Можно передать дату или дату со временем. Примеры:    * `2019-06-20`   * `2019-06-20T23:59:59`   * `2019-06-20T00:00:00.12345`   * `2017-03-25T00:00:00` 
   *
   * @generated from field: string dateTo = 2;
   */
  dateTo = "";

  constructor(data?: PartialMessage<ApiV1PaidStorageGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1PaidStorageGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dateFrom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dateTo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1PaidStorageGetRequest {
    return new ApiV1PaidStorageGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1PaidStorageGetRequest {
    return new ApiV1PaidStorageGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1PaidStorageGetRequest {
    return new ApiV1PaidStorageGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1PaidStorageGetRequest | PlainMessage<ApiV1PaidStorageGetRequest> | undefined, b: ApiV1PaidStorageGetRequest | PlainMessage<ApiV1PaidStorageGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1PaidStorageGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1PaidStorageTasksTaskIdDownloadGetRequest
 */
export class ApiV1PaidStorageTasksTaskIdDownloadGetRequest extends Message<ApiV1PaidStorageTasksTaskIdDownloadGetRequest> {
  /**
   * ID задания на генерацию 
   *
   * @generated from field: string taskId = 1;
   */
  taskId = "";

  constructor(data?: PartialMessage<ApiV1PaidStorageTasksTaskIdDownloadGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1PaidStorageTasksTaskIdDownloadGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "taskId", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1PaidStorageTasksTaskIdDownloadGetRequest {
    return new ApiV1PaidStorageTasksTaskIdDownloadGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1PaidStorageTasksTaskIdDownloadGetRequest {
    return new ApiV1PaidStorageTasksTaskIdDownloadGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1PaidStorageTasksTaskIdDownloadGetRequest {
    return new ApiV1PaidStorageTasksTaskIdDownloadGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1PaidStorageTasksTaskIdDownloadGetRequest | PlainMessage<ApiV1PaidStorageTasksTaskIdDownloadGetRequest> | undefined, b: ApiV1PaidStorageTasksTaskIdDownloadGetRequest | PlainMessage<ApiV1PaidStorageTasksTaskIdDownloadGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1PaidStorageTasksTaskIdDownloadGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1PaidStorageTasksTaskIdDownloadGetResponse
 */
export class ApiV1PaidStorageTasksTaskIdDownloadGetResponse extends Message<ApiV1PaidStorageTasksTaskIdDownloadGetResponse> {
  /**
   * @generated from field: repeated wb.analytics.v1.ResponsePaidStorageInner data = 1;
   */
  data: ResponsePaidStorageInner[] = [];

  constructor(data?: PartialMessage<ApiV1PaidStorageTasksTaskIdDownloadGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1PaidStorageTasksTaskIdDownloadGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "message", T: ResponsePaidStorageInner, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1PaidStorageTasksTaskIdDownloadGetResponse {
    return new ApiV1PaidStorageTasksTaskIdDownloadGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1PaidStorageTasksTaskIdDownloadGetResponse {
    return new ApiV1PaidStorageTasksTaskIdDownloadGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1PaidStorageTasksTaskIdDownloadGetResponse {
    return new ApiV1PaidStorageTasksTaskIdDownloadGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1PaidStorageTasksTaskIdDownloadGetResponse | PlainMessage<ApiV1PaidStorageTasksTaskIdDownloadGetResponse> | undefined, b: ApiV1PaidStorageTasksTaskIdDownloadGetResponse | PlainMessage<ApiV1PaidStorageTasksTaskIdDownloadGetResponse> | undefined): boolean {
    return proto3.util.equals(ApiV1PaidStorageTasksTaskIdDownloadGetResponse, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV1PaidStorageTasksTaskIdStatusGetRequest
 */
export class ApiV1PaidStorageTasksTaskIdStatusGetRequest extends Message<ApiV1PaidStorageTasksTaskIdStatusGetRequest> {
  /**
   * ID задания на генерацию 
   *
   * @generated from field: string taskId = 1;
   */
  taskId = "";

  constructor(data?: PartialMessage<ApiV1PaidStorageTasksTaskIdStatusGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV1PaidStorageTasksTaskIdStatusGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "taskId", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV1PaidStorageTasksTaskIdStatusGetRequest {
    return new ApiV1PaidStorageTasksTaskIdStatusGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV1PaidStorageTasksTaskIdStatusGetRequest {
    return new ApiV1PaidStorageTasksTaskIdStatusGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV1PaidStorageTasksTaskIdStatusGetRequest {
    return new ApiV1PaidStorageTasksTaskIdStatusGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV1PaidStorageTasksTaskIdStatusGetRequest | PlainMessage<ApiV1PaidStorageTasksTaskIdStatusGetRequest> | undefined, b: ApiV1PaidStorageTasksTaskIdStatusGetRequest | PlainMessage<ApiV1PaidStorageTasksTaskIdStatusGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV1PaidStorageTasksTaskIdStatusGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportDetailHistoryPostRequest
 */
export class ApiV2NmReportDetailHistoryPostRequest extends Message<ApiV2NmReportDetailHistoryPostRequest> {
  /**
   *
   *
   * @generated from field: wb.analytics.v1.NmReportDetailHistoryReq nmReportDetailHistoryReq = 1;
   */
  nmReportDetailHistoryReq?: NmReportDetailHistoryReq;

  constructor(data?: PartialMessage<ApiV2NmReportDetailHistoryPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportDetailHistoryPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmReportDetailHistoryReq", kind: "message", T: NmReportDetailHistoryReq },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportDetailHistoryPostRequest {
    return new ApiV2NmReportDetailHistoryPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportDetailHistoryPostRequest {
    return new ApiV2NmReportDetailHistoryPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportDetailHistoryPostRequest {
    return new ApiV2NmReportDetailHistoryPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportDetailHistoryPostRequest | PlainMessage<ApiV2NmReportDetailHistoryPostRequest> | undefined, b: ApiV2NmReportDetailHistoryPostRequest | PlainMessage<ApiV2NmReportDetailHistoryPostRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportDetailHistoryPostRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportDetailPostRequest
 */
export class ApiV2NmReportDetailPostRequest extends Message<ApiV2NmReportDetailPostRequest> {
  /**
   *
   *
   * @generated from field: wb.analytics.v1.NmReportDetailReq nmReportDetailReq = 1;
   */
  nmReportDetailReq?: NmReportDetailReq;

  constructor(data?: PartialMessage<ApiV2NmReportDetailPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportDetailPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmReportDetailReq", kind: "message", T: NmReportDetailReq },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportDetailPostRequest {
    return new ApiV2NmReportDetailPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportDetailPostRequest {
    return new ApiV2NmReportDetailPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportDetailPostRequest {
    return new ApiV2NmReportDetailPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportDetailPostRequest | PlainMessage<ApiV2NmReportDetailPostRequest> | undefined, b: ApiV2NmReportDetailPostRequest | PlainMessage<ApiV2NmReportDetailPostRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportDetailPostRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportDownloadsFileDownloadIdGetRequest
 */
export class ApiV2NmReportDownloadsFileDownloadIdGetRequest extends Message<ApiV2NmReportDownloadsFileDownloadIdGetRequest> {
  /**
   * ID отчёта
   *
   * @generated from field: string downloadId = 1;
   */
  downloadId = "";

  constructor(data?: PartialMessage<ApiV2NmReportDownloadsFileDownloadIdGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportDownloadsFileDownloadIdGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "downloadId", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportDownloadsFileDownloadIdGetRequest {
    return new ApiV2NmReportDownloadsFileDownloadIdGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsFileDownloadIdGetRequest {
    return new ApiV2NmReportDownloadsFileDownloadIdGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsFileDownloadIdGetRequest {
    return new ApiV2NmReportDownloadsFileDownloadIdGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportDownloadsFileDownloadIdGetRequest | PlainMessage<ApiV2NmReportDownloadsFileDownloadIdGetRequest> | undefined, b: ApiV2NmReportDownloadsFileDownloadIdGetRequest | PlainMessage<ApiV2NmReportDownloadsFileDownloadIdGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportDownloadsFileDownloadIdGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportDownloadsFileDownloadIdGetResponse
 */
export class ApiV2NmReportDownloadsFileDownloadIdGetResponse extends Message<ApiV2NmReportDownloadsFileDownloadIdGetResponse> {
  /**
   * @generated from field: string data = 1;
   */
  data = "";

  constructor(data?: PartialMessage<ApiV2NmReportDownloadsFileDownloadIdGetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportDownloadsFileDownloadIdGetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "data", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportDownloadsFileDownloadIdGetResponse {
    return new ApiV2NmReportDownloadsFileDownloadIdGetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsFileDownloadIdGetResponse {
    return new ApiV2NmReportDownloadsFileDownloadIdGetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsFileDownloadIdGetResponse {
    return new ApiV2NmReportDownloadsFileDownloadIdGetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportDownloadsFileDownloadIdGetResponse | PlainMessage<ApiV2NmReportDownloadsFileDownloadIdGetResponse> | undefined, b: ApiV2NmReportDownloadsFileDownloadIdGetResponse | PlainMessage<ApiV2NmReportDownloadsFileDownloadIdGetResponse> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportDownloadsFileDownloadIdGetResponse, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportDownloadsGetRequest
 */
export class ApiV2NmReportDownloadsGetRequest extends Message<ApiV2NmReportDownloadsGetRequest> {
  /**
   * ID отчёта
   *
   * @generated from field: repeated string filterLeft_Square_BracketdownloadIdsRight_Square_Bracket = 1;
   */
  filterLeftSquareBracketdownloadIdsRightSquareBracket: string[] = [];

  constructor(data?: PartialMessage<ApiV2NmReportDownloadsGetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportDownloadsGetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "filterLeft_Square_BracketdownloadIdsRight_Square_Bracket", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportDownloadsGetRequest {
    return new ApiV2NmReportDownloadsGetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsGetRequest {
    return new ApiV2NmReportDownloadsGetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsGetRequest {
    return new ApiV2NmReportDownloadsGetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportDownloadsGetRequest | PlainMessage<ApiV2NmReportDownloadsGetRequest> | undefined, b: ApiV2NmReportDownloadsGetRequest | PlainMessage<ApiV2NmReportDownloadsGetRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportDownloadsGetRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportDownloadsPostRequest
 */
export class ApiV2NmReportDownloadsPostRequest extends Message<ApiV2NmReportDownloadsPostRequest> {
  /**
   * @generated from field: wb.analytics.v1.ApiV2NmReportDownloadsPostReq apiV2NmReportDownloadsPostReq = 1;
   */
  apiV2NmReportDownloadsPostReq?: ApiV2NmReportDownloadsPostReq;

  constructor(data?: PartialMessage<ApiV2NmReportDownloadsPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportDownloadsPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "apiV2NmReportDownloadsPostReq", kind: "message", T: ApiV2NmReportDownloadsPostReq },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportDownloadsPostRequest {
    return new ApiV2NmReportDownloadsPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsPostRequest {
    return new ApiV2NmReportDownloadsPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsPostRequest {
    return new ApiV2NmReportDownloadsPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportDownloadsPostRequest | PlainMessage<ApiV2NmReportDownloadsPostRequest> | undefined, b: ApiV2NmReportDownloadsPostRequest | PlainMessage<ApiV2NmReportDownloadsPostRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportDownloadsPostRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportDownloadsRetryPostRequest
 */
export class ApiV2NmReportDownloadsRetryPostRequest extends Message<ApiV2NmReportDownloadsRetryPostRequest> {
  /**
   * @generated from field: wb.analytics.v1.NmReportRetryReportReq nmReportRetryReportReq = 1;
   */
  nmReportRetryReportReq?: NmReportRetryReportReq;

  constructor(data?: PartialMessage<ApiV2NmReportDownloadsRetryPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportDownloadsRetryPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmReportRetryReportReq", kind: "message", T: NmReportRetryReportReq },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportDownloadsRetryPostRequest {
    return new ApiV2NmReportDownloadsRetryPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsRetryPostRequest {
    return new ApiV2NmReportDownloadsRetryPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportDownloadsRetryPostRequest {
    return new ApiV2NmReportDownloadsRetryPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportDownloadsRetryPostRequest | PlainMessage<ApiV2NmReportDownloadsRetryPostRequest> | undefined, b: ApiV2NmReportDownloadsRetryPostRequest | PlainMessage<ApiV2NmReportDownloadsRetryPostRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportDownloadsRetryPostRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportGroupedHistoryPostRequest
 */
export class ApiV2NmReportGroupedHistoryPostRequest extends Message<ApiV2NmReportGroupedHistoryPostRequest> {
  /**
   *
   *
   * @generated from field: wb.analytics.v1.NmReportGroupedHistoryReq nmReportGroupedHistoryReq = 1;
   */
  nmReportGroupedHistoryReq?: NmReportGroupedHistoryReq;

  constructor(data?: PartialMessage<ApiV2NmReportGroupedHistoryPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportGroupedHistoryPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmReportGroupedHistoryReq", kind: "message", T: NmReportGroupedHistoryReq },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportGroupedHistoryPostRequest {
    return new ApiV2NmReportGroupedHistoryPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportGroupedHistoryPostRequest {
    return new ApiV2NmReportGroupedHistoryPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportGroupedHistoryPostRequest {
    return new ApiV2NmReportGroupedHistoryPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportGroupedHistoryPostRequest | PlainMessage<ApiV2NmReportGroupedHistoryPostRequest> | undefined, b: ApiV2NmReportGroupedHistoryPostRequest | PlainMessage<ApiV2NmReportGroupedHistoryPostRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportGroupedHistoryPostRequest, a, b);
  }
}

/**
 * @generated from message wb.analytics.v1.ApiV2NmReportGroupedPostRequest
 */
export class ApiV2NmReportGroupedPostRequest extends Message<ApiV2NmReportGroupedPostRequest> {
  /**
   *
   *
   * @generated from field: wb.analytics.v1.NmReportGroupedReq nmReportGroupedReq = 1;
   */
  nmReportGroupedReq?: NmReportGroupedReq;

  constructor(data?: PartialMessage<ApiV2NmReportGroupedPostRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "wb.analytics.v1.ApiV2NmReportGroupedPostRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nmReportGroupedReq", kind: "message", T: NmReportGroupedReq },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApiV2NmReportGroupedPostRequest {
    return new ApiV2NmReportGroupedPostRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApiV2NmReportGroupedPostRequest {
    return new ApiV2NmReportGroupedPostRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApiV2NmReportGroupedPostRequest {
    return new ApiV2NmReportGroupedPostRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApiV2NmReportGroupedPostRequest | PlainMessage<ApiV2NmReportGroupedPostRequest> | undefined, b: ApiV2NmReportGroupedPostRequest | PlainMessage<ApiV2NmReportGroupedPostRequest> | undefined): boolean {
    return proto3.util.equals(ApiV2NmReportGroupedPostRequest, a, b);
  }
}

