/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1SlashingSlashing from "./slashing.pb"
export class Slashing {
  static ListStakerSlashingHistory(req: EigenlayerSidecarV1SlashingSlashing.ListStakerSlashingHistoryRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1SlashingSlashing.ListStakerSlashingHistoryResponse> {
    return fm.fetchReq<EigenlayerSidecarV1SlashingSlashing.ListStakerSlashingHistoryRequest, EigenlayerSidecarV1SlashingSlashing.ListStakerSlashingHistoryResponse>(`/v1/slashing/stakers/${req["stakerAddress"]}/slashing-history?${fm.renderURLSearchParams(req, ["stakerAddress"])}`, {...initReq, method: "GET"})
  }
  static ListOperatorSlashingHistory(req: EigenlayerSidecarV1SlashingSlashing.ListOperatorSlashingHistoryRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1SlashingSlashing.ListOperatorSlashingHistoryResponse> {
    return fm.fetchReq<EigenlayerSidecarV1SlashingSlashing.ListOperatorSlashingHistoryRequest, EigenlayerSidecarV1SlashingSlashing.ListOperatorSlashingHistoryResponse>(`/v1/slashing/operators/${req["operatorAddress"]}/slashing-history?${fm.renderURLSearchParams(req, ["operatorAddress"])}`, {...initReq, method: "GET"})
  }
  static ListAvsSlashingHistory(req: EigenlayerSidecarV1SlashingSlashing.ListAvsSlashingHistoryRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1SlashingSlashing.ListAvsSlashingHistoryResponse> {
    return fm.fetchReq<EigenlayerSidecarV1SlashingSlashing.ListAvsSlashingHistoryRequest, EigenlayerSidecarV1SlashingSlashing.ListAvsSlashingHistoryResponse>(`/v1/slashing/avss/${req["avsAddress"]}/slashing-history?${fm.renderURLSearchParams(req, ["avsAddress"])}`, {...initReq, method: "GET"})
  }
  static ListAvsOperatorSetSlashingHistory(req: EigenlayerSidecarV1SlashingSlashing.ListAvsOperatorSetSlashingHistoryRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1SlashingSlashing.ListAvsOperatorSetSlashingHistoryResponse> {
    return fm.fetchReq<EigenlayerSidecarV1SlashingSlashing.ListAvsOperatorSetSlashingHistoryRequest, EigenlayerSidecarV1SlashingSlashing.ListAvsOperatorSetSlashingHistoryResponse>(`/v1/slashing/avss/${req["avsAddress"]}/operator-sets/${req["operatorSetId"]}/slashing-history?${fm.renderURLSearchParams(req, ["avsAddress", "operatorSetId"])}`, {...initReq, method: "GET"})
  }
}