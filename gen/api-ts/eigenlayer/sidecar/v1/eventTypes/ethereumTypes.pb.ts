/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../../../google/protobuf/timestamp.pb"
export type TransactionLog = {
  transactionHash?: string
  transactionIndex?: string
  logIndex?: string
  blockNumber?: string
  address?: string
  eventName?: string
  arguments?: Uint8Array
  outputData?: Uint8Array
}

export type Transaction = {
  transactionHash?: string
  transactionIndex?: string
  blockNumber?: string
  fromAddress?: string
  toAddress?: string
  contractAddress?: string
  logs?: TransactionLog[]
}

export type Block = {
  blockNumber?: string
  blockHash?: string
  parentHash?: string
  blockTime?: GoogleProtobufTimestamp.Timestamp
  transactions?: Transaction[]
}