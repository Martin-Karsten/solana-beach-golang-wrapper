package solanabeach

import (
	"encoding/json"

	"github.com/fatih/structs"
)

type Validtor struct {
	Validator struct {
		ActivatedStake   int    `json:"activatedStake"`
		StakePercentage  int    `json:"stakePercentage"`
		Commission       int    `json:"commission"`
		EpochCredits     []int  `json:"epochCredits"`
		EpochVoteAccount bool   `json:"epochVoteAccount"`
		LastVote         int    `json:"lastVote"`
		NodePubkey       string `json:"nodePubkey"`
		RootSlot         int    `json:"rootSlot"`
		VotePubkey       string `json:"votePubkey"`
		BlockProduction  struct {
			LeaderSlots  int `json:"leaderSlots"`
			SkippedSlots int `json:"skippedSlots"`
		} `json:"blockProduction"`
		DelegatingStakeAccounts []struct {
			Pubkey   string `json:"pubkey"`
			Lamports int    `json:"lamports"`
			Data     struct {
				State int `json:"state"`
				Meta  struct {
					RentExemptReserve int `json:"rent_exempt_reserve"`
					Authorized        struct {
						Staker     string `json:"staker"`
						Withdrawer string `json:"withdrawer"`
					} `json:"authorized"`
				} `json:"meta"`
				Lockup struct {
					UnixTimestamp int    `json:"unix_timestamp"`
					Epoch         int    `json:"epoch"`
					Custodian     string `json:"custodian"`
				} `json:"lockup"`
				Stake struct {
					Delegation struct {
						VoterPubkey        string `json:"voter_pubkey"`
						Stake              int    `json:"stake"`
						ActivationEpoch    int    `json:"activation_epoch"`
						DeactivationEpoch  int    `json:"deactivation_epoch"`
						WarmupCooldownRate int    `json:"warmup_cooldown_rate"`
						ValidatorInfo      struct {
							Name       string `json:"name"`
							Image      string `json:"image"`
							Website    string `json:"website"`
							NodePubkey string `json:"nodePubkey"`
						} `json:"validatorInfo"`
					} `json:"delegation"`
					CreditsObserved int `json:"credits_observed"`
				} `json:"stake"`
			} `json:"data"`
		} `json:"delegatingStakeAccounts"`
		DelegatorCount int `json:"delegatorCount"`
		Location       struct {
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
		Moniker    string `json:"moniker"`
		Website    string `json:"website"`
		PictureURL string `json:"pictureURL"`
		Version    string `json:"version"`
		Details    string `json:"details"`
		Asn        struct {
			Code         string `json:"code"`
			Organization string `json:"organization"`
		} `json:"asn"`
	} `json:"validator"`
	Slots [][]struct {
		RelativeSlot   int  `json:"relativeSlot"`
		AbsoluteSlot   int  `json:"absoluteSlot"`
		ConfirmedBlock bool `json:"confirmedBlock"`
	} `json:"slots"`
	Historic []struct {
		Stake      int    `json:"stake"`
		Delegators int    `json:"delegators"`
		Timestamp  string `json:"timestamp"`
	} `json:"historic"`
	LatestBlocks []struct {
		Blocknumber int `json:"blocknumber"`
		Blocktime   struct {
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
			Totalvaluemoved   int `json:"totalvaluemoved"`
		} `json:"metrics"`
		Proposer string `json:"proposer"`
	} `json:"latestBlocks"`
}

type ValidatorHistory struct {
	ID         int    `json:"id"`
	Validator  string `json:"validator"`
	Stake      int    `json:"stake"`
	Delegators int    `json:"delegators"`
	Timestamp  string `json:"timestamp"`
}

type ValidtorSlot struct {
	Validator string `json:"validator"`
	Slots     []struct {
		RelativeSlot   int  `json:"relativeSlot"`
		AbsoluteSlot   int  `json:"absoluteSlot"`
		ConfirmedBlock bool `json:"confirmedBlock"`
	} `json:"slots"`
}

type ValidatorDelegator struct {
	Pubkey   string `json:"pubkey"`
	Lamports int    `json:"lamports"`
	Data     struct {
		State int `json:"state"`
		Meta  struct {
			RentExemptReserve int `json:"rent_exempt_reserve"`
			Authorized        struct {
				Staker     string `json:"staker"`
				Withdrawer string `json:"withdrawer"`
			} `json:"authorized"`
		} `json:"meta"`
		Lockup struct {
			UnixTimestamp int    `json:"unix_timestamp"`
			Epoch         int    `json:"epoch"`
			Custodian     string `json:"custodian"`
		} `json:"lockup"`
		Stake struct {
			Delegation struct {
				VoterPubkey        string `json:"voter_pubkey"`
				Stake              int    `json:"stake"`
				ActivationEpoch    int    `json:"activation_epoch"`
				DeactivationEpoch  int    `json:"deactivation_epoch"`
				WarmupCooldownRate int    `json:"warmup_cooldown_rate"`
				ValidatorInfo      struct {
					Name       string `json:"name"`
					Image      string `json:"image"`
					Website    string `json:"website"`
					NodePubkey string `json:"nodePubkey"`
				} `json:"validatorInfo"`
			} `json:"delegation"`
			CreditsObserved int `json:"credits_observed"`
		} `json:"stake"`
	} `json:"data"`
}

type ValidatorDelegatorParams struct {
	Limit  string
	offset string
}

func FetchValidator(votePubkey string) (Validtor, error) {
	var result Validtor

	b, err := getResponseBody(`validator/`+votePubkey, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchValidatorHistory(votePubkey string) (ValidatorHistory, error) {
	var result ValidatorHistory

	b, err := getResponseBody(`validator/`+votePubkey+`/history`, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchValidatorSlots(votePubkey string) (ValidtorSlot, error) {
	var result ValidtorSlot

	b, err := getResponseBody(`validator/`+votePubkey+`/slots`, nil)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}

func FetchValidatorDelegators(votePubkey string, options ValidatorDelegatorParams) ([]ValidatorDelegator, error) {
	var emptyStruct ValidatorDelegatorParams
	var result []ValidatorDelegator
	var params = make(map[string]interface{})
	if options != emptyStruct {
		params = structs.Map(options)
	}
	b, err := getResponseBody(`validator/`+votePubkey+`/delegators`, params)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &result)

	return result, err
}
