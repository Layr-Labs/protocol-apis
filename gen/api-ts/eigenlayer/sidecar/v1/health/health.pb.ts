/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

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