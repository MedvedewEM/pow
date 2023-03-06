package app

import (
	"encoding/json"
	"net/http"

	"github.com/MedvedewEM/pow/internal/components"
	"github.com/MedvedewEM/pow/pkg/api"
	"github.com/gorilla/mux"
)

func NewHandler(
	auth mux.MiddlewareFunc,
	challenge components.Challenger,
	wisdom components.Wisdomer,
) http.Handler {
	router := mux.NewRouter()
	router.Use(ContentTypeJSON)

	router.HandleFunc("/api/v1/please", challengePleaseHandler(challenge))

	authRouter := router.PathPrefix("/api/v1").Subrouter()
	authRouter.Use(auth)
	authRouter.HandleFunc("/wisdom/word", wisdomWordHandler(wisdom))

	return router
}

func challengePleaseHandler(challenge components.Challenger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var res interface{}

		id, suffix, err := challenge.TryGenerate()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			res = api.InternalError{
				Error:       err.Error(),
				Description: "Can not generate challenge, please try again later...",
			}
		} else {
			res = api.ChallengePleaseResponse{
				ID:     id,
				Suffix: suffix,
			}
		}

		json.NewEncoder(w).Encode(res)
	}
}

func wisdomWordHandler(wisdom components.Wisdomer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word := wisdom.Word()

		res := api.WisdomWordResponse{
			Word: word,
		}

		json.NewEncoder(w).Encode(res)
	}
}
