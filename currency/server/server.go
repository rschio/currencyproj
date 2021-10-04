package server

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/rschio/currencyproj/currency/currency/proto"
	"github.com/rschio/currencyproj/currency/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	db storage.Storage
	pb.UnimplementedExchangerServer
}

func New(db storage.Storage) *Server {
	return &Server{db: db}
}

func (s *Server) Exchange(ctx context.Context, req *pb.ExchangeRequest,
) (*pb.ExchangeResponse, error) {
	if !isValidExchange(req.From, req.To) {
		return nil, status.Error(codes.InvalidArgument, "invalid exchange")
	}
	if req.Rate <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid rate")
	}

	roundedString := fmt.Sprintf("%.2f", req.Amount*req.Rate)
	value, err := strconv.ParseFloat(roundedString, 64)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	from := symbol[req.From.String()]
	to := symbol[req.To.String()]

	err = s.db.CreateExchange(ctx, storage.Exchange{
		Amount:         req.Amount,
		Rate:           req.Rate,
		From:           from,
		To:             to,
		ConvertedValue: value,
		CurrencySymbol: to,
	})
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.ExchangeResponse{
		ConvertedValue: value,
		CurrencySymbol: to,
	}, nil
}

func isValidExchange(from, to pb.Currency) bool {
	switch from {
	case pb.Currency_BRL:
		switch to {
		case pb.Currency_USD, pb.Currency_EUR:
			return true
		default:
			return false
		}
	case pb.Currency_USD:
		if to == pb.Currency_BRL {
			return true
		}
		return false
	case pb.Currency_EUR:
		if to == pb.Currency_BRL {
			return true
		}
		return false
	case pb.Currency_BTC:
		switch to {
		case pb.Currency_USD, pb.Currency_BRL:
			return true
		default:
			return false
		}
	default:
		return false
	}
}

var symbol = map[string]string{
	"USD": "$",
	"BRL": "R$",
	"EUR": "€",
	"BTC": "₿",
}
