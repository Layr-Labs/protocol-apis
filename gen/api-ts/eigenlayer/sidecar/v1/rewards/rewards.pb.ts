/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../../../google/protobuf/timestamp.pb"
import * as GoogleProtobufWrappers from "../../../../google/protobuf/wrappers.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

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

export type GetRewardsRootRequest = {
  blockHeight?: string
}

export type GetRewardsRootResponse = {
  rewardsRoot?: DistributionRoot
}

export type GenerateRewardsRequest = {
  cutoffDate?: string
  respondWithRewardsData?: boolean
  waitForComplete?: boolean
}


type BaseGenerateRewardsResponse = {
  cutoffDate?: string
  queued?: boolean
}

export type GenerateRewardsResponse = BaseGenerateRewardsResponse
  & OneOf<{ rewards: Reward }>

export type GenerateStakerOperatorsRequest = {
  cutoffDate?: string
  waitForComplete?: boolean
}

export type GenerateStakerOperatorsResponse = {
  queued?: boolean
}

export type BackfillStakerOperatorsRequest = {
  waitForComplete?: boolean
}

export type BackfillStakerOperatorsResponse = {
  queued?: boolean
}

export type GenerateRewardsRootRequest = {
  cutoffDate?: string
}

export type GenerateRewardsRootResponse = {
  rewardsRoot?: string
  rewardsCalcEndDate?: string
}

export type GetRewardsForSnapshotRequest = {
  snapshot?: string
}

export type GetRewardsForSnapshotResponse = {
  rewards?: Reward[]
}

export type GetAttributableRewardsForSnapshotRequest = {
  snapshot?: string
}

export type GetAttributableRewardsForSnapshotResponse = {
  rewards?: AttributableReward[]
}

export type GetAttributableRewardsForDistributionRootRequest = {
  distributionRoot?: string
}

export type GetAttributableRewardsForDistributionRootResponse = {
  rewards?: AttributableReward[]
}

export type GetRewardsByAvsForDistributionRootRequest = {
  rootIndex?: string
}

export type GetRewardsByAvsForDistributionRootResponse = {
  rewards?: AvsReward[]
}


type BaseGenerateClaimProofRequest = {
  earnerAddress?: string
  tokens?: string[]
}

export type GenerateClaimProofRequest = BaseGenerateClaimProofRequest
  & OneOf<{ rootIndex: GoogleProtobufWrappers.Int64Value }>

export type GenerateClaimProofResponse = {
  proof?: Proof
}

export type GetClaimableRewardsRequest = {
  earnerAddress?: string
  blockHeight?: string
}

export type GetClaimableRewardsResponse = {
  rewards?: Reward[]
}

export type TotalClaimedReward = {
  earner?: string
  token?: string
  amount?: string
}


type BaseGetTotalClaimedRewardsRequest = {
  earnerAddress?: string
}

export type GetTotalClaimedRewardsRequest = BaseGetTotalClaimedRewardsRequest
  & OneOf<{ blockHeight: string }>

export type GetTotalClaimedRewardsResponse = {
  rewards?: TotalClaimedReward[]
}


type BaseGetAvailableRewardsTokensRequest = {
  earnerAddress?: string
}

export type GetAvailableRewardsTokensRequest = BaseGetAvailableRewardsTokensRequest
  & OneOf<{ blockHeight: string }>

export type GetAvailableRewardsTokensResponse = {
  tokens?: string[]
}

export type SummarizedEarnerReward = {
  token?: string
  earned?: string
  active?: string
  claimed?: string
  claimable?: string
}


type BaseGetSummarizedRewardsForEarnerRequest = {
  earnerAddress?: string
}

export type GetSummarizedRewardsForEarnerRequest = BaseGetSummarizedRewardsForEarnerRequest
  & OneOf<{ blockHeight: string }>

export type GetSummarizedRewardsForEarnerResponse = {
  rewards?: SummarizedEarnerReward[]
}

export type ClaimedReward = {
  earner?: string
  claimer?: string
  token?: string
  amount?: string
  distributionRoot?: string
  blockNumber?: string
  recipient?: string
}

export type GetClaimedRewardsByBlockRequest = {
  blockHeight?: string
}

export type GetClaimedRewardsByBlockResponse = {
  rewards?: ClaimedReward[]
}


type BaseListDistributionRootsRequest = {
}

export type ListDistributionRootsRequest = BaseListDistributionRootsRequest
  & OneOf<{ blockHeight: string }>

export type ListDistributionRootsResponse = {
  distributionRoots?: DistributionRoot[]
}

export type ListClaimedRewardsByBlockRangeRequest = {
  earnerAddress?: string
  startBlockHeight?: string
  endBlockHeight?: string
}

export type ListClaimedRewardsByBlockRangeResponse = {
  rewards?: ClaimedReward[]
}