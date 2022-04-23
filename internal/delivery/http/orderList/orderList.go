package orderList

import (
	"log"
	"net/http"
	httpHelper "request-order/internal/delivery/http"
	"request-order/pkg/response"
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

	sNum := r.URL.Query()["num"][0]
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
