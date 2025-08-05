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

export class Health {
  static Check(req: HealthCheckRequest, initReq?: fm.InitReq): Promise<HealthCheckResponse> {
    return fm.fetchReq<HealthCheckRequest, HealthCheckResponse>(`/grpc.health.v1.Health/Check`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Watch(req: HealthCheckRequest, entityNotifier?: fm.NotifyStreamEntityArrival<HealthCheckResponse>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<HealthCheckRequest, HealthCheckResponse>(`/grpc.health.v1.Health/Watch`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}