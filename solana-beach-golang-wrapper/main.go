package main

import (
	"encoding/json"

	solana_beach "github.com/Martin-Karsten/solana-beach-golang-wrapper/solanabeach"
)

func main() {
	solana_beach.FetchLatestBlocks(solana_beach.LatestBlocksParams{Limit: "2"})
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
