package server

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	pb "github.com/rschio/currencyproj/currency/currency/proto"
	"github.com/rschio/currencyproj/currency/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestExchange(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name string
		req  *pb.ExchangeRequest
		want *pb.ExchangeResponse
	}{
		{name: "Round problem",
			req: &pb.ExchangeRequest{
				Amount: 23.32,
				Rate:   0.18,
				From:   pb.Currency_BRL,
				To:     pb.Currency_USD,
			},
			want: &pb.ExchangeResponse{
				ConvertedValue: 4.20,
				CurrencySymbol: "$",
			},
		},
		{name: "BTC-BRL",
			req: &pb.ExchangeRequest{
				Amount: 10.55,
				Rate:   200000.0,
				From:   pb.Currency_BTC,
				To:     pb.Currency_BRL,
			},
			want: &pb.ExchangeResponse{
				ConvertedValue: 2110000,
				CurrencySymbol: "R$",
			},
		},
		{name: "USD-BRL",
			req: &pb.ExchangeRequest{
				Amount: 10,
				Rate:   4.50,
				From:   pb.Currency_USD,
				To:     pb.Currency_BRL,
			},
			want: &pb.ExchangeResponse{
				ConvertedValue: 45,
				CurrencySymbol: "R$",
			},
		},
	}

	s := New(mockDB{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Exchange(ctx, tt.req)
			if err != nil {
				t.Error(err)
			}
			ignoreUnexported := cmpopts.IgnoreUnexported(pb.ExchangeResponse{})
			if !cmp.Equal(got, tt.want, ignoreUnexported) {
				t.Error(cmp.Diff(got, tt.want, ignoreUnexported))
			}
		})
	}
}

func TestExchangeInvalidRate(t *testing.T) {
	ctx := context.Background()
	req := &pb.ExchangeRequest{
		Amount: 10,
		Rate:   0,
		From:   pb.Currency_USD,
		To:     pb.Currency_BRL,
	}
	s := New(mockDB{})
	want := codes.InvalidArgument
	_, err := s.Exchange(ctx, req)
	if got := status.Convert(err).Code(); got != want {
		t.Fatalf("got code: %v, want: %v", got, want)
	}
}

func TestExchangeInvalidCurrencies(t *testing.T) {
	ctx := context.Background()
	req := &pb.ExchangeRequest{
		Amount: 10,
		Rate:   1,
		From:   pb.Currency_USD,
		To:     pb.Currency_USD,
	}
	s := New(mockDB{})
	want := codes.InvalidArgument
	_, err := s.Exchange(ctx, req)
	if got := status.Convert(err).Code(); got != want {
		t.Fatalf("got code: %v, want: %v", got, want)
	}
}

type mockDB struct{}

func (mockDB) Close() error { return nil }
func (mockDB) CreateExchange(context.Context, storage.Exchange) error {
	return nil
}
