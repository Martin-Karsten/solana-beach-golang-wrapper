package solanabeach

import (
	"encoding/json"
	"time"

	"github.com/fatih/structs"
)

type StakeHistory struct {
	Epoch        int `json:"epoch"`
	Effective    int `json:"effective"`
	Activating   int `json:"activating"`
	Deactivating int `json:"deactivating"`
}

type FetchStakeHistoryParams struct {
	Limit  string
	Offset string
}

type Health struct {
	Boundaries struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"boundaries"`
	Window      int `json:"window"`
	NetworkLag  int `json:"networkLag"`
	CurrentSlot int `json:"currentSlot"`
}

type NetworkStatus struct {
	LastSyncedSlot    int `json:"lastSyncedSlot"`
	LastNetworkSlot   int `json:"lastNetworkSlot"`
	NetworkLag        int `json:"networkLag"`
	LastEpochInfoSync int `json:"lastEpochInfoSync"`
	LaggingBehind     int `json:"laggingBehind"`
}

type StakingApy struct {
	Apy float64 `json:"apy"`
}

type Epoch struct {
	Epoch int `json:"epoch"`
}

type Alias struct {
	MintAddress struct {
		Name        string `json:"name"`
		Ticker      string `json:"ticker"`
		CmcID       string `json:"cmcId"`
		CoingeckoID string `json:"coingeckoId"`
		Logo        string `json:"logo"`
		Meta        struct {
			URL string `json:"url"`
		} `json:"meta"`
	} `json:"mintAddress"`
}

type NonValidator struct {
	Pubkey struct {
		Address string `json:"address"`
	} `json:"pubkey"`
	FeatureSet   int    `json:"featureSet"`
	Gossip       string `json:"gossip"`
	RPC          string `json:"rpc"`
	ShredVersion int    `json:"shredVersion"`
	Tpu          string `json:"tpu"`
	Version      string `json:"version"`
	Location     struct {
		Range    []int  `json:"range"`
		Country  string `json:"country"`
		Region   string `json:"region"`
		Eu       string `json:"eu"`
		Timezone string `json:"timezone"`
		City     string `json:"city"`
		Ll       []int  `json:"ll"`
		Metro    int    `json:"metro"`
		Area     int    `json:"area"`
	} `json:"location"`
	Asn struct {
		Code         int    `json:"code"`
		Organization string `json:"organization"`
	} `json:"asn"`
	Validator bool `json:"validator"`
}

type MarketChartData struct {
	Timestamp time.Time `json:"timestamp"`
	Price     float64   `json:"price"`
	Volume24H float64   `json:"volume_24h"`
	MarketCap float64   `json:"market_cap"`
}

func FetchStakeHistory(options FetchStakeHistoryParams) ([]StakeHistory, error) {
	var emptyStruct FetchStakeHistoryParams
	var result []StakeHistory

	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody("stake-history", params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchHealth() (Health, error) {
	var result Health

	b, err := getResponseBody("health", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchNetworkStatus() (NetworkStatus, error) {
	var result NetworkStatus

	b, err := getResponseBody("network-status", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchStakingApy() (float64, error) {
	var result StakingApy

	b, err := getResponseBody("staking-apy", nil)
	if err != nil {
		return 0, err
	}
	json.Unmarshal(b, &result)

	return result.Apy, err
}

func FetchEpochHistory() (Epoch, error) {
	var result Epoch

	b, err := getResponseBody("epoch-history", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

/*
 *	Fetch alias data. Aliases represent an address to metadata mapping that includes human readable address names and token metadata (such as ticker, website, ...).
 *	It includes all tokens from the official Solana token list, most of the known programs and DEX markets.
 */
func FetchAlias() (Alias, error) {
	var result Alias

	b, err := getResponseBody("alias", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchNonValidators() ([]NonValidator, error) {
	var result []NonValidator

	b, err := getResponseBody("non-validators", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchClusterTime() (time.Time, error) {
	var result int

	b, err := getResponseBody("cluster-time", nil)
	if err != nil {
		return time.Time{}, err
	}
	json.Unmarshal(b, &result)

	return time.Unix(int64(result), 0), err
}

func FetchMarketChartData() ([]MarketChartData, error) {
	var result []MarketChartData

	b, err := getResponseBody("market-chart-data", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}
