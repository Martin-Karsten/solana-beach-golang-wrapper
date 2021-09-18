package solana_beach

func GetLatestTransactions() string {
	return getResponse(`latest-transactions`)
}
