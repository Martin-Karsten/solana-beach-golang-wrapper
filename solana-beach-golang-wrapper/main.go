package main

import (
	"encoding/json"

	solanabeach "github.com/Martin-Karsten/solana-beach-golang-wrapper/solanabeach"
)

func main() {
	t, err := solanabeach.FetchAccountTokenTransfers(`ARcSeHfseT2u1Qz6g3czMndza5C1ofv7fBnbCewKX3RV`, solanabeach.AccountTokenTransfersParams{Limit: "2"})
	if err != nil {
		println(err.Error())
	}

	println(prettyPrint(t))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
