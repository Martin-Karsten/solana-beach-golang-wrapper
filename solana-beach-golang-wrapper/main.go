package main

import (
	"encoding/json"

	solanabeach "github.com/Martin-Karsten/solana-beach-golang-wrapper/solanabeach"
)

func main() {
	solanabeach.FetchBlock("9999")
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
