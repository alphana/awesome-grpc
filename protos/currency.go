package server

import (
	"awesome-grpc/protos/protos"
	"context"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	logger hclog.Logger
}

func (currency Currency) GetRate(ctx context.Context, request *protos.RateRequest) (*protos.RateResponse, error) {
	currency.logger.Info("Handle GetRate", "base", request.GetBase(), "destination", request.GetDestination())
	response := protos.RateResponse{Rate: 0.0}
	return &response, nil
}

func NewCurrency(logger hclog.Logger) *Currency {
	return &Currency{logger: logger}
}
