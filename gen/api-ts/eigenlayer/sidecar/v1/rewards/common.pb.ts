/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../../../google/protobuf/timestamp.pb"

export enum RewardType {
  REWARD_TYPE_AVS = "REWARD_TYPE_AVS",
  REWARD_TYPE_FOR_ALL = "REWARD_TYPE_FOR_ALL",
  REWARD_TYPE_FOR_ALL_EARNERS = "REWARD_TYPE_FOR_ALL_EARNERS",
}

export type EarnerLeaf = {
  earner?: string
  earnerTokenRoot?: Uint8Array
}

export type TokenLeaf = {
  token?: string
  cumulativeEarnings?: string
}

export type Proof = {
  root?: Uint8Array
  rootIndex?: number
  earnerIndex?: number
  earnerTreeProof?: Uint8Array
  earnerLeaf?: EarnerLeaf
  tokenIndices?: number[]
  tokenTreeProofs?: Uint8Array[]
  tokenLeaves?: TokenLeaf[]
}

export type Reward = {
  earner?: string
  token?: string
  amount?: string
  snapshot?: string
}

export type AttributableReward = {
  earner?: string
  operator?: string
  avs?: string
  token?: string
  strategy?: string
  multiplier?: string
  amount?: string
  shares?: string
  rewardHash?: string
  snapshot?: string
  rewardType?: RewardType
}

export type AvsReward = {
  earner?: string
  avs?: string
  token?: string
  amount?: string
  rewardHash?: string
  snapshot?: string
  rewardType?: RewardType
}

export type DistributionRoot = {
  root?: string
  rootIndex?: string
  rewardsCalculationEnd?: GoogleProtobufTimestamp.Timestamp
  rewardsCalculationEndUnit?: string
  activatedAt?: GoogleProtobufTimestamp.Timestamp
  activatedAtUnit?: string
  createdAtBlockNumber?: string
  transactionHash?: string
  blockHeight?: string
  logIndex?: string
  disabled?: boolean
}

export type TotalClaimedReward = {
  earner?: string
  token?: string
  amount?: string
}

export type SummarizedEarnerReward = {
  token?: string
  earned?: string
  active?: string
  claimed?: string
  claimable?: string
}

export type ClaimedReward = {
  earner?: string
  claimer?: string
  token?: string
  amount?: string
  distributionRoot?: string
  blockNumber?: string
  recipient?: string
  transactionHash?: string
  logIndex?: string
}

export type RewardAmount = {
  token?: string
  amount?: string
}

export type HistoricalRewardHistoricalRewardAmount = {
  amount?: string
  snapshot?: string
}

export type HistoricalReward = {
  token?: string
  amounts?: HistoricalRewardHistoricalRewardAmount[]
}