/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
export type SlashingEventStrategy = {
  strategy?: string
  wadSlashed?: string
  totalSharesSlashed?: string
}

export type SlashingEventOperatorSet = {
  id?: string
  avs?: string
}

export type SlashingEvent = {
  description?: string
  operator?: string
  transactionHash?: string
  logIndex?: string
  blockNumber?: string
  strategies?: SlashingEventStrategy[]
  operatorSet?: SlashingEventOperatorSet
}

export type SlashedStaker = {
  staker?: string
  transactionHash?: string
  logIndex?: string
  blockNumber?: string
  strategy?: string
  shares?: string
}

export type StakerSlashingEvent = {
  staker?: string
  slashingEvent?: SlashingEvent
}