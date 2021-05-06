package main

type status struct {
	ErdAppVersion                           string `json:"erd_app_version"`
	ErdAverageBlockTxCount                  string `json:"erd_average_block_tx_count"`
	ErdChainID                              string `json:"erd_chain_id"`
	ErdConnectedNodes                       int    `json:"erd_connected_nodes"`
	ErdConsensusGroupSize                   int    `json:"erd_consensus_group_size"`
	ErdConsensusProcessedProposedBlock      int    `json:"erd_consensus_processed_proposed_block"`
	ErdConsensusReceivedProposedBlock       int    `json:"erd_consensus_received_proposed_block"`
	ErdConsensusRoundState                  string `json:"erd_consensus_round_state"`
	ErdConsensusState                       string `json:"erd_consensus_state"`
	ErdCountAcceptedBlocks                  int    `json:"erd_count_accepted_blocks"`
	ErdCountConsensus                       int    `json:"erd_count_consensus"`
	ErdCountConsensusAcceptedBlocks         int    `json:"erd_count_consensus_accepted_blocks"`
	ErdCountLeader                          int    `json:"erd_count_leader"`
	ErdCPULoadPercent                       int    `json:"erd_cpu_load_percent"`
	ErdCrossCheckBlockHeight                string `json:"erd_cross_check_block_height"`
	ErdCurrentBlockHash                     string `json:"erd_current_block_hash"`
	ErdCurrentBlockSize                     int    `json:"erd_current_block_size"`
	ErdCurrentRound                         int    `json:"erd_current_round"`
	ErdCurrentRoundTimestamp                int    `json:"erd_current_round_timestamp"`
	ErdDenomination                         int    `json:"erd_denomination"`
	ErdDevRewards                           string `json:"erd_dev_rewards"`
	ErdEpochForEconomicsData                int    `json:"erd_epoch_for_economics_data"`
	ErdEpochNumber                          int    `json:"erd_epoch_number"`
	ErdForkChoiceCount                      int    `json:"erd_fork_choice_count"`
	ErdGasPerDataByte                       int    `json:"erd_gas_per_data_byte"`
	ErdGasPriceModifier                     string `json:"erd_gas_price_modifier"`
	ErdHighestFinalNonce                    int    `json:"erd_highest_final_nonce"`
	ErdInflation                            string `json:"erd_inflation"`
	ErdIsSyncing                            int    `json:"erd_is_syncing"`
	ErdLastBlockTxCount                     int    `json:"erd_last_block_tx_count"`
	ErdLatestTagSoftwareVersion             string `json:"erd_latest_tag_software_version"`
	ErdLeaderPercentage                     string `json:"erd_leader_percentage"`
	ErdLiveValidatorNodes                   int    `json:"erd_live_validator_nodes"`
	ErdMemHeapInuse                         int    `json:"erd_mem_heap_inuse"`
	ErdMemLoadPercent                       int    `json:"erd_mem_load_percent"`
	ErdMemStackInuse                        int    `json:"erd_mem_stack_inuse"`
	ErdMemTotal                             int64  `json:"erd_mem_total"`
	ErdMemUsedGolang                        int64  `json:"erd_mem_used_golang"`
	ErdMemUsedSys                           int64  `json:"erd_mem_used_sys"`
	ErdMetaConsensusGroupSize               int    `json:"erd_meta_consensus_group_size"`
	ErdMinGasLimit                          int    `json:"erd_min_gas_limit"`
	ErdMinGasPrice                          int    `json:"erd_min_gas_price"`
	ErdMinTransactionVersion                int    `json:"erd_min_transaction_version"`
	ErdMiniBlocksSize                       int    `json:"erd_mini_blocks_size"`
	ErdNetworkRecvBps                       int    `json:"erd_network_recv_bps"`
	ErdNetworkRecvBpsPeak                   int    `json:"erd_network_recv_bps_peak"`
	ErdNetworkRecvBytesInEpochPerHost       int64  `json:"erd_network_recv_bytes_in_epoch_per_host"`
	ErdNetworkRecvPercent                   int    `json:"erd_network_recv_percent"`
	ErdNetworkSentBps                       int    `json:"erd_network_sent_bps"`
	ErdNetworkSentBpsPeak                   int    `json:"erd_network_sent_bps_peak"`
	ErdNetworkSentBytesInEpochPerHost       int64  `json:"erd_network_sent_bytes_in_epoch_per_host"`
	ErdNetworkSentPercent                   int    `json:"erd_network_sent_percent"`
	ErdNodeDisplayName                      string `json:"erd_node_display_name"`
	ErdNodeType                             string `json:"erd_node_type"`
	ErdNonce                                int    `json:"erd_nonce"`
	ErdNonceAtEpochStart                    int    `json:"erd_nonce_at_epoch_start"`
	ErdNonceForTps                          int    `json:"erd_nonce_for_tps"`
	ErdNoncesPassedInCurrentEpoch           int    `json:"erd_nonces_passed_in_current_epoch"`
	ErdNumAccountsStateCheckpoints          int    `json:"erd_num_accounts_state_checkpoints"`
	ErdNumConnectedPeers                    int    `json:"erd_num_connected_peers"`
	ErdNumConnectedPeersClassification      string `json:"erd_num_connected_peers_classification"`
	ErdNumMetachainNodes                    int    `json:"erd_num_metachain_nodes"`
	ErdNumMiniBlocks                        int    `json:"erd_num_mini_blocks"`
	ErdNumNodesInShard                      int    `json:"erd_num_nodes_in_shard"`
	ErdNumPeerStateCheckpoints              int    `json:"erd_num_peer_state_checkpoints"`
	ErdNumShardHeadersFromPool              int    `json:"erd_num_shard_headers_from_pool"`
	ErdNumShardHeadersProcessed             int    `json:"erd_num_shard_headers_processed"`
	ErdNumShardsWithoutMeta                 int    `json:"erd_num_shards_without_meta"`
	ErdNumTransactionsProcessed             int    `json:"erd_num_transactions_processed"`
	ErdNumTransactionsProcessedTpsBenchmark int    `json:"erd_num_transactions_processed_tps_benchmark"`
	ErdNumTxBlock                           int    `json:"erd_num_tx_block"`
	ErdNumValidators                        int    `json:"erd_num_validators"`
	ErdPeakTps                              int    `json:"erd_peak_tps"`
	ErdPeerType                             string `json:"erd_peer_type"`
	ErdProbableHighestNonce                 int    `json:"erd_probable_highest_nonce"`
	ErdPublicKeyBlockSign                   string `json:"erd_public_key_block_sign"`
	ErdRewardsTopUpGradientPoint            string `json:"erd_rewards_top_up_gradient_point"`
	ErdRoundAtEpochStart                    int    `json:"erd_round_at_epoch_start"`
	ErdRoundDuration                        int    `json:"erd_round_duration"`
	ErdRoundTime                            int    `json:"erd_round_time"`
	ErdRoundsPassedInCurrentEpoch           int    `json:"erd_rounds_passed_in_current_epoch"`
	ErdRoundsPerEpoch                       int    `json:"erd_rounds_per_epoch"`
	ErdShardConsensusGroupSize              int    `json:"erd_shard_consensus_group_size"`
	ErdShardID                              int    `json:"erd_shard_id"`
	ErdStartTime                            int    `json:"erd_start_time"`
	ErdSynchronizedRound                    int    `json:"erd_synchronized_round"`
	ErdTopUpFactor                          string `json:"erd_top_up_factor"`
	ErdTotalFees                            string `json:"erd_total_fees"`
	ErdTotalSupply                          string `json:"erd_total_supply"`
	ErdTxPoolLoad                           int    `json:"erd_tx_pool_load"`
}

type statusReply struct {
	Data struct {
		Metrics status `json:"metrics"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}
