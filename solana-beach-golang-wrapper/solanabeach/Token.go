package solanabeach

import (
	"encoding/json"
	"errors"

	"github.com/fatih/structs"
)

type TokenResponse struct {
	TotalPages int `json:"totalPages"`
	Data       []struct {
		Pubkey      string      `json:"pubkey"`
		Lamports    int         `json:"lamports"`
		Supply      int         `json:"supply"`
		Decimals    int         `json:"decimals"`
		Initialized bool        `json:"initialized"`
		Holders     interface{} `json:"holders"`
		Name        string      `json:"name"`
		Ticker      string      `json:"ticker"`
		Cmcid       int         `json:"cmcid"`
		Pricedata   struct {
		} `json:"pricedata"`
		Logo string `json:"logo"`
		Meta struct {
			URL             string `json:"url"`
			SwapPrice       int    `json:"swapPrice"`
			SwapVolume      int    `json:"swapVolume"`
			MarketPrice     int    `json:"marketPrice"`
			SolanaPrice     int    `json:"solanaPrice"`
			MarketVolume    int    `json:"marketVolume"`
			SolanaVolume    int    `json:"solanaVolume"`
			MintAuthority   string `json:"mintAuthority"`
			FreezeAuthority string `json:"freezeAuthority"`
		} `json:"meta"`
	} `json:"data"`
}

type TokenHolder struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"pubkey"`
	Lamports int `json:"lamports"`
	Mint     struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"mint"`
	Owner struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"owner"`
	Amount int `json:"amount"`
	State  int `json:"state"`
	Meta   struct {
		Delegate struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"delegate"`
		Native          bool `json:"native"`
		DelegatedAmount int  `json:"delegatedAmount"`
		CloseAuthority  struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"closeAuthority"`
	} `json:"meta"`
}

type TokenHoldersParams struct {
	Limit  string
	Offset string
}

type TokenTransfer []struct {
	ID              int    `json:"id"`
	Blocknumber     int    `json:"blocknumber"`
	Transactionhash string `json:"transactionhash"`
	Mint            struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"mint"`
	Amount int `json:"amount"`
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
	Decimals  int `json:"decimals"`
	Timestamp struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"timestamp"`
	Valid            bool `json:"valid"`
	Innerinstruction bool `json:"innerinstruction"`
	Index            int  `json:"index"`
	Owner            struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"owner"`
}

type TokenTransfersParams struct {
	Limit  string
	Offset string
	Cursor string
}

type Token struct {
	Name      string `json:"name"`
	Ticker    string `json:"ticker"`
	Logo      string `json:"logo"`
	Cmcid     int    `json:"cmcid"`
	Pricedata struct {
		Price           int `json:"price"`
		Volume24H       int `json:"volume_24h"`
		PercentChange1H int `json:"percent_change_1h"`
		LastUpdated     int `json:"last_updated"`
	} `json:"pricedata"`
	Pubkey      string `json:"pubkey"`
	Lamports    int    `json:"lamports"`
	Supply      int    `json:"supply"`
	Decimals    int    `json:"decimals"`
	Initialized bool   `json:"initialized"`
	Holders     int    `json:"holders"`
	Meta        struct {
		FreezeAuthority string `json:"freezeAuthority"`
		MintAuthority   string `json:"mintAuthority"`
		URL             string `json:"url"`
		SwapVolume      int    `json:"swapVolume"`
		SwapPrice       int    `json:"swapPrice"`
		MarketVolume    int    `json:"marketVolume"`
		MarketPrice     int    `json:"marketPrice"`
	} `json:"meta"`
	Ondemand bool `json:"ondemand"`
}

type TokenSortByParam string
type TokenSortDirectionParam string

const (
	SortByName         TokenSortByParam = "name"
	SortByPrice        TokenSortByParam = "price"
	SortByVolume       TokenSortByParam = "volume"
	SortByChange       TokenSortByParam = "change"
	SortByHolders      TokenSortByParam = "holders"
	SortBySwapVolume   TokenSortByParam = "swapVolume"
	SortBySwapPrice    TokenSortByParam = "swapPrice"
	SortByMarketVolume TokenSortByParam = "marketVolume"
	SortByMarketPrice  TokenSortByParam = "marketPrice"
)

const (
	SortDirAsc  TokenSortDirectionParam = "asc"
	SortDirDesc TokenSortDirectionParam = "desc"
)

var TokenSortByParams = [9]TokenSortByParam{SortByName, SortByPrice, SortByVolume, SortByChange, SortByHolders, SortBySwapVolume, SortBySwapPrice, SortByMarketVolume, SortByMarketPrice}
var TokenSortDirectionParams = [2]TokenSortDirectionParam{SortDirAsc, SortDirDesc}

func (t TokenSortByParam) String() string {
	return string(t)
}

func (t TokenSortDirectionParam) String() string {
	return string(t)
}

func containsParams(params interface{}, element interface{}) bool {
	switch t := params.(type) {
	case [9]TokenSortByParam:
		for _, i := range t {
			if i == element {
				return true
			}
		}
	case [2]TokenSortDirectionParam:
		for _, i := range t {
			if i == element {
				return true
			}
		}
	}
	return false
}

func stringParams(params interface{}) string {
	result := ""
	switch t := params.(type) {
	case [9]TokenSortByParam:
		{
			for _, str := range t {
				if len(result) == 0 {
					result += str.String()
				} else {
					result += ", " + str.String()
				}
			}
		}
	case [2]TokenSortDirectionParam:
		{
			for _, str := range t {
				if len(result) == 0 {
					result += str.String()
				} else {
					result += ", " + str.String()
				}
			}
		}
	}

	return result
}

type FetchTokensParams struct {
	Limit  string
	Offset string
	Sort   TokenSortByParam
	Dir    TokenSortDirectionParam
}

func FetchTokens(options FetchTokensParams) (TokenResponse, error) {
	var emptyStruct FetchTokensParams
	var result TokenResponse

	if options.Sort == "" {
		options.Sort = SortByName
	}
	if options.Dir == "" {
		options.Dir = SortDirAsc
	}

	if !containsParams(TokenSortByParams, options.Sort) {
		errorMessage := "TokenSortByParam must be [" + stringParams(TokenSortByParams) + "]"
		return result, errors.New(errorMessage)
	}
	if !containsParams(TokenSortDirectionParams, options.Dir) {
		errorMessage := "TokenSortByParam must be " + stringParams(TokenSortDirectionParams)
		return result, errors.New(errorMessage)
	}

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody(`tokens`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchTokenHolders(mint string, options TokenHoldersParams) ([]TokenHolder, error) {
	var emptyStruct TokenHoldersParams
	var result []TokenHolder

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody("token/"+mint+"holders", params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchTokenTransfers(mint string, options TokenTransfersParams) ([]TokenTransfer, error) {
	var emptyStruct TokenTransfersParams
	var result []TokenTransfer

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody("token/"+mint+"transfers", params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchToken(pubkey string) (Token, error) {
	var result Token

	b, err := getResponseBody("token", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}
