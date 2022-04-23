package orderList

import (
	"context"
	orderlist "request-order/internal/entity/orderList"
)

type Data interface {
	GetROHeader(ctx context.Context) ([]orderlist.Orders, error)

	GetRODHeader(ctx context.Context, sNumber string) (orderlist.ReqOrderHeader, error)
	GetRODDetail(ctx context.Context, sNumber string) ([]orderlist.ReqOrderDetail, error)
	GetRODProcod(ctx context.Context, sNumber string) ([]orderlist.ROProcode, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
