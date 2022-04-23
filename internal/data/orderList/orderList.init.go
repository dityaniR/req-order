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
	qGetROHeader = `SELECT roh.* FROM th_reqprod roh`

	getROHDetails  = "GetRODHeader"
	qGetROHDetails = `select * from th_reqprod 
						where Req_Number = ?`

	getRODDetails = "GetRODDetail"
	qGetRDDetails = `select rod.* from td_reqprod rod
						where Req_Number = ?`

	getROProcods  = "GetRODProcod"
	qGetROProcods = `select a.* from ro_realproses a
						inner join (select max(RO_Date) RO_Date, RO_Procod from	ro_realproses 
									group by RO_Procod) b
						on a.RO_Date = b.RO_Date
						and a.RO_Procod = b.RO_Procod
						left join td_reqprod c
						on req_prodcode = a.RO_Procod 
						where Req_Number = ?`
)

var (
	readStmt = []statement{
		{getROHeader, qGetROHeader},
		{getROHDetails, qGetROHDetails},
		{getRODDetails, qGetRDDetails},
		{getROProcods, qGetROProcods},
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
