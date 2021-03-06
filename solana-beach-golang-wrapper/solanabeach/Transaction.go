package solanabeach

import (
	"encoding/json"

	"github.com/fatih/structs"
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
		Name   string  `json:"name"`
		Weight float32 `json:"weight"`
		Index  int     `json:"index"`
	} `json:"mostImportantInstruction"`
	Smart []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"smart"`
	Ondemand bool `json:"ondemand"`
}

type LatestTransactionsParams struct {
	Limit  string
	Cursor string
}

func FetchTransationByHash(hash string) (Transaction, error) {
	var result Transaction

	t, err := getResponseBody(`transaction/`+hash, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(t, &result)

	return result, err
}

func FetchLatestTransactionHashes(address string) (Transaction, error) {
	var result Transaction

	t, err := getResponseBody(`transaction-hashes/`+address, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(t, &result)

	return result, err
}

func FetchLatestTransactionsByAddress(address string) (Transaction, error) {
	var result Transaction

	t, err := getResponseBody(`transactions`, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(t, &result)

	return result, err
}

func FetchLatestTransactions(options LatestTransactionsParams) ([]Transaction, error) {
	var emptyStruct LatestTransactionsParams
	var result []Transaction
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	t, err := getResponseBody(`latest-transactions`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(t, &result)

	return result, err
}

func FetchLatestTransactionsInBlock(number string) (Transaction, error) {
	var result Transaction

	t, err := getResponseBody(`block-transactions`, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(t, &result)

	return result, err
}
