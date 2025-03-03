/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1EventsEvents from "./events.pb"
export class Events {
  static StreamEigenStateChanges(req: EigenlayerSidecarV1EventsEvents.StreamEigenStateChangesRequest, entityNotifier?: fm.NotifyStreamEntityArrival<EigenlayerSidecarV1EventsEvents.StreamEigenStateChangesResponse>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<EigenlayerSidecarV1EventsEvents.StreamEigenStateChangesRequest, EigenlayerSidecarV1EventsEvents.StreamEigenStateChangesResponse>(`/events/v1/stream-eigen-state-changes`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static StreamIndexedBlocks(req: EigenlayerSidecarV1EventsEvents.StreamIndexedBlocksRequest, entityNotifier?: fm.NotifyStreamEntityArrival<EigenlayerSidecarV1EventsEvents.StreamIndexedBlocksResponse>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<EigenlayerSidecarV1EventsEvents.StreamIndexedBlocksRequest, EigenlayerSidecarV1EventsEvents.StreamIndexedBlocksResponse>(`/events/v1/stream-indexed-blocks`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}