package main

import (
	"fmt"
	"net/http"

	"github.com/nishankx/RSS-Aggregator/internal/database"
	"github.com/nishankx/RSS-Aggregator/internal/database/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User) 

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, "Auth error: "+err.Error())
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}