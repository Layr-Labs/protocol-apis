/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
export type GetBlockHeightRequest = {
  verified?: boolean
}

export type GetBlockHeightResponse = {
  blockNumber?: string
  blockHash?: string
}

export type GetStateRootRequest = {
  blockNumber?: string
}

export type GetStateRootResponse = {
  ethBlockNumber?: string
  ethBlockHash?: string
  stateRoot?: string
}

export type AboutRequest = {
}

export type AboutResponse = {
  version?: string
  commit?: string
  chain?: string
}