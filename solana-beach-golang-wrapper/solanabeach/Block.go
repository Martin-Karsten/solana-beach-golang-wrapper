package solana_beach

func GetLatestBlocks() string {
	return getResponse(`latest-blocks`)
}
