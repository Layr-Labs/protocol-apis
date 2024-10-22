/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
export type GetDailyOperatorStrategyAprsRequest = {
  operatorAddress?: string
  date?: string
}

export type OperatorStrategyApr = {
  strategy?: string
  apr?: string
}

export type GetDailyOperatorStrategyAprsResponse = {
  aprs?: OperatorStrategyApr[]
}

export type GetDailyAprForEarnerStrategyRequest = {
  earnerAddress?: string
  strategy?: string
  date?: string
}

export type GetDailyAprForEarnerStrategyResponse = {
  apr?: string
}

export class Aprs {
  static GetDailyOperatorStrategyAprs(req: GetDailyOperatorStrategyAprsRequest, initReq?: fm.InitReq): Promise<GetDailyOperatorStrategyAprsResponse> {
    return fm.fetchReq<GetDailyOperatorStrategyAprsRequest, GetDailyOperatorStrategyAprsResponse>(`/aprs/v1/operators/${req["operatorAddress"]}/daily-aprs/${req["date"]}?${fm.renderURLSearchParams(req, ["operatorAddress", "date"])}`, {...initReq, method: "GET"})
  }
  static GetDailyAprForEarnerStrategy(req: GetDailyAprForEarnerStrategyRequest, initReq?: fm.InitReq): Promise<GetDailyAprForEarnerStrategyResponse> {
    return fm.fetchReq<GetDailyAprForEarnerStrategyRequest, GetDailyAprForEarnerStrategyResponse>(`/aprs/v1/earners/${req["earnerAddress"]}/strategies/${req["strategy"]}/daily-apr/${req["date"]}?${fm.renderURLSearchParams(req, ["earnerAddress", "strategy", "date"])}`, {...initReq, method: "GET"})
  }
}