/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1CommonTypes from "../common/types.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

type BaseGetRegisteredAvsForOperatorRequest = {
  operatorAddress?: string
}

export type GetRegisteredAvsForOperatorRequest = BaseGetRegisteredAvsForOperatorRequest
  & OneOf<{ blockHeight: string }>

export type GetRegisteredAvsForOperatorResponse = {
  avsAddresses?: string[]
}


type BaseGetDelegatedStrategiesForOperatorRequest = {
  operatorAddress?: string
}

export type GetDelegatedStrategiesForOperatorRequest = BaseGetDelegatedStrategiesForOperatorRequest
  & OneOf<{ blockHeight: string }>

export type GetDelegatedStrategiesForOperatorResponse = {
  strategyAddresses?: string[]
}


type BaseGetOperatorDelegatedStakeForStrategyRequest = {
  operatorAddress?: string
  strategyAddress?: string
}

export type GetOperatorDelegatedStakeForStrategyRequest = BaseGetOperatorDelegatedStakeForStrategyRequest
  & OneOf<{ blockHeight: string }>

export type GetOperatorDelegatedStakeForStrategyResponse = {
  shares?: string
  avsAddresses?: string[]
}


type BaseGetDelegatedStakersForOperatorRequest = {
  operatorAddress?: string
}

export type GetDelegatedStakersForOperatorRequest = BaseGetDelegatedStakersForOperatorRequest
  & OneOf<{ blockHeight: string }>
  & OneOf<{ pagination: EigenlayerSidecarV1CommonTypes.Pagination }>


type BaseGetDelegatedStakersForOperatorResponse = {
  stakerAddresses?: string[]
}

export type GetDelegatedStakersForOperatorResponse = BaseGetDelegatedStakersForOperatorResponse
  & OneOf<{ nextPage: EigenlayerSidecarV1CommonTypes.Pagination }>


type BaseGetStakerSharesRequest = {
  stakerAddress?: string
}

export type GetStakerSharesRequest = BaseGetStakerSharesRequest
  & OneOf<{ blockHeight: string }>


type BaseStakerShare = {
  strategy?: string
  shares?: string
  avsAddresses?: string[]
}

export type StakerShare = BaseStakerShare
  & OneOf<{ operatorAddress: string }>

export type GetStakerSharesResponse = {
  shares?: StakerShare[]
}

export class Protocol {
  static GetRegisteredAvsForOperator(req: GetRegisteredAvsForOperatorRequest, initReq?: fm.InitReq): Promise<GetRegisteredAvsForOperatorResponse> {
    return fm.fetchReq<GetRegisteredAvsForOperatorRequest, GetRegisteredAvsForOperatorResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/registered-avs?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static GetDelegatedStrategiesForOperator(req: GetDelegatedStrategiesForOperatorRequest, initReq?: fm.InitReq): Promise<GetDelegatedStrategiesForOperatorResponse> {
    return fm.fetchReq<GetDelegatedStrategiesForOperatorRequest, GetDelegatedStrategiesForOperatorResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/delegated-strategies?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static GetOperatorDelegatedStakeForStrategy(req: GetOperatorDelegatedStakeForStrategyRequest, initReq?: fm.InitReq): Promise<GetOperatorDelegatedStakeForStrategyResponse> {
    return fm.fetchReq<GetOperatorDelegatedStakeForStrategyRequest, GetOperatorDelegatedStakeForStrategyResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/strategies/${req["strategyAddress"]}/delegated-stake?${fm.renderURLSearchParams(req, ["operatorAddress", "strategyAddress"])}`, {...initReq, method: "GET"})
  }
  static GetDelegatedStakersForOperator(req: GetDelegatedStakersForOperatorRequest, initReq?: fm.InitReq): Promise<GetDelegatedStakersForOperatorResponse> {
    return fm.fetchReq<GetDelegatedStakersForOperatorRequest, GetDelegatedStakersForOperatorResponse>(`/protocol/v1/operators/${req["operatorAddress"]}/delegated-stakers?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static GetStakerShares(req: GetStakerSharesRequest, initReq?: fm.InitReq): Promise<GetStakerSharesResponse> {
    return fm.fetchReq<GetStakerSharesRequest, GetStakerSharesResponse>(`/protocol/v1/stakers/${req["stakerAddress"]}/shares?${fm.renderURLSearchParams(req, ["stakerAddress"])}`, {...initReq, method: "GET"})
  }
}