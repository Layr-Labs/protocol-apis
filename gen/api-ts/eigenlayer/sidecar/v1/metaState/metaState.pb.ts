/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../../../google/protobuf/timestamp.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);
export type TransactionMetadata = {
  transactionHash?: string
  logIndex?: string
  blockHeight?: string
}

export type GenerationReservationCreated = {
  avs?: string
  operatorSetId?: string
  transactionMetadata?: TransactionMetadata
}

export type KeyRotationScheduledTrigger = {
  activateAt?: GoogleProtobufTimestamp.Timestamp
}


type BaseMetaStateChange = {
}

export type MetaStateChange = BaseMetaStateChange
  & OneOf<{ generationReservationCreated: GenerationReservationCreated; keyRotationScheduledTrigger: KeyRotationScheduledTrigger }>