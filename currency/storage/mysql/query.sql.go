// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package mysql

import (
	"context"
)

const createExchange = `-- name: CreateExchange :exec
INSERT INTO exchanges (
	amount,
	rate,
	from_currency,
	to_currency,
	converted_value,
	currency_symbol
) VALUES (
	?, ?, ?, ?, ?, ?
)
`

type CreateExchangeParams struct {
	Amount         float64
	Rate           float64
	FromCurrency   string
	ToCurrency     string
	ConvertedValue float64
	CurrencySymbol string
}

func (q *Queries) CreateExchange(ctx context.Context, arg CreateExchangeParams) error {
	_, err := q.db.ExecContext(ctx, createExchange,
		arg.Amount,
		arg.Rate,
		arg.FromCurrency,
		arg.ToCurrency,
		arg.ConvertedValue,
		arg.CurrencySymbol,
	)
	return err
}
