package solanabeach

import (
	"encoding/json"

	"github.com/fatih/structs"
)

type Account struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name,omitempty"`
		Logo    string `json:"logo,omitempty"`
		Ticker  string `json:"ticker,omitempty"`
		CmcID   string `json:"cmcId,omitempty"`
	} `json:"pubkey"`
	Balance int `json:"balance"`
}

type Wealth struct {
	Wealth struct {
		Distribution struct {
			Zero10 struct {
				Value      string `json:"value"`
				Percentage string `json:"percentage"`
				UsdValue   string `json:"usdValue"`
			} `json:"0_10"`
			One150 struct {
				Value      string `json:"value"`
				Percentage string `json:"percentage"`
				UsdValue   string `json:"usdValue"`
			} `json:"11_50"`
			Five1100 struct {
				Value      string `json:"value"`
				Percentage string `json:"percentage"`
				UsdValue   string `json:"usdValue"`
			} `json:"51_100"`
			Rest struct {
				Value      string `json:"value"`
				Percentage string `json:"percentage"`
				UsdValue   string `json:"usdValue"`
			} `json:"rest"`
			Total struct {
				System string `json:"system"`
				Stake  string `json:"stake"`
				Spl    string `json:"spl"`
				Sum    string `json:"sum"`
			} `json:"total"`
		} `json:"distribution"`
		GroupedBalances []struct {
			Span struct {
				From int `json:"from"`
				To   int `json:"to"`
			} `json:"span"`
			Addresses struct {
				Count      int    `json:"count"`
				Percentage string `json:"percentage"`
			} `json:"addresses"`
			Balances struct {
				Value      string `json:"value"`
				Percentage string `json:"percentage"`
				UsdValue   string `json:"usdValue"`
			} `json:"balances"`
		} `json:"groupedBalances"`
		Price int `json:"price"`
	} `json:"wealth"`
}

type AccountParams struct {
	Limit  string
	Cursor string
}

type AccountTransaction struct {
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

type AccountTransactionsParams struct {
	Limit  string
	Offset string
	Cursor string
}

type StakeAccount struct {
	Pubkey   string `json:"pubkey"`
	Lamports int    `json:"lamports"`
	Data     struct {
		State int `json:"state"`
		Meta  struct {
			RentExemptReserve int `json:"rent_exempt_reserve"`
			Authorized        struct {
				Staker     string `json:"staker"`
				Withdrawer string `json:"withdrawer"`
			} `json:"authorized"`
		} `json:"meta"`
		Lockup struct {
			UnixTimestamp int    `json:"unix_timestamp"`
			Epoch         int    `json:"epoch"`
			Custodian     string `json:"custodian"`
		} `json:"lockup"`
		Stake struct {
			Delegation struct {
				VoterPubkey        string `json:"voter_pubkey"`
				Stake              int    `json:"stake"`
				ActivationEpoch    int    `json:"activation_epoch"`
				DeactivationEpoch  int    `json:"deactivation_epoch"`
				WarmupCooldownRate int    `json:"warmup_cooldown_rate"`
				ValidatorInfo      struct {
					Name       string `json:"name"`
					Image      string `json:"image"`
					Website    string `json:"website"`
					NodePubkey string `json:"nodePubkey"`
				} `json:"validatorInfo"`
			} `json:"delegation"`
			CreditsObserved int `json:"credits_observed"`
		} `json:"stake"`
	} `json:"data"`
}

type StakeAccountsParams struct {
	Limit  string
	Offset string
	Cursor string
}

type StakeReward []struct {
	Epoch         int    `json:"epoch"`
	EffectiveSlot int    `json:"effectiveSlot"`
	Amount        string `json:"amount"`
	PostBalance   string `json:"postBalance"`
	PercentChange int    `json:"percentChange"`
	Apr           int    `json:"apr"`
}

type StakeRewardsParams struct {
	Offset string
}

type AccountToken struct {
	Mint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"mint"`
	Amount  int `json:"amount"`
	Address struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"address"`
	Decimals int `json:"decimals"`
}

type AccountTokensParams struct {
	Limit  string
	Offset string
}

type Tokentransfer struct {
	Blocknumber int    `json:"blocknumber"`
	Txhash      string `json:"txhash"`
	Mint        struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"mint"`
	Valid  bool `json:"valid"`
	Amount int  `json:"amount"`
	Source struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"source"`
	Destination struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"destination"`
	Inner     bool `json:"inner"`
	Txindex   int  `json:"txindex"`
	Timestamp struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"timestamp"`
	Decimals int `json:"decimals"`
}

type AccountTokenTransfersParams struct {
	Limit  string
	Offset string
	Cursor string
	Inner  string
}

type SerumInstruction struct {
	ID              int    `json:"id"`
	Blocknumber     int    `json:"blocknumber"`
	Transactionhash string `json:"transactionhash"`
	Market          struct {
		Name   string `json:"name"`
		Pubkey struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"pubkey"`
		QuoteMint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"quoteMint"`
		BaseMint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"baseMint"`
	} `json:"market"`
	Owner struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"owner"`
	Instruction     string `json:"instruction"`
	Instructiondata struct {
	} `json:"instructiondata"`
	Valid     bool `json:"valid"`
	Index     int  `json:"index"`
	Txindex   int  `json:"txindex"`
	Timestamp struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"timestamp"`
	Ondemand bool `json:"ondemand"`
}

type SerumInstructionsParams struct {
	Limit  string
	Offset string
	Cursor string
}

type SerumOder struct {
	ID              int    `json:"id"`
	Blocknumber     int    `json:"blocknumber"`
	Transactionhash string `json:"transactionhash"`
	Market          struct {
		Pubkey struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"pubkey"`
		QuoteMint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"quoteMint"`
		BaseMint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"baseMint"`
		CurrentPrice  int `json:"currentPrice"`
		BaseDecimals  int `json:"baseDecimals"`
		QuoteDecimals int `json:"quoteDecimals"`
	} `json:"market"`
	Owner struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"owner"`
	Openorders struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"openorders"`
	Side        string `json:"side"`
	Type        string `json:"type"`
	Pricelimit  int    `json:"pricelimit"`
	Maxquantity int    `json:"maxquantity"`
	Valid       bool   `json:"valid"`
	Status      string `json:"status"`
	Processedat int    `json:"processedat"`
	Txindex     int    `json:"txindex"`
	Timestamp   struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"timestamp"`
	Ondemand bool `json:"ondemand"`
	Price    int  `json:"price"`
	Qty      int  `json:"qty"`
}

type SerumOrdersParams struct {
	Limit  string
	Offset string
	Cursor string
}

type SwapInstruction struct {
	ID              int    `json:"id"`
	Blocknumber     int    `json:"blocknumber"`
	Transactionhash string `json:"transactionhash"`
	Tokenswap       struct {
		Pubkey struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"pubkey"`
		Name   string `json:"name"`
		TokenA struct {
			Mint struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"mint"`
			Decimals int `json:"decimals"`
		} `json:"tokenA"`
		TokenB struct {
			Mint struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"mint"`
			Decimals int `json:"decimals"`
		} `json:"tokenB"`
		PoolToken struct {
			Decimals int `json:"decimals"`
		} `json:"poolToken"`
	} `json:"tokenswap"`
	Owner struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"owner"`
	Instruction     string `json:"instruction"`
	Instructiondata struct {
	} `json:"instructiondata"`
	Valid     bool `json:"valid"`
	Index     int  `json:"index"`
	Txindex   int  `json:"txindex"`
	Timestamp struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"timestamp"`
	Ondemand bool `json:"ondemand"`
}

type SwapInstructionsParams struct {
	Limit  string
	Offset string
	Cursor string
}

type WalletAccount struct {
	Type  string `json:"type"`
	Value struct {
		Base struct {
			Address struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"address"`
			Balance    int  `json:"balance"`
			Executable bool `json:"executable"`
			Owner      struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"owner"`
			RentEpoch int `json:"rentEpoch"`
			DataSize  int `json:"dataSize"`
		} `json:"base"`
		Extended struct {
		} `json:"extended"`
	} `json:"value"`
}

func FetchAccounts(options AccountParams) ([]Account, error) {
	var emptyStruct AccountParams
	var result []Account
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`accounts`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchWealthDistribution() (Wealth, error) {
	var result Wealth
	w, err := getResponseBody(`wealth`, nil)
	if err != nil {
		return result, err
	}

	json.Unmarshal(w, &result)

	return result, err
}

func FetchAccountTransactions(pubkey string, options AccountTransactionsParams) ([]AccountTransaction, error) {
	var emptyStruct AccountTransactionsParams
	var result []AccountTransaction
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+pubkey+`/transactions`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchStakeAccounts(pubkey string, options StakeAccountsParams) ([]StakeAccount, error) {
	var emptyStruct StakeAccountsParams
	var result []StakeAccount
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+pubkey+`/stakes`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchStakeRewards(stakePubkey string, options StakeRewardsParams) ([]StakeReward, error) {
	var emptyStruct StakeRewardsParams
	var result []StakeReward
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+stakePubkey+`/stake-rewards`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchAccountTokens(pubkey string, options AccountTokensParams) ([]AccountToken, error) {
	var emptyStruct AccountTokensParams
	var result []AccountToken
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+pubkey+`/tokens`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchAccountTokenTransfers(pubkey string, options AccountTokenTransfersParams) ([]Tokentransfer, error) {
	var emptyStruct AccountTokenTransfersParams
	var result []Tokentransfer
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+pubkey+`/token-transfers`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchAccountSerumInstructios(pubkey string, options SerumInstructionsParams) ([]SerumInstruction, error) {
	var emptyStruct SerumInstructionsParams
	var result []SerumInstruction
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+pubkey+`/serum-instructions`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchAccountSerumOrders(pubkey string, options SerumOrdersParams) ([]SerumOder, error) {
	var emptyStruct SerumOrdersParams
	var result []SerumOder
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+pubkey+`/serum-orders`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchAccountSwapInstructions(pubkey string, options SwapInstructionsParams) ([]SwapInstruction, error) {
	var emptyStruct SwapInstructionsParams
	var result []SwapInstruction
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	a, err := getResponseBody(`/account/`+pubkey+`/swap-instructions`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}

func FetchWalletAccount(pubkey string) (WalletAccount, error) {
	var result WalletAccount

	a, err := getResponseBody(`/account/`+pubkey, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(a, &result)

	return result, err
}
