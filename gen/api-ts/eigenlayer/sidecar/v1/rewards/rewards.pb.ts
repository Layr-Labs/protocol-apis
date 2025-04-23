/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufWrappers from "../../../../google/protobuf/wrappers.pb"
import * as EigenlayerSidecarV1CommonTypes from "../common/types.pb"
import * as EigenlayerSidecarV1RewardsCommon from "./common.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);
export type GetRewardsRootRequest = {
  blockHeight?: string
}

export type GetRewardsRootResponse = {
  rewardsRoot?: EigenlayerSidecarV1RewardsCommon.DistributionRoot
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
  & OneOf<{ rewards: EigenlayerSidecarV1RewardsCommon.Reward }>

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


type BaseGetRewardsForSnapshotRequest = {
  snapshot?: string
}

export type GetRewardsForSnapshotRequest = BaseGetRewardsForSnapshotRequest
  & OneOf<{ earner: string }>

export type GetRewardsForSnapshotResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.Reward[]
}

export type GetRewardsForDistributionRootRequest = {
  rootIndex?: string
}

export type GetRewardsForDistributionRootResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.Reward[]
}

export type GetAttributableRewardsForSnapshotRequest = {
  snapshot?: string
}

export type GetAttributableRewardsForSnapshotResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.AttributableReward[]
}

export type GetAttributableRewardsForDistributionRootRequest = {
  distributionRoot?: string
}

export type GetAttributableRewardsForDistributionRootResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.AttributableReward[]
}

export type GetRewardsByAvsForDistributionRootRequest = {
  rootIndex?: string
  pagination?: EigenlayerSidecarV1CommonTypes.Pagination
}

export type GetRewardsByAvsForDistributionRootResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.AvsReward[]
  nextPage?: EigenlayerSidecarV1CommonTypes.Pagination
}


type BaseGenerateClaimProofRequest = {
  earnerAddress?: string
  tokens?: string[]
}

export type GenerateClaimProofRequest = BaseGenerateClaimProofRequest
  & OneOf<{ rootIndex: GoogleProtobufWrappers.Int64Value }>

export type GenerateClaimProofResponse = {
  proof?: EigenlayerSidecarV1RewardsCommon.Proof
}

export type GetClaimableRewardsRequest = {
  earnerAddress?: string
  blockHeight?: string
}

export type GetClaimableRewardsResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.Reward[]
}


type BaseGetTotalClaimedRewardsRequest = {
  earnerAddress?: string
}

export type GetTotalClaimedRewardsRequest = BaseGetTotalClaimedRewardsRequest
  & OneOf<{ blockHeight: string }>

export type GetTotalClaimedRewardsResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.TotalClaimedReward[]
}


type BaseGetAvailableRewardsTokensRequest = {
  earnerAddress?: string
}

export type GetAvailableRewardsTokensRequest = BaseGetAvailableRewardsTokensRequest
  & OneOf<{ blockHeight: string }>

export type GetAvailableRewardsTokensResponse = {
  tokens?: string[]
}


type BaseGetSummarizedRewardsForEarnerRequest = {
  earnerAddress?: string
}

export type GetSummarizedRewardsForEarnerRequest = BaseGetSummarizedRewardsForEarnerRequest
  & OneOf<{ blockHeight: string }>

export type GetSummarizedRewardsForEarnerResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.SummarizedEarnerReward[]
}

export type GetClaimedRewardsByBlockRequest = {
  blockHeight?: string
}

export type GetClaimedRewardsByBlockResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.ClaimedReward[]
}


type BaseListDistributionRootsRequest = {
}

export type ListDistributionRootsRequest = BaseListDistributionRootsRequest
  & OneOf<{ blockHeight: string }>

export type ListDistributionRootsResponse = {
  distributionRoots?: EigenlayerSidecarV1RewardsCommon.DistributionRoot[]
}

export type ListClaimedRewardsByBlockRangeRequest = {
  earnerAddress?: string
  startBlockHeight?: string
  endBlockHeight?: string
}

export type ListClaimedRewardsByBlockRangeResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.ClaimedReward[]
}


type BaseListEarnerLifetimeRewardsRequest = {
  earnerAddress?: string
}

export type ListEarnerLifetimeRewardsRequest = BaseListEarnerLifetimeRewardsRequest
  & OneOf<{ blockHeight: string }>
  & OneOf<{ pagination: EigenlayerSidecarV1CommonTypes.Pagination }>


type BaseListEarnerLifetimeRewardsResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.RewardAmount[]
}

export type ListEarnerLifetimeRewardsResponse = BaseListEarnerLifetimeRewardsResponse
  & OneOf<{ nextPage: EigenlayerSidecarV1CommonTypes.Pagination }>


type BaseListEarnerHistoricalRewardsRequest = {
  earnerAddress?: string
  tokens?: string[]
}

export type ListEarnerHistoricalRewardsRequest = BaseListEarnerHistoricalRewardsRequest
  & OneOf<{ startBlockHeight: string }>
  & OneOf<{ endBlockHeight: string }>
  & OneOf<{ pagination: EigenlayerSidecarV1CommonTypes.Pagination }>


type BaseListEarnerHistoricalRewardsResponse = {
  rewards?: EigenlayerSidecarV1RewardsCommon.HistoricalReward[]
}

export type ListEarnerHistoricalRewardsResponse = BaseListEarnerHistoricalRewardsResponse
  & OneOf<{ nextPage: EigenlayerSidecarV1CommonTypes.Pagination }>