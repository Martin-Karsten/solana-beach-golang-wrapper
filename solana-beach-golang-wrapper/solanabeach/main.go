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
