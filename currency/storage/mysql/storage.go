package mysql

import (
	"context"
	"database/sql"

	"github.com/rschio/currencyproj/currency/storage"
)

type MySQL struct {
	db *sql.DB
	q  *Queries
}

// NewMySQL creates a new MySQL.
func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db, q: New(db)}
}

func (m *MySQL) Close() error { return m.db.Close() }

// CreateExchange inserts a exchange into the database.
func (m *MySQL) CreateExchange(ctx context.Context, e storage.Exchange) error {
	params := CreateExchangeParams{
		Amount:         e.Amount,
		Rate:           e.Rate,
		FromCurrency:   e.From,
		ToCurrency:     e.To,
		ConvertedValue: e.ConvertedValue,
		CurrencySymbol: e.CurrencySymbol,
	}
	return m.q.CreateExchange(ctx, params)
}
