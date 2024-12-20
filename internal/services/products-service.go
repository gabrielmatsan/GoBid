package services

import (
	"context"
	"time"

	"github.com/gabrielmatsan/GoBid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (ps *ProductService) CreateProduct(
	ctx context.Context,
	sellerId uuid.UUID,
	productName,
	description string,
	price float64,
	auctionEnd time.Time,
) (uuid.UUID, error) {

	id, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		SellerID:    sellerId,
		ProductName: productName,
		Description: description,
		Price:       price,
		AuctionEnd:  auctionEnd,
	})

	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}
