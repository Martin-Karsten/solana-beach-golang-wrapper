package main

import (
	solana_beach "github.com/Martin-Karsten/solana-beach-golang-wrapper/solanabeach"
)

func main() {
	solana_beach.FetchLatestTransactions(solana_beach.LatestTransactionsParams{Limit: "1"})
}
