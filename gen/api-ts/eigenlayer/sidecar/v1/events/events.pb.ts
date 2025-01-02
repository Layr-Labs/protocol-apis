/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1EventTypesEigenState from "../eventTypes/eigenState.pb"
import * as EigenlayerSidecarV1EventTypesEthereumTypes from "../eventTypes/ethereumTypes.pb"
export type StreamEigenStateChangesRequest = {
}

export type StreamEigenStateChangesResponse = {
  blockNumber?: string
  stateRoot?: EigenlayerSidecarV1EventTypesEigenState.StateRoot
  changes?: EigenlayerSidecarV1EventTypesEigenState.EigenStateChange[]
}

export type StreamIndexedBlocksRequest = {
  includeStateChanges?: boolean
}

export type StreamIndexedBlocksResponse = {
  block?: EigenlayerSidecarV1EventTypesEthereumTypes.Block
  stateRoot?: EigenlayerSidecarV1EventTypesEigenState.StateRoot
  changes?: EigenlayerSidecarV1EventTypesEigenState.EigenStateChange[]
}

export class Events {
  static StreamEigenStateChanges(req: StreamEigenStateChangesRequest, entityNotifier?: fm.NotifyStreamEntityArrival<StreamEigenStateChangesResponse>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<StreamEigenStateChangesRequest, StreamEigenStateChangesResponse>(`/events/v1/stream-eigen-state-changes`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static StreamIndexedBlocks(req: StreamIndexedBlocksRequest, entityNotifier?: fm.NotifyStreamEntityArrival<StreamIndexedBlocksResponse>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<StreamIndexedBlocksRequest, StreamIndexedBlocksResponse>(`/events/v1/stream-indexed-blocks`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}