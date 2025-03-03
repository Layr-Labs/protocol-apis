/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../../fetch.pb"
import * as EigenlayerSidecarV1RewardsRewards from "./rewards.pb"
export class Rewards {
  static GetRewardsRoot(req: EigenlayerSidecarV1RewardsRewards.GetRewardsRootRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetRewardsRootResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetRewardsRootRequest, EigenlayerSidecarV1RewardsRewards.GetRewardsRootResponse>(`/rewards/v1/blocks/${req["blockHeight"]}/rewards-root?${fm.renderURLSearchParams(req, ["blockHeight"])}`, {...initReq, method: "GET"})
  }
  static GenerateRewards(req: EigenlayerSidecarV1RewardsRewards.GenerateRewardsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GenerateRewardsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GenerateRewardsRequest, EigenlayerSidecarV1RewardsRewards.GenerateRewardsResponse>(`/rewards/v1/generate-rewards`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GenerateStakerOperators(req: EigenlayerSidecarV1RewardsRewards.GenerateStakerOperatorsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GenerateStakerOperatorsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GenerateStakerOperatorsRequest, EigenlayerSidecarV1RewardsRewards.GenerateStakerOperatorsResponse>(`/rewards/v1/generate-staker-operators`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static BackfillStakerOperators(req: EigenlayerSidecarV1RewardsRewards.BackfillStakerOperatorsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.BackfillStakerOperatorsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.BackfillStakerOperatorsRequest, EigenlayerSidecarV1RewardsRewards.BackfillStakerOperatorsResponse>(`/rewards/v1/backfill-staker-operators`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GenerateRewardsRoot(req: EigenlayerSidecarV1RewardsRewards.GenerateRewardsRootRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GenerateRewardsRootResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GenerateRewardsRootRequest, EigenlayerSidecarV1RewardsRewards.GenerateRewardsRootResponse>(`/rewards/v1/generate-rewards-root`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetRewardsForSnapshot(req: EigenlayerSidecarV1RewardsRewards.GetRewardsForSnapshotRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetRewardsForSnapshotResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetRewardsForSnapshotRequest, EigenlayerSidecarV1RewardsRewards.GetRewardsForSnapshotResponse>(`/rewards/v1/rewards/${req["snapshot"]}?${fm.renderURLSearchParams(req, ["snapshot"])}`, {...initReq, method: "GET"})
  }
  static GetAttributableRewardsForSnapshot(req: EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForSnapshotRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForSnapshotResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForSnapshotRequest, EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForSnapshotResponse>(`/rewards/v1/attributable-rewards/${req["snapshot"]}?${fm.renderURLSearchParams(req, ["snapshot"])}`, {...initReq, method: "GET"})
  }
  static GetAttributableRewardsForDistributionRoot(req: EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForDistributionRootRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForDistributionRootResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForDistributionRootRequest, EigenlayerSidecarV1RewardsRewards.GetAttributableRewardsForDistributionRootResponse>(`/rewards/v1/attributable-rewards-by-root/${req["distributionRoot"]}?${fm.renderURLSearchParams(req, ["distributionRoot"])}`, {...initReq, method: "GET"})
  }
  static GetRewardsByAvsForDistributionRoot(req: EigenlayerSidecarV1RewardsRewards.GetRewardsByAvsForDistributionRootRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetRewardsByAvsForDistributionRootResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetRewardsByAvsForDistributionRootRequest, EigenlayerSidecarV1RewardsRewards.GetRewardsByAvsForDistributionRootResponse>(`/rewards/v1/avs-rewards-by-root/${req["rootIndex"]}?${fm.renderURLSearchParams(req, ["rootIndex"])}`, {...initReq, method: "GET"})
  }
  static GenerateClaimProof(req: EigenlayerSidecarV1RewardsRewards.GenerateClaimProofRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GenerateClaimProofResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GenerateClaimProofRequest, EigenlayerSidecarV1RewardsRewards.GenerateClaimProofResponse>(`/rewards/v1/claim-proof`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetClaimableRewards(req: EigenlayerSidecarV1RewardsRewards.GetClaimableRewardsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetClaimableRewardsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetClaimableRewardsRequest, EigenlayerSidecarV1RewardsRewards.GetClaimableRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/claimable-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetTotalClaimedRewards(req: EigenlayerSidecarV1RewardsRewards.GetTotalClaimedRewardsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetTotalClaimedRewardsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetTotalClaimedRewardsRequest, EigenlayerSidecarV1RewardsRewards.GetTotalClaimedRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/total-claimed-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetAvailableRewardsTokens(req: EigenlayerSidecarV1RewardsRewards.GetAvailableRewardsTokensRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetAvailableRewardsTokensResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetAvailableRewardsTokensRequest, EigenlayerSidecarV1RewardsRewards.GetAvailableRewardsTokensResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/available-rewards-tokens?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetSummarizedRewardsForEarner(req: EigenlayerSidecarV1RewardsRewards.GetSummarizedRewardsForEarnerRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetSummarizedRewardsForEarnerResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetSummarizedRewardsForEarnerRequest, EigenlayerSidecarV1RewardsRewards.GetSummarizedRewardsForEarnerResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/summarized-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static GetClaimedRewardsByBlock(req: EigenlayerSidecarV1RewardsRewards.GetClaimedRewardsByBlockRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.GetClaimedRewardsByBlockResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.GetClaimedRewardsByBlockRequest, EigenlayerSidecarV1RewardsRewards.GetClaimedRewardsByBlockResponse>(`/rewards/v1/blocks/${req["blockHeight"]}/claimed-rewards?${fm.renderURLSearchParams(req, ["blockHeight"])}`, {...initReq, method: "GET"})
  }
  static ListClaimedRewardsByBlockRange(req: EigenlayerSidecarV1RewardsRewards.ListClaimedRewardsByBlockRangeRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.ListClaimedRewardsByBlockRangeResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.ListClaimedRewardsByBlockRangeRequest, EigenlayerSidecarV1RewardsRewards.ListClaimedRewardsByBlockRangeResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/claimed-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static ListDistributionRoots(req: EigenlayerSidecarV1RewardsRewards.ListDistributionRootsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.ListDistributionRootsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.ListDistributionRootsRequest, EigenlayerSidecarV1RewardsRewards.ListDistributionRootsResponse>(`/rewards/v1/distribution-roots?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static ListEarnerLifetimeRewards(req: EigenlayerSidecarV1RewardsRewards.ListEarnerLifetimeRewardsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.ListEarnerLifetimeRewardsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.ListEarnerLifetimeRewardsRequest, EigenlayerSidecarV1RewardsRewards.ListEarnerLifetimeRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/lifetime-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
  static ListEarnerHistoricalRewards(req: EigenlayerSidecarV1RewardsRewards.ListEarnerHistoricalRewardsRequest, initReq?: fm.InitReq): Promise<EigenlayerSidecarV1RewardsRewards.ListEarnerHistoricalRewardsResponse> {
    return fm.fetchReq<EigenlayerSidecarV1RewardsRewards.ListEarnerHistoricalRewardsRequest, EigenlayerSidecarV1RewardsRewards.ListEarnerHistoricalRewardsResponse>(`/rewards/v1/earners/${req["earnerAddress"]}/historical-rewards?${fm.renderURLSearchParams(req, ["earnerAddress"])}`, {...initReq, method: "GET"})
  }
}