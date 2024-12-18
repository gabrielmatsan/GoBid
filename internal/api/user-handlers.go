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
	}

	_ = utils.EncodeJson(w, r, http.StatusCreated, map[string]any{"user_id": id})

}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("TO DO - NOT IMPLEMENTED")
}

func (api *Api) handleLogOutUser(w http.ResponseWriter, r *http.Request) {
	panic("TO DO - NOT IMPLEMENTED")
}
