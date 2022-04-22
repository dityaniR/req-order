package orderList

import (
	"log"
	"net/http"
	httpHelper "request-order/internal/delivery/http"
	"request-order/pkg/response"
)

// func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
// 	resp := response.Response{}
// 	defer resp.RenderJSON(w, r)

// 	ctx := r.Context()

// 	result, err := h.service.GetRO(ctx)
// 	if err != nil {
// 		resp = httpHelper.ParseErrorCode(err.Error())
// 		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
// 		return
// 	}

// 	resp.Data = result
// 	resp.Metadata = "metadata"
// 	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
// }

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
