package boot

import (
	"log"
	"net/http"
	"request-order/docs"
	"request-order/pkg/httpclient"

	"request-order/internal/config"
	"request-order/internal/viper"

	"github.com/jmoiron/sqlx"

	server "request-order/internal/delivery/http"

	authData "request-order/internal/data/auth"
	authService "request-order/internal/service/auth"

	skeletonData "request-order/internal/data/skeleton"
	skeletonHandler "request-order/internal/delivery/http/skeleton"
	skeletonService "request-order/internal/service/skeleton"

	orderData "request-order/internal/data/orderList"
	orderHandler "request-order/internal/delivery/http/orderList"
	orderService "request-order/internal/service/orderList"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG1] Failed to initialize config: %v", err)
	}

	errViper := viper.Init()
	if errViper != nil {
		log.Fatalf("[CONFIG2] Failed to initialize config: %v", errViper)
	}

	v := viper.GetConn()
	db := v.ServerRO

	cfg := config.Get()
	// Open MySQL DB Connection
	db, err = sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	//
	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	httpc := httpclient.NewClient()
	ad := authData.New(httpc, cfg.API.Auth)
	as := authService.New(ad)

	// Diganti dengan domain yang anda buat
	sd := skeletonData.New(db)
	ss := skeletonService.New(sd, as)
	sh := skeletonHandler.New(ss)

	od := orderData.New(db)
	os := orderService.New(od)
	oh := orderHandler.New(os)

	s := server.Server{
		Skeleton: sh,
		Orders:   oh,
	}

	viper.ViperActive()

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
