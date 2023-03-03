package validators

import (
	"log"
	"regexp"
)

func ETH_validator(request Validate_wallet_request) (result bool) {
	eth_exp, err := regexp.Compile("^(0x)[a-zA-Z0-9]{40}$")
	if err != nil {
		log.Printf("I can't read regular expression. Do something!")
	}
	regexp_result := eth_exp.MatchString(request.Wallet)
	return regexp_result
}
