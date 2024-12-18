package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gabrielmatsan/GoBid/internal/api"
	"github.com/gabrielmatsan/GoBid/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx,
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT")),
	)

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: services.NewUserService(pool),
	}

	api.BindRoutes()

	fmt.Println("Starting server on :5080", os.Getenv("PORT"))

	if err := http.ListenAndServe("localhost:5080", api.Router); err != nil {
		panic(err)
	}
}
