package main

import (
	solana_beach "github.com/Martin-Karsten/solana-beach-golang-wrapper/solanabeach"
)

func main() {
	solana_beach.FetchLatestBlocks(solana_beach.LatestBlocksParams{Limit: "1"})
}
