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
	qGetRDDetails = `select distinct Req_ProdCode,RO_Name,Req_Qty,ifnull(RO_OrderUnit,1) RO_OrderUnit,
						SellPackName,RO_Stock,KetOr,RO_Hold,Req_OrderLimit,Remain,
						(ifnull(RO_SellPrice,0)*0.75)/ifnull(MedPack,0) RO_NettPrice,
						(ifnull(RO_SellPrice,0)*0.75)/ifnull(MedPack,0) * ifnull(Req_Qty,0) RO_TotalNettPrice,
						RO_MinOrder,RO_MaxOrder,RO_LocalProcod,Req_UserID 
					from td_reqprod
					left join ro_realproses
					on ro_number = Req_Number 
						and ro_procod = Req_ProdCode 
					where req_number = ?`

	getROProcods  = "GetRODProcods"
	qGetROProcods = `select distinct RO_Procod Req_ProdCode,RO_Name,
						(ifnull(RO_SellPrice,0)*0.75)/ifnull(MedPack,0) RO_NettPrice
					from ro_realproses
					where RO_Number is not null   #karena belum ada perhitungan
					group by RO_Procod,RO_Name
					order by RO_Procod,RO_Name desc `

	//belumada perhitungan
	getROProcod  = "GetRODProcod"
	qGetROProcod = `select distinct RO_Procod Req_ProdCode,RO_Name,RO_QtyOrder Req_Qty,ifnull(RO_OrderUnit,1) RO_OrderUnit,
						SellPackName,RO_Stock,KetOr,RO_Hold,RO_Limit Req_OrderLimit,Remain,
						(ifnull(RO_SellPrice,0)*0.75)/ifnull(MedPack,0) RO_NettPrice,
						(ifnull(RO_SellPrice,0)*0.75)/ifnull(MedPack,0) * ifnull(RO_QtyOrder,0) RO_TotalNettPrice,
						RO_MinOrder,RO_MaxOrder,RO_LocalProcod,'' Req_UserID 
					from ro_realproses
					where ro_procod = ?
					order by RO_Date desc 
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
