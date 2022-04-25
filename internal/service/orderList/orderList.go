package orderList

import (
	"context"
	orderlist "request-order/internal/entity/orderList"
	"request-order/pkg/errors"
)

func (s Service) GetRO(ctx context.Context) ([]orderlist.ReqOrderHeader, error) {
	var roheaders []orderlist.ReqOrderHeader

	headers, err := s.data.GetROHeader(ctx)
	if err != nil {
		return headers, errors.Wrap(err, "[SERVICE][GetROHeader]")
	}

	roheaders = headers
	return roheaders, err
}

//terakhir pakai ini yg header
// func (s Service) GetROdHeader(ctx context.Context, sNumber string) (orderlist.ReqOrderHeader, error) {
// 	detheader, err := s.data.GetRODHeader(ctx, sNumber)
// 	println("service : " + sNumber)
// 	if err != nil {
// 		return detheader, errors.Wrap(err, "[SERVICE][GetHeaderPerNumber]")
// 	}

// 	return detheader, nil
// }

//trail di gabung
func (s Service) GetROdHeader(ctx context.Context, sNumber string) (orderlist.ReqOrderHeader, error) {
	var (
		detheader  orderlist.ReqOrderHeader
		detdetails []orderlist.ReqOrderDetail

		result orderlist.ReqOrderHeader
		err    error
	)

	detheader, err = s.data.GetRODHeader(ctx, sNumber)
	println("service : " + sNumber)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GetHeaderPerNumber]")
	}

	detdetails, err = s.data.GetRODDetail(ctx, sNumber)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GetDetailPerNumber]")
	}

	detheader.ReqDetail = detdetails

	result = detheader
	return result, nil
}

func (s Service) GetROProcod(ctx context.Context, sProcod string) (orderlist.ROProcode, error) {
	var (
		result orderlist.ROProcode
		err    error
	)
	result, err = s.data.GetRODProcod(ctx, sProcod)
	println("service procod: " + sProcod)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GetROProcod]")
	}
	return result, nil
}

func (s Service) GetROProcods(ctx context.Context) ([]orderlist.ROProcode, error) {
	var rprocodes []orderlist.ROProcode

	procods, err := s.data.GetRODProcods(ctx)
	if err != nil {
		return procods, errors.Wrap(err, "[SERVICE][GetROHeader]")
	}

	rprocodes = procods
	return rprocodes, err
}