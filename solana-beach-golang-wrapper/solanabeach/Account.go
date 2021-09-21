package solana_beach

import (
	"encoding/json"
	"fmt"
)

type Account []struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"pubkey"`
	Balance int `json:"balance"`
}

func FetchAccounts(limit string) Account {
	var result Account

	if err := json.Unmarshal(getResponse(`accounts`, nil), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(prettyPrint(result))

	return result
}
