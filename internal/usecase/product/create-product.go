package product

import (
	"context"
	"time"

	"github.com/gabrielmatsan/GoBid/internal/validator"
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	AuctionEnd  time.Time `json:"auction_end"`
}

func (req CreateProductRequest) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.ProductName), "product_name", "Product name must be provided")

	eval.CheckField(validator.NotBlank(req.Description), "description", "Description must be provided")
	eval.CheckField(validator.MinChars(req.Description, 10) && validator.MaxChars(req.Description, 255), "description", "Description must have a length between 10 and 255 characters")

	eval.CheckField(validator.MinPrice(req.Price), "price", "Price must be greater than 0")

	eval.CheckField(validator.MinAuctionDuration(req.AuctionEnd, time.Now()), "auction_end", "Auction end must be at least 2 hours in the future")

	return eval

}
