/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
export type Task = {
  taskId?: string
  metadata?: Uint8Array
  payload?: Uint8Array
}

export type TaskResult = {
  taskId?: string
  result?: Uint8Array
}

export type HealthRequest = {
}

export type HealthResponse = {
  status?: string
}

export class PerformerService {
  static ExecuteTask(req: Task, initReq?: fm.InitReq): Promise<TaskResult> {
    return fm.fetchReq<Task, TaskResult>(`/eigenlayer.hourglass.v1.performer.PerformerService/ExecuteTask`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Health(req: HealthRequest, initReq?: fm.InitReq): Promise<HealthResponse> {
    return fm.fetchReq<HealthRequest, HealthResponse>(`/eigenlayer.hourglass.v1.performer.PerformerService/Health`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}