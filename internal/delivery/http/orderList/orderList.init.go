package orderList

import (
	"context"
	orderlist "request-order/internal/entity/orderList"
)

type Service interface {
	GetRO(ctx context.Context) ([]orderlist.ReqOrderHeader, error)

	//trialGabung
	//GetROdHeader(ctx context.Context, sNumber string) (orderlist.ListDetailsRO, error)
	GetROdHeader(ctx context.Context, sNumber string) (orderlist.ReqOrderHeader, error)
	// GetRODetDetail(ctx context.Context, sNumber string) ([]orderlist.ReqOrderDetail, error)
	// GetRODetProcod(ctx context.Context, sNumber string) ([]orderlist.ROProcode, error)
	GetROProcods(ctx context.Context) ([]orderlist.ROProcode, error)
	GetROProcod(ctx context.Context, sProcod string) (orderlist.ReqOrderDetail, error)
}

type (
	Handler struct {
		service Service
	}
)

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}
