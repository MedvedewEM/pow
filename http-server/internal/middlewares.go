package app

import (
	"net/http"

	"github.com/MedvedewEM/pow/internal/components"
)

func NewAuthMiddleware(ch components.Challenger) authMiddleware {
	return authMiddleware{ch: ch}
}

type authMiddleware struct {
	ch components.Challenger
}

func (amw *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("POWID")
		token := r.Header.Get("POWToken")

		err := amw.ch.TryVerify(id, token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		w.Header().Set("Content-Type", "application/json")
	})
}
