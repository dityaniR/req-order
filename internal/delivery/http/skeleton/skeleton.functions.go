package skeleton

import (
	"log"
	"net/http"
	httpHelper "request-order/internal/delivery/http"
	"request-order/pkg/response"
)

// GetSkeleton godoc
// @Summary Get entries of all skeletons
// @Description Get entries of all skeletons
// @Tags Skeleton
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200
// @Router /v1/profiles [get]
func (h *Handler) GetSkeleton(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	err := h.skeletonSvc.GetSkeleton(ctx)
	//

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		//
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
