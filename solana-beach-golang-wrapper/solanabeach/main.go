package solana_beach

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func getResponse(slug string, params map[string]interface{}) []byte {
	req, err := http.NewRequest("GET", url+slug, nil)
	req.Header.Add("Authorization", bearer)

	if len(params) > 0 && params != nil {
		q := req.URL.Query()

		for param, value := range params {
			if value != "" {
				q.Add(strings.ToLower(param), value.(string))
			}
		}

		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	return body
}
