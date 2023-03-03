package handlers

import (
	"context"
	"fmt"
	api "github.com/deliryo-io/wallet_api/internal/protoc"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
)

const query_get_wallets = "select  coin, network, array_agg(wallet) as wallets from production.wallets where user_id = $1 group by (coin, network) order by coin;"

type Query_row_wallets struct {
	Coin    string
	Network string
	Wallets []string
}

type Response_get_wallets struct {
	coins []Coin
}

type Coin struct {
	coin_nm  string
	networks []Network
}
type Network struct {
	network_nm string
	wallets    []string
}

func (srv GRPCserver) GetWallets(ctx context.Context, req *api.UserId) (*api.GetWalletsResult, error) {
	var query_row_wallets Query_row_wallets
	var first_coin_flg = true
	var response_get_wallets api.GetWalletsResult
	var last_considered_coin string

	log.Println("Start getting wallets for user_id:", req.GetUserId())
	rows, err := srv.Pool.Query(ctx, query_get_wallets, req.UserId)
	if err != nil {
		log.Printf("Error by getting rows: %v, user_id: %v", err, req.GetUserId())
		return &api.GetWalletsResult{Status: api.GetWalletsStatus_INTERNAL_ERROR, Coins: nil}, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); {
		err := rows.Scan(&query_row_wallets.Coin, &query_row_wallets.Network, &query_row_wallets.Wallets)
		if err != nil {
			return Exclude_DB_errors(err)
		}

		if first_coin_flg {
			response_get_wallets.Status = api.GetWalletsStatus_OK
			var networks []*api.Networks
			response_get_wallets.Coins = append(response_get_wallets.Coins, &api.Coins{Coin: query_row_wallets.Coin, Networks: append(networks, &api.Networks{
				Network: query_row_wallets.Network,
				Wallets: query_row_wallets.Wallets,
			})})
			first_coin_flg = false
			last_considered_coin = query_row_wallets.Coin
			i++
		} else if query_row_wallets.Coin == last_considered_coin {
			response_get_wallets.Coins[i-1].Networks = append(response_get_wallets.Coins[i-1].Networks, &api.Networks{
				Network: query_row_wallets.Network,
				Wallets: query_row_wallets.Wallets,
			})
		} else {
			var networks []*api.Networks
			response_get_wallets.Coins = append(response_get_wallets.Coins, &api.Coins{Coin: query_row_wallets.Coin, Networks: append(networks, &api.Networks{
				Network: query_row_wallets.Network,
				Wallets: query_row_wallets.Wallets,
			})})
			last_considered_coin = query_row_wallets.Coin
			i++
		}
	}
	return &response_get_wallets, nil
}

func Exclude_DB_errors(err error) (*api.GetWalletsResult, error) {
	if pgErr, ok := err.(*pgconn.PgError); ok { //обработка ошибок бд
		newErr := fmt.Sprintf("SQL Error: %s, Detail: %s, Code: %s, SQLState: %%", pgErr.Message, pgErr.Detail, pgErr.Code, pgErr.SQLState())
		log.Printf("Error in database: %v ", newErr)
		return &api.GetWalletsResult{
			Status: api.GetWalletsStatus_INTERNAL_ERROR,
			Coins:  nil,
		}, err
	} else if err == UpdatedNoRowsErr { //Ошибка отсутствия строчки в базе
		log.Printf("Row is not found.")
		return &api.GetWalletsResult{
			Status: api.GetWalletsStatus_EMPTY,
			Coins:  nil,
		}, err
	} else {
		log.Printf("Unknown error in database_tools connection! Error: %v", err)
		return &api.GetWalletsResult{
			Status: api.GetWalletsStatus_INTERNAL_ERROR,
			Coins:  nil,
		}, err
	}
}
