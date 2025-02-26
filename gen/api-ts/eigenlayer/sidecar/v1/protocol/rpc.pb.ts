/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1ProtocolProtocol from "./protocol.pb"
export class Protocol {
  static GetRegisteredAvsForOperator(req: EigenlayerSidecarV1ProtocolProtocol.GetRegisteredAvsForOperatorRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.GetRegisteredAvsForOperatorResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.GetRegisteredAvsForOperatorRequest, EigenlayerSidecarV1ProtocolProtocol.GetRegisteredAvsForOperatorResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/registered-avs?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static GetDelegatedStrategiesForOperator(req: EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStrategiesForOperatorRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStrategiesForOperatorResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStrategiesForOperatorRequest, EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStrategiesForOperatorResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/delegated-strategies?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static GetOperatorDelegatedStakeForStrategy(req: EigenlayerSidecarV1ProtocolProtocol.GetOperatorDelegatedStakeForStrategyRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.GetOperatorDelegatedStakeForStrategyResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.GetOperatorDelegatedStakeForStrategyRequest, EigenlayerSidecarV1ProtocolProtocol.GetOperatorDelegatedStakeForStrategyResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/strategies/${req["strategyAddress"]}/delegated-stake?${fm.renderURLSearchParams(req, ["operatorAddress", "strategyAddress"])}`, {...initReq, method: "GET"})
  }
  static GetDelegatedStakersForOperator(req: EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStakersForOperatorRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStakersForOperatorResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStakersForOperatorRequest, EigenlayerSidecarV1ProtocolProtocol.GetDelegatedStakersForOperatorResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/delegated-stakers?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static GetStakerShares(req: EigenlayerSidecarV1ProtocolProtocol.GetStakerSharesRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.GetStakerSharesResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.GetStakerSharesRequest, EigenlayerSidecarV1ProtocolProtocol.GetStakerSharesResponse>(`/protocol/v1/stakers/${req["stakerAddress"]}/shares?${fm.renderURLSearchParams(req, ["stakerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetEigenStateChanges(req: EigenlayerSidecarV1ProtocolProtocol.GetEigenStateChangesRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.GetEigenStateChangesResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.GetEigenStateChangesRequest, EigenlayerSidecarV1ProtocolProtocol.GetEigenStateChangesResponse>(`/protocol/v1/eigen-state-changes?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}