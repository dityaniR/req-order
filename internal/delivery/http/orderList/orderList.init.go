package orderList

import (
	"context"
	orderlist "request-order/internal/entity/orderList"
)

type Service interface {
	GetRO(ctx context.Context) ([]orderlist.Orders, error)

	GetRODetHeader(ctx context.Context, sNumber string) ([]orderlist.ReqOrderHeader, error)
	GetRODetDetail(ctx context.Context, sNumber string) ([]orderlist.ReqOrderDetail, error)
	GetRODetProcod(ctx context.Context, sNumber string) ([]orderlist.ROProcode, error)
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
