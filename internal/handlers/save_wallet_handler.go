package handlers

import (
	"context"
	"fmt"
	api "github.com/deliryo-io/wallet_api/internal/protoc"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
)

const query_check_wallet = "select user_id from production.wallets where user_id = $1 AND wallet = $2 AND coin = $3 AND network = $4;"
const query_save_wallet = "insert into production.wallets(user_id, wallet, coin, network) VALUES($1,$2,$3,$4);"

func (srv GRPCserver) SaveWallet(ctx context.Context, req *api.WalletInfo) (*api.SaveWalletStatus, error) {
	var user_id int32
	check_row := srv.Pool.QueryRow(ctx, query_check_wallet, req.GetUserId(), req.GetWallet(), req.GetCoin(), req.GetNetwork())
	check_err := check_row.Scan(&user_id)
	if check_err != nil {
		if check_err.Error() != UpdatedNoRowsErr.Error() {
			log.Printf("Unknown error in database_tools connection! Error: %v", check_err)
			return &api.SaveWalletStatus{
				Status: api.SaveWalletStates_ERROR,
			}, check_err
		}
	} else {
		return &api.SaveWalletStatus{
			Status: api.SaveWalletStates_WAS_SAVED,
		}, nil
	}

	row, err := srv.Pool.Query(ctx, query_save_wallet, req.GetUserId(), req.GetWallet(), req.GetCoin(), req.GetNetwork())
	defer row.Close()
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok { //обработка ошибок бд
			newErr := fmt.Sprintf("SQL Error: %s, Detail: %s, Code: %s, SQLState: %%", pgErr.Message, pgErr.Detail, pgErr.Code, pgErr.SQLState())
			log.Printf("Error in database: %v ", newErr)
			return &api.SaveWalletStatus{
				Status: api.SaveWalletStates_ERROR,
			}, err
		} else {
			log.Printf("Unknown error in database_tools connection! Error: %v", err)
			return &api.SaveWalletStatus{
				Status: api.SaveWalletStates_ERROR,
			}, err
		}
	}
	return &api.SaveWalletStatus{
		Status: api.SaveWalletStates_SUCCESS,
	}, nil
}
