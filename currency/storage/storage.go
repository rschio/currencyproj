package storage

import "context"

type Storage interface {
	Close() error
	CreateExchange(context.Context, Exchange) error
}

type Exchange struct {
	Amount         float64
	Rate           float64
	From           string
	To             string
	ConvertedValue float64
	CurrencySymbol string
}
