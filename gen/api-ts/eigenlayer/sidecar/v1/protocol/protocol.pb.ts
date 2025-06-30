/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as EigenlayerSidecarV1CommonTypes from "../common/types.pb"
import * as EigenlayerSidecarV1EigenStateEigenState from "../eigenState/eigenState.pb"
import * as EigenlayerSidecarV1ProtocolCommon from "./common.pb"

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
  & OneOf<{ strategy: string }>
  & OneOf<{ startBlock: string }>
  & OneOf<{ endBlock: string }>

export type GetStakerSharesResponse = {
  shares?: EigenlayerSidecarV1ProtocolCommon.StakerShare[]
}

export type GetEigenStateChangesRequest = {
  blockHeight?: string
}

export type GetEigenStateChangesResponse = {
  changes?: EigenlayerSidecarV1EigenStateEigenState.EigenStateChange[]
}

export type ListStrategiesRequest = {
}

export type ListStrategiesResponse = {
  strategies?: EigenlayerSidecarV1ProtocolCommon.Strategy[]
}


type BaseListStakerStrategiesRequest = {
  stakerAddress?: string
}

export type ListStakerStrategiesRequest = BaseListStakerStrategiesRequest
  & OneOf<{ blockHeight: string }>

export type ListStakerStrategiesResponse = {
  stakerStrategies?: EigenlayerSidecarV1ProtocolCommon.StakerStrategy[]
}


type BaseGetStrategyForStakerRequest = {
  stakerAddress?: string
  strategyAddress?: string
}

export type GetStrategyForStakerRequest = BaseGetStrategyForStakerRequest
  & OneOf<{ blockHeight: string }>

export type GetStrategyForStakerResponse = {
  stakerStrategy?: EigenlayerSidecarV1ProtocolCommon.StakerStrategy
}

export type ListStakerQueuedWithdrawalsRequest = {
  stakerAddress?: string
}

export type ListStakerQueuedWithdrawalsResponse = {
  queuedWithdrawals?: EigenlayerSidecarV1ProtocolCommon.QueuedStrategyWithdrawal[]
}

export type ListStrategyQueuedWithdrawalsRequest = {
  strategyAddress?: string
}

export type ListStrategyQueuedWithdrawalsResponse = {
  withdrawals?: EigenlayerSidecarV1ProtocolCommon.StakerWithdrawal[]
}

export type ListOperatorQueuedWithdrawalsRequest = {
  operatorAddress?: string
}

export type ListOperatorQueuedWithdrawalsResponse = {
  strategies?: EigenlayerSidecarV1ProtocolCommon.QueueStakerStrategyWithdrawal[]
}

export type ListOperatorStrategyQueuedWithdrawalsRequest = {
  operatorAddress?: string
  strategyAddress?: string
}

export type ListOperatorStrategyQueuedWithdrawalsResponse = {
  withdrawals?: EigenlayerSidecarV1ProtocolCommon.StakerWithdrawal[]
}