package solana_beach

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()
var base_url = "https://api.solanabeach.io/v1/"
var bearer = "Bearer " + os.Getenv("TOKEN")

func getResponse(slug string) string {
	req, err := http.NewRequest("GET", `${base_url}{slug}`, nil)

	req.Header.Add("Authorization", bearer)

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
	return string([]byte(body))
}
