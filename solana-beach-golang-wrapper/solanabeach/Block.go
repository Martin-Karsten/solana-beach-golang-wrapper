package solanabeach

import (
	"encoding/json"

	"github.com/fatih/structs"
)

type TopPrograms struct {
	Window   int `json:"window"`
	Programs []struct {
		ProgramID struct {
			Name    string `json:"name"`
			Address string `json:"address"`
			Logo    string `json:"logo,omitempty"`
			Ticker  string `json:"ticker,omitempty"`
			CmcID   string `json:"cmcId,omitempty"`
		} `json:"programId"`
		Count        int `json:"count"`
		Instructions struct {
			Vote                         int `json:"Vote,omitempty"`
			VoteSwitch                   int `json:"VoteSwitch,omitempty"`
			Withdraw                     int `json:"Withdraw,omitempty"`
			UpdateCommission             int `json:"UpdateCommission,omitempty"`
			NewOrderV3                   int `json:"NewOrderV3,omitempty"`
			ConsumeEvents                int `json:"ConsumeEvents,omitempty"`
			CancelOrderByClientIDV2      int `json:"CancelOrderByClientIdV2,omitempty"`
			SettleFunds                  int `json:"SettleFunds,omitempty"`
			MatchOrders                  int `json:"MatchOrders,omitempty"`
			CancelOrderV2                int `json:"CancelOrderV2,omitempty"`
			Transfer                     int `json:"Transfer,omitempty"`
			CreateAccount                int `json:"CreateAccount,omitempty"`
			AdvanceNonceAccount          int `json:"AdvanceNonceAccount,omitempty"`
			CreateAccountWithSeed        int `json:"CreateAccountWithSeed,omitempty"`
			MintTo                       int `json:"MintTo,omitempty"`
			Approve                      int `json:"Approve,omitempty"`
			CloseAccount                 int `json:"CloseAccount,omitempty"`
			SetAuthority                 int `json:"SetAuthority,omitempty"`
			InitializeMint               int `json:"InitializeMint,omitempty"`
			InitializeAccount            int `json:"InitializeAccount,omitempty"`
			TransferChecked              int `json:"TransferChecked,omitempty"`
			Revoke                       int `json:"Revoke,omitempty"`
			MintToChecked                int `json:"MintToChecked,omitempty"`
			Burn                         int `json:"Burn,omitempty"`
			BurnChecked                  int `json:"BurnChecked,omitempty"`
			Memo                         int `json:"Memo,omitempty"`
			CreateAssociatedTokenAccount int `json:"CreateAssociatedTokenAccount,omitempty"`
			Deactivate                   int `json:"Deactivate,omitempty"`
			Delegate                     int `json:"Delegate,omitempty"`
			Initialize                   int `json:"Initialize,omitempty"`
			Authorize                    int `json:"Authorize,omitempty"`
			Split                        int `json:"Split,omitempty"`
			Swap                         int `json:"Swap,omitempty"`
		} `json:"instructions,omitempty"`
	} `json:"programs"`
}

type ProgramID struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Logo    string `json:"logo,omitempty"`
	Ticker  string `json:"ticker,omitempty"`
	CmcID   string `json:"cmcId,omitempty"`
}

type Instructions struct {
	Vote                         int `json:"Vote,omitempty"`
	VoteSwitch                   int `json:"VoteSwitch,omitempty"`
	Withdraw                     int `json:"Withdraw,omitempty"`
	UpdateCommission             int `json:"UpdateCommission,omitempty"`
	NewOrderV3                   int `json:"NewOrderV3,omitempty"`
	ConsumeEvents                int `json:"ConsumeEvents,omitempty"`
	CancelOrderByClientIDV2      int `json:"CancelOrderByClientIdV2,omitempty"`
	SettleFunds                  int `json:"SettleFunds,omitempty"`
	MatchOrders                  int `json:"MatchOrders,omitempty"`
	CancelOrderV2                int `json:"CancelOrderV2,omitempty"`
	Transfer                     int `json:"Transfer,omitempty"`
	CreateAccount                int `json:"CreateAccount,omitempty"`
	AdvanceNonceAccount          int `json:"AdvanceNonceAccount,omitempty"`
	CreateAccountWithSeed        int `json:"CreateAccountWithSeed,omitempty"`
	MintTo                       int `json:"MintTo,omitempty"`
	Approve                      int `json:"Approve,omitempty"`
	CloseAccount                 int `json:"CloseAccount,omitempty"`
	SetAuthority                 int `json:"SetAuthority,omitempty"`
	InitializeMint               int `json:"InitializeMint,omitempty"`
	InitializeAccount            int `json:"InitializeAccount,omitempty"`
	TransferChecked              int `json:"TransferChecked,omitempty"`
	Revoke                       int `json:"Revoke,omitempty"`
	MintToChecked                int `json:"MintToChecked,omitempty"`
	Burn                         int `json:"Burn,omitempty"`
	BurnChecked                  int `json:"BurnChecked,omitempty"`
	Memo                         int `json:"Memo,omitempty"`
	CreateAssociatedTokenAccount int `json:"CreateAssociatedTokenAccount,omitempty"`
	Deactivate                   int `json:"Deactivate,omitempty"`
	Delegate                     int `json:"Delegate,omitempty"`
	Initialize                   int `json:"Initialize,omitempty"`
	Authorize                    int `json:"Authorize,omitempty"`
	Split                        int `json:"Split,omitempty"`
	Swap                         int `json:"Swap,omitempty"`
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
	Programstats []struct {
		Count     int `json:"count"`
		ProgramID struct {
			Name    string `json:"name"`
			Address string `json:"address"`
			Logo    string `json:"logo,omitempty"`
			Ticker  string `json:"ticker,omitempty"`
			CmcID   string `json:"cmcId,omitempty"`
		} `json:"programId"`
		Instructions struct {
			Vote                         int `json:"Vote,omitempty"`
			VoteSwitch                   int `json:"VoteSwitch,omitempty"`
			Withdraw                     int `json:"Withdraw,omitempty"`
			UpdateCommission             int `json:"UpdateCommission,omitempty"`
			NewOrderV3                   int `json:"NewOrderV3,omitempty"`
			ConsumeEvents                int `json:"ConsumeEvents,omitempty"`
			CancelOrderByClientIDV2      int `json:"CancelOrderByClientIdV2,omitempty"`
			SettleFunds                  int `json:"SettleFunds,omitempty"`
			MatchOrders                  int `json:"MatchOrders,omitempty"`
			CancelOrderV2                int `json:"CancelOrderV2,omitempty"`
			Transfer                     int `json:"Transfer,omitempty"`
			CreateAccount                int `json:"CreateAccount,omitempty"`
			AdvanceNonceAccount          int `json:"AdvanceNonceAccount,omitempty"`
			CreateAccountWithSeed        int `json:"CreateAccountWithSeed,omitempty"`
			MintTo                       int `json:"MintTo,omitempty"`
			Approve                      int `json:"Approve,omitempty"`
			CloseAccount                 int `json:"CloseAccount,omitempty"`
			SetAuthority                 int `json:"SetAuthority,omitempty"`
			InitializeMint               int `json:"InitializeMint,omitempty"`
			InitializeAccount            int `json:"InitializeAccount,omitempty"`
			TransferChecked              int `json:"TransferChecked,omitempty"`
			Revoke                       int `json:"Revoke,omitempty"`
			MintToChecked                int `json:"MintToChecked,omitempty"`
			Burn                         int `json:"Burn,omitempty"`
			BurnChecked                  int `json:"BurnChecked,omitempty"`
			Memo                         int `json:"Memo,omitempty"`
			CreateAssociatedTokenAccount int `json:"CreateAssociatedTokenAccount,omitempty"`
			Deactivate                   int `json:"Deactivate,omitempty"`
			Delegate                     int `json:"Delegate,omitempty"`
			Initialize                   int `json:"Initialize,omitempty"`
			Authorize                    int `json:"Authorize,omitempty"`
			Split                        int `json:"Split,omitempty"`
			Swap                         int `json:"Swap,omitempty"`
		} `json:"instructions,omitempty"`
	} `json:"programstats"`
	Rewards []struct {
		Pubkey      string      `json:"pubkey,omitempty"`
		Lamports    int         `json:"lamports,omitempty"`
		Commission  interface{} `json:"commission,omitempty"`
		RewardType  string      `json:"rewardType,omitempty"`
		PostBalance int         `json:"postBalance,omitempty"`
	} `json:"rewards"`
	Proposer     string `json:"proposer"`
	ProposerData struct {
		Name       string `json:"name"`
		Image      string `json:"image"`
		Website    string `json:"website"`
		NodePubkey string `json:"nodePubkey"`
	} `json:"proposerData"`
	Ondemand bool `json:"ondemand"`
}

type Programstat struct {
	Count     int `json:"count"`
	ProgramID struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Logo    string `json:"logo,omitempty"`
		Ticker  string `json:"ticker,omitempty"`
		CmcID   string `json:"cmcId,omitempty"`
	} `json:"programId"`
}

type LatestBlocksParams struct {
	Limit  string
	Cursor string
}

func FetchBlock(number string) (Block, error) {
	var result Block

	b, err := getResponseBody(`block/`+number, nil)
	if err != nil {
		return result, err
	}

	json.Unmarshal(b, &result)

	return result, err
}

func FetchBlockByHash(hash string) (Block, error) {
	var result Block

	b, err := getResponseBody(`block-hash/`+hash, nil)
	if err != nil {
		return result, err
	}

	json.Unmarshal(b, &result)

	return result, err
}

func FetchLatestBlocks(options LatestBlocksParams) ([]Block, error) {
	var emptyStruct LatestBlocksParams
	var result []Block
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}

	b, err := getResponseBody(`latest-blocks`, params)
	if err != nil {
		return result, err
	}

	json.Unmarshal(b, &result)

	return result, err
}

func FetchTopPrograms() (TopPrograms, error) {
	var result TopPrograms

	b, err := getResponseBody(`top-programs`, nil)
	if err != nil {
		return result, err
	}

	json.Unmarshal(b, &result)

	return result, err
}
