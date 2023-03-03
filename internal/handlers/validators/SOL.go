package validators

import (
	"log"
	"regexp"
)

func SOL_validator(request Validate_wallet_request) (result bool) {
	//https://www.reddit.com/r/solana/comments/soi03d/comment/hwwo48l/
	sol_exp, err := regexp.Compile("^[1-9A-HJ-NP-Za-km-z]{32,44}$")
	if err != nil {
		log.Printf("I can't read regular expression. Do something!")
	}
	regexp_result := sol_exp.MatchString(request.Wallet)
	return regexp_result
}
