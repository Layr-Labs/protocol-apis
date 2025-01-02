/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);
export type StateRoot = {
  ethBlockNumber?: string
  ethBlockHash?: string
  stateRoot?: string
}

export type AvsOperatorStateChange = {
}

export type OperatorShareDelta = {
}

export type RewardSubmission = {
}

export type StakerDelegationchange = {
}

export type StakerShareDelta = {
}

export type SubmittedDistributionRoot = {
}

export type DisabledDistributionRoot = {
}


type BaseEigenStateChange = {
}

export type EigenStateChange = BaseEigenStateChange
  & OneOf<{ avsOperatorStateChange: AvsOperatorStateChange; operatorShareDelta: OperatorShareDelta; rewardSubmission: RewardSubmission; stakerDelegationChange: StakerDelegationchange; stakerShareDelta: StakerShareDelta; submittedDistributionRoot: SubmittedDistributionRoot; disabledDistributionRoot: DisabledDistributionRoot }>