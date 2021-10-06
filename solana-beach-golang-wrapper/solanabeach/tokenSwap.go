package solanabeach

import (
	"encoding/json"
	"errors"

	"github.com/fatih/structs"
)

type TokenSwapsResponse struct {
	TotalPages int `json:"totalPages"`
	Data       []struct {
		Pubkey struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"pubkey"`
		Name     string `json:"name"`
		Poolmint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"poolmint"`
		Tokenamint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"tokenamint"`
		Tokenbmint struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"tokenbmint"`
		Tokena struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"tokena"`
		Tokenb struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"tokenb"`
		Tokenprogram struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Logo    string `json:"logo"`
			Ticker  string `json:"ticker"`
			CmcID   string `json:"cmcId"`
		} `json:"tokenprogram"`
		Meta struct {
			Supply struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"supply"`
			Liquidity struct {
				Total        int `json:"total"`
				UsdTokenA    int `json:"usdTokenA"`
				UsdTokenB    int `json:"usdTokenB"`
				NativeTokenA struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"nativeTokenA"`
				NativeTokenB struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"nativeTokenB"`
			} `json:"liquidity"`
			Volume struct {
				Total        int `json:"total"`
				UsdTokenA    int `json:"usdTokenA"`
				UsdTokenB    int `json:"usdTokenB"`
				NativeTokenA struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"nativeTokenA"`
				NativeTokenB struct {
					Amount   int `json:"amount"`
					Decimals int `json:"decimals"`
				} `json:"nativeTokenB"`
			} `json:"volume"`
			Prices struct {
				AperB int `json:"AperB"`
				BperA int `json:"BperA"`
			} `json:"prices"`
			Volumes   []int `json:"volumes"`
			Volume24H int   `json:"volume24h"`
		} `json:"meta"`
	} `json:"data"`
}

type FetchTokenSwapsParams struct {
	Limit  string
	Offset string
	Sort   TokenSwapSortByParam
	Dir    TokenSwapSortDirectionParam
}

type TokenSwapsMint struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"pubkey"`
	Name     string `json:"name"`
	Poolmint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"poolmint"`
	Tokenamint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenamint"`
	Tokenbmint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenbmint"`
	Tokena struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokena"`
	Tokenb struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenb"`
	Tokenprogram struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenprogram"`
	Meta struct {
		Supply struct {
			Amount   int `json:"amount"`
			Decimals int `json:"decimals"`
		} `json:"supply"`
		Liquidity struct {
			Total        int `json:"total"`
			UsdTokenA    int `json:"usdTokenA"`
			UsdTokenB    int `json:"usdTokenB"`
			NativeTokenA struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenA"`
			NativeTokenB struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenB"`
		} `json:"liquidity"`
		Volume struct {
			Total        int `json:"total"`
			UsdTokenA    int `json:"usdTokenA"`
			UsdTokenB    int `json:"usdTokenB"`
			NativeTokenA struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenA"`
			NativeTokenB struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenB"`
		} `json:"volume"`
		Prices struct {
			AperB int `json:"AperB"`
			BperA int `json:"BperA"`
		} `json:"prices"`
		Volumes   []int `json:"volumes"`
		Volume24H int   `json:"volume24h"`
	} `json:"meta"`
}

type FetchTokenSwapsMintParams struct {
	Limit  string
	Offset string
}

type TokenSwap struct {
	Pubkey struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"pubkey"`
	Name     string `json:"name"`
	Poolmint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"poolmint"`
	Tokenamint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenamint"`
	Tokenbmint struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenbmint"`
	Tokena struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokena"`
	Tokenb struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenb"`
	Tokenprogram struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"tokenprogram"`
	Meta struct {
		Supply struct {
			Amount   int `json:"amount"`
			Decimals int `json:"decimals"`
		} `json:"supply"`
		Liquidity struct {
			Total        int `json:"total"`
			UsdTokenA    int `json:"usdTokenA"`
			UsdTokenB    int `json:"usdTokenB"`
			NativeTokenA struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenA"`
			NativeTokenB struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenB"`
		} `json:"liquidity"`
		Volume struct {
			Total        int `json:"total"`
			UsdTokenA    int `json:"usdTokenA"`
			UsdTokenB    int `json:"usdTokenB"`
			NativeTokenA struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenA"`
			NativeTokenB struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
			} `json:"nativeTokenB"`
		} `json:"volume"`
		Prices struct {
			AperB int `json:"AperB"`
			BperA int `json:"BperA"`
		} `json:"prices"`
		Volumes   []int `json:"volumes"`
		Volume24H int   `json:"volume24h"`
	} `json:"meta"`
}

type TokenSwapSortByParam string
type TokenSwapSortDirectionParam string

const (
	TokenSwapSortByName      TokenSwapSortByParam = "name"
	TokenSwapSortByPrice     TokenSwapSortByParam = "price"
	TokenSwapSortByVolume    TokenSwapSortByParam = "volume"
	TokenSwapSortByLiquidity TokenSwapSortByParam = "liquidity"
)

const (
	TokenSwapSortDirAsc  TokenSwapSortDirectionParam = "asc"
	TokenSwapSortDirDesc TokenSwapSortDirectionParam = "desc"
)

var TokenSwapSortByParams = [4]TokenSwapSortByParam{TokenSwapSortByName, TokenSwapSortByPrice, TokenSwapSortByVolume, TokenSwapSortByLiquidity}
var TokenSwapSortDirectionParams = [2]TokenSwapSortDirectionParam{TokenSwapSortDirAsc, TokenSwapSortDirDesc}

func (t TokenSwapSortByParam) String() string {
	return string(t)
}

func (t TokenSwapSortDirectionParam) String() string {
	return string(t)
}

func FetchTokenSwaps(options FetchTokenSwapsParams) (TokenSwapsResponse, error) {
	var emptyStruct FetchTokenSwapsParams
	var result TokenSwapsResponse

	if options.Sort == "" {
		options.Sort = TokenSwapSortByName
	}
	if options.Dir == "" {
		options.Dir = TokenSwapSortDirAsc
	}

	if !containsParams(TokenSwapSortByParams, options.Sort) {
		errorMessage := "TokenSwapSortByParam must be [" + stringParams(TokenSwapSortByParams) + "]"
		return result, errors.New(errorMessage)
	}
	if !containsParams(TokenSwapSortDirectionParams, options.Dir) {
		errorMessage := "TokenSwapSortByParam must be " + stringParams(TokenSwapSortDirectionParams)
		return result, errors.New(errorMessage)
	}

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody(`token-swaps`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchTokenSwapsByMint(mint string, options FetchTokenSwapsMintParams) ([]TokenSwapsMint, error) {
	var emptyStruct FetchTokenSwapsMintParams
	var result []TokenSwapsMint

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody("token-swaps/"+mint, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchTokenSwap(pubkey string) (TokenSwap, error) {
	var result TokenSwap

	b, err := getResponseBody("token-swaps/"+pubkey, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}
