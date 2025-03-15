package repo

import (
	"authentication/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/url"
)

type Connection struct {
	*pgx.Conn
}

type Repository interface {
	Close(ctx context.Context) error
}

func dsnParser(dbConfig *utils.DatabaseConfig) string {
	query := url.Values{}

	for key, value := range dbConfig.Params {
		query.Add(key, value)
	}

	baseUrl := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(dbConfig.Username, dbConfig.Password),
		Host:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		Path:     dbConfig.DBName,
		RawQuery: query.Encode(),
	}

	return baseUrl.String()
}

func newConnection(dbConfig *utils.DatabaseConfig, ctx context.Context) *Connection {
	connection, err := pgx.Connect(ctx, dsnParser(dbConfig))
	if err != nil {
		panic(err)
	}

	return &Connection{connection}
}
