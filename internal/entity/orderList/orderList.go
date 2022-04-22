package orderlist

import "time"

// List Order
type Orders struct {
	OutCode        string `db:"Req_OutCode" json:"outCode"`
	Number         string `db:"Req_Number" json:"number"`
	Date           string `db:"Req_Date" json:"date"`
	ConfirmYN      string `db:"Req_ConfirmYN" json:"confirmYN"`
	Printed        int    `db:"Req_Printed" json:"printed"`
	TotalNetPrice  int    `db:"Req_TotalNetPrice" json:"totalNetPrice"`
	TotalQtyDetail int    `db:"Req_TotalQtyDetail" json:"totalQtyDetail"`
	ConfirmBy      int    `db:"Req_ConfirmBy" json:"confirmBy"`
	LastUpdate     string `db:"Req_LastUpdate" json:"lastUpdate"`
	UserID         int    `db:"Req_UserID" json:"userID"`
	RecordNum      int    `db:"Req_RecordNum" json:"recordNum"`
	ActiveYN       string `db:"Req_ActiveYN" json:"activeYN"`
	Type           string `db:"Req_Type" json:"type"`
}

// List Req Orders Header
type ReqOrderHeader struct {
	Number         string `db:"Req_Number" json:"number"`
	Date           string `db:"Req_Date" json:"date"`
	ConfirmYN      string `db:"Req_ConfirmYN" json:"confirmYN"`
	Printed        int    `db:"Req_Printed" json:"printed"`
	TotalNetPrice  int    `db:"Req_TotalNetPrice" json:"totalNetPrice"`
	TotalQtyDetail int    `db:"Req_TotalQtyDetail" json:"totalQtyDetail"`
	ConfirmBy      int    `db:"Req_ConfirmBy" json:"confirmBy"`
	LastUpdate     string `db:"Req_LastUpdate" json:"lastUpdate"`
	UserID         int    `db:"Req_UserID" json:"userID"`
	RecordNum      int    `db:"Req_RecordNum" json:"recordNum"`
	ActiveYN       string `db:"Req_ActiveYN" json:"activeYN"`
	Type           string `db:"Req_Type" json:"type"`
}

// List Req Orders Detail
type ReqOrderDetail struct {
	Number           string    `db:"Req_Number" json:"number"`
	ReqProdCode      string    `db:"Req_ProdCode" json:"ReqProdCode"`
	ReqQty           int       `db:"Req_Qty" json:"ReqQty"`
	ReqOrderLimit    int       `db:"Req_OrderLimit" json:"ReqOrderLimit"`
	ReqFlag          int       `db:"Req_Flag" json:"ReqFlag"`
	ReqHONum         string    `db:"Req_HONum" json:"ReqHONum"`
	ReqHoDate        time.Time `db:"Req_HoDate" json:"ReqHoDate"`
	ReqPurchNum      string    `db:"Req_PurchNum" json:"ReqPurchNum"`
	ReqPurchDate     time.Time `db:"Req_PurchDate" json:"ReqPurchDate"`
	ReqAlocNum       string    `db:"Req_AlocNum" json:"ReqAlocNum"`
	ReqAlocDate      time.Time `db:"Req_AlocDate" json:"ReqAlocDate"`
	ReqTrfNum        string    `db:"Req_TrfNum" json:"ReqTrfNum"`
	ReqTrfDate       time.Time `db:"Req_TrfDate" json:"ReqTrfDate"`
	ReqOutletRcvNum  string    `db:"Req_OutletRcvNum" json:"ReqOutletRcvNum"`
	ReqOutletRcvDate time.Time `db:"Req_OutletRcvDate" json:"ReqOutletRcvDate"`
	ReqCancelDate    time.Time `db:"Req_CancelDate" json:"ReqCancelDate"`
	ReqCancelCode    int       `db:"Req_CancelCode" json:"ReqCancelCode"`
	ReqReorder       int       `db:"Req_Reorder" json:"ReqReorder"`
	ReqLastUpdate    time.Time `db:"Req_LastUpdate" json:"ReqLastUpdate"`
	ReqUserID        int       `db:"Req_UserID" json:"ReqUserID"`
	ReqActiveYN      string    `db:"Req_ActiveYN" json:"ReqActiveYN"`
	ReqRecYN         string    `db:"Req_RecYN" json:"ReqRecYN"`
	ReqRecQty        int       `db:"Req_RecQty" json:"ReqRecQty"`
	ReqROQty         int       `db:"Req_ProdCode" json:"ReqROQty"`
}

// List RO Procode
type ROProcode struct {
	RODate           time.Time `db:"RO_Date" json:"RODate"`
	ROProcod         string    `db:"RO_Procod" json:"ROProcod"`
	ROName           string    `db:"RO_Name" json:"ROName"`
	ROB0             float64   `db:"RO_B0" json:"ROB0"`
	ROB1             float64   `db:"RO_B1" json:"ROB1"`
	ROB2             float64   `db:"RO_B2" json:"ROB2"`
	ROAverage        float64   `db:"RO_Average" json:"ROAverage"`
	ROLimit          float64   `db:"RO_Limit" json:"ROLimit"`
	ROMedPack        int       `db:"RO_MedPack" json:"ROMedPack"`
	ROSellUnit       int       `db:"RO_SellUnit" json:"ROSellUnit"`
	ROSellPrice      float64   `db:"RO_SellPrice" json:"ROSellPrice"`
	ROQtyOrder       int       `db:"RO_QtyOrder" json:"ROQtyOrder"`
	ROStatus         string    `db:"RO_Status" json:"ROStatus"`
	ROHold           float64   `db:"RO_Hold" json:"ROHold"`
	ROStock          float64   `db:"RO_Stock" json:"ROStock"`
	ROMinOrder       float64   `db:"RO_MinOrder" json:"ROMinOrder"`
	ROMaxOrder       float64   `db:"RO_MaxOrder" json:"ROMaxOrder"`
	RONettPrice      float64   `db:"RO_NettPrice" json:"RONettPrice"`
	RONettPriceTotal float64   `db:"RO_NettPriceTotal" json:"RONettPriceTotal"`
	ROFlag           string    `db:"RO_Flag" json:"ROFlag"`
	ROStatusA        int       `db:"RO_StatusA" json:"ROStatusA"`
	ROFlagA          string    `db:"RO_FlagA" json:"ROFlagA"`
	RONP             string    `db:"RO_NP" json:"RONP"`
	ROSub            string    `db:"RO_Sub" json:"ROSub"`
	ROTampil         int       `db:"RO_Tampil" json:"ROTampil"`
	MedPack          int       `db:"MedPack" json:"MedPack"`
	ProOrderPack     int       `db:"Pro_OrderPack" json:"ProOrderPack"`
	OrderPackName    string    `db:"OrderPackName" json:"OrderPackName"`
	ProSellPack      int       `db:"Pro_SellPack" json:"ProSellPack"`
	SellPackName     string    `db:"SellPackName" json:"SellPackName"`
	Remain           float64   `db:"Remain" json:"Remain"`
	KetOr            string    `db:"KetOr" json:"KetOr"`
	ROQtyOutlet      int       `db:"RO_QtyOutlet" json:"ROQtyOutlet"`
	ROTTT            string    `db:"RO_TTT" json:"ROTTT"`
}

//List Details
type ListDetailsRO struct {
	Header      ReqOrderHeader   `json:"ReqOrderHeader"`
	Details     []ReqOrderDetail `json:"ReqOrderDetail"`
	DetProcodes []ROProcode      `json:"ROProcode"`
}
