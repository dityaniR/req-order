package http

import (
	"net/http"

	"request-order/pkg/grace"

	"github.com/rs/cors"
)

// SkeletonHandler ...
type SkeletonHandler interface {
	GetSkeleton(w http.ResponseWriter, r *http.Request)
}

type OrderHandler interface {
	GetOrder(w http.ResponseWriter, r *http.Request)
	GetRODetHeader(w http.ResponseWriter, r *http.Request)
	//GetRODetDetail(w http.ResponseWriter, r *http.Request)
	GetROProCode(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server   *http.Server
	Skeleton SkeletonHandler
	Orders   OrderHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
