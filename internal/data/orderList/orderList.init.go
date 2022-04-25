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
	qGetROHeader = `select Req_Number,Req_ConfirmYN,Req_Date,Req_ConfirmBy,Req_TotalNettPrice
						From TH_ReqProd`

	getROHDetails  = "GetRODHeader"
	qGetROHDetails = `select Req_Number,Req_ConfirmYN,Req_Date,Req_ConfirmBy,Req_TotalNettPrice
						From TH_ReqProd 
						where Req_Number = ?`

	getRODDetails = "GetRODDetail"
	qGetRDDetails = `select a.Req_Number,Req_ProdCode,RO_Name,Req_Qty,Req_OrderLimit,Remain,
						RO_NettPrice,Req_TotalNettPrice
					from Td_ReqProd a
					left join TH_ReqProd b
					on a.Req_Number = b.Req_Number
					left join (select a.* from RO_RealProses a
								inner join (select max(RO_Date) RO_Date,RO_Procod 
											from RO_RealProses
											group by RO_Procod) b
								on a.RO_Date = b.RO_Date
								and a.RO_Procod = b.RO_Procod) c
					on RO_Procod = Req_ProdCode
					where a.Req_Number = ?`

	getROProcods  = "GetRODProcods"
	qGetROProcods = `select a.RO_Procod,a.RO_Name,a.RO_QtyOrder,a.RO_MaxOrder,
						a.RO_Stock,a.RO_Hold,a.KetOr,a.OrderPackName,
						a.RO_NettPrice,a.RO_NettPriceTotal
					from RO_RealProses a
					inner join (select max(RO_Date) RO_Date,RO_Procod 
								from RO_RealProses
								group by RO_Procod) b
					on a.RO_Date = b.RO_Date
					and a.RO_Procod = b.RO_Procod`

	getROProcod  = "GetRODProcod"
	qGetROProcod = `select a.RO_Procod,a.RO_Name,a.RO_QtyOrder,a.RO_MaxOrder,
						a.RO_Stock,a.RO_Hold,a.KetOr,a.OrderPackName,
						a.RO_NettPrice,a.RO_NettPriceTotal
					from RO_RealProses a
					inner join (select max(RO_Date) RO_Date,RO_Procod 
								from RO_RealProses
								group by RO_Procod) b
					on a.RO_Date = b.RO_Date
					and a.RO_Procod = b.RO_Procod
					where a.RO_Procod =?`
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
