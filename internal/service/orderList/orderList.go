package orderList

import (
	"context"
	orderlist "request-order/internal/entity/orderList"
	"request-order/pkg/errors"
)

func (s Service) GetRO(ctx context.Context) ([]orderlist.Orders, error) {
	var roheaders []orderlist.Orders

	headers, err := s.data.GetROHeader(ctx)
	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetROHeader]")
	}

	roheaders = headers
	return roheaders, err
}

func (s Service) GetRODetHeader(ctx context.Context, sNumber string) ([]orderlist.ReqOrderHeader, error) {
	var roheaders []orderlist.ReqOrderHeader

	headers, err := s.data.GetRODHeader(ctx, sNumber)
	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetRODetHeader]")
	}

	roheaders = headers
	return roheaders, err
}

func (s Service) GetRODetDetail(ctx context.Context, sNumber string) ([]orderlist.ReqOrderDetail, error) {
	var rodetails []orderlist.ReqOrderDetail

	details, err := s.data.GetRODDetail(ctx, sNumber)
	if err != nil {
		return details, errors.Wrap(err, "[SERVICE][GetRODetDetail]")
	}

	rodetails = details
	return rodetails, err
}

func (s Service) GetRODetProcod(ctx context.Context, sNumber string) ([]orderlist.ROProcode, error) {
	var roprocods []orderlist.ROProcode

	procods, err := s.data.GetRODProcod(ctx, sNumber)
	if err != nil {
		return procods, errors.Wrap(err, "[SERVICE][GetRODetProcod]")
	}

	roprocods = procods
	return roprocods, err
}
