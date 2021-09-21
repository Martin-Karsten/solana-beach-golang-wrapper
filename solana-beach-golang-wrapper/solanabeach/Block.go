package solana_beach

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
)

type Programstats struct {
	Count     int `json:"count"`
	ProgramID struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Ticker  string `json:"ticker"`
		CmcID   string `json:"cmcId"`
	} `json:"programId"`
	Instructions struct {
	} `json:"instructions"`
}

type TopPrograms struct {
	Window   int            `json:"window"`
	Programs []Programstats `json:"programs"`
}

type Block struct {
	Blocknumber       int    `json:"blocknumber"`
	Blockhash         string `json:"blockhash"`
	Previousblockhash string `json:"previousblockhash"`
	Parentslot        int    `json:"parentslot"`
	Blocktime         struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"blocktime"`
	Metrics struct {
		Txcount           int `json:"txcount"`
		Failedtxs         int `json:"failedtxs"`
		Totalfees         int `json:"totalfees"`
		Instructions      int `json:"instructions"`
		Sucessfultxs      int `json:"sucessfultxs"`
		Innerinstructions int `json:"innerinstructions"`
	} `json:"metrics"`
	Programstats []Programstats `json:"programstats"`
	Rewards      interface{}    `json:"rewards"`
	Proposer     string         `json:"proposer"`
	ProposerData struct {
		Name       string `json:"name"`
		Image      string `json:"image"`
		Website    string `json:"website"`
		NodePubkey string `json:"nodePubkey"`
	} `json:"proposerData"`
	Ondemand bool `json:"ondemand"`
}

type LatestBlocksParams struct {
	Limit  string
	Cursor string
}

func FetchBlockHash(hash string) Block {
	var result Block

	if err := json.Unmarshal(getResponse(`block-hash/`+hash, nil), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func FetchBlock(number string) Block {
	var result Block

	if err := json.Unmarshal(getResponse(`block/`+number, nil), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func FetchLatestBlocks(options LatestBlocksParams) []Block {
	var emptyStruct LatestBlocksParams
	var result []Block
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	if err := json.Unmarshal(getResponse(`latest-blocks`, params), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func FetchTopPrograms() TopPrograms {
	var result TopPrograms

	if err := json.Unmarshal(getResponse(`top-programs`, nil), &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}
