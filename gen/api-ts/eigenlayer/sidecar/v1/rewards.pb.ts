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

export class Rewards {
  static GetRewardsRoot(req: GetRewardsRootRequest, initReq?: fm.InitReq): Promise<GetRewardsRootResponse> {
    return fm.fetchReq<GetRewardsRootRequest, GetRewardsRootResponse>(`/rewards/v1/rewards-roots/${req["blockHeight"]}?${fm.renderURLSearchParams(req, ["blockHeight"])}`, {...initReq, method: "GET"})
  }
  static GenerateRewards(req: GenerateRewardsRequest, initReq?: fm.InitReq): Promise<GenerateRewardsResponse> {
    return fm.fetchReq<GenerateRewardsRequest, GenerateRewardsResponse>(`/rewards/v1/generate-rewards`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetRewardsForSnapshot(req: GetRewardsForSnapshotRequest, initReq?: fm.InitReq): Promise<GetRewardsForSnapshotResponse> {
    return fm.fetchReq<GetRewardsForSnapshotRequest, GetRewardsForSnapshotResponse>(`/rewards/v1/rewards/${req["snapshot"]}?${fm.renderURLSearchParams(req, ["snapshot"])}`, {...initReq, method: "GET"})
  }
  static GenerateClaimProof(req: GenerateClaimProofRequest, initReq?: fm.InitReq): Promise<GenerateClaimProofResponse> {
    return fm.fetchReq<GenerateClaimProofRequest, GenerateClaimProofResponse>(`/rewards/v1/claim-proof`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAvailableRewards(req: GetAvailableRewardsRequest, initReq?: fm.InitReq): Promise<GetAvailableRewardsResponse> {
    return fm.fetchReq<GetAvailableRewardsRequest, GetAvailableRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/available-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetTotalClaimedRewards(req: GetAvailableRewardsRequest, initReq?: fm.InitReq): Promise<GetAvailableRewardsResponse> {
    return fm.fetchReq<GetAvailableRewardsRequest, GetAvailableRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/total-claimed-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetAvailableRewardsTokens(req: GetAvailableRewardsTokensRequest, initReq?: fm.InitReq): Promise<GetAvailableRewardsTokensResponse> {
    return fm.fetchReq<GetAvailableRewardsTokensRequest, GetAvailableRewardsTokensResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/available-rewards-tokens?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
}