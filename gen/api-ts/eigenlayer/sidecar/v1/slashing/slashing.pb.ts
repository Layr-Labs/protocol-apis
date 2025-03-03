/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as EigenlayerSidecarV1SlashingCommon from "./common.pb"
export type ListStakerSlashingHistoryRequest = {
  stakerAddress?: string
}

export type ListStakerSlashingHistoryResponse = {
  stakerSlashingEvents?: EigenlayerSidecarV1SlashingCommon.StakerSlashingEvent[]
}

export type ListOperatorSlashingHistoryRequest = {
  operatorAddress?: string
}

export type ListOperatorSlashingHistoryResponse = {
  slashingEvents?: EigenlayerSidecarV1SlashingCommon.SlashingEvent[]
}

export type ListAvsSlashingHistoryRequest = {
  avsAddress?: string
}

export type ListAvsSlashingHistoryResponse = {
  slashingEvents?: EigenlayerSidecarV1SlashingCommon.SlashingEvent[]
}

export type ListAvsOperatorSetSlashingHistoryRequest = {
  avsAddress?: string
  operatorSetId?: string
}

export type ListAvsOperatorSetSlashingHistoryResponse = {
  slashingEvents?: EigenlayerSidecarV1SlashingCommon.SlashingEvent[]
}