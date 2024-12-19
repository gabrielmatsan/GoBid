package api

import (
	"errors"
	"net/http"

	"github.com/gabrielmatsan/GoBid/internal/services"
	"github.com/gabrielmatsan/GoBid/internal/usecase/user"
	"github.com/gabrielmatsan/GoBid/internal/utils"
)

func (api *Api) handleSignUpUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[user.CreateUserRequest](r)

	if err != nil {
		_ = utils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(), data.Username, data.Email, data.Password, data.Bio)

	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUsername) {
			_ = utils.EncodeJson(w, r, http.StatusConflict, map[string]any{"error": "Duplicated email or username"})
			return
		}

		_ = utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "Internal server error"})
		return
	}

	_ = utils.EncodeJson(w, r, http.StatusCreated, map[string]any{"user_id": id})

}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[user.LoginUserRequest](r)

	if err != nil {
		_ = utils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)

	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			_ = utils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{"error": "Invalid credentials"})
			return
		}

		_ = utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "Internal server error"})
		return
	}

	err = api.Sessions.RenewToken(r.Context())

	if err != nil {
		_ = utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "Internal server error"})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id)

	utils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "Login successfuly"})

}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())

	if err != nil {
		_ = utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "Internal server error"})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")

	utils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "logged out successfuly"})
}
