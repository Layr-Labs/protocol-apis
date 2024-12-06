/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"

export enum HealthCheckResponseServingStatus {
  UNKNOWN = "UNKNOWN",
  SERVING = "SERVING",
  NOT_SERVING = "NOT_SERVING",
  SERVICE_UNKNOWN = "SERVICE_UNKNOWN",
}

export type HealthCheckRequest = {
  service?: string
}

export type HealthCheckResponse = {
  status?: HealthCheckResponseServingStatus
}

export type ReadyRequest = {
}

export type ReadyResponse = {
  ready?: boolean
}

export class Health {
  static HealthCheck(req: HealthCheckRequest, initReq?: fm.InitReq): Promise<HealthCheckResponse> {
    return fm.fetchReq<HealthCheckRequest, HealthCheckResponse>(`/v1/health?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ReadyCheck(req: ReadyRequest, initReq?: fm.InitReq): Promise<ReadyResponse> {
    return fm.fetchReq<ReadyRequest, ReadyResponse>(`/v1/ready?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}