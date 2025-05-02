/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"

export enum PerformerStatus {
  READY_FOR_SYNC = "READY_FOR_SYNC",
  READY_FOR_TASK = "READY_FOR_TASK",
}

export type HealthCheckRequest = {
}

export type HealthCheckResponse = {
  status?: PerformerStatus
}

export type StartSyncRequest = {
  payload?: Uint8Array
}

export type StartSyncResponse = {
}

export type TaskRequest = {
  taskId?: Uint8Array
  payload?: Uint8Array
}

export type TaskResponse = {
  taskId?: Uint8Array
  result?: Uint8Array
}

export class PerformerService {
  static HealthCheck(req: HealthCheckRequest, initReq?: fm.InitReq): Promise<HealthCheckResponse> {
    return fm.fetchReq<HealthCheckRequest, HealthCheckResponse>(`/eigenlayer.hourglass.v1.performer.PerformerService/HealthCheck`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static StartSync(req: StartSyncRequest, initReq?: fm.InitReq): Promise<StartSyncResponse> {
    return fm.fetchReq<StartSyncRequest, StartSyncResponse>(`/eigenlayer.hourglass.v1.performer.PerformerService/StartSync`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static ExecuteTask(req: TaskRequest, initReq?: fm.InitReq): Promise<TaskResponse> {
    return fm.fetchReq<TaskRequest, TaskResponse>(`/eigenlayer.hourglass.v1.performer.PerformerService/ExecuteTask`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}