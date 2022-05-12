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
	Number         null.String      `db:"Req_Number" json:"Number"`
	ConfirmYN      null.String      `db:"Req_ConfirmYN" json:"ConfirmYN"`
	Date           null.String      `db:"Req_ConfirmDate" json:"ConfirmDate"`
	ConfirmBy      null.Int         `db:"Req_ConfirmBy" json:"ConfirmBy"`
	TotalNettPrice null.Float       `db:"Req_TotalNettPrice" json:"TotalNettPrice"`
	ReqDetail      []ReqOrderDetail `json:"ReqDetail"`
}

// List Req Orders Detail
type ReqOrderDetail struct {
	ReqProdCode       null.String `db:"Req_ProdCode" json:"ReqProdCode"`
	ReqName           null.String `db:"Req_Name" json:"ReqName"`
	ReqQty            null.Float  `db:"Req_Qty" json:"ReqQty"`
	ReqOrderUnit      null.Int    `db:"Req_OrderUnit" json:"ReqOrderUnit"`
	ReqSellPackName   null.String `db:"Req_SellPackName" json:"ReqSellPackName"`
	ReqStock          null.Float  `db:"Req_Stock" json:"ReqStock"`
	ReqKetOr          null.String `db:"Req_KetOr" json:"ReqKetOr"`
	ReqHold           null.Int    `db:"Req_Hold" json:"ReqHold"`
	ReqOrderLimit     null.Float  `db:"Req_OrderLimit" json:"ReqOrderLimit"`
	ReqRemain         null.Float  `db:"Req_Remain" json:"ReqRemain"`
	ReqNettPrice      null.Float  `db:"Req_NettPrice" json:"ReqNettPrice"`
	ReqNettPriceTotal null.Float  `db:"Req_NettPriceTotal" json:"ReqNettPriceTotal"`
	ReqMinOrder       null.Float  `db:"Req_MinOrder" json:"ReqMinOrder"`
	ReqMaxOrder       null.Float  `db:"Req_MaxOrder" json:"ReqMaxOrder"`
	ReqLocalProcod    null.String `db:"Req_LocalProcod" json:"ReqLocalProcod"`
	ReqPurchNum       null.String `db:"Req_PurchNum" json:"ReqPurchNum"`
	ReqPurchDate      null.String `db:"Req_PurchDate" json:"ReqPurchDate"`
	ReqAlocNum        null.String `db:"Req_AlocNum" json:"ReqAlocNum"`
	ReqAlocDate       null.String `db:"Req_AlocDate" json:"ReqAlocDate"`
	ReqTrfNum         null.String `db:"Req_TrfNum" json:"ReqTrfNum"`
	ReqTrfDate        null.String `db:"Req_TrfDate" json:"ReqTrfDate"`
	ReqOutletRcvNum   null.String `db:"Req_OutletRcvNum" json:"ReqOutletRcvNum"`
	ReqOutletRcvDate  null.String `db:"Req_OutletRcvDate" json:"ReqOutletRcvDate"`
	ReqCancelCode     null.String `db:"Req_CancelCode" json:"ReqCancelCode"`
	ReqCancelDate     null.String `db:"Req_CancelDate" json:"ReqCancelDate"`
	ReqUserID         null.String `db:"Req_UserID" json:"ReqUserID"`
	ReqLastUpdate     null.String `db:"Req_LastUpdate" json:"ReqLastUpdate"`
}

// List RO Procode
type ROProcode struct {
	ProdCode  null.String `db:"Req_ProdCode" json:"ProdCode"`
	Name      null.String `db:"Req_Name" json:"Name"`
	NettPrice null.Float  `db:"Req_NettPrice" json:"NettPrice"`
}

// //List Details
// type ListDetailsRO struct {
// 	ReqOrderHeader interface{} `json:"ReqOrderHeader"`
// 	ReqOrderDetail interface{} `json:"ReqOrderDetail"`
// 	ROProcode      interface{} `json:"ROProcode"`
// }
