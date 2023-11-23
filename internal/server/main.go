package server

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/s992/logger/internal/generated/db"
	"github.com/s992/logger/internal/generated/proto/logger/v1/loggerv1connect"
)

type ServerConfig struct {
	ClientFiles embed.FS
	Port        int
	Queries     *db.Queries
}

func Run(config *ServerConfig) error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))

	tagService := NewTagService(config.Queries)
	tagSvcPath, tagSvcHandler := loggerv1connect.NewTagServiceHandler(tagService)
	r.Mount(tagSvcPath, tagSvcHandler)

	logLineService := NewLogLineService(config.Queries)
	logLineSvcPath, logLineSvcHandler := loggerv1connect.NewLogLineServiceHandler(logLineService)
	r.Mount(logLineSvcPath, logLineSvcHandler)

	r.Handle("/*", SPAHandler(config.ClientFiles))

	addr := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Server running at %s\n", addr)

	return http.ListenAndServe(addr, r)
}
