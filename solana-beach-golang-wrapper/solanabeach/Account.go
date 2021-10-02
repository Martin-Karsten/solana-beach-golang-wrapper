package solanabeach

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

func FetchAccounts(limit string) (Account, error) {
	var result Account

	a, err := getResponseBody(`accounts`, nil)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(a, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(prettyPrint(result))

	return result, err
}
