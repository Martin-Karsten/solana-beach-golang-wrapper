package solana_beach

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	TransactionHash string `json:"transactionHash"`
	BlockNumber     int    `json:"blockNumber"`
	Index           int    `json:"index"`
	Accounts        []struct {
		Account struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"account"`
		Writable bool `json:"writable"`
		Signer   bool `json:"signer"`
		Program  bool `json:"program"`
	} `json:"accounts"`
	Header struct {
		NumRequiredSignatures       int `json:"numRequiredSignatures"`
		NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
		NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
	} `json:"header"`
	Instructions []struct {
		Parsed struct {
		} `json:"parsed"`
		Raw struct {
			Data           string `json:"data"`
			Accounts       []int  `json:"accounts"`
			ProgramIDIndex int    `json:"programIdIndex"`
		} `json:"raw"`
	} `json:"instructions"`
	RecentBlockhash string   `json:"recentBlockhash"`
	Signatures      []string `json:"signatures"`
	Meta            struct {
		Err struct {
		} `json:"err"`
		Fee    int `json:"fee"`
		Status struct {
		} `json:"status"`
		LogMessages  []string `json:"logMessages"`
		PreBalances  []int    `json:"preBalances"`
		PostBalances []int    `json:"postBalances"`
	} `json:"meta"`
	Valid     bool `json:"valid"`
	Blocktime struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"blocktime"`
	MostImportantInstruction struct {
		Name   string `json:"name"`
		Weight int    `json:"weight"`
		Index  int    `json:"index"`
	} `json:"mostImportantInstruction"`
	Smart []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"smart"`
	Ondemand bool `json:"ondemand"`
}

func FetchTransationByHash(hash string) Transaction {
	var result Transaction

	if err := json.Unmarshal(getResponse(`transaction/`+hash), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(prettyPrint(result))

	return result
}

func FetchLatestTransactionHashes(address string) Transaction {
	var result Transaction

	if err := json.Unmarshal(getResponse(`transaction-hashes/`+address), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(prettyPrint(result))

	return result
}

func FetchLatestTransactionsByAddress(address string) Transaction {
	var result Transaction

	if err := json.Unmarshal(getResponse(`transactions`+address), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(prettyPrint(result))

	return result
}

func FetchLatestTransactions() []Transaction {
	var result []Transaction

	if err := json.Unmarshal(getResponse(`latest-transactions`), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(prettyPrint(result))

	return result
}

func FetchLatestTransactionsInBlock(number string) Transaction {
	var result Transaction

	if err := json.Unmarshal(getResponse(`block-transactions`+number), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(prettyPrint(result))

	return result
}