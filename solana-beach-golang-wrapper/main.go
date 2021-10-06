package main

import (
	"encoding/json"
	"fmt"

	solanabeach "github.com/Martin-Karsten/solana-beach-golang-wrapper/solanabeach"
)

func main() {
	t, err := solanabeach.FetchMarketChartData()
	if err != nil {
		println(err.Error())
	}

	fmt.Println(prettyPrint(t))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
