package api

import (
	"net/http"

	"github.com/gabrielmatsan/GoBid/internal/utils"
	"github.com/gorilla/csrf"
)

func (api *Api) HandlerGetCsrfToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	utils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"csrf_token": token,
	})
}

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			utils.EncodeJson(w, r, http.StatusUnauthorized, map[string]interface{}{"error": "Must be logged in"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
