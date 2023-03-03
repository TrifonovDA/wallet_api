package handlers

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GRPCserver struct {
	Pool *pgxpool.Pool
}

var UpdatedNoRowsErr error = errors.New("no rows in result set")
