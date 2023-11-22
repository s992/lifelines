package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/s992/logger/internal/generated/db"
	"github.com/s992/logger/internal/generated/proto/logger/v1/loggerv1connect"
)

type ServerConfig struct {
	Port    int
	Queries *db.Queries
}

func Run(config *ServerConfig) error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tagService := NewTagService(config.Queries)
	tagSvcPath, tagSvcHandler := loggerv1connect.NewTagServiceHandler(tagService)
	r.Mount(tagSvcPath, tagSvcHandler)

	logLineService := NewLogLineService(config.Queries)
	logLineSvcPath, logLineSvcHandler := loggerv1connect.NewLogLineServiceHandler(logLineService)
	r.Mount(logLineSvcPath, logLineSvcHandler)

	addr := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Server running at %s\n", addr)

	return http.ListenAndServe(addr, r)
}
