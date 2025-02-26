/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1SidecarSidecar from "./sidecar.pb"
export class Rpc {
  static GetBlockHeight(req: EigenlayerSidecarV1SidecarSidecar.GetBlockHeightRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1SidecarSidecar.GetBlockHeightResponse> {
    return fm.fetchReq<EigenlayerSidecarV1SidecarSidecar.GetBlockHeightRequest, EigenlayerSidecarV1SidecarSidecar.GetBlockHeightResponse>(`/rpc/v1/latest-block?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetStateRoot(req: EigenlayerSidecarV1SidecarSidecar.GetStateRootRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1SidecarSidecar.GetStateRootResponse> {
    return fm.fetchReq<EigenlayerSidecarV1SidecarSidecar.GetStateRootRequest, EigenlayerSidecarV1SidecarSidecar.GetStateRootResponse>(`/rpc/v1/state-roots/${req["blockNumber"]}?${fm.renderURLSearchParams(req, ["blockNumber"])}`, {...initReq, method: "GET"})
  }
  static About(req: EigenlayerSidecarV1SidecarSidecar.AboutRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1SidecarSidecar.AboutResponse> {
    return fm.fetchReq<EigenlayerSidecarV1SidecarSidecar.AboutRequest, EigenlayerSidecarV1SidecarSidecar.AboutResponse>(`/rpc/v1/about?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}