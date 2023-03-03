package handlers

import (
	"context"
	validators "github.com/deliryo-io/wallet_api/internal/handlers/validators"
	api "github.com/deliryo-io/wallet_api/internal/protoc"
)

//type Validate_wallet_request struct {
//	Wallet  string
//	Network string
//}

var Parser_map = map[string]func(request validators.Validate_wallet_request) (result bool){
	"TRC_20": validators.TRC_20_validator,
	"BTC":    validators.BTC_validator,
	"SOL":    validators.SOL_validator,
	"ETH":    validators.ETH_validator,
}

func (srv GRPCserver) ValidateWallet(ctx context.Context, req *api.ValidateWalletRequest) (*api.ValidateResult, error) {

	validate_wallet_request := validators.Validate_wallet_request{
		Wallet:  req.GetWallet(),
		Network: req.GetNetwork(),
	}

	validation_func, ok := Parser_map[req.GetNetwork()]
	if !ok {
		return &api.ValidateResult{
			InternalStatus: api.ValidateStates_VALIDATOR_ERROR,
			Result:         false,
		}, nil
	}
	result := validation_func(validate_wallet_request)

	return &api.ValidateResult{
		InternalStatus: api.ValidateStates_CHECKED,
		Result:         result,
	}, nil
}
