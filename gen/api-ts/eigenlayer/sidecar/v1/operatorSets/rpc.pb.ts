/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1OperatorSetsOperatorSets from "./operatorSets.pb"
export class OperatorSets {
  static ListOperatorsForStaker(req: EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStakerRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStakerResponse> {
    return fm.fetchReq<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStakerRequest, EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStakerResponse>(`/v1/operatorSets/stakers/${req["stakerAddress"]}/operators?${fm.renderURLSearchParams(req, ["stakerAddress"])}`, {...initReq, method: "GET"})
  }
  static ListOperatorsForStrategy(req: EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStrategyRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStrategyResponse> {
    return fm.fetchReq<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStrategyRequest, EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForStrategyResponse>(`/v1/operatorSets/strategies/${req["strategyAddress"]}/operators?${fm.renderURLSearchParams(req, ["strategyAddress"])}`, {...initReq, method: "GET"})
  }
  static ListOperatorsForAvs(req: EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForAvsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForAvsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForAvsRequest, EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForAvsResponse>(`/v1/operatorSets/avss/${req["avsAddress"]}/operators?${fm.renderURLSearchParams(req, ["avsAddress"])}`, {...initReq, method: "GET"})
  }
  static ListOperatorsForBlockRange(req: EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForBlockRangeRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForBlockRangeResponse> {
    return fm.fetchReq<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForBlockRangeRequest, EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorsForBlockRangeResponse>(`/v1/operatorSets/blockRange/${req["startBlock"]}/${req["endBlock"]}/operators?${fm.renderURLSearchParams(req, ["startBlock", "endBlock"])}`, {...initReq, method: "GET"})
  }
  static ListOperatorSets(req: EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorSetsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorSetsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorSetsRequest, EigenlayerSidecarV1OperatorSetsOperatorSets.ListOperatorSetsResponse>(`/v1/operatorSets?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}