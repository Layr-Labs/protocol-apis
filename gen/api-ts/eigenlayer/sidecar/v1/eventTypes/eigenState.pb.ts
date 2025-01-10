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

export enum RewardSubmissionRewardType {
  AVS = "AVS",
  ALL_STAKERS = "ALL_STAKERS",
  ALL_EARNERS = "ALL_EARNERS",
}

export type StateRoot = {
  ethBlockNumber?: string
  ethBlockHash?: string
  stateRoot?: string
}

export type TransactionMetadata = {
  logIndex?: string
  transactionHash?: string
  blockHeight?: string
}

export type AvsOperatorStateChange = {
  avs?: string
  operator?: string
  registered?: boolean
  transactionMetadata?: TransactionMetadata
}

export type OperatorShareDelta = {
  operator?: string
  staker?: string
  strategy?: string
  shares?: string
  blockTime?: GoogleProtobufTimestamp.Timestamp
  blockDate?: string
  transactionMetadata?: TransactionMetadata
}

export type RewardSubmission = {
  avs?: string
  rewardHash?: string
  token?: string
  amount?: string
  strategy?: string
  strategyIndex?: string
  multiplier?: string
  startTimestamp?: GoogleProtobufTimestamp.Timestamp
  endTimestamp?: GoogleProtobufTimestamp.Timestamp
  duration?: string
  rewardType?: RewardSubmissionRewardType
  transactionMetadata?: TransactionMetadata
}

export type StakerDelegationChange = {
  staker?: string
  operator?: string
  delegated?: boolean
  transactionMetadata?: TransactionMetadata
}

export type StakerShareDelta = {
  staker?: string
  strategy?: string
  shares?: string
  strategyIndex?: string
  blockTime?: GoogleProtobufTimestamp.Timestamp
  blockDate?: string
  transactionMetadata?: TransactionMetadata
}

export type SubmittedDistributionRoot = {
  root?: string
  rootIndex?: string
  rewardsCalculationEnd?: GoogleProtobufTimestamp.Timestamp
  rewardsCalculationEndUnit?: string
  activatedAt?: GoogleProtobufTimestamp.Timestamp
  activatedAtUnit?: string
  createdAtBlockNumber?: string
  transactionMetadata?: TransactionMetadata
}

export type DisabledDistributionRoot = {
  rootIndex?: string
  transactionMetadata?: TransactionMetadata
}


type BaseEigenStateChange = {
}

export type EigenStateChange = BaseEigenStateChange
  & OneOf<{ avsOperatorStateChange: AvsOperatorStateChange; operatorShareDelta: OperatorShareDelta; rewardSubmission: RewardSubmission; stakerDelegationChange: StakerDelegationChange; stakerShareDelta: StakerShareDelta; submittedDistributionRoot: SubmittedDistributionRoot; disabledDistributionRoot: DisabledDistributionRoot }>