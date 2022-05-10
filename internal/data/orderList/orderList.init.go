package orderList

import (
	"context"
	"log"

	"request-order/internal/viper"

	"github.com/jmoiron/sqlx"
)

type (
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	statement struct {
		key   string
		query string
	}
)

func (d *Data) UpdateConn() {
	v := viper.GetConn()
	d.db = v.ServerRO
}

// Query List to Prepare
const (
	getROHeader  = "GetROHeader"
	qGetROHeader = `select Req_Number,Req_ConfirmYN,Req_ConfirmDate,Req_ConfirmBy,Req_TotalNettPrice
						From TH_ReqProd`

	getROHDetails  = "GetRODHeader"
	qGetROHDetails = `select Req_Number,Req_ConfirmYN,Req_ConfirmDate,Req_ConfirmBy,Req_TotalNettPrice
						From TH_ReqProd
						where Req_Number = ?`

	getRODDetails = "GetRODDetail"
	qGetRDDetails = `select distinct Req_ProdCode,TRIM(Req_Name) Req_Name,Req_Qty,ifnull(Req_OrderUnit,1) Req_OrderUnit,
						TRIM(Req_SellPackName) Req_SellPackName,Req_Stock,TRIM(Req_KetOr) Req_KetOr,Req_Hold,Req_OrderLimit,Req_Remain,
						Req_NettPrice,Req_NettPriceTotal,
						Req_MinOrder,Req_MaxOrder,ifnull(Req_LocalProcod,'N') Req_LocalProcod,Req_UserID 
					from td_reqprod
					where req_number = ?`

	//nanti ambil datanya dari mh_product
	getROProcods  = "GetRODProcods"
	qGetROProcods = `select distinct Req_ProdCode,ifnull(Req_Name,'') Req_Name,
						ifnull(Req_NettPrice,0) Req_NettPrice
					from td_reqprod
					order by Req_ProdCode,Req_Name`

	//belumada perhitungan
	getROProcod  = "GetRODProcod"
	qGetROProcod = `#belum ada perhitungan					
					select distinct Req_ProdCode,Req_Name,0 Req_Qty,ifnull(Req_OrderUnit,1) Req_OrderUnit,
						TRIM(Req_SellPackName) Req_SellPackName,Req_Stock,TRIM(Req_KetOr) Req_KetOr,Req_Hold,Req_OrderLimit,Req_Remain,
						Req_NettPrice,0 Req_NettPriceTotal,
						Req_MinOrder,Req_MaxOrder,ifnull(Req_LocalProcod,'N') Req_LocalProcod,'' Req_UserID 
					from td_reqprod
					where Req_ProdCode = ?
					order by Req_LastUpdate desc
					limit 1`
)

var (
	readStmt = []statement{
		{getROHeader, qGetROHeader},
		{getROHDetails, qGetROHDetails},
		{getRODDetails, qGetRDDetails},
		{getROProcods, qGetROProcods},
		{getROProcod, qGetROProcod},
	}
	upsertStmt = []statement{}
	deleteStmt = []statement{}
)

func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range upsertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize upsert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}
