package main

import (
	"encoding/json"
	"fmt"

	solanabeach "github.com/Martin-Karsten/solana-beach-golang-wrapper/solanabeach"
)

func main() {
	t, err := solanabeach.FetchSupply()
	if err != nil {
		println(err.Error())
	}

	fmt.Println(t.NonCirculating)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
