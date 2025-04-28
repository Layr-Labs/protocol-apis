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

type BaseStakerShare = {
  strategy?: string
  shares?: string
  avsAddresses?: string[]
}

export type StakerShare = BaseStakerShare
  & OneOf<{ operatorAddress: string }>

export type Strategy = {
  strategy?: string
  totalStake?: string
  rewardedTokens?: string[]
}

export type StakerStrategy = {
  stakedAmount?: string
  strategy?: Strategy
}

export type QueuedStrategyWithdrawal = {
  strategy?: string
  amount?: string
  blockNumber?: string
  transactionHash?: string
  logIndex?: string
}

export type StakerWithdrawal = {
  staker?: string
  amount?: string
  blockNumber?: string
}

export type QueueStakerStrategyWithdrawal = {
  strategy?: string
  withdrawals?: StakerWithdrawal[]
}