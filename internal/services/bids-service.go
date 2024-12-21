package services

import (
	"context"
	"errors"

	"github.com/gabrielmatsan/GoBid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BidsService struct {
	pool    *pgxpool.Pool
	queries pgstore.Queries
}

var ErrBidIsLowerThanAnotherBid = errors.New("this bid is lower than another bid")

func NewBidsService(pool *pgxpool.Pool) BidsService {
	return BidsService{
		pool:    pool,
		queries: *pgstore.New(pool),
	}
}

func (bs *BidsService) PlaceBid(ctx context.Context, product_id, bidder_id uuid.UUID, amount float64) (pgstore.Bid, error) {
	// Amount > previous bid amount
	// First Amount > price

	product, err := bs.queries.GetProductByID(ctx, product_id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Bid{}, err
		}
	}

	highestBid, err := bs.queries.GetHighestBidByProductId(ctx, product_id)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Bid{}, err
		}
	}

	if product.Price >= amount || highestBid.BidAmount >= amount {
		return pgstore.Bid{}, ErrBidIsLowerThanAnotherBid
	}

	highestBid, err = bs.queries.CreateBid(ctx, pgstore.CreateBidParams{
		ProductID: product_id,
		BidderID:  bidder_id,
		BidAmount: amount,
	})

	if err != nil {
		return pgstore.Bid{}, err

	}

	return highestBid, nil
}
