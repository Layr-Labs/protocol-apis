/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as EigenlayerSidecarV1OperatorSetsCommon from "./common.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

type BaseListOperatorsForStakerRequest = {
  stakerAddress?: string
}

export type ListOperatorsForStakerRequest = BaseListOperatorsForStakerRequest
  & OneOf<{ blockHeight: string }>

export type ListOperatorsForStakerResponse = {
  operators?: EigenlayerSidecarV1OperatorSetsCommon.Operator[]
}


type BaseListOperatorsForStrategyRequest = {
  strategyAddress?: string
}

export type ListOperatorsForStrategyRequest = BaseListOperatorsForStrategyRequest
  & OneOf<{ blockHeight: string }>

export type ListOperatorsForStrategyResponse = {
  operators?: EigenlayerSidecarV1OperatorSetsCommon.Operator[]
}


type BaseListOperatorsForAvsRequest = {
  avsAddress?: string
}

export type ListOperatorsForAvsRequest = BaseListOperatorsForAvsRequest
  & OneOf<{ blockHeight: string }>

export type ListOperatorsForAvsResponse = {
  operators?: EigenlayerSidecarV1OperatorSetsCommon.Operator[]
}


type BaseListOperatorsForBlockRangeRequest = {
  startBlock?: string
  endBlock?: string
}

export type ListOperatorsForBlockRangeRequest = BaseListOperatorsForBlockRangeRequest
  & OneOf<{ avsAddress: string }>
  & OneOf<{ strategyAddress: string }>
  & OneOf<{ stakerAddress: string }>

export type ListOperatorsForBlockRangeResponse = {
  operators?: EigenlayerSidecarV1OperatorSetsCommon.Operator[]
}

export type ListOperatorSetsRequest = {
}

export type ListOperatorSetsResponse = {
  operatorSets?: EigenlayerSidecarV1OperatorSetsCommon.OperatorSet[]
}