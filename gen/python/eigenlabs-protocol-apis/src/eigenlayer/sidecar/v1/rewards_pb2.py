# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: eigenlayer/sidecar/v1/rewards.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'eigenlayer/sidecar/v1/rewards.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#eigenlayer/sidecar/v1/rewards.proto\x12\x1d\x65igenlayer.sidecar.rewards.v1\x1a\x1cgoogle/api/annotations.proto\"P\n\nEarnerLeaf\x12\x16\n\x06\x65\x61rner\x18\x01 \x01(\tR\x06\x65\x61rner\x12*\n\x11\x65\x61rner_token_root\x18\x02 \x01(\tR\x0f\x65\x61rnerTokenRoot\"R\n\tTokenLeaf\x12\x14\n\x05token\x18\x01 \x01(\tR\x05token\x12/\n\x13\x63umulative_earnings\x18\x02 \x01(\tR\x12\x63umulativeEarnings\"\xf1\x02\n\x05Proof\x12\x12\n\x04root\x18\x01 \x01(\tR\x04root\x12\x1d\n\nroot_index\x18\x02 \x01(\rR\trootIndex\x12!\n\x0c\x65\x61rner_index\x18\x03 \x01(\rR\x0b\x65\x61rnerIndex\x12*\n\x11\x65\x61rner_tree_proof\x18\x04 \x01(\tR\x0f\x65\x61rnerTreeProof\x12J\n\x0b\x65\x61rner_leaf\x18\x05 \x01(\x0b\x32).eigenlayer.sidecar.rewards.v1.EarnerLeafR\nearnerLeaf\x12!\n\x0cleaf_indices\x18\x06 \x03(\rR\x0bleafIndices\x12*\n\x11token_tree_proofs\x18\x07 \x03(\tR\x0ftokenTreeProofs\x12K\n\x0ctoken_leaves\x18\x08 \x03(\x0b\x32(.eigenlayer.sidecar.rewards.v1.TokenLeafR\x0btokenLeaves\"j\n\x06Reward\x12\x16\n\x06\x65\x61rner\x18\x01 \x01(\tR\x06\x65\x61rner\x12\x14\n\x05token\x18\x02 \x01(\tR\x05token\x12\x16\n\x06\x61mount\x18\x03 \x01(\tR\x06\x61mount\x12\x1a\n\x08snapshot\x18\x04 \x01(\tR\x08snapshot\"\xf7\x02\n\x12\x41ttributableReward\x12\x16\n\x06\x65\x61rner\x18\x01 \x01(\tR\x06\x65\x61rner\x12\x1a\n\x08operator\x18\x02 \x01(\tR\x08operator\x12\x10\n\x03\x61vs\x18\x03 \x01(\tR\x03\x61vs\x12\x14\n\x05token\x18\x04 \x01(\tR\x05token\x12\x1a\n\x08strategy\x18\x05 \x01(\tR\x08strategy\x12\x1e\n\nmultiplier\x18\x06 \x01(\tR\nmultiplier\x12\x16\n\x06\x61mount\x18\x07 \x01(\tR\x06\x61mount\x12\x16\n\x06shares\x18\x08 \x01(\tR\x06shares\x12\x1f\n\x0breward_hash\x18\t \x01(\tR\nrewardHash\x12\x1a\n\x08snapshot\x18\n \x01(\tR\x08snapshot\"\\\n\x0breward_type\x12\x13\n\x0fREWARD_TYPE_AVS\x10\x00\x12\x17\n\x13REWARD_TYPE_FOR_ALL\x10\x01\x12\x1f\n\x1bREWARD_TYPE_FOR_ALL_EARNERS\x10\x02\"\xc4\x01\n\x10\x44istributionRoot\x12\x12\n\x04root\x18\x01 \x01(\tR\x04root\x12!\n\x0c\x62lock_height\x18\x02 \x01(\x04R\x0b\x62lockHeight\x12:\n\x19\x63\x61lculation_end_timestamp\x18\x03 \x01(\x04R\x17\x63\x61lculationEndTimestamp\x12!\n\x0c\x61\x63tivated_at\x18\x04 \x01(\x04R\x0b\x61\x63tivatedAt\x12\x1a\n\x08\x64isabled\x18\x05 \x01(\x08R\x08\x64isabled\":\n\x15GetRewardsRootRequest\x12!\n\x0c\x62lock_height\x18\x01 \x01(\x04R\x0b\x62lockHeight\"l\n\x16GetRewardsRootResponse\x12R\n\x0crewards_root\x18\x01 \x01(\x0b\x32/.eigenlayer.sidecar.rewards.v1.DistributionRootR\x0brewardsRoot\"t\n\x16GenerateRewardsRequest\x12\x1f\n\x0b\x63utoff_date\x18\x01 \x01(\tR\ncutoffDate\x12\x39\n\x19respond_with_rewards_data\x18\x02 \x01(\x08R\x16respondWithRewardsData\"\x8c\x01\n\x17GenerateRewardsResponse\x12\x1f\n\x0b\x63utoff_date\x18\x01 \x01(\tR\ncutoffDate\x12\x44\n\x07rewards\x18\x02 \x01(\x0b\x32%.eigenlayer.sidecar.rewards.v1.RewardH\x00R\x07rewards\x88\x01\x01\x42\n\n\x08_rewards\"A\n\x1eGenerateStakerOperatorsRequest\x12\x1f\n\x0b\x63utoff_date\x18\x01 \x01(\tR\ncutoffDate\"!\n\x1fGenerateStakerOperatorsResponse\" \n\x1e\x42\x61\x63kfillStakerOperatorsRequest\"!\n\x1f\x42\x61\x63kfillStakerOperatorsResponse\"=\n\x1aGenerateRewardsRootRequest\x12\x1f\n\x0b\x63utoff_date\x18\x01 \x01(\tR\ncutoffDate\"s\n\x1bGenerateRewardsRootResponse\x12!\n\x0crewards_root\x18\x01 \x01(\tR\x0brewardsRoot\x12\x31\n\x15rewards_calc_end_date\x18\x02 \x01(\tR\x12rewardsCalcEndDate\":\n\x1cGetRewardsForSnapshotRequest\x12\x1a\n\x08snapshot\x18\x01 \x01(\tR\x08snapshot\"`\n\x1dGetRewardsForSnapshotResponse\x12?\n\x07rewards\x18\x01 \x03(\x0b\x32%.eigenlayer.sidecar.rewards.v1.RewardR\x07rewards\"F\n(GetAttributableRewardsForSnapshotRequest\x12\x1a\n\x08snapshot\x18\x01 \x01(\tR\x08snapshot\"x\n)GetAttributableRewardsForSnapshotResponse\x12K\n\x07rewards\x18\x01 \x03(\x0b\x32\x31.eigenlayer.sidecar.rewards.v1.AttributableRewardR\x07rewards\"_\n0GetAttributableRewardsForDistributionRootRequest\x12+\n\x11\x64istribution_root\x18\x01 \x01(\tR\x10\x64istributionRoot\"\x80\x01\n1GetAttributableRewardsForDistributionRootResponse\x12K\n\x07rewards\x18\x01 \x03(\x0b\x32\x31.eigenlayer.sidecar.rewards.v1.AttributableRewardR\x07rewards\"Z\n\x19GenerateClaimProofRequest\x12%\n\x0e\x65\x61rner_address\x18\x01 \x01(\tR\rearnerAddress\x12\x16\n\x06tokens\x18\x02 \x03(\tR\x06tokens\"X\n\x1aGenerateClaimProofResponse\x12:\n\x05proof\x18\x01 \x01(\x0b\x32$.eigenlayer.sidecar.rewards.v1.ProofR\x05proof\"C\n\x1aGetAvailableRewardsRequest\x12%\n\x0e\x65\x61rner_address\x18\x01 \x01(\tR\rearnerAddress\"^\n\x1bGetAvailableRewardsResponse\x12?\n\x07rewards\x18\x01 \x03(\x0b\x32%.eigenlayer.sidecar.rewards.v1.RewardR\x07rewards\"\x7f\n\x1dGetTotalClaimedRewardsRequest\x12%\n\x0e\x65\x61rner_address\x18\x01 \x01(\tR\rearnerAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"a\n\x1eGetTotalClaimedRewardsResponse\x12?\n\x07rewards\x18\x01 \x03(\x0b\x32%.eigenlayer.sidecar.rewards.v1.RewardR\x07rewards\"I\n GetAvailableRewardsTokensRequest\x12%\n\x0e\x65\x61rner_address\x18\x01 \x01(\tR\rearnerAddress\";\n!GetAvailableRewardsTokensResponse\x12\x16\n\x06tokens\x18\x01 \x03(\tR\x06tokens\"~\n\x16SummarizedEarnerReward\x12\x14\n\x05token\x18\x01 \x01(\tR\x05token\x12\x16\n\x06\x65\x61rned\x18\x02 \x01(\tR\x06\x65\x61rned\x12\x1c\n\tclaimable\x18\x03 \x01(\tR\tclaimable\x12\x18\n\x07\x63laimed\x18\x04 \x01(\tR\x07\x63laimed\"\x86\x01\n$GetSummarizedRewardsForEarnerRequest\x12%\n\x0e\x65\x61rner_address\x18\x01 \x01(\tR\rearnerAddress\x12&\n\x0c\x62lock_height\x18\x02 \x01(\x04H\x00R\x0b\x62lockHeight\x88\x01\x01\x42\x0f\n\r_block_height\"x\n%GetSummarizedRewardsForEarnerResponse\x12O\n\x07rewards\x18\x01 \x03(\x0b\x32\x35.eigenlayer.sidecar.rewards.v1.SummarizedEarnerRewardR\x07rewards\"\x9c\x01\n\rClaimedReward\x12\x16\n\x06\x65\x61rner\x18\x01 \x01(\tR\x06\x65\x61rner\x12\x18\n\x07\x63laimer\x18\x02 \x01(\tR\x07\x63laimer\x12\x14\n\x05token\x18\x03 \x01(\tR\x05token\x12\x16\n\x06\x61mount\x18\x04 \x01(\tR\x06\x61mount\x12+\n\x11\x64istribution_root\x18\x05 \x01(\tR\x10\x64istributionRoot\"D\n\x1fGetClaimedRewardsByBlockRequest\x12!\n\x0c\x62lock_height\x18\x01 \x01(\x04R\x0b\x62lockHeight\"j\n GetClaimedRewardsByBlockResponse\x12\x46\n\x07rewards\x18\x01 \x03(\x0b\x32,.eigenlayer.sidecar.rewards.v1.ClaimedRewardR\x07rewards2\x90\x17\n\x07Rewards\x12\xb5\x01\n\x0eGetRewardsRoot\x12\x34.eigenlayer.sidecar.rewards.v1.GetRewardsRootRequest\x1a\x35.eigenlayer.sidecar.rewards.v1.GetRewardsRootResponse\"6\x82\xd3\xe4\x93\x02\x30\x12./rewards/v1/blocks/{block_height}/rewards-root\x12\xa9\x01\n\x0fGenerateRewards\x12\x35.eigenlayer.sidecar.rewards.v1.GenerateRewardsRequest\x1a\x36.eigenlayer.sidecar.rewards.v1.GenerateRewardsResponse\"\'\x82\xd3\xe4\x93\x02!\"\x1c/rewards/v1/generate-rewards:\x01*\x12\xca\x01\n\x17GenerateStakerOperators\x12=.eigenlayer.sidecar.rewards.v1.GenerateStakerOperatorsRequest\x1a>.eigenlayer.sidecar.rewards.v1.GenerateStakerOperatorsResponse\"0\x82\xd3\xe4\x93\x02*\"%/rewards/v1/generate-staker-operators:\x01*\x12\xca\x01\n\x17\x42\x61\x63kfillStakerOperators\x12=.eigenlayer.sidecar.rewards.v1.BackfillStakerOperatorsRequest\x1a>.eigenlayer.sidecar.rewards.v1.BackfillStakerOperatorsResponse\"0\x82\xd3\xe4\x93\x02*\"%/rewards/v1/backfill-staker-operators:\x01*\x12\xba\x01\n\x13GenerateRewardsRoot\x12\x39.eigenlayer.sidecar.rewards.v1.GenerateRewardsRootRequest\x1a:.eigenlayer.sidecar.rewards.v1.GenerateRewardsRootResponse\",\x82\xd3\xe4\x93\x02&\"!/rewards/v1/generate-rewards-root:\x01*\x12\xba\x01\n\x15GetRewardsForSnapshot\x12;.eigenlayer.sidecar.rewards.v1.GetRewardsForSnapshotRequest\x1a<.eigenlayer.sidecar.rewards.v1.GetRewardsForSnapshotResponse\"&\x82\xd3\xe4\x93\x02 \x12\x1e/rewards/v1/rewards/{snapshot}\x12\xeb\x01\n!GetAttributableRewardsForSnapshot\x12G.eigenlayer.sidecar.rewards.v1.GetAttributableRewardsForSnapshotRequest\x1aH.eigenlayer.sidecar.rewards.v1.GetAttributableRewardsForSnapshotResponse\"3\x82\xd3\xe4\x93\x02-\x12+/rewards/v1/attributable-rewards/{snapshot}\x12\x94\x02\n)GetAttributableRewardsForDistributionRoot\x12O.eigenlayer.sidecar.rewards.v1.GetAttributableRewardsForDistributionRootRequest\x1aP.eigenlayer.sidecar.rewards.v1.GetAttributableRewardsForDistributionRootResponse\"D\x82\xd3\xe4\x93\x02>\x12</rewards/v1/attributable-rewards-by-root/{distribution_root}\x12\xad\x01\n\x12GenerateClaimProof\x12\x38.eigenlayer.sidecar.rewards.v1.GenerateClaimProofRequest\x1a\x39.eigenlayer.sidecar.rewards.v1.GenerateClaimProofResponse\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/rewards/v1/claim-proof:\x01*\x12\xcc\x01\n\x13GetAvailableRewards\x12\x39.eigenlayer.sidecar.rewards.v1.GetAvailableRewardsRequest\x1a:.eigenlayer.sidecar.rewards.v1.GetAvailableRewardsResponse\">\x82\xd3\xe4\x93\x02\x38\x12\x36/rewards/v1/earners/{earner_address}/available-rewards\x12\xd9\x01\n\x16GetTotalClaimedRewards\x12<.eigenlayer.sidecar.rewards.v1.GetTotalClaimedRewardsRequest\x1a=.eigenlayer.sidecar.rewards.v1.GetTotalClaimedRewardsResponse\"B\x82\xd3\xe4\x93\x02<\x12:/rewards/v1/earners/{earner_address}/total-claimed-rewards\x12\xe5\x01\n\x19GetAvailableRewardsTokens\x12?.eigenlayer.sidecar.rewards.v1.GetAvailableRewardsTokensRequest\x1a@.eigenlayer.sidecar.rewards.v1.GetAvailableRewardsTokensResponse\"E\x82\xd3\xe4\x93\x02?\x12=/rewards/v1/earners/{earner_address}/available-rewards-tokens\x12\xeb\x01\n\x1dGetSummarizedRewardsForEarner\x12\x43.eigenlayer.sidecar.rewards.v1.GetSummarizedRewardsForEarnerRequest\x1a\x44.eigenlayer.sidecar.rewards.v1.GetSummarizedRewardsForEarnerResponse\"?\x82\xd3\xe4\x93\x02\x39\x12\x37/rewards/v1/earners/{earner_address}/summarized-rewards\x12\xd6\x01\n\x18GetClaimedRewardsByBlock\x12>.eigenlayer.sidecar.rewards.v1.GetClaimedRewardsByBlockRequest\x1a?.eigenlayer.sidecar.rewards.v1.GetClaimedRewardsByBlockResponse\"9\x82\xd3\xe4\x93\x02\x33\x12\x31/rewards/v1/blocks/{block_height}/claimed-rewardsB\x95\x02\n!com.eigenlayer.sidecar.rewards.v1B\x0cRewardsProtoP\x01ZKgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/rewards/v1\xa2\x02\x03\x45SR\xaa\x02\x1d\x45igenlayer.Sidecar.Rewards.V1\xca\x02\x1d\x45igenlayer\\Sidecar\\Rewards\\V1\xe2\x02)Eigenlayer\\Sidecar\\Rewards\\V1\\GPBMetadata\xea\x02 Eigenlayer::Sidecar::Rewards::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'eigenlayer.sidecar.v1.rewards_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n!com.eigenlayer.sidecar.rewards.v1B\014RewardsProtoP\001ZKgithub.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/sidecar/rewards/v1\242\002\003ESR\252\002\035Eigenlayer.Sidecar.Rewards.V1\312\002\035Eigenlayer\\Sidecar\\Rewards\\V1\342\002)Eigenlayer\\Sidecar\\Rewards\\V1\\GPBMetadata\352\002 Eigenlayer::Sidecar::Rewards::V1'
  _globals['_REWARDS'].methods_by_name['GetRewardsRoot']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetRewardsRoot']._serialized_options = b'\202\323\344\223\0020\022./rewards/v1/blocks/{block_height}/rewards-root'
  _globals['_REWARDS'].methods_by_name['GenerateRewards']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GenerateRewards']._serialized_options = b'\202\323\344\223\002!\"\034/rewards/v1/generate-rewards:\001*'
  _globals['_REWARDS'].methods_by_name['GenerateStakerOperators']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GenerateStakerOperators']._serialized_options = b'\202\323\344\223\002*\"%/rewards/v1/generate-staker-operators:\001*'
  _globals['_REWARDS'].methods_by_name['BackfillStakerOperators']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['BackfillStakerOperators']._serialized_options = b'\202\323\344\223\002*\"%/rewards/v1/backfill-staker-operators:\001*'
  _globals['_REWARDS'].methods_by_name['GenerateRewardsRoot']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GenerateRewardsRoot']._serialized_options = b'\202\323\344\223\002&\"!/rewards/v1/generate-rewards-root:\001*'
  _globals['_REWARDS'].methods_by_name['GetRewardsForSnapshot']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetRewardsForSnapshot']._serialized_options = b'\202\323\344\223\002 \022\036/rewards/v1/rewards/{snapshot}'
  _globals['_REWARDS'].methods_by_name['GetAttributableRewardsForSnapshot']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetAttributableRewardsForSnapshot']._serialized_options = b'\202\323\344\223\002-\022+/rewards/v1/attributable-rewards/{snapshot}'
  _globals['_REWARDS'].methods_by_name['GetAttributableRewardsForDistributionRoot']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetAttributableRewardsForDistributionRoot']._serialized_options = b'\202\323\344\223\002>\022</rewards/v1/attributable-rewards-by-root/{distribution_root}'
  _globals['_REWARDS'].methods_by_name['GenerateClaimProof']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GenerateClaimProof']._serialized_options = b'\202\323\344\223\002\034\"\027/rewards/v1/claim-proof:\001*'
  _globals['_REWARDS'].methods_by_name['GetAvailableRewards']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetAvailableRewards']._serialized_options = b'\202\323\344\223\0028\0226/rewards/v1/earners/{earner_address}/available-rewards'
  _globals['_REWARDS'].methods_by_name['GetTotalClaimedRewards']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetTotalClaimedRewards']._serialized_options = b'\202\323\344\223\002<\022:/rewards/v1/earners/{earner_address}/total-claimed-rewards'
  _globals['_REWARDS'].methods_by_name['GetAvailableRewardsTokens']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetAvailableRewardsTokens']._serialized_options = b'\202\323\344\223\002?\022=/rewards/v1/earners/{earner_address}/available-rewards-tokens'
  _globals['_REWARDS'].methods_by_name['GetSummarizedRewardsForEarner']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetSummarizedRewardsForEarner']._serialized_options = b'\202\323\344\223\0029\0227/rewards/v1/earners/{earner_address}/summarized-rewards'
  _globals['_REWARDS'].methods_by_name['GetClaimedRewardsByBlock']._loaded_options = None
  _globals['_REWARDS'].methods_by_name['GetClaimedRewardsByBlock']._serialized_options = b'\202\323\344\223\0023\0221/rewards/v1/blocks/{block_height}/claimed-rewards'
  _globals['_EARNERLEAF']._serialized_start=100
  _globals['_EARNERLEAF']._serialized_end=180
  _globals['_TOKENLEAF']._serialized_start=182
  _globals['_TOKENLEAF']._serialized_end=264
  _globals['_PROOF']._serialized_start=267
  _globals['_PROOF']._serialized_end=636
  _globals['_REWARD']._serialized_start=638
  _globals['_REWARD']._serialized_end=744
  _globals['_ATTRIBUTABLEREWARD']._serialized_start=747
  _globals['_ATTRIBUTABLEREWARD']._serialized_end=1122
  _globals['_ATTRIBUTABLEREWARD_REWARD_TYPE']._serialized_start=1030
  _globals['_ATTRIBUTABLEREWARD_REWARD_TYPE']._serialized_end=1122
  _globals['_DISTRIBUTIONROOT']._serialized_start=1125
  _globals['_DISTRIBUTIONROOT']._serialized_end=1321
  _globals['_GETREWARDSROOTREQUEST']._serialized_start=1323
  _globals['_GETREWARDSROOTREQUEST']._serialized_end=1381
  _globals['_GETREWARDSROOTRESPONSE']._serialized_start=1383
  _globals['_GETREWARDSROOTRESPONSE']._serialized_end=1491
  _globals['_GENERATEREWARDSREQUEST']._serialized_start=1493
  _globals['_GENERATEREWARDSREQUEST']._serialized_end=1609
  _globals['_GENERATEREWARDSRESPONSE']._serialized_start=1612
  _globals['_GENERATEREWARDSRESPONSE']._serialized_end=1752
  _globals['_GENERATESTAKEROPERATORSREQUEST']._serialized_start=1754
  _globals['_GENERATESTAKEROPERATORSREQUEST']._serialized_end=1819
  _globals['_GENERATESTAKEROPERATORSRESPONSE']._serialized_start=1821
  _globals['_GENERATESTAKEROPERATORSRESPONSE']._serialized_end=1854
  _globals['_BACKFILLSTAKEROPERATORSREQUEST']._serialized_start=1856
  _globals['_BACKFILLSTAKEROPERATORSREQUEST']._serialized_end=1888
  _globals['_BACKFILLSTAKEROPERATORSRESPONSE']._serialized_start=1890
  _globals['_BACKFILLSTAKEROPERATORSRESPONSE']._serialized_end=1923
  _globals['_GENERATEREWARDSROOTREQUEST']._serialized_start=1925
  _globals['_GENERATEREWARDSROOTREQUEST']._serialized_end=1986
  _globals['_GENERATEREWARDSROOTRESPONSE']._serialized_start=1988
  _globals['_GENERATEREWARDSROOTRESPONSE']._serialized_end=2103
  _globals['_GETREWARDSFORSNAPSHOTREQUEST']._serialized_start=2105
  _globals['_GETREWARDSFORSNAPSHOTREQUEST']._serialized_end=2163
  _globals['_GETREWARDSFORSNAPSHOTRESPONSE']._serialized_start=2165
  _globals['_GETREWARDSFORSNAPSHOTRESPONSE']._serialized_end=2261
  _globals['_GETATTRIBUTABLEREWARDSFORSNAPSHOTREQUEST']._serialized_start=2263
  _globals['_GETATTRIBUTABLEREWARDSFORSNAPSHOTREQUEST']._serialized_end=2333
  _globals['_GETATTRIBUTABLEREWARDSFORSNAPSHOTRESPONSE']._serialized_start=2335
  _globals['_GETATTRIBUTABLEREWARDSFORSNAPSHOTRESPONSE']._serialized_end=2455
  _globals['_GETATTRIBUTABLEREWARDSFORDISTRIBUTIONROOTREQUEST']._serialized_start=2457
  _globals['_GETATTRIBUTABLEREWARDSFORDISTRIBUTIONROOTREQUEST']._serialized_end=2552
  _globals['_GETATTRIBUTABLEREWARDSFORDISTRIBUTIONROOTRESPONSE']._serialized_start=2555
  _globals['_GETATTRIBUTABLEREWARDSFORDISTRIBUTIONROOTRESPONSE']._serialized_end=2683
  _globals['_GENERATECLAIMPROOFREQUEST']._serialized_start=2685
  _globals['_GENERATECLAIMPROOFREQUEST']._serialized_end=2775
  _globals['_GENERATECLAIMPROOFRESPONSE']._serialized_start=2777
  _globals['_GENERATECLAIMPROOFRESPONSE']._serialized_end=2865
  _globals['_GETAVAILABLEREWARDSREQUEST']._serialized_start=2867
  _globals['_GETAVAILABLEREWARDSREQUEST']._serialized_end=2934
  _globals['_GETAVAILABLEREWARDSRESPONSE']._serialized_start=2936
  _globals['_GETAVAILABLEREWARDSRESPONSE']._serialized_end=3030
  _globals['_GETTOTALCLAIMEDREWARDSREQUEST']._serialized_start=3032
  _globals['_GETTOTALCLAIMEDREWARDSREQUEST']._serialized_end=3159
  _globals['_GETTOTALCLAIMEDREWARDSRESPONSE']._serialized_start=3161
  _globals['_GETTOTALCLAIMEDREWARDSRESPONSE']._serialized_end=3258
  _globals['_GETAVAILABLEREWARDSTOKENSREQUEST']._serialized_start=3260
  _globals['_GETAVAILABLEREWARDSTOKENSREQUEST']._serialized_end=3333
  _globals['_GETAVAILABLEREWARDSTOKENSRESPONSE']._serialized_start=3335
  _globals['_GETAVAILABLEREWARDSTOKENSRESPONSE']._serialized_end=3394
  _globals['_SUMMARIZEDEARNERREWARD']._serialized_start=3396
  _globals['_SUMMARIZEDEARNERREWARD']._serialized_end=3522
  _globals['_GETSUMMARIZEDREWARDSFOREARNERREQUEST']._serialized_start=3525
  _globals['_GETSUMMARIZEDREWARDSFOREARNERREQUEST']._serialized_end=3659
  _globals['_GETSUMMARIZEDREWARDSFOREARNERRESPONSE']._serialized_start=3661
  _globals['_GETSUMMARIZEDREWARDSFOREARNERRESPONSE']._serialized_end=3781
  _globals['_CLAIMEDREWARD']._serialized_start=3784
  _globals['_CLAIMEDREWARD']._serialized_end=3940
  _globals['_GETCLAIMEDREWARDSBYBLOCKREQUEST']._serialized_start=3942
  _globals['_GETCLAIMEDREWARDSBYBLOCKREQUEST']._serialized_end=4010
  _globals['_GETCLAIMEDREWARDSBYBLOCKRESPONSE']._serialized_start=4012
  _globals['_GETCLAIMEDREWARDSBYBLOCKRESPONSE']._serialized_end=4118
  _globals['_REWARDS']._serialized_start=4121
  _globals['_REWARDS']._serialized_end=7081
# @@protoc_insertion_point(module_scope)
