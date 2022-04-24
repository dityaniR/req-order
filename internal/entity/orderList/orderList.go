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
	Date           null.String `db:"Req_Date" json:"Date"`
	ConfirmYN      null.String `db:"Req_ConfirmYN" json:"ConfirmYN"`
	Printed        null.Int    `db:"Req_Printed" json:"Printed"`
	TotalNettPrice null.Float  `db:"Req_TotalNettPrice" json:"TotalNettPrice"`
	TotQtyDetail   null.Int    `db:"Req_TotQtyDetail" json:"TotQtyDetail"`
	ConfirmBy      null.Int    `db:"Req_ConfirmBy" json:"ConfirmBy"`
	LastUpdate     null.String `db:"Req_LastUpdate" json:"LastUpdate"`
	UserId         null.Int    `db:"Req_UserId" json:"UserId"`
	RecordNum      null.Int    `db:"Req_RecordNum" json:"RecordNum"`
	ActiveYN       null.String `db:"Req_ActiveYN" json:"ActiveYN"`
	Type           null.String `db:"Req_Type" json:"Type"`
}

// List Req Orders Detail
type ReqOrderDetail struct {
	Number           null.String `db:"Req_Number" json:"number"`
	ReqProdCode      null.String `db:"Req_ProdCode" json:"ReqProdCode"`
	ReqQty           null.Int    `db:"Req_Qty" json:"ReqQty"`
	ReqOrderLimit    null.Int    `db:"Req_OrderLimit" json:"ReqOrderLimit"`
	ReqFlag          null.Int    `db:"Req_Flag" json:"ReqFlag"`
	ReqHONum         null.String `db:"Req_HONum" json:"ReqHONum"`
	ReqHoDate        null.String `db:"Req_HoDate" json:"ReqHoDate"`
	ReqPurchNum      null.String `db:"Req_PurchNum" json:"ReqPurchNum"`
	ReqPurchDate     null.String `db:"Req_PurchDate" json:"ReqPurchDate"`
	ReqAlocNum       null.String `db:"Req_AlocNum" json:"ReqAlocNum"`
	ReqAlocDate      null.String `db:"Req_AlocDate" json:"ReqAlocDate"`
	ReqTrfNum        null.String `db:"Req_TrfNum" json:"ReqTrfNum"`
	ReqTrfDate       null.String `db:"Req_TrfDate" json:"ReqTrfDate"`
	ReqOutletRcvNum  null.String `db:"Req_OutletRcvNum" json:"ReqOutletRcvNum"`
	ReqOutletRcvDate null.String `db:"Req_OutletRcvDate" json:"ReqOutletRcvDate"`
	ReqCancelDate    null.String `db:"Req_CancelDate" json:"ReqCancelDate"`
	ReqCancelCode    null.Int    `db:"Req_CancelCode" json:"ReqCancelCode"`
	ReqReorder       null.Int    `db:"Req_Reorder" json:"ReqReorder"`
	ReqLastUpdate    null.String `db:"Req_LastUpdate" json:"ReqLastUpdate"`
	ReqUserID        null.Int    `db:"Req_UserID" json:"ReqUserID"`
	ReqActiveYN      null.String `db:"Req_ActiveYN" json:"ReqActiveYN"`
	ReqRecYN         null.String `db:"Req_RecYN" json:"ReqRecYN"`
	ReqRecQty        null.Int    `db:"Req_RecQty" json:"ReqRecQty"`
	ReqROQty         null.Int    `db:"Req_ROQty" json:"ReqROQty"`
}

// List RO Procode
type ROProcode struct {
	RODate           null.String `db:"RO_Date" json:"RODate"`
	ROProcod         null.String `db:"RO_Procod" json:"ROProcod"`
	ROName           null.String `db:"RO_Name" json:"ROName"`
	ROB0             null.Float  `db:"RO_B0" json:"ROB0"`
	ROB1             null.Float  `db:"RO_B1" json:"ROB1"`
	ROB2             null.Float  `db:"RO_B2" json:"ROB2"`
	ROAverage        null.Float  `db:"RO_Average" json:"ROAverage"`
	ROLimit          null.Float  `db:"RO_Limit" json:"ROLimit"`
	ROMedPack        null.Int    `db:"RO_MedPack" json:"ROMedPack"`
	ROSellUnit       null.Int    `db:"RO_SellUnit" json:"ROSellUnit"`
	ROSellPrice      null.Float  `db:"RO_SellPrice" json:"ROSellPrice"`
	ROQtyOrder       null.Int    `db:"RO_QtyOrder" json:"ROQtyOrder"`
	ROStatus         null.String `db:"RO_Status" json:"ROStatus"`
	ROHold           null.Float  `db:"RO_Hold" json:"ROHold"`
	ROStock          null.Float  `db:"RO_Stock" json:"ROStock"`
	ROMinOrder       null.Float  `db:"RO_MinOrder" json:"ROMinOrder"`
	ROMaxOrder       null.Float  `db:"RO_MaxOrder" json:"ROMaxOrder"`
	RONettPrice      null.Float  `db:"RO_NettPrice" json:"RONettPrice"`
	RONettPriceTotal null.Float  `db:"RO_NettPriceTotal" json:"RONettPriceTotal"`
	ROFlag           null.String `db:"RO_Flag" json:"ROFlag"`
	ROStatusA        null.Int    `db:"RO_StatusA" json:"ROStatusA"`
	ROFlagA          null.String `db:"RO_FlagA" json:"ROFlagA"`
	RONP             null.String `db:"RO_NP" json:"RONP"`
	ROSub            null.String `db:"RO_Sub" json:"ROSub"`
	ROTampil         null.Int    `db:"RO_Tampil" json:"ROTampil"`
	MedPack          null.Int    `db:"MedPack" json:"MedPack"`
	ProOrderPack     null.Float  `db:"Pro_OrderPack" json:"ProOrderPack"`
	OrderPackName    null.String `db:"OrderPackName" json:"OrderPackName"`
	ProSellPack      null.Int    `db:"Pro_SellPack" json:"ProSellPack"`
	SellPackName     null.String `db:"SellPackName" json:"SellPackName"`
	Remain           null.Float  `db:"Remain" json:"Remain"`
	KetOr            null.String `db:"KetOr" json:"KetOr"`
	ROQtyOutlet      null.Int    `db:"RO_QtyOutlet" json:"ROQtyOutlet"`
	ROTTT            null.String `db:"RO_TTT" json:"ROTTT"`
}

//List Details
type ListDetailsRO struct {
	ReqOrderHeader interface{} `json:"ReqOrderHeader"`
	ReqOrderDetail interface{} `json:"ReqOrderDetail"`
	ROProcode      interface{} `json:"ROProcode"`
}
