package validators

import (
	"log"
	"regexp"
)

func TRC_20_validator(request Validate_wallet_request) (result bool) {
	//https://stackoverflow.com/questions/69465635/how-to-programmatically-validate-tether-trc20-wallet-addresses-in-php
	trc_20_exp, err := regexp.Compile("^T[A-Za-z1-9]{33}$")
	if err != nil {
		log.Printf("I can't read regular expression. Do something!")
	}
	regexp_result := trc_20_exp.MatchString(request.Wallet)
	return regexp_result
}
