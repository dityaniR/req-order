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
func (s Service) GetROdHeader(ctx context.Context, sNumber string) (orderlist.ListDetailsRO, error) {
	var (
		detheader  orderlist.ReqOrderHeader
		detdetails []orderlist.ReqOrderDetail
		detprocod  []orderlist.ROProcode

		result orderlist.ListDetailsRO
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

	detprocod, err = s.data.GetRODProcod(ctx, sNumber)
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GetDetailPerNumber]")
	}

	result.ReqOrderHeader = detheader
	result.ReqOrderDetail = detdetails
	result.ROProcode = detprocod

	return result, nil
}

//pisah
// func (s Service) GetRODetDetail(ctx context.Context, sNumber string) ([]orderlist.ReqOrderDetail, error) {
// 	// var rodetails []orderlist.ReqOrderDetail

// 	// details, err := s.data.GetRODDetail(ctx, sNumber)
// 	// if err != nil {
// 	// 	return details, errors.Wrap(err, "[SERVICE][GetRODetDetail]")
// 	// }

// 	// rodetails = details
// 	// return rodetails, err
// 	rodetails, err := s.data.GetRODDetail(ctx, sNumber)
// 	if err != nil {
// 		return rodetails, errors.Wrap(err, "[SERVICE][GetRODDetail]")
// 	}
// 	return rodetails, nil
// }

// func (s Service) GetRODetProcod(ctx context.Context, sNumber string) ([]orderlist.ROProcode, error) {
// 	var roprocods []orderlist.ROProcode

// 	procods, err := s.data.GetRODProcod(ctx, sNumber)
// 	if err != nil {
// 		return procods, errors.Wrap(err, "[SERVICE][GetRODetProcod]")
// 	}

// 	roprocods = procods
// 	return roprocods, err
// }
