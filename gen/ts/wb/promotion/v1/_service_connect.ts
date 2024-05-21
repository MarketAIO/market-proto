//
//Описание API Продвижение
//
//Синхронизация данных из бд происходит раз в 3 минуты.  <br>Изменение статуса происходит раз в 1 минуту. Внутри этого интервала будет сохранено последнее действие по изменению статуса. <br>Изменение ставки происходит раз в 30 секунд. Внутри этого интервала будет сохранено последнее действие по изменению ставки.
//
//The version of the OpenAPI document: 
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file wb/promotion/v1/_service.proto (package wb.promotion.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import {
    AdvV0ActiveGetRequest,
    AdvV0AllcpmPostRequest,
    AdvV0AllcpmPostResponse,
    AdvV0CpmGetRequest,
    AdvV0CpmGetResponse,
    AdvV0CpmPostRequest,
    AdvV0DeleteGetRequest,
    AdvV0ParamsMenuGetRequest,
    AdvV0ParamsMenuGetResponse,
    AdvV0ParamsSetGetRequest,
    AdvV0ParamsSetGetResponse,
    AdvV0ParamsSubjectGetRequest,
    AdvV0ParamsSubjectGetResponse,
    AdvV0PauseGetRequest,
    AdvV0RenamePostRequest,
    AdvV0StartGetRequest,
    AdvV0StopGetRequest,
    AdvV1AdvertGetRequest,
    AdvV1AdvertPausePostRequest,
    AdvV1AdvertsGetRequest,
    AdvV1AdvertsGetResponse,
    AdvV1AdvertStartPostRequest,
    AdvV1AdvertStopPostRequest,
    AdvV1AutoActivePostRequest,
    AdvV1AutoGetnmtoaddGetRequest,
    AdvV1AutoGetnmtoaddGetResponse,
    AdvV1AutoSetExcludedPostRequest,
    AdvV1AutoStatGetRequest,
    AdvV1AutoUpdatenmPostRequest,
    AdvV1BudgetDepositPostRequest,
    AdvV1BudgetGetRequest,
    AdvV1ItemCpmChangePostRequest,
    AdvV1PaymentsGetRequest,
    AdvV1PaymentsGetResponse,
    AdvV1PromotionAdvertsPostRequest,
    AdvV1SaveAdPostRequest,
    AdvV1SaveAdPostResponse,
    AdvV1SeacatStatGetRequest,
    AdvV1SearchSetExcludedPostRequest,
    AdvV1SearchSetPhrasePostRequest,
    AdvV1SearchSetPlusGetRequest,
    AdvV1SearchSetPlusPostRequest,
    AdvV1SearchSetPlusPostResponse,
    AdvV1SearchSetStrongPostRequest,
    AdvV1SearchSupplierProductsGetRequest,
    AdvV1SearchSupplierProductsGetResponse,
    AdvV1SearchSupplierSubjectsGetResponse,
    AdvV1StatsPostRequest,
    AdvV1StatsPostResponse,
    AdvV1StatWordsGetRequest,
    AdvV1SupplierSubjectsGetResponse,
    AdvV1UpdGetRequest,
    AdvV1UpdGetResponse,
    AdvV1UpdIntervalsGetResponse,
    AdvV2AutoDailyWordsGetRequest,
    AdvV2AutoStatWordsGetRequest,
    AdvV2FullstatsPostRequest,
    AdvV2SeacatSaveAdPostRequest,
    AdvV2SeacatSaveAdPostResponse,
    AdvV2SupplierNmsPostRequest,
    AdvV2SupplierNmsPostResponse
} from "./_service_pb.js";
import {Empty, MethodKind} from "@bufbuild/protobuf";
import {AdvV1AdvertGet200Response} from "./adv_v1_advert_get200_response_pb.js";
import {AdvV1AutoStatGet200Response} from "./adv_v1_auto_stat_get200_response_pb.js";
import {AdvV1BalanceGet200Response} from "./adv_v1_balance_get200_response_pb.js";
import {AdvV1BudgetDepositPost200Response} from "./adv_v1_budget_deposit_post200_response_pb.js";
import {AdvV1BudgetGet200Response} from "./adv_v1_budget_get200_response_pb.js";
import {AdvV1CountGet200Response} from "./adv_v1_count_get200_response_pb.js";
import {AdvV1PromotionAdvertsPost200Response} from "./adv_v1_promotion_adverts_post200_response_pb.js";
import {AdvV1PromotionCountGet200Response} from "./adv_v1_promotion_count_get200_response_pb.js";
import {AdvV1SeacatStatGet200Response} from "./adv_v1_seacat_stat_get200_response_pb.js";
import {AdvV1StatWordsGet200Response} from "./adv_v1_stat_words_get200_response_pb.js";
import {AdvV2AutoDailyWordsGet200Response} from "./adv_v2_auto_daily_words_get200_response_pb.js";
import {AdvV2AutoStatWordsGet200Response} from "./adv_v2_auto_stat_words_get200_response_pb.js";
import {AdvV2FullstatsPost200Response} from "./adv_v2_fullstats_post200_response_pb.js";

/**
 * @generated from service wb.promotion.v1.DefaultService
 */
export const DefaultService = {
  typeName: "wb.promotion.v1.DefaultService",
  methods: {
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0ActiveGet
     */
    advV0ActiveGet: {
      name: "AdvV0ActiveGet",
      I: AdvV0ActiveGetRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0AllcpmPost
     */
    advV0AllcpmPost: {
      name: "AdvV0AllcpmPost",
      I: AdvV0AllcpmPostRequest,
      O: AdvV0AllcpmPostResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0CpmGet
     */
    advV0CpmGet: {
      name: "AdvV0CpmGet",
      I: AdvV0CpmGetRequest,
      O: AdvV0CpmGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0CpmPost
     */
    advV0CpmPost: {
      name: "AdvV0CpmPost",
      I: AdvV0CpmPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0DeleteGet
     */
    advV0DeleteGet: {
      name: "AdvV0DeleteGet",
      I: AdvV0DeleteGetRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0ParamsMenuGet
     */
    advV0ParamsMenuGet: {
      name: "AdvV0ParamsMenuGet",
      I: AdvV0ParamsMenuGetRequest,
      O: AdvV0ParamsMenuGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0ParamsSetGet
     */
    advV0ParamsSetGet: {
      name: "AdvV0ParamsSetGet",
      I: AdvV0ParamsSetGetRequest,
      O: AdvV0ParamsSetGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0ParamsSubjectGet
     */
    advV0ParamsSubjectGet: {
      name: "AdvV0ParamsSubjectGet",
      I: AdvV0ParamsSubjectGetRequest,
      O: AdvV0ParamsSubjectGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0PauseGet
     */
    advV0PauseGet: {
      name: "AdvV0PauseGet",
      I: AdvV0PauseGetRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0RenamePost
     */
    advV0RenamePost: {
      name: "AdvV0RenamePost",
      I: AdvV0RenamePostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0StartGet
     */
    advV0StartGet: {
      name: "AdvV0StartGet",
      I: AdvV0StartGetRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV0StopGet
     */
    advV0StopGet: {
      name: "AdvV0StopGet",
      I: AdvV0StopGetRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AdvertGet
     */
    advV1AdvertGet: {
      name: "AdvV1AdvertGet",
      I: AdvV1AdvertGetRequest,
      O: AdvV1AdvertGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AdvertPausePost
     */
    advV1AdvertPausePost: {
      name: "AdvV1AdvertPausePost",
      I: AdvV1AdvertPausePostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AdvertStartPost
     */
    advV1AdvertStartPost: {
      name: "AdvV1AdvertStartPost",
      I: AdvV1AdvertStartPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AdvertStopPost
     */
    advV1AdvertStopPost: {
      name: "AdvV1AdvertStopPost",
      I: AdvV1AdvertStopPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AdvertsGet
     */
    advV1AdvertsGet: {
      name: "AdvV1AdvertsGet",
      I: AdvV1AdvertsGetRequest,
      O: AdvV1AdvertsGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AutoActivePost
     */
    advV1AutoActivePost: {
      name: "AdvV1AutoActivePost",
      I: AdvV1AutoActivePostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AutoGetnmtoaddGet
     */
    advV1AutoGetnmtoaddGet: {
      name: "AdvV1AutoGetnmtoaddGet",
      I: AdvV1AutoGetnmtoaddGetRequest,
      O: AdvV1AutoGetnmtoaddGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AutoSetExcludedPost
     */
    advV1AutoSetExcludedPost: {
      name: "AdvV1AutoSetExcludedPost",
      I: AdvV1AutoSetExcludedPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AutoStatGet
     */
    advV1AutoStatGet: {
      name: "AdvV1AutoStatGet",
      I: AdvV1AutoStatGetRequest,
      O: AdvV1AutoStatGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1AutoUpdatenmPost
     */
    advV1AutoUpdatenmPost: {
      name: "AdvV1AutoUpdatenmPost",
      I: AdvV1AutoUpdatenmPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1BalanceGet
     */
    advV1BalanceGet: {
      name: "AdvV1BalanceGet",
      I: Empty,
      O: AdvV1BalanceGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1BudgetDepositPost
     */
    advV1BudgetDepositPost: {
      name: "AdvV1BudgetDepositPost",
      I: AdvV1BudgetDepositPostRequest,
      O: AdvV1BudgetDepositPost200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1BudgetGet
     */
    advV1BudgetGet: {
      name: "AdvV1BudgetGet",
      I: AdvV1BudgetGetRequest,
      O: AdvV1BudgetGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1CountGet
     */
    advV1CountGet: {
      name: "AdvV1CountGet",
      I: Empty,
      O: AdvV1CountGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1ItemCpmChangePost
     */
    advV1ItemCpmChangePost: {
      name: "AdvV1ItemCpmChangePost",
      I: AdvV1ItemCpmChangePostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1PaymentsGet
     */
    advV1PaymentsGet: {
      name: "AdvV1PaymentsGet",
      I: AdvV1PaymentsGetRequest,
      O: AdvV1PaymentsGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1PromotionAdvertsPost
     */
    advV1PromotionAdvertsPost: {
      name: "AdvV1PromotionAdvertsPost",
      I: AdvV1PromotionAdvertsPostRequest,
      O: AdvV1PromotionAdvertsPost200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1PromotionCountGet
     */
    advV1PromotionCountGet: {
      name: "AdvV1PromotionCountGet",
      I: Empty,
      O: AdvV1PromotionCountGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SaveAdPost
     */
    advV1SaveAdPost: {
      name: "AdvV1SaveAdPost",
      I: AdvV1SaveAdPostRequest,
      O: AdvV1SaveAdPostResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SeacatStatGet
     */
    advV1SeacatStatGet: {
      name: "AdvV1SeacatStatGet",
      I: AdvV1SeacatStatGetRequest,
      O: AdvV1SeacatStatGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SearchSetExcludedPost
     */
    advV1SearchSetExcludedPost: {
      name: "AdvV1SearchSetExcludedPost",
      I: AdvV1SearchSetExcludedPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SearchSetPhrasePost
     */
    advV1SearchSetPhrasePost: {
      name: "AdvV1SearchSetPhrasePost",
      I: AdvV1SearchSetPhrasePostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SearchSetPlusGet
     */
    advV1SearchSetPlusGet: {
      name: "AdvV1SearchSetPlusGet",
      I: AdvV1SearchSetPlusGetRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SearchSetPlusPost
     */
    advV1SearchSetPlusPost: {
      name: "AdvV1SearchSetPlusPost",
      I: AdvV1SearchSetPlusPostRequest,
      O: AdvV1SearchSetPlusPostResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SearchSetStrongPost
     */
    advV1SearchSetStrongPost: {
      name: "AdvV1SearchSetStrongPost",
      I: AdvV1SearchSetStrongPostRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SearchSupplierProductsGet
     */
    advV1SearchSupplierProductsGet: {
      name: "AdvV1SearchSupplierProductsGet",
      I: AdvV1SearchSupplierProductsGetRequest,
      O: AdvV1SearchSupplierProductsGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SearchSupplierSubjectsGet
     */
    advV1SearchSupplierSubjectsGet: {
      name: "AdvV1SearchSupplierSubjectsGet",
      I: Empty,
      O: AdvV1SearchSupplierSubjectsGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1StatWordsGet
     */
    advV1StatWordsGet: {
      name: "AdvV1StatWordsGet",
      I: AdvV1StatWordsGetRequest,
      O: AdvV1StatWordsGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1StatsPost
     */
    advV1StatsPost: {
      name: "AdvV1StatsPost",
      I: AdvV1StatsPostRequest,
      O: AdvV1StatsPostResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1SupplierSubjectsGet
     */
    advV1SupplierSubjectsGet: {
      name: "AdvV1SupplierSubjectsGet",
      I: Empty,
      O: AdvV1SupplierSubjectsGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1UpdGet
     */
    advV1UpdGet: {
      name: "AdvV1UpdGet",
      I: AdvV1UpdGetRequest,
      O: AdvV1UpdGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV1UpdIntervalsGet
     */
    advV1UpdIntervalsGet: {
      name: "AdvV1UpdIntervalsGet",
      I: Empty,
      O: AdvV1UpdIntervalsGetResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV2AutoDailyWordsGet
     */
    advV2AutoDailyWordsGet: {
      name: "AdvV2AutoDailyWordsGet",
      I: AdvV2AutoDailyWordsGetRequest,
      O: AdvV2AutoDailyWordsGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV2AutoStatWordsGet
     */
    advV2AutoStatWordsGet: {
      name: "AdvV2AutoStatWordsGet",
      I: AdvV2AutoStatWordsGetRequest,
      O: AdvV2AutoStatWordsGet200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV2FullstatsPost
     */
    advV2FullstatsPost: {
      name: "AdvV2FullstatsPost",
      I: AdvV2FullstatsPostRequest,
      O: AdvV2FullstatsPost200Response,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV2SeacatSaveAdPost
     */
    advV2SeacatSaveAdPost: {
      name: "AdvV2SeacatSaveAdPost",
      I: AdvV2SeacatSaveAdPostRequest,
      O: AdvV2SeacatSaveAdPostResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc wb.promotion.v1.DefaultService.AdvV2SupplierNmsPost
     */
    advV2SupplierNmsPost: {
      name: "AdvV2SupplierNmsPost",
      I: AdvV2SupplierNmsPostRequest,
      O: AdvV2SupplierNmsPostResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
