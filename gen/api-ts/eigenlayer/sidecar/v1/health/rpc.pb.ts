/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1HealthHealth from "./health.pb"
export class Health {
  static HealthCheck(req: EigenlayerSidecarV1HealthHealth.HealthCheckRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1HealthHealth.HealthCheckResponse> {
    return fm.fetchReq<EigenlayerSidecarV1HealthHealth.HealthCheckRequest, EigenlayerSidecarV1HealthHealth.HealthCheckResponse>(`/v1/health?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ReadyCheck(req: EigenlayerSidecarV1HealthHealth.ReadyRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1HealthHealth.ReadyResponse> {
    return fm.fetchReq<EigenlayerSidecarV1HealthHealth.ReadyRequest, EigenlayerSidecarV1HealthHealth.ReadyResponse>(`/v1/ready?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}