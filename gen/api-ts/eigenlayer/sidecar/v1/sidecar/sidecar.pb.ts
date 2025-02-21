/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
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

export class Rpc {
  static GetBlockHeight(req: GetBlockHeightRequest, initReq?: fm.InitReq): Promise<GetBlockHeightResponse> {
    return fm.fetchReq<GetBlockHeightRequest, GetBlockHeightResponse>(`/rpc/v1/latest-block?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetStateRoot(req: GetStateRootRequest, initReq?: fm.InitReq): Promise<GetStateRootResponse> {
    return fm.fetchReq<GetStateRootRequest, GetStateRootResponse>(`/rpc/v1/state-roots/${req["blockNumber"]}?${fm.renderURLSearchParams(req, ["blockNumber"])}`, {...initReq, method: "GET"})
  }
  static About(req: AboutRequest, initReq?: fm.InitReq): Promise<AboutResponse> {
    return fm.fetchReq<AboutRequest, AboutResponse>(`/rpc/v1/about?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}