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
  static ListStrategies(req: EigenlayerSidecarV1ProtocolProtocol.ListStrategiesRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.ListStrategiesResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.ListStrategiesRequest, EigenlayerSidecarV1ProtocolProtocol.ListStrategiesResponse>(`/protocol/v1/strategies?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ListStakerStrategies(req: EigenlayerSidecarV1ProtocolProtocol.ListStakerStrategiesRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.ListStakerStrategiesResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.ListStakerStrategiesRequest, EigenlayerSidecarV1ProtocolProtocol.ListStakerStrategiesResponse>(`/protocol/v1/stakers/${req["stakerAddress"]}/strategies?${fm.renderURLSearchParams(req, ["stakerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetStrategyForStaker(req: EigenlayerSidecarV1ProtocolProtocol.GetStrategyForStakerRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.GetStrategyForStakerResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.GetStrategyForStakerRequest, EigenlayerSidecarV1ProtocolProtocol.GetStrategyForStakerResponse>(`/protocol/v1/stakers/${req["stakerAddress"]}/strategies/${req["strategyAddress"]}?${fm.renderURLSearchParams(req, ["stakerAddress", "strategyAddress"])}`, {...initReq, method: "GET"})
  }
  static ListStakerQueuedWithdrawals(req: EigenlayerSidecarV1ProtocolProtocol.ListStakerQueuedWithdrawalsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.ListStakerQueuedWithdrawalsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.ListStakerQueuedWithdrawalsRequest, EigenlayerSidecarV1ProtocolProtocol.ListStakerQueuedWithdrawalsResponse>(`/protocol/v1/stakers/${req["stakerAddress"]}/queued-withdrawals?${fm.renderURLSearchParams(req, ["stakerAddress"])}`, {...initReq, method: "GET"})
  }
  static ListStrategyQueuedWithdrawals(req: EigenlayerSidecarV1ProtocolProtocol.ListStrategyQueuedWithdrawalsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.ListStrategyQueuedWithdrawalsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.ListStrategyQueuedWithdrawalsRequest, EigenlayerSidecarV1ProtocolProtocol.ListStrategyQueuedWithdrawalsResponse>(`/protocol/v1/strategies/${req["strategyAddress"]}/queued-withdrawals?${fm.renderURLSearchParams(req, ["strategyAddress"])}`, {...initReq, method: "GET"})
  }
  static ListOperatorQueuedWithdrawals(req: EigenlayerSidecarV1ProtocolProtocol.ListOperatorQueuedWithdrawalsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.ListOperatorQueuedWithdrawalsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.ListOperatorQueuedWithdrawalsRequest, EigenlayerSidecarV1ProtocolProtocol.ListOperatorQueuedWithdrawalsResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/queued-withdrawals?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static ListOperatorStrategyQueuedWithdrawals(req: EigenlayerSidecarV1ProtocolProtocol.ListOperatorStrategyQueuedWithdrawalsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1ProtocolProtocol.ListOperatorStrategyQueuedWithdrawalsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1ProtocolProtocol.ListOperatorStrategyQueuedWithdrawalsRequest, EigenlayerSidecarV1ProtocolProtocol.ListOperatorStrategyQueuedWithdrawalsResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/strategies/${req["strategyAddress"]}/queued-withdrawals?${fm.renderURLSearchParams(req, ["operatorAddress", "strategyAddress"])}`, {...initReq, method: "GET"})
  }
}