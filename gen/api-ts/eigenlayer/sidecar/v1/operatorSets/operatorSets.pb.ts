/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as EigenlayerSidecarV1OperatorSetsCommon from "./common.pb"
export type ListOperatorsForStakerRequest = {
  stakerAddress?: string
}

export type ListOperatorsForStakerResponse = {
  operators?: EigenlayerSidecarV1OperatorSetsCommon.Operator[]
}

export type ListOperatorsForStrategyRequest = {
  strategyAddress?: string
}

export type ListOperatorsForStrategyResponse = {
  operators?: EigenlayerSidecarV1OperatorSetsCommon.Operator[]
}

export type ListOperatorsForAvsRequest = {
  avsAddress?: string
}

export type ListOperatorsForAvsResponse = {
  operators?: EigenlayerSidecarV1OperatorSetsCommon.Operator[]
}