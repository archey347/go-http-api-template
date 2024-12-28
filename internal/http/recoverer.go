package http

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Recoverer struct {
	log *slog.Logger
}

func (rec *Recoverer) Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()

			if err != nil {
				rec.log.Error("panic", "error", fmt.Sprintf("%v", err))
				WriteErrorResponse(w, http.StatusInternalServerError, "internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
