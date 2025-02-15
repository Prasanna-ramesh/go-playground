package main

import (
	"github.com/Prasanna-ramesh/go-playground/custom-terraform-provider/mock_upstream_server/user"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
)

const (
	PORT = ":4000"
)

func main() {
	router := chi.NewRouter()
	user.AddRoutes(router)
	router.Get("/healthy", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("I'm healthy"))
	})

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	logger.Info("Starting HTTP server", "port", PORT)
	if err := http.ListenAndServe(PORT, router); err != nil {
		logger.Error("Failed to start server", "error", err)
	}
}
