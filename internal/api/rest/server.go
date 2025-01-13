package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lallison21/library_rest_service/internal/config/config"
	"net/http"
)

type Rest struct {
	Server *http.Server
}

func New(cfg *config.Server) *Rest {
	gin.SetMode(cfg.GinMode)

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	server := &http.Server{
		Addr: address,
	}

	return &Rest{
		Server: server,
	}
}

func (r *Rest) RegisterAuthHandler() {
	router := gin.New()
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		panic(err)
	}

	router.Group("/auth")

	r.Server.Handler = router
}
