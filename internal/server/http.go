package server

import (
	"github.com/oliver258/eagle/internal/routers"
	"github.com/oliver258/eagle/pkg/app"
	"github.com/oliver258/eagle/pkg/transport/http"
)

// NewHTTPServer creates a HTTP server
func NewHTTPServer(c *app.ServerConfig) *http.Server {
	router := routers.NewRouter()

	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWriteTimeout(c.WriteTimeout),
	)

	srv.Handler = router
	// NOTE: register svc to http server

	return srv
}
