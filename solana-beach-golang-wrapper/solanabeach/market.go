package solanabeach

import (
	"encoding/json"
	"errors"

	"github.com/fatih/structs"
)

type MarketsResponse struct {
	TotalPages int `json:"totalPages"`
	Data       []struct {
		Pubkey struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"pubkey"`
		Basemint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"basemint"`
		Quotemint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"quotemint"`
		Meta struct {
			Asks struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"asks"`
			Bids struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"bids"`
			BaseLotSize        string `json:"baseLotSize"`
			QuoteLotSize       string `json:"quoteLotSize"`
			BaseTokenDecimals  int    `json:"baseTokenDecimals"`
			QuoteTokenDecimals int    `json:"quoteTokenDecimals"`
			BaseVault          struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"baseVault"`
			QuoteVault struct {
				Address string `json:"address"`
				Name    string `json:"name"`
				Logo    string `json:"logo"`
				Ticker  string `json:"ticker"`
				CmcID   string `json:"cmcId"`
			} `json:"quoteVault"`
			Name        string `json:"name"`
			MarketPrice int    `json:"marketPrice"`
			OrderBooks  struct {
				Asks [][]int `json:"asks"`
				Bids [][]int `json:"bids"`
			} `json:"orderBooks"`
			Liquidity struct {
				Total    int `json:"total"`
				UsdBase  int `json:"usdBase"`
				UsdQuote int `json:"usdQuote"`
				Base     struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"base"`
				Quote struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"quote"`
			} `json:"liquidity"`
			Volume struct {
				Total    int `json:"total"`
				UsdBase  int `json:"usdBase"`
				UsdQuote int `json:"usdQuote"`
				Base     struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"base"`
				Quote struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"quote"`
			} `json:"volume"`
			Volumes   []int `json:"volumes"`
			Volume24H int   `json:"volume24h"`
		} `json:"meta"`
	} `json:"data"`
}

type FetchMarketsParams struct {
	Limit  string
	offset string
	Sort   MarketSortByParam
	Dir    MarketSortDirectionParam
}

type MarketBaseMint struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"pubkey"`
	Basemint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"basemint"`
	Quotemint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"quotemint"`
	Meta struct {
		Asks struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"asks"`
		Bids struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"bids"`
		BaseLotSize        string `json:"baseLotSize"`
		QuoteLotSize       string `json:"quoteLotSize"`
		BaseTokenDecimals  int    `json:"baseTokenDecimals"`
		QuoteTokenDecimals int    `json:"quoteTokenDecimals"`
		BaseVault          struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"baseVault"`
		QuoteVault struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"quoteVault"`
		Name        string `json:"name"`
		MarketPrice int    `json:"marketPrice"`
		OrderBooks  struct {
			Asks [][]int `json:"asks"`
			Bids [][]int `json:"bids"`
		} `json:"orderBooks"`
		Liquidity struct {
			Total    int `json:"total"`
			UsdBase  int `json:"usdBase"`
			UsdQuote int `json:"usdQuote"`
			Base     struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"base"`
			Quote struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"quote"`
		} `json:"liquidity"`
		Volume struct {
			Total    int `json:"total"`
			UsdBase  int `json:"usdBase"`
			UsdQuote int `json:"usdQuote"`
			Base     struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"base"`
			Quote struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"quote"`
		} `json:"volume"`
		Volumes   []int `json:"volumes"`
		Volume24H int   `json:"volume24h"`
	} `json:"meta"`
}

type FetchMarketsByBaseMintParams struct {
	Limit  string
	offset string
}

type QuoteMint struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"pubkey"`
	Basemint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"basemint"`
	Quotemint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"quotemint"`
	Meta struct {
		Asks struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"asks"`
		Bids struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"bids"`
		BaseLotSize        string `json:"baseLotSize"`
		QuoteLotSize       string `json:"quoteLotSize"`
		BaseTokenDecimals  int    `json:"baseTokenDecimals"`
		QuoteTokenDecimals int    `json:"quoteTokenDecimals"`
		BaseVault          struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"baseVault"`
		QuoteVault struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"quoteVault"`
		Name        string `json:"name"`
		MarketPrice int    `json:"marketPrice"`
		OrderBooks  struct {
			Asks [][]int `json:"asks"`
			Bids [][]int `json:"bids"`
		} `json:"orderBooks"`
		Liquidity struct {
			Total    int `json:"total"`
			UsdBase  int `json:"usdBase"`
			UsdQuote int `json:"usdQuote"`
			Base     struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"base"`
			Quote struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"quote"`
		} `json:"liquidity"`
		Volume struct {
			Total    int `json:"total"`
			UsdBase  int `json:"usdBase"`
			UsdQuote int `json:"usdQuote"`
			Base     struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"base"`
			Quote struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"quote"`
		} `json:"volume"`
		Volumes   []int `json:"volumes"`
		Volume24H int   `json:"volume24h"`
	} `json:"meta"`
}

type FetchMarketsByQuoteMintParams struct {
	Limit  string
	offset string
}

type Market struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"pubkey"`
	Basemint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"basemint"`
	Quotemint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"quotemint"`
	Meta struct {
		Asks struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"asks"`
		Bids struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"bids"`
		BaseLotSize        string `json:"baseLotSize"`
		QuoteLotSize       string `json:"quoteLotSize"`
		BaseTokenDecimals  int    `json:"baseTokenDecimals"`
		QuoteTokenDecimals int    `json:"quoteTokenDecimals"`
		BaseVault          struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"baseVault"`
		QuoteVault struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"quoteVault"`
		Name        string `json:"name"`
		MarketPrice int    `json:"marketPrice"`
		OrderBooks  struct {
			Asks [][]int `json:"asks"`
			Bids [][]int `json:"bids"`
		} `json:"orderBooks"`
		Liquidity struct {
			Total    int `json:"total"`
			UsdBase  int `json:"usdBase"`
			UsdQuote int `json:"usdQuote"`
			Base     struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"base"`
			Quote struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"quote"`
		} `json:"liquidity"`
		Volume struct {
			Total    int `json:"total"`
			UsdBase  int `json:"usdBase"`
			UsdQuote int `json:"usdQuote"`
			Base     struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"base"`
			Quote struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"quote"`
		} `json:"volume"`
		Volumes   []int `json:"volumes"`
		Volume24H int   `json:"volume24h"`
	} `json:"meta"`
}

type MarketSortByParam string
type MarketSortDirectionParam string

const (
	MarketSortByName      MarketSortByParam = "name"
	MarketSortByPrice     MarketSortByParam = "price"
	MarketSortByVolume    MarketSortByParam = "volume"
	MarketSortByLiquidity MarketSortByParam = "liquidity"
)

const (
	MarketSortDirAsc  MarketSortDirectionParam = "asc"
	MarketSortDirDesc MarketSortDirectionParam = "desc"
)

var MarketSortByParams = [4]MarketSortByParam{MarketSortByName, MarketSortByPrice, MarketSortByVolume, MarketSortByLiquidity}
var MarketSortDirectionParams = [2]MarketSortDirectionParam{MarketSortDirAsc, MarketSortDirDesc}

func (t MarketSortByParam) String() string {
	return string(t)
}

func (t MarketSortDirectionParam) String() string {
	return string(t)
}

func FetchMarkets(options FetchMarketsParams) (MarketsResponse, error) {
	var emptyStruct FetchMarketsParams
	var result MarketsResponse

	if options.Sort == "" {
		options.Sort = MarketSortByName
	}
	if options.Dir == "" {
		options.Dir = MarketSortDirAsc
	}

	if !containsParams(MarketSortByParams, options.Sort) {
		errorMessage := "TokenSortByParam must be [" + stringParams(MarketSortByParams) + "]"
		return result, errors.New(errorMessage)
	}
	if !containsParams(MarketSortDirectionParams, options.Dir) {
		errorMessage := "TokenSortByParam must be " + stringParams(MarketSortDirectionParams)
		return result, errors.New(errorMessage)
	}

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody(`markets`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchMarketsByBaseMint(baseMint string, options FetchMarketsByBaseMintParams) ([]MarketBaseMint, error) {
	var emptyStruct FetchMarketsByBaseMintParams
	var result []MarketBaseMint

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody("markets/base/"+baseMint, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchMarketsByQuoteMint(quoteMint string, options FetchMarketsByQuoteMintParams) ([]QuoteMint, error) {
	var emptyStruct FetchMarketsByQuoteMintParams
	var result []QuoteMint

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody("markets/quote/"+quoteMint, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchMarket(pubkey string) (Market, error) {
	var result Market

	b, err := getResponseBody("market/"+pubkey, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}
