package app

import (
	"database/sql"
	"go.uber.org/zap"
	"net/http"
)

func NewServer(logger *zap.SugaredLogger, db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	var handler http.Handler = mux

	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infow("request", "method", r.Method, "url", r.URL.String(), "remote", r.RemoteAddr)
		mux.ServeHTTP(w, r)
	})

	addRoutes(logger, db, mux)

	return handler
}
