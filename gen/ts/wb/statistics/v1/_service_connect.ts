//
//Описание API Статистики
//
//С помощью этих методов можно получить отчёты.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file wb/statistics/v1/_service.proto (package wb.statistics.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import {
    ApiV1SupplierIncomesGetRequest,
    ApiV1SupplierIncomesGetResponse,
    ApiV1SupplierOrdersGetRequest,
    ApiV1SupplierOrdersGetResponse,
    ApiV1SupplierSalesGetRequest,
    ApiV1SupplierSalesGetResponse,
    ApiV1SupplierStocksGetRequest,
    ApiV1SupplierStocksGetResponse,
    ApiV5SupplierReportDetailByPeriodGetRequest,
    ApiV5SupplierReportDetailByPeriodGetResponse
} from "./_service_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * @generated from service wb.statistics.v1.DefaultService
 */
export const DefaultService = {
  typeName: "wb.statistics.v1.DefaultService",
  methods: {
    /**
     * @generated from rpc wb.statistics.v1.DefaultService.ApiV1SupplierIncomesGet
     */
    apiV1SupplierIncomesGet: {
      name: "ApiV1SupplierIncomesGet",
      I: ApiV1SupplierIncomesGetRequest,
      O: ApiV1SupplierIncomesGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.statistics.v1.DefaultService.ApiV1SupplierOrdersGet
     */
    apiV1SupplierOrdersGet: {
      name: "ApiV1SupplierOrdersGet",
      I: ApiV1SupplierOrdersGetRequest,
      O: ApiV1SupplierOrdersGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.statistics.v1.DefaultService.ApiV1SupplierSalesGet
     */
    apiV1SupplierSalesGet: {
      name: "ApiV1SupplierSalesGet",
      I: ApiV1SupplierSalesGetRequest,
      O: ApiV1SupplierSalesGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.statistics.v1.DefaultService.ApiV1SupplierStocksGet
     */
    apiV1SupplierStocksGet: {
      name: "ApiV1SupplierStocksGet",
      I: ApiV1SupplierStocksGetRequest,
      O: ApiV1SupplierStocksGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.statistics.v1.DefaultService.ApiV5SupplierReportDetailByPeriodGet
     */
    apiV5SupplierReportDetailByPeriodGet: {
      name: "ApiV5SupplierReportDetailByPeriodGet",
      I: ApiV5SupplierReportDetailByPeriodGetRequest,
      O: ApiV5SupplierReportDetailByPeriodGetResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
