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

export type CompletedSlashingWithdrawal = {
  withdrawalRoot?: string
  transactionMetadata?: TransactionMetadata
}

export type QueuedSlashingWithdrawal = {
  operator?: string
  withdrawalRoot?: string
  targetBlock?: string
  stakerOptOutWindowBlocks?: string
  operatorSetId?: string
  avs?: string
  transactionMetadata?: TransactionMetadata
}

export type SlashedOperator = {
  operator?: string
  operatorSetId?: string
  avs?: string
  transactionMetadata?: TransactionMetadata
}

export type SlashedOperatorShares = {
  operator?: string
  strategy?: string
  shares?: string
  transactionMetadata?: TransactionMetadata
}

export type DefaultOperatorSplit = {
  operator?: string
  oldOperatorBasisPoints?: string
  newOperatorBasisPoints?: string
  transactionMetadata?: TransactionMetadata
}

export type OperatorAllocation = {
  operator?: string
  operatorSetId?: string
  avs?: string
  strategy?: string
  shares?: string
  transactionMetadata?: TransactionMetadata
}

export type OperatorAllocationDelay = {
  operator?: string
  delay?: string
  effectiveBlock?: string
  transactionMetadata?: TransactionMetadata
}

export type OperatorAVSSplit = {
  operator?: string
  avs?: string
  operatorBasisPoints?: string
  avsBasisPoints?: string
  startTimestamp?: GoogleProtobufTimestamp.Timestamp
  endTimestamp?: GoogleProtobufTimestamp.Timestamp
  transactionMetadata?: TransactionMetadata
}

export type OperatorPISplit = {
  operator?: string
  operatorBasisPoints?: string
  piBasisPoints?: string
  startTimestamp?: GoogleProtobufTimestamp.Timestamp
  endTimestamp?: GoogleProtobufTimestamp.Timestamp
  transactionMetadata?: TransactionMetadata
}

export type OperatorSet = {
  operatorSetId?: string
  avs?: string
  transactionMetadata?: TransactionMetadata
}

export type OperatorSetOperatorRegistration = {
  operator?: string
  operatorSetId?: string
  avs?: string
  registered?: boolean
  transactionMetadata?: TransactionMetadata
}

export type OperatorSetSplit = {
  operatorSetId?: string
  avs?: string
  operatorSetBasisPoints?: string
  avsBasisPoints?: string
  startTimestamp?: GoogleProtobufTimestamp.Timestamp
  endTimestamp?: GoogleProtobufTimestamp.Timestamp
  transactionMetadata?: TransactionMetadata
}

export type OperatorSetStrategyRegistration = {
  operatorSetId?: string
  avs?: string
  strategy?: string
  registered?: boolean
  transactionMetadata?: TransactionMetadata
}

export type OperatorDirectedRewardSubmission = {
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
  recipient?: string
  transactionMetadata?: TransactionMetadata
}

export type OperatorDirectedOperatorSetRewardSubmission = {
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
  operatorSetId?: string
  transactionMetadata?: TransactionMetadata
}

export type EncumberedMagnitude = {
  operator?: string
  strategy?: string
  encumberedMagnitude?: string
  transactionMetadata?: TransactionMetadata
}

export type OperatorMaxMagnitude = {
  operator?: string
  strategy?: string
  maxMagnitude?: string
  transactionMetadata?: TransactionMetadata
}


type BaseEigenStateChange = {
}

export type EigenStateChange = BaseEigenStateChange
  & OneOf<{ avsOperatorStateChange: AvsOperatorStateChange; operatorShareDelta: OperatorShareDelta; rewardSubmission: RewardSubmission; stakerDelegationChange: StakerDelegationChange; stakerShareDelta: StakerShareDelta; submittedDistributionRoot: SubmittedDistributionRoot; disabledDistributionRoot: DisabledDistributionRoot; completedSlashingWithdrawal: CompletedSlashingWithdrawal; queuedSlashingWithdrawal: QueuedSlashingWithdrawal; slashedOperator: SlashedOperator; slashedOperatorShares: SlashedOperatorShares; defaultOperatorSplit: DefaultOperatorSplit; operatorAllocation: OperatorAllocation; operatorAllocationDelay: OperatorAllocationDelay; operatorAvsSplit: OperatorAVSSplit; operatorPiSplit: OperatorPISplit; operatorSet: OperatorSet; operatorSetOperatorRegistration: OperatorSetOperatorRegistration; operatorSetSplit: OperatorSetSplit; operatorSetStrategyRegistration: OperatorSetStrategyRegistration; operatorDirectedRewardSubmission: OperatorDirectedRewardSubmission; operatorDirectedOperatorSetRewardSubmission: OperatorDirectedOperatorSetRewardSubmission; encumberedMagnitude: EncumberedMagnitude; operatorMaxMagnitude: OperatorMaxMagnitude }>