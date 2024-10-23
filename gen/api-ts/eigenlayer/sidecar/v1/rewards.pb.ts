/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

export enum AttributableRewardreward_type {
  REWARD_TYPE_AVS = "REWARD_TYPE_AVS",
  REWARD_TYPE_FOR_ALL = "REWARD_TYPE_FOR_ALL",
  REWARD_TYPE_FOR_ALL_EARNERS = "REWARD_TYPE_FOR_ALL_EARNERS",
}

export type EarnerLeaf = {
  earner?: string
  earnerTokenRoot?: string
}

export type TokenLeaf = {
  token?: string
  cumulativeEarnings?: string
}

export type Proof = {
  root?: string
  rootIndex?: number
  earnerIndex?: number
  earnerTreeProof?: string
  earnerLeaf?: EarnerLeaf
  leafIndices?: number[]
  tokenTreeProofs?: string[]
  tokenLeaves?: TokenLeaf[]
}

export type Reward = {
  token?: string
  amount?: string
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
}

export type GetRewardsRootRequest = {
  blockHeight?: string
}

export type GetRewardsRootResponse = {
  rewardsRoot?: string
}


type BaseGenerateRewardsRequest = {
  snapshot?: string
}

export type GenerateRewardsRequest = BaseGenerateRewardsRequest
  & OneOf<{ respondWithRewardsData: boolean }>


type BaseGenerateRewardsResponse = {
  snapshot?: string
}

export type GenerateRewardsResponse = BaseGenerateRewardsResponse
  & OneOf<{ rewards: Reward }>

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

export type GenerateClaimProofRequest = {
  earnerAddress?: string
  tokens?: string[]
}

export type GenerateClaimProofResponse = {
  proof?: Proof
}

export type GetAvailableRewardsRequest = {
  earnerAddress?: string
}

export type GetAvailableRewardsResponse = {
  rewards?: Reward[]
}


type BaseGetTotalClaimedRewardsRequest = {
  earnerAddress?: string
}

export type GetTotalClaimedRewardsRequest = BaseGetTotalClaimedRewardsRequest
  & OneOf<{ blockHeight: string }>

export type GetTotalClaimedRewardsResponse = {
  rewards?: Reward[]
}

export type GetAvailableRewardsTokensRequest = {
  earnerAddress?: string
}

export type GetAvailableRewardsTokensResponse = {
  tokens?: string[]
}

export type SummarizedEarnerReward = {
  token?: string
  earned?: string
  claimable?: string
  claimed?: string
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
  token?: string
  amount?: string
}

export type GetClaimedRewardsByBlockRequest = {
  blockHeight?: string
}

export type GetClaimedRewardsByBlockResponse = {
  rewards?: ClaimedReward[]
}

export class Rewards {
  static GetRewardsRoot(req: GetRewardsRootRequest, initReq?: fm.InitReq): Promise<GetRewardsRootResponse> {
    return fm.fetchReq<GetRewardsRootRequest, GetRewardsRootResponse>(`/rewards/v1/blocks/${req["blockHeight"]}/rewards-root?${fm.renderURLSearchParams(req, ["blockHeight"])}`, {...initReq, method: "GET"})
  }
  static GenerateRewards(req: GenerateRewardsRequest, initReq?: fm.InitReq): Promise<GenerateRewardsResponse> {
    return fm.fetchReq<GenerateRewardsRequest, GenerateRewardsResponse>(`/rewards/v1/generate-rewards`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetRewardsForSnapshot(req: GetRewardsForSnapshotRequest, initReq?: fm.InitReq): Promise<GetRewardsForSnapshotResponse> {
    return fm.fetchReq<GetRewardsForSnapshotRequest, GetRewardsForSnapshotResponse>(`/rewards/v1/rewards/${req["snapshot"]}?${fm.renderURLSearchParams(req, ["snapshot"])}`, {...initReq, method: "GET"})
  }
  static GetAttributableRewardsForSnapshot(req: GetAttributableRewardsForSnapshotRequest, initReq?: fm.InitReq): Promise<GetAttributableRewardsForSnapshotResponse> {
    return fm.fetchReq<GetAttributableRewardsForSnapshotRequest, GetAttributableRewardsForSnapshotResponse>(`/rewards/v1/attributable-rewards/${req["snapshot"]}?${fm.renderURLSearchParams(req, ["snapshot"])}`, {...initReq, method: "GET"})
  }
  static GenerateClaimProof(req: GenerateClaimProofRequest, initReq?: fm.InitReq): Promise<GenerateClaimProofResponse> {
    return fm.fetchReq<GenerateClaimProofRequest, GenerateClaimProofResponse>(`/rewards/v1/claim-proof`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAvailableRewards(req: GetAvailableRewardsRequest, initReq?: fm.InitReq): Promise<GetAvailableRewardsResponse> {
    return fm.fetchReq<GetAvailableRewardsRequest, GetAvailableRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/available-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetTotalClaimedRewards(req: GetTotalClaimedRewardsRequest, initReq?: fm.InitReq): Promise<GetTotalClaimedRewardsResponse> {
    return fm.fetchReq<GetTotalClaimedRewardsRequest, GetTotalClaimedRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/total-claimed-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetAvailableRewardsTokens(req: GetAvailableRewardsTokensRequest, initReq?: fm.InitReq): Promise<GetAvailableRewardsTokensResponse> {
    return fm.fetchReq<GetAvailableRewardsTokensRequest, GetAvailableRewardsTokensResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/available-rewards-tokens?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetSummarizedRewardsForEarner(req: GetSummarizedRewardsForEarnerRequest, initReq?: fm.InitReq): Promise<GetSummarizedRewardsForEarnerResponse> {
    return fm.fetchReq<GetSummarizedRewardsForEarnerRequest, GetSummarizedRewardsForEarnerResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/summarized-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetClaimedRewardsByBlock(req: GetClaimedRewardsByBlockRequest, initReq?: fm.InitReq): Promise<GetClaimedRewardsByBlockResponse> {
    return fm.fetchReq<GetClaimedRewardsByBlockRequest, GetClaimedRewardsByBlockResponse>(`/rewards/v1/blocks/${req["blockHeight"]}/claimed-rewards?${fm.renderURLSearchParams(req, ["blockHeight"])}`, {...initReq, method: "GET"})
  }
}