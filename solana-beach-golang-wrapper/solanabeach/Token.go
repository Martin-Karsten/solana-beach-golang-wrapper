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

type FetchTokensParams struct {
	Limit  string
	Offset string
	Sort   TokenSortByParam
	Dir    TokenSortDirectionParam
}

const (
	TokenSortByName         TokenSortByParam = "name"
	TokenSortByPrice        TokenSortByParam = "price"
	TokenSortByVolume       TokenSortByParam = "volume"
	TokenSortByChange       TokenSortByParam = "change"
	TokenSortByHolders      TokenSortByParam = "holders"
	TokenSortBySwapVolume   TokenSortByParam = "swapVolume"
	TokenSortBySwapPrice    TokenSortByParam = "swapPrice"
	TokenSortByMarketVolume TokenSortByParam = "marketVolume"
	TokenSortByMarketPrice  TokenSortByParam = "marketPrice"
)

const (
	TokenSortDirAsc  TokenSortDirectionParam = "asc"
	TokenSortDirDesc TokenSortDirectionParam = "desc"
)

var tokenSortByParams = [9]TokenSortByParam{TokenSortByName, TokenSortByPrice, TokenSortByVolume, TokenSortByChange, TokenSortByHolders, TokenSortBySwapVolume, TokenSortBySwapPrice, TokenSortByMarketVolume, TokenSortByMarketPrice}
var tokenSortDirectionParams = [2]TokenSortDirectionParam{TokenSortDirAsc, TokenSortDirDesc}

func (t TokenSortByParam) String() string {
	return string(t)
}

func (t TokenSortDirectionParam) String() string {
	return string(t)
}

func FetchTokens(options FetchTokensParams) (TokenResponse, error) {
	var emptyStruct FetchTokensParams
	var result TokenResponse

	if options.Sort == "" {
		options.Sort = TokenSortByName
	}
	if options.Dir == "" {
		options.Dir = TokenSortDirAsc
	}

	if !containsParams(tokenSortByParams, options.Sort) {
		errorMessage := "TokenSortByParam must be [" + stringParams(tokenSortByParams) + "]"
		return result, errors.New(errorMessage)
	}
	if !containsParams(tokenSortDirectionParams, options.Dir) {
		errorMessage := "TokenSortByParam must be " + stringParams(tokenSortDirectionParams)
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
