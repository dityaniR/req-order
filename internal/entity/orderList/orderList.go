package orderlist

import (
	//"time"

	"gopkg.in/guregu/null.v3"
)

// List Order
// type Orders struct {
// 	OutCode        string `db:"Req_OutCode" json:"outCode"`
// 	Number         string `db:"Req_Number" json:"number"`
// 	Date           string `db:"Req_Date" json:"date"`
// 	ConfirmYN      string `db:"Req_ConfirmYN" json:"confirmYN"`
// 	Printed        int    `db:"Req_Printed" json:"printed"`
// 	TotalNetPrice  int    `db:"Req_TotalNetPrice" json:"totalNetPrice"`
// 	TotalQtyDetail int    `db:"Req_TotalQtyDetail" json:"totalQtyDetail"`
// 	ConfirmBy      int    `db:"Req_ConfirmBy" json:"confirmBy"`
// 	LastUpdate     string `db:"Req_LastUpdate" json:"lastUpdate"`
// 	UserID         int    `db:"Req_UserID" json:"userID"`
// 	RecordNum      int    `db:"Req_RecordNum" json:"recordNum"`
// 	ActiveYN       string `db:"Req_ActiveYN" json:"activeYN"`
// 	Type           string `db:"Req_Type" json:"type"`
// }

// List Req Orders Header
type ReqOrderHeader struct {
	Number         null.String `db:"Req_Number" json:"Number"`
	ConfirmYN      null.String `db:"Req_ConfirmYN" json:"ConfirmYN"`
	Date           null.String `db:"Req_Date" json:"Date"`
	ConfirmBy      null.Int    `db:"Req_ConfirmBy" json:"ConfirmBy"`
	TotalNettPrice null.Float  `db:"Req_TotalNettPrice" json:"TotalNettPrice"`
	ReqDetail	   []ReqOrderDetail `json:"ReqDetail"`
}

// List Req Orders Detail
type ReqOrderDetail struct {
	Number           	null.String `db:"Req_Number" json:"number"`
	Req_ProdCode      	null.String `db:"Req_ProdCode" json:"ReqProdCode"`
	RO_Name 			null.String `db:"RO_Name" json:"ROName"`
	ReqQty           	null.Int    `db:"Req_Qty" json:"ReqQty"` 	
	ReqOrderLimit   	null.Int    `db:"Req_OrderLimit" json:"ReqOrderLimit"`
	Remain 				null.Float `db:"Remain" json:"Remain"`
	RO_NettPrice 		null.Float `db:"RO_NettPrice" json:"RONettPrice"`
	Req_TotalNettPrice  null.Float `db:"Req_TotalNettPrice" json:"ReqTotalNettPrice"`
}

// List RO Procode
type ROProcode struct {
	ROProcod         null.String `db:"RO_Procod" json:"ROProcod"`
	ROName           null.String `db:"RO_Name" json:"ROName"`
	ROQtyOrder       null.Int    `db:"RO_QtyOrder" json:"ROQtyOrder"`
	ROMaxOrder       null.Float  `db:"RO_MaxOrder" json:"ROMaxOrder"`
	ROStock          null.Float  `db:"RO_Stock" json:"ROStock"`
	ROHold         	 null.Float  `db:"RO_Hold" json:"ROHold"`
	KetOr            null.String `db:"KetOr" json:"KetOr"`
	OrderPackName    null.String `db:"OrderPackName" json:"OrderPackName"`
	RONettPrice      null.Float  `db:"RO_NettPrice" json:"RONettPrice"`
	RONettPriceTotal null.Float  `db:"RO_NettPriceTotal" json:"RONettPriceTotal"`	
}

// //List Details
// type ListDetailsRO struct {
// 	ReqOrderHeader interface{} `json:"ReqOrderHeader"`
// 	ReqOrderDetail interface{} `json:"ReqOrderDetail"`
// 	ROProcode      interface{} `json:"ROProcode"`
// }
