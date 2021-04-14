package postgres

import (
	"context"
	"fmt"
	"github.com/d-sense/go-scalable-rest-api-example/config/flags"
	"github.com/jackc/pgx"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type Client struct {
	conn *pgx.Conn
	db   *gorm.DB
}

// New is a postgress database constructor
func New(
	ctx context.Context,
	cfg flags.Configuration,
) *Client {
	url := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Database)

	var client *Client

	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatalf("could not create postgres connection, err=%v ", err)
	}

	client = &Client{db: db}

	return client
}

func (c *Client) Close() error {
	return c.db.Close()
}
