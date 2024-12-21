package api

import (
	"net/http"

	"github.com/gabrielmatsan/GoBid/internal/usecase/product"
	"github.com/gabrielmatsan/GoBid/internal/utils"
	"github.com/google/uuid"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[product.CreateProductRequest](r)

	if err != nil {
		utils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	userId, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)

	if !ok {
		utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "Unexpected Internal Server Error: Try again later"})
		return
	}

	id, err := api.ProductService.CreateProduct(r.Context(), userId, data.ProductName, data.Description, data.Price, data.AuctionEnd)

	if err != nil {
		utils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "Unexpected Internal Server Error: Try again later"})
		return
	}

	utils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"message":    "Product created successfuly",
		"product_id": id,
	})
}
