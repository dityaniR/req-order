package orderList

import (
	"context"
	orderlist "request-order/internal/entity/orderList"
	"request-order/pkg/errors"
)

func (d Data) GetROHeader(ctx context.Context) ([]orderlist.Orders, error) {
	headers := []orderlist.Orders{}

	d.UpdateConn()

	rows, err := d.stmt[getROHeader].QueryxContext(ctx)
	if err != nil {
		return headers, errors.Wrap(err, "[DATA][GetROHeader1]")
	}

	defer rows.Close()

	for rows.Next() {
		header := orderlist.Orders{}
		if err = rows.StructScan(&header); err != nil {
			return headers, errors.Wrap(err, "[DATA][GetROHeader2]")
		}
		headers = append(headers, header)
	}

	return headers, nil
}

func (d Data) GetRODHeader(ctx context.Context, sNumber string) (orderlist.ReqOrderHeader, error) {
	dHeader := orderlist.ReqOrderHeader{}

	d.UpdateConn()

	if err := d.stmt[getROHDetails].QueryRowxContext(ctx, sNumber).StructScan(&dHeader); err != nil {
		return dHeader, errors.Wrap(err, "[DATA][GetRODHeader]")
	}

	println("data : " + sNumber)
	return dHeader, nil
}

func (d Data) GetRODDetail(ctx context.Context, sNumber string) ([]orderlist.ReqOrderDetail, error) {
	dDetail := []orderlist.ReqOrderDetail{}

	d.UpdateConn()

	rows, err := d.stmt[getRODDetails].QueryxContext(ctx, sNumber)
	if err != nil {
		return dDetail, errors.Wrap(err, "[DATA][GetRODDetail1]")
	}

	defer rows.Close()

	for rows.Next() {
		detail := orderlist.ReqOrderDetail{}
		if err = rows.StructScan(&detail); err != nil {
			return dDetail, errors.Wrap(err, "[DATA][GetRODDetail2]")
		}
		dDetail = append(dDetail, detail)
	}

	return dDetail, nil
}

func (d Data) GetRODProcod(ctx context.Context, sNumber string) ([]orderlist.ROProcode, error) {
	dROProcod := []orderlist.ROProcode{}

	d.UpdateConn()

	rows, err := d.stmt[getROProcods].QueryxContext(ctx, sNumber)
	if err != nil {
		return dROProcod, errors.Wrap(err, "[DATA][GetRODProcod1]")
	}

	defer rows.Close()

	for rows.Next() {
		dProcod := orderlist.ROProcode{}
		if err = rows.StructScan(&dProcod); err != nil {
			return dROProcod, errors.Wrap(err, "[DATA][GetRODProcod2]")
		}
		dROProcod = append(dROProcod, dProcod)
	}

	return dROProcod, nil
}
