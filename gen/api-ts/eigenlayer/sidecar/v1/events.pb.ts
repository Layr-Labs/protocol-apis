/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
import * as EigenlayerSidecarEventsV1EigenState from "./events/eigenState.pb"
export type StreamEigenStateChangesRequest = {
}

export type StreamEigenStateChangesResponse = {
  blockNumber?: string
  stateRoot?: EigenlayerSidecarEventsV1EigenState.StateRoot
  changes?: EigenlayerSidecarEventsV1EigenState.EigenStateChange[]
}

export class Events {
  static StreamEigenStateChanges(req: StreamEigenStateChangesRequest, entityNotifier?: fm.NotifyStreamEntityArrival<StreamEigenStateChangesResponse>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<StreamEigenStateChangesRequest, StreamEigenStateChangesResponse>(`/events/v1/stream-eigen-state-changes`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}