// Code generated by type.sh. DO NOT EDIT.
// source: ../source/base/base.json ../source/runtime/runtime.json ../source/treasury/source.json

package source

var BaseType = `{"\u0026[u8]":"Bytes","(AccountId, Balance)":{"type":"struct","type_mapping":[["account","AccountId"],["balance","Balance"]]},"(BalanceOf\u003cT, I\u003e, BidKind\u003cAccountId, BalanceOf\u003cT, I\u003e\u003e)":{"type":"struct","type_mapping":[["balance","Balance"],["bidkind","BidKind"]]},"\u003cAuthorityId as RuntimeAppPublic\u003e::Signature":"Signature","AbridgedCandidateReceipt":{"type":"struct","type_mapping":[["parachainIndex","ParaId"],["relayParent","Hash"],["headData","HeadData"],["collator","CollatorId"],["signature","CollatorSignature"],["povBlockHash","Hash"],["commitments","CandidateCommitments"]]},"AccountData":{"type":"struct","type_mapping":[["free","Balance"],["reserved","Balance"],["miscFrozen","Balance"],["feeFrozen","Balance"]]},"AccountInfo\u003cIndex, AccountData\u003e":{"type":"struct","type_mapping":[["nonce","Index"],["refcount","RefCount"],["data","AccountData"]]},"AccountStatus":{"type":"struct","type_mapping":[["validity","AccountValidity"],["freeBalance","Balance"],["lockedBalance","Balance"],["signature","Vec\u003cu8\u003e"],["vat","Permill"]]},"AccountValidity":{"type":"enum","value_list":["Invalid","Initiated","Pending","ValidLow","ValidHigh","Completed"]},"AccountVote":{"type":"enum","type_mapping":[["Standard","AccountVoteStandard"],["Split","AccountVoteSplit"]]},"AccountVoteSplit":{"type":"struct","type_mapping":[["aye","Balance"],["nay","Balance"]]},"AccountVoteStandard":{"type":"struct","type_mapping":[["vote","Vote"],["balance","Balance"]]},"ActiveEraInfo":{"type":"struct","type_mapping":[["index","EraIndex"],["start","MomentOf"]]},"ActiveRecovery":{"type":"struct","type_mapping":[["created","BlockNumber"],["deposit","Balance"],["friends","Vec\u003cAccountId\u003e"]]},"AllowedSlots":{"type":"enum","value_list":["PrimarySlots","PrimaryAndSecondaryPlainSlots","PrimaryAndSecondaryVRFSlots"]},"Announcement":"ProxyAnnouncement","Approvals":"[bool; 4]","AssetOptions":{"type":"struct","type_mapping":[["initalIssuance","Compact\u003cBalance\u003e"],["permissions","PermissionLatest"]]},"AttestedCandidate":{"type":"struct","type_mapping":[["candidate","AbridgedCandidateReceipt"],["validityVotes","Vec\u003cValidityAttestation\u003e"],["validatorIndices","BitVec"]]},"AuctionIndex":"U32","AuthIndex":"U32","AuthorityIndex":"u64","AuthorityList":"Vec\u003cNextAuthority\u003e","AuthoritySignature":"Signature","AuthorityWeight":"u64","BabeAuthorityWeight":"u64","BabeBlockWeight":"u32","BabeEquivocationProof":{"type":"struct","type_mapping":[["offender","AuthorityId"],["slotNumber","SlotNumber"],["firstHeader","Header"],["secondHeader","Header"]]},"BabeWeight":"u64","BalanceLock":{"type":"struct","type_mapping":[["id","LockIdentifier"],["amount","Balance"],["until","u32"],["reasons","WithdrawReasons"]]},"BalanceLock\u003cBalance, BlockNumber\u003e":{"type":"struct","type_mapping":[["id","LockIdentifier"],["amount","Balance"],["reasons","Reasons"]]},"BalanceLock\u003cBalance\u003e":{"type":"struct","type_mapping":[["id","LockIdentifier"],["amount","Balance"],["reasons","Reasons"]]},"BalanceOf":"Balance","BalanceStatus":{"type":"enum","value_list":["Free","Reserved"]},"BalanceUpload":{"type":"struct","type_mapping":[["col1","AccountId"],["col2","u64"]]},"Bid":{"type":"struct","type_mapping":[["who","AccountId"],["kind","BidKind"],["value","Balance"]]},"BidKind":{"type":"enum","type_mapping":[["Deposit","Balance"],["Vouch","BidKindVouch"]]},"BidKindVouch":{"type":"struct","type_mapping":[["account","AccountId"],["amount","Balance"]]},"Bidder":{"type":"enum","type_mapping":[["New","NewBidder"],["Existing","ParaId"]]},"BlockAttestations":{"type":"struct","type_mapping":[["receipt","CandidateReceipt"],["valid","Vec\u003cAccountId\u003e"],["invalid","Vec\u003cAccountId\u003e"]]},"Bounty":{"type":"struct","type_mapping":[["proposer","AccountId"],["value","Balance"],["fee","Balance"],["curatorDeposit","Balance"],["bond","Balance"],["status","BountyStatus"]]},"BountyIndex":"u32","BountyStatus":{"type":"enum","type_mapping":[["Proposed","Null"],["Approved","Null"],["Funded","Null"],["CuratorProposed","BountyStatusCuratorProposed"],["Active","BountyStatusActive"],["PendingPayout","BountyStatusPendingPayout"]]},"BountyStatusActive":{"type":"struct","type_mapping":[["curator","AccountId"],["updateDue","BlockNumber"]]},"BountyStatusCuratorProposed":{"type":"struct","type_mapping":[["curator","AccountId"]]},"BountyStatusPendingPayout":{"type":"struct","type_mapping":[["curator","AccountId"],["beneficiary","AccountId"],["unlockAt","BlockNumber"]]},"Box\u003c\u003cT as Trait\u003cI\u003e\u003e::Proposal\u003e":"BoxProposal","Box\u003cCall\u003e":"BoxProposal","CallHash":"H256","CallHashOf":"H256","CandidateCommitments":{"type":"struct","type_mapping":[["fees","Balance"],["upwardMessages","Vec\u003cUpwardMessage\u003e"],["erasureRoot","Hash"],["newValidationCode","Option\u003cValidationCode\u003e"],["processedDownwardMessages","u32"]]},"CandidateReceipt":{"type":"struct","type_mapping":[["parachainIndex","ParaId"],["collator","AccountId"],["signature","CollatorSignature"],["headData","HeadData"],["balanceUploads","Vec\u003cBalanceUpload\u003e"],["egressQueueRoots","Vec\u003cEgressQueueRoot\u003e"],["fees","u64"],["blockDataHash","Hash"]]},"ChangesTrieConfiguration":{"type":"struct","type_mapping":[["digestInterval","u32"],["digestLevels","u32"]]},"CodeHash":"Hash","CollatorId":"H256","CollatorSignature":"Signature","CompactAssignments":{"type":"struct","type_mapping":[["votes1","Vec\u003c(NominatorIndex, ValidatorIndex)\u003e"],["votes2","Vec\u003c(NominatorIndex, [CompactScore; 1], ValidatorIndex)\u003e"],["votes3","Vec\u003c(NominatorIndex, [CompactScore; 2], ValidatorIndex)\u003e"],["votes4","Vec\u003c(NominatorIndex, [CompactScore; 3], ValidatorIndex)\u003e"],["votes5","Vec\u003c(NominatorIndex, [CompactScore; 4], ValidatorIndex)\u003e"],["votes6","Vec\u003c(NominatorIndex, [CompactScore; 5], ValidatorIndex)\u003e"],["votes7","Vec\u003c(NominatorIndex, [CompactScore; 6], ValidatorIndex)\u003e"],["votes8","Vec\u003c(NominatorIndex, [CompactScore; 7], ValidatorIndex)\u003e"],["votes9","Vec\u003c(NominatorIndex, [CompactScore; 8], ValidatorIndex)\u003e"],["votes10","Vec\u003c(NominatorIndex, [CompactScore; 9], ValidatorIndex)\u003e"],["votes11","Vec\u003c(NominatorIndex, [CompactScore; 10], ValidatorIndex)\u003e"],["votes12","Vec\u003c(NominatorIndex, [CompactScore; 11], ValidatorIndex)\u003e"],["votes13","Vec\u003c(NominatorIndex, [CompactScore; 12], ValidatorIndex)\u003e"],["votes14","Vec\u003c(NominatorIndex, [CompactScore; 13], ValidatorIndex)\u003e"],["votes15","Vec\u003c(NominatorIndex, [CompactScore; 14], ValidatorIndex)\u003e"],["votes16","Vec\u003c(NominatorIndex, [CompactScore; 15], ValidatorIndex)\u003e"]]},"CompactAssignmentsLatest":{"type":"struct","type_mapping":[["votes1","Vec\u003c(NominatorIndexCompact, ValidatorIndexCompact)\u003e"],["votes2","Vec\u003c(NominatorIndexCompact, CompactScoreCompact, ValidatorIndexCompact)\u003e"],["votes3","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 2], ValidatorIndexCompact)\u003e"],["votes4","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 3], ValidatorIndexCompact)\u003e"],["votes5","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 4], ValidatorIndexCompact)\u003e"],["votes6","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 5], ValidatorIndexCompact)\u003e"],["votes7","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 6], ValidatorIndexCompact)\u003e"],["votes8","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 7], ValidatorIndexCompact)\u003e"],["votes9","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 8], ValidatorIndexCompact)\u003e"],["votes10","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 9], ValidatorIndexCompact)\u003e"],["votes11","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 10], ValidatorIndexCompact)\u003e"],["votes12","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 11], ValidatorIndexCompact)\u003e"],["votes13","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 12], ValidatorIndexCompact)\u003e"],["votes14","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 13], ValidatorIndexCompact)\u003e"],["votes15","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 14], ValidatorIndexCompact)\u003e"],["votes16","Vec\u003c(NominatorIndexCompact, [CompactScoreCompact; 15], ValidatorIndexCompact)\u003e"]]},"CompactAssignmentsVote":{"type":"struct","type_mapping":[["accountId1","AccountId"],["scores","Vec\u003cCompactScore\u003e"],["accountId2","AccountId"]]},"CompactScore":{"type":"struct","type_mapping":[["validatorIndex","ValidatorIndex"],["offchainAccuracy","OffchainAccuracy"]]},"CompactScoreCompact":{"type":"struct","type_mapping":[["validatorIndex","ValidatorIndexCompact"],["offchainAccuracy","OffchainAccuracyCompact"]]},"ConsensusEngineId":"[u8; 4]","ContractCallRequest":{"type":"struct","type_mapping":[["origin","AccountId"],["dest","AccountId"],["value","Balance"],["gasLimit","u64"],["inputData","Bytes"]]},"ContractExecResult":{"type":"enum","type_mapping":[["Success","ContractExecResultSuccess"],["Error","Null"]]},"ContractExecResultSuccess":{"type":"struct","type_mapping":[["flags","u32"],["data","Vec\u003cu8\u003e"],["gasConsumed","u64"]]},"ContractExecResultSuccessTo255":{"type":"struct","type_mapping":[["status","u8"],["data","Raw"]]},"ContractExecResultTo255":{"type":"enum","type_mapping":[["Success","ContractExecResultSuccessTo255"],["Error","Null"]]},"ContractStorageKey":"[u8; 32]","Conviction":{"type":"enum","value_list":["None","Locked1x","Locked2x","Locked3x","Locked4x","Locked5x","Locked6x"]},"CreatedBlock":{"type":"struct","type_mapping":[["hash","BlockHash"],["aux","ImportedAux"]]},"DefunctVoter":{"type":"struct","type_mapping":[["who","AccountId"],["voteCount","Compact\u003cu32\u003e"],["candidateCount","Compact\u003cu32\u003e"]]},"Delegations":{"type":"struct","type_mapping":[["votes","Balance"],["capital","Balance"]]},"Digest":{"type":"struct","type_mapping":[["logs","Vec\u003cDigestItem\u003cHash\u003e\u003e"]]},"DigestItem":{"type":"enum","type_mapping":[["Other","Vec\u003cu8\u003e"],["AuthoritiesChange","Vec\u003cAuthorityId\u003e"],["ChangesTrieRoot","Hash"],["SealV0","SealV0"],["Consensus","Consensus"],["Seal","Seal"],["PreRuntime","PreRuntime"]]},"DigestOf":"Digest","DispatchClass":{"type":"enum","value_list":["Normal","Operational","Mandatory"]},"DispatchError":{"type":"enum","type_mapping":[["Other","Null"],["CannotLookup","Null"],["BadOrigin","Null"],["Module","DispatchErrorModule"]]},"DispatchErrorModule":{"type":"struct","type_mapping":[["index","u8"],["error","u8"]]},"DispatchInfo":{"type":"struct","type_mapping":[["weight","Weight"],["class","DispatchClass"],["paysFee","bool"]]},"DispatchResult":{"type":"enum","type_mapping":[["Error","DispatchError"],["Ok","Null"]]},"DoNotConstruct":"Null","DoubleVoteReport":{"type":"struct","type_mapping":[["identity","ValidatorId"],["first","(Statement, ValidatorSignature)"],["second","(Statement, ValidatorSignature)"],["proof","MembershipProof"],["signingContext","SigningContext"]]},"DownwardMessage":{"type":"enum","type_mapping":[["TransferInto","(AccountId, Balance, Remark)"],["Opaque","Vec\u003cu8\u003e"]]},"EgressQueueRoot":{"type":"struct","type_mapping":[["col1","ParaId"],["col2","Hash"]]},"ElectionCompute":{"type":"enum","value_list":["OnChain","Signed","Authority"]},"ElectionResult":{"type":"struct","type_mapping":[["compute","ElectionCompute"],["slotStake","Balance"],["electedStashes","Vec\u003cAccountId\u003e"],["exposures","Vec\u003c(AccountId, Exposure)\u003e"]]},"ElectionScore":"[u128; 3]","ElectionSize":{"type":"struct","type_mapping":[["validators","ValidatorIndex"],["nominators","NominatorIndex"]]},"ElectionStatus":{"type":"enum","type_mapping":[["Close","Null"],["Open","BlockNumber"]]},"EpochAuthorship":{"type":"struct","type_mapping":[["primary","Vec\u003cu64\u003e"],["secondary","Vec\u003cu64\u003e"],["secondary_vrf","Vec\u003cu64\u003e"]]},"Equivocation":"GrandpaEquivocation","EquivocationProof\u003cHash, BlockNumber\u003e":"GrandpaEquivocationProof","EquivocationProof\u003cHeader\u003e":"BabeEquivocationProof","EraRewardPoints":{"type":"struct","type_mapping":[["total","RewardPoint"],["individual","Vec\u003c(AccountId, RewardPoint)\u003e"]]},"EthereumAccountId":"EthereumAddress","EthereumLookupSource":"EthereumAddress","EventIndex":"u32","ExtrinsicMetadata":{"type":"struct","type_mapping":[["version","u8"],["signedExtensions","Vec\u003cBytes\u003e"]]},"ExtrinsicsWeight":{"type":"struct","type_mapping":[["normal","Weight"],["operational","Weight"]]},"Forcing":{"type":"enum","value_list":["NotForcing","ForceNew","ForceNone"]},"FullIdentification":{"type":"struct","type_mapping":[["total","Compact\u003cBalance\u003e"],["own","Compact\u003cBalance\u003e"],["others","Vec\u003cIndividualExposure\u003e"]]},"GlobalValidationSchedule":{"type":"struct","type_mapping":[["maxCodeSize","u32"],["maxHeadDataSize","u32"],["blockNumber","BlockNumber"]]},"GrandpaEquivocation":{"type":"enum","type_mapping":[["Prevote","GrandpaEquivocationValue"],["Precommit","GrandpaEquivocationValue"]]},"GrandpaEquivocationProof":{"type":"struct","type_mapping":[["setId","SetId"],["equivocation","GrandpaEquivocation"]]},"GrandpaEquivocationValue":{"type":"struct","type_mapping":[["roundNumber","u64"],["identity","AuthorityId"],["first","(GrandpaPrevote, AuthoritySignature)"],["second","(GrandpaPrevote, AuthoritySignature)"]]},"GrandpaPrevote":{"type":"struct","type_mapping":[["targetHash","Hash"],["targetNumber","BlockNumber"]]},"HeadData":"Bytes","Header":{"type":"struct","type_mapping":[["parent_hash","H256"],["number","Compact\u003cBlockNumber\u003e"],["state_root","H256"],["extrinsics_root","H256"],["digest","Digest"]]},"Heartbeat":{"type":"struct","type_mapping":[["blockNumber","BlockNumber"],["networkState","OpaqueNetworkState"],["sessionIndex","SessionIndex"],["authorityIndex","AuthIndex"]]},"IdentificationTuple":{"type":"struct","type_mapping":[["validatorId","ValidatorId"],["exposure","FullIdentification"]]},"IdentityFields":{"bit_length":64,"type":"set","value_list":["Display","Legal","Web","Riot","Email","PgpFingerprint","Image","Twitter"]},"IdentityInfo":{"type":"struct","type_mapping":[["additional","Vec\u003cIdentityInfoAdditional\u003e"],["display","Data"],["legal","Data"],["web","Data"],["riot","Data"],["email","Data"],["pgpFingerprint","Option\u003cH160\u003e"],["image","Data"],["twitter","Data"]]},"IdentityInfoAdditional":{"type":"struct","type_mapping":[["field","Data"],["value","Data"]]},"ImportedAux":{"type":"struct","type_mapping":[["headerOnly","bool"],["clearJustificationRequests","bool"],["needsJustification","bool"],["badJustification","bool"],["needsFinalityProof","bool"],["isNewBest","bool"]]},"IncludedBlocks":{"type":"struct","type_mapping":[["actualNumber","BlockNumber"],["session","SessionIndex"],["randomSeed","H256"],["activeParachains","Vec\u003cParaId\u003e"],["paraBlocks","Vec\u003cHash\u003e"]]},"IncomingParachain":{"type":"enum","type_mapping":[["Unset","NewBidder"],["Fixed","IncomingParachainFixed"],["Deploy","IncomingParachainDeploy"]]},"IncomingParachainDeploy":{"type":"struct","type_mapping":[["code","ValidationCode"],["initialHeadData","HeadData"]]},"IncomingParachainFixed":{"type":"struct","type_mapping":[["codeHash","Hash"],["codeSize","u32"],["initialHeadData","Bytes"]]},"Index":"U32","IndividualExposure":{"type":"struct","type_mapping":[["who","AccountId"],["value","Compact\u003cBalance\u003e"]]},"Judgement":{"type":"enum","type_mapping":[["Unknown","Null"],["FeePaid","Balance"],["Reasonable","Null"],["KnownGood","Null"],["OutOfDate","Null"],["LowQuality","Null"],["Erroneous","Null"]]},"KeyOwnerProof":{"type":"struct","type_mapping":[["session","SessionIndex"],["trieNodes","Vec\u003cBytes\u003e"],["validatorCount","ValidatorCount"]]},"KeyValue":{"type":"struct","type_mapping":[["key","Vec\u003cu8\u003e"],["value","Vec\u003cu8\u003e"]]},"Keys":{"type":"struct","type_mapping":[["grandpa","AccountId"],["babe","AccountId"],["im_online","AccountId"],["authority_discovery","AccountId"]]},"Kind":"[u8; 16]","LastRuntimeUpgradeInfo":{"type":"struct","type_mapping":[["specVersion","Compact\u003cu32\u003e"],["specName","Bytes"]]},"LeasePeriod":"BlockNumber","LeasePeriodOf":"BlockNumber","Linkage":{"type":"struct","type_mapping":[["previous","Option\u003cAccountId\u003e"],["next","Option\u003cAccountId\u003e"]]},"LocalValidationData":{"type":"struct","type_mapping":[["parentHead","HeadData"],["balance","Balance"],["codeUpgradeAllowed","Option\u003cBlockNumber\u003e"]]},"LookupSource":"AccountId","MaybeRandomness":"Option\u003cRandomness\u003e","MaybeVrf":"[u8; 32]","MemberCount":"u32","MembershipProof":{"type":"struct","type_mapping":[["session","sessionIndex"],["trieNodes","Vec\u003cVec\u003cu8\u003e\u003e"],["validatorCount","ValidatorCount"]]},"MetadataVersion":{"type":"enum","value_list":["MetadataV0Decoder","MetadataV1Decoder","MetadataV2Decoder","MetadataV3Decoder","MetadataV4Decoder","MetadataV5Decoder","MetadataV6Decoder","MetadataV7Decoder","MetadataV8Decoder","MetadataV9Decoder","MetadataV10Decoder","MetadataV11Decoder","MetadataV12Decoder"]},"ModuleId":"LockIdentifier","MomentOf":"Moment","MoreAttestations":"Null","Multiplier":"u64","Multisig":{"type":"struct","type_mapping":[["when","Timepoint"],["deposit","Balance"],["depositor","AccountId"],["approvals","Vec\u003cAccountId\u003e"]]},"NewBidder":{"type":"struct","type_mapping":[["who","AccountId"],["sub","SubId"]]},"NextAuthority":{"type":"struct","type_mapping":[["AuthorityId","AuthorityId"],["weight","AuthorityWeight"]]},"NextConfigDescriptor":{"type":"enum","type_mapping":[["V0","Null"],["V1","NextConfigDescriptorV1"]]},"NextConfigDescriptorV1":{"type":"struct","type_mapping":[["c","(u64, u64)"],["allowedSlots","AllowedSlots"]]},"Nominations":{"type":"struct","type_mapping":[["targets","Vec\u003cAccountId\u003e"],["submittedIn","EraIndex"],["suppressed","bool"]]},"NominatorIndex":"u32","NominatorIndexCompact":"Compact\u003cNominatorIndex\u003e","OffchainAccuracy":"u16","OffchainAccuracyCompact":"Compact\u003cOffchainAccuracy\u003e","OffenceDetails\u003cAccountId, IdentificationTuple\u003e":{"type":"struct","type_mapping":[["offender","Offender"],["reporters","Vec\u003cReporter\u003e"]]},"Offender":{"type":"struct","type_mapping":[["col1","ValidatorId"],["col2","Exposure"]]},"OpaqueMultiaddr":"Bytes","OpaqueNetworkState":{"type":"struct","type_mapping":[["peerId","OpaquePeerId"],["externalAddresses","Vec\u003cOpaqueMultiaddr\u003e"]]},"OpaquePeerId":"Bytes","OpaqueTimeSlot":"Bytes","OpenTip":{"type":"struct","type_mapping":[["reason","Hash"],["who","AccountId"],["finder","AccountId"],["deposit","Balance"],["closes","Option\u003cBlockNumber\u003e"],["tips","Vec\u003cOpenTipTip\u003e"],["findersFee","bool"]]},"OpenTip\u003cAccountId, BalanceOf, BlockNumber, Hash\u003e":{"type":"struct","type_mapping":[["reason","Hash"],["who","AccountId"],["finder","Option\u003cOpenTipFinder\u003e"],["closes","Option\u003cBlockNumber\u003e"],["tips","Vec\u003cOpenTipTip\u003e"]]},"OpenTipTip":"(AccountId, Balance)","Origin":"Null","Owner":{"type":"enum","type_mapping":[["None","Null"],["Address","AccountId"]]},"PalletVersion":{"type":"struct","type_mapping":[["major","u16"],["minor","u8"],["patch","u8"]]},"ParaInfo":{"type":"struct","type_mapping":[["scheduling","ParaScheduling"]]},"ParaPastCodeMeta":{"type":"struct","type_mapping":[["upgrade_times","Vec\u003cBlockNumber\u003e"],["last_pruned","Option\u003cBlockNumber\u003e"]]},"ParaScheduling":{"type":"enum","value_list":["Always","Dynamic"]},"ParachainDispatchOrigin":{"type":"enum","value_list":["Signed","Parachain","Root"]},"Pays":{"type":"enum","value_list":["Yes","No"]},"PerU16":"u16","Perbill":"u32","Percent":"u8","Period":"(BlockNumber, u32)","Permill":"u32","PermissionLatest":"PermissionsV1","PermissionVersions":{"type":"enum","type_mapping":[["V1","PermissionsV1"]]},"PermissionsV1":{"type":"struct","type_mapping":[["update","Owner"],["mint","Owner"],["burn","Owner"]]},"PhantomData":"Null","Phase":{"type":"enum","type_mapping":[["ApplyExtrinsic","u32"],["Finalization","Null"],["Initialization","Null"]]},"PhragmenScore":"[u128; 3]","Points":"u32","PrefabWasmModule":{"type":"struct","type_mapping":[["scheduleVersion","Compact\u003cu32\u003e"],["initial","Compact\u003cu32\u003e"],["maximum","Compact\u003cu32\u003e"],["_reserved","PrefabWasmModuleReserved"],["code","Bytes"]]},"PrefabWasmModuleReserved":"Option\u003cNull\u003e","PreimageStatus":{"type":"enum","type_mapping":[["Missing","BlockNumber"],["Available","PreimageStatusAvailable"]]},"PreimageStatusAvailable":{"type":"struct","type_mapping":[["data","Bytes"],["provider","AccountId"],["deposit","Balance"],["since","BlockNumber"],["expiry","Option\u003cBlockNumber\u003e"]]},"PriorLock":{"type":"struct","type_mapping":[["blockNumber","BlockNumber"],["balance","Balance"]]},"Priority":"u8","Proposal":"BoxProposal","ProposalCategory":{"type":"enum","value_list":["Signaling"]},"ProposalContents":"Bytes","ProposalIndex":"u32","ProposalStage":{"type":"enum","value_list":["PreVoting","Voting","Completed"]},"ProposalTitle":"Bytes","ProxyAnnouncement":{"type":"struct","type_mapping":[["real","AccountId"],["callHash","Hash"],["height","BlockNumber"]]},"ProxyDefinition":{"type":"struct","type_mapping":[["delegate","AccountId"],["proxyType","ProxyType"],["delay","BlockNumber"]]},"ProxyState":{"type":"struct","type_mapping":[["Open","AccountId"],["Active","AccountId"]]},"ProxyType":{"type":"enum","value_list":["Any","NonTransfer","Governance","Staking"]},"Randomness":"Hash","RawBabePreDigestCompat":{"type":"enum","type_mapping":[["Zero","u32"],["One","u32"],["Two","u32"],["Three","u32"]]},"Reasons":{"type":"enum","value_list":["Fee","Misc","All"]},"RecoveryConfig":{"type":"struct","type_mapping":[["delayPeriod","BlockNumber"],["deposit","Balance"],["friends","Vec\u003cAccountId\u003e"],["threshold","u16"]]},"RefCount":"u8","RefCount1":"u32","ReferendumInfo":{"type":"struct","type_mapping":[["end","BlockNumber"],["proposal","Proposal"],["threshold","VoteThreshold"],["delay","BlockNumber"]]},"ReferendumInfo\u003cBlockNumber, Hash, Balance\u003e":{"type":"enum","type_mapping":[["Ongoing","ReferendumStatus"],["Finished","ReferendumInfoFinished"]]},"ReferendumInfo\u003cBlockNumber, Hash, BalanceOf\u003e":{"type":"enum","type_mapping":[["Ongoing","ReferendumStatus"],["Finished","ReferendumInfoFinished"]]},"ReferendumInfo\u003cBlockNumber, Hash\u003e":{"type":"struct","type_mapping":[["end","BlockNumber"],["proposalHash","Hash"],["threshold","VoteThreshold"],["delay","BlockNumber"]]},"ReferendumInfo\u003cBlockNumber, Proposal\u003e":{"type":"struct","type_mapping":[["end","BlockNumber"],["proposal","Proposal"],["threshold","VoteThreshold"],["delay","BlockNumber"]]},"ReferendumInfoFinished":{"type":"struct","type_mapping":[["approved","bool"],["end","BlockNumber"]]},"ReferendumStatus":{"type":"struct","type_mapping":[["end","BlockNumber"],["proposalHash","Hash"],["threshold","VoteThreshold"],["delay","BlockNumber"],["tally","Tally"]]},"RegistrarIndex":"u32","RegistrarInfo":{"type":"struct","type_mapping":[["account","AccountId"],["fee","Balance"],["fields","IdentityFields"]]},"Registration":{"type":"struct","type_mapping":[["judgements","Vec\u003cRegistrationJudgement\u003e"],["deposit","Balance"],["info","IdentityInfo"]]},"RegistrationJudgement":{"type":"struct","type_mapping":[["col1","RegistrarIndex"],["col2","Judgement"]]},"RelayChainBlockNumber":"BlockNumber","Releases":"ReleasesBalances","ReleasesBalances":{"type":"enum","value_list":["V1_0_0","V2_0_0"]},"Remark":"[u8; 32]","Renouncing":{"type":"enum","type_mapping":[["Member","Null"],["RunnerUp","Null"],["Candidate","Compact\u003cu32\u003e"]]},"ReportIdOf":"Hash","Reporter":"AccountId","RewardDestination":{"type":"enum","type_mapping":[["Staked","Null"],["Stash","Null"],["Controller","Null"],["Account","AccountId"]]},"RewardPoint":"u32","RoundNumber":"U64","RpcMethods":{"type":"struct","type_mapping":[["version","u32"],["methods","Vec\u003cText\u003e"]]},"RuntimeDbWeight":{"type":"struct","type_mapping":[["read","Weight"],["write","Weight"]]},"SchedulePeriod":"(BlockNumber, u32)","SchedulePriority":"u8","Scheduled":{"type":"struct","type_mapping":[["maybeId","Option\u003cBytes\u003e"],["priority","SchedulePriority"],["call","Call"],["maybePeriodic","Option\u003cSchedulePeriod\u003e"]]},"Scheduling":{"type":"enum","value_list":["Always","Dynamic"]},"SessionIndex":"U32","SessionKeysPolkadot":{"type":"struct","type_mapping":[["grandpa","AccountId"],["babe","AccountId"],["im_online","AccountId"],["authority_discovery","AccountId"],["parachains","AccountId"]]},"SetId":"U64","SigningContext":{"type":"struct","type_mapping":[["sessionIndex","sessionIndex"],["parentHash","Hash"]]},"SlashingSpans":{"type":"struct","type_mapping":[["spanIndex","SpanIndex"],["lastStart","EraIndex"],["lastNonzeroSlash","EraIndex"],["prior","Vec\u003cEraIndex\u003e"]]},"SlotRange":{"type":"enum","value_list":["ZeroZero","ZeroOne","ZeroTwo","ZeroThree","OneOne","OneTwo","OneThree","TwoTwo","TwoThree","ThreeThree"]},"SpanIndex":"u32","SpanRecord":{"type":"struct","type_mapping":[["slashed","Balance"],["paidOut","Balance"]]},"StakingLedger\u003cAccountId, BalanceOf\u003e":{"type":"struct","type_mapping":[["stash","AccountId"],["total","Compact\u003cBalance\u003e"],["active","Compact\u003cBalance\u003e"],["unlocking","Vec\u003cUnlockChunk\u003cBalance\u003e\u003e"],["lastReward","Option\u003cEraIndex\u003e"]]},"StatementKind":{"type":"enum","value_list":["Regular","Saft"]},"Status":"BalanceStatus","StorageFunctionType":{"type":"enum","value_list":["PlainType","MapType","DoubleMapType"]},"StorageKey":"Bytes","StorageModify":{"type":"enum","value_list":["Optional","Default"]},"StoredPendingChange":{"type":"struct","type_mapping":[["scheduled_at","u32"],["forced","u32"]]},"StoredState":{"type":"enum","value_list":["Live","PendingPause","Paused","PendingResume"]},"StrikeCount":"u32","SubId":"u32","Tally":{"type":"struct","type_mapping":[["ayes","Balance"],["nays","Balance"],["turnout","Balance"]]},"TaskAddress":{"type":"struct","type_mapping":[["blockNumber","BlockNumber"],["index","u32"]]},"Text":"Bytes","Timepoint":{"type":"struct","type_mapping":[["height","BlockNumber"],["index","u32"]]},"TransactionPriority":"u64","TreasuryProposal":{"type":"struct","type_mapping":[["proposer","AccountId"],["value","Balance"],["beneficiary","AccountId"],["bond","Balance"]]},"UnappliedSlash":{"type":"struct","type_mapping":[["validator","AccountId"],["own","Balance"],["others","Vec\u003cUnappliedSlashOther\u003e"],["reporters","Vec\u003cAccountId\u003e"],["payout","Balance"]]},"UnappliedSlash\u003cAccountId, BalanceOf\u003e":{"type":"struct","type_mapping":[["validator","AccountId"],["own","Balance"],["others","Vec\u003cUnappliedSlashOther\u003e"],["reporters","Vec\u003cAccountId\u003e"],["payout","Balance"]]},"UnappliedSlashOther":{"type":"struct","type_mapping":[["account","AccountId"],["amount","Balance"]]},"UnlockChunk":{"type":"struct","type_mapping":[["value","Compact\u003cBalance\u003e"],["era","Compact\u003cEraIndex\u003e"]]},"UpwardMessage":{"type":"struct","type_mapping":[["origin","ParachainDispatchOrigin"],["data","Vec\u003cu8\u003e"]]},"ValidationCode":"Bytes","ValidationFunctionParams":{"type":"struct","type_mapping":[["maxCodeSize","u32"],["relayChainHeight","RelayChainBlockNumber"],["codeUpgradeAllowed","Option\u003cRelayChainBlockNumber\u003e"]]},"ValidatorCount":"u32","ValidatorId":"AccountId","ValidatorIndex":"u16","ValidatorIndexCompact":"Compact\u003cValidatorIndex\u003e","ValidatorPrefs":{"type":"struct","type_mapping":[["Commission","Compact\u003cBalance\u003e"]]},"ValidatorSignature":"Signature","ValidityAttestation":{"type":"enum","type_mapping":[["Never","Null"],["Implicit","CollatorSignature"],["Explicit","CollatorSignature"]]},"VestingInfo":{"type":"struct","type_mapping":[["locked","Balance"],["perBlock","Balance"],["startingBlock","BlockNumber"]]},"VoteStage":{"type":"enum","value_list":["PreVoting","Commit","Voting","Completed"]},"VoteThreshold":{"type":"enum","value_list":["SuperMajorityApprove","SuperMajorityAgainst","SimpleMajority"]},"Votes":{"type":"struct","type_mapping":[["index","ProposalIndex"],["threshold","MemberCount"],["ayes","Vec\u003cAccountId\u003e"],["nays","Vec\u003cAccountId\u003e"]]},"Voting":{"type":"enum","type_mapping":[["Direct","VotingDirect"],["Delegating","VotingDelegating"]]},"VotingDelegating":{"type":"struct","type_mapping":[["balance","Balance"],["target","AccountId"],["conviction","Conviction"],["delegations","Delegations"],["prior","PriorLock"]]},"VotingDirect":{"type":"struct","type_mapping":[["votes","Vec\u003cVotingDirectVote\u003e"],["delegations","Delegations"],["prior","PriorLock"]]},"VotingDirectVote":{"type":"struct","type_mapping":[["referendumIndex","ReferendumIndex"],["accountVote","AccountVote"]]},"VouchingStatus":{"type":"enum","value_list":["Vouching","Banned"]},"VrfData":"[u8; 32]","VrfOutput":"[u8; 32]","VrfProof":"[u8; 64]","Weight":"u32","WeightMultiplier":"u64","WinningData":"Vec\u003cWinningDataEntry\u003e","WinningDataEntry":{"type":"struct","type_mapping":[["col1","AccountId"],["col2","ParaIdOf"],["col3","BalanceOf"]]},"WithdrawReasons":{"bit_length":8,"type":"set","value_list":["TransactionPayment","Transfer","Reserve","Fee","Tip"]},"gas":"u64","schedule::Period\u003cBlockNumber\u003e":"(BlockNumber, u32)","schedule::Priority":"u8","schnorrkel::Randomness":"Hash","schnorrkel::RawVRFOutput":"[u8; 32]","slashing::SlashingSpans":"SlashingSpans","slashing::SpanIndex":"SpanIndex","slashing::SpanRecord\u003cBalanceOf\u003e":"SpanRecord","weights::ExtrinsicsWeight":"ExtrinsicsWeight"}`
