package solanabeach

import "encoding/json"

type Inflation struct {
	Epoch      int `json:"epoch"`
	Foundation int `json:"foundation"`
	Total      int `json:"total"`
	Validator  int `json:"validator"`
}

type Supply struct {
	Total          int `json:"total"`
	Circulating    int `json:"circulating"`
	NonCirculating int `json:"nonCirculating"`
}

func FetchInflation() (Inflation, error) {
	var result Inflation

	b, err := getResponseBody("inflation", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchSupply() (Supply, error) {
	var result Supply

	b, err := getResponseBody("supply", nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}
