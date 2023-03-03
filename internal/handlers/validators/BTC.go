package validators

import (
	"log"
	"regexp"
)

func BTC_validator(request Validate_wallet_request) (result bool) {
	//https://ihateregex.io/expr/bitcoin-address/
	btc_exp, err := regexp.Compile("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{24,39}$")
	if err != nil {
		log.Printf("I can't read regular expression. Do something!")
	}
	regexp_result := btc_exp.MatchString(request.Wallet)
	return regexp_result
}
