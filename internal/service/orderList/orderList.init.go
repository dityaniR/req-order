package orderList

import (
	"context"
	orderlist "request-order/internal/entity/orderList"
)

type Data interface {
	GetROHeader(ctx context.Context) ([]orderlist.ReqOrderHeader, error)

	GetRODHeader(ctx context.Context, sNumber string) (orderlist.ReqOrderHeader, error)
	GetRODDetail(ctx context.Context, sNumber string) ([]orderlist.ReqOrderDetail, error)
	GetRODProcods(ctx context.Context, sNumber string) ([]orderlist.ROProcode, error)
	GetRODProcod(ctx context.Context, sProcod string) (orderlist.ROProcode, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
