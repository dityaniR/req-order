package orderList

import (
	"log"
	"net/http"
	httpHelper "request-order/internal/delivery/http"
	"request-order/pkg/response"

	"github.com/gorilla/mux"
)

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.service.GetRO(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) GetRODetHeader(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	params := mux.Vars(r)
	sNum := params["id"]

	// sNum := r.URL.Query()["num"][0]
	result, err = h.service.GetROdHeader(ctx, sNum)
	//result, err = h.service.GetRODetDetail(ctx, sNum)
	println("delivery : " + sNum)

	if err != nil {
		resp.Error = response.Error{
			Status: true,
			Msg:    err.Error(),
			Code:   500,
		}
		return
	}

	resp.Data = result
	resp.Metadata = metadata
}

// func (h *Handler) GetROProCode(w http.ResponseWriter, r *http.Request) {
// 	resp := response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	ctx := r.Context()

// 	sProCode := r.URL.Query()["procode"][0]
// 	result, err := h.service.GetROProcod(ctx, sProCode)
// 	if err != nil {
// 		resp = httpHelper.ParseErrorCode(err.Error())
// 		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
// 		return
// 	}

// 	resp.Data = result
// 	resp.Metadata = "metadata"
// 	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
// }

func (h *Handler) GetROProCode(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	params := mux.Vars(r)
	sProcod := params["procod"]

	// sNum := r.URL.Query()["num"][0]
	result, err = h.service.GetROProcod(ctx, sProcod)
	//result, err = h.service.GetRODetDetail(ctx, sNum)
	println("delivery : " + sProcod)

	if err != nil {
		resp.Error = response.Error{
			Status: true,
			Msg:    err.Error(),
			Code:   500,
		}
		return
	}

	resp.Data = result
	resp.Metadata = metadata
}

func (h *Handler) GetROProCodes(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.service.GetROProcods(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}
