package solanabeach

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()
var url = "https://api.solanabeach.io/v1/"
var bearer = "Bearer " + os.Getenv("TOKEN")

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func getResponseBody(slug string, params map[string]interface{}) ([]byte, error) {
	req, err := http.NewRequest("GET", url+slug, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)

	if len(params) > 0 && params != nil {
		q := req.URL.Query()

		for param, value := range params {
			if value != "" {
				switch tValue := value.(type) {
				case TokenSortByParam:
					q.Add(strings.ToLower(param), tValue.String())
				case TokenSortDirectionParam:
					q.Add(strings.ToLower(param), tValue.String())
				case MarketSortByParam:
					q.Add(strings.ToLower(param), tValue.String())
				case MarketSortDirectionParam:
					q.Add(strings.ToLower(param), tValue.String())
				default:
					q.Add(strings.ToLower(param), value.(string))
				}
			}
		}

		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
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
	case [4]MarketSortByParam:
		for _, i := range t {
			if i == element {
				return true
			}
		}
	case [2]MarketSortDirectionParam:
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
	case [4]MarketSortByParam:
		{
			for _, str := range t {
				if len(result) == 0 {
					result += str.String()
				} else {
					result += ", " + str.String()
				}
			}
		}
	case [2]MarketSortDirectionParam:
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
