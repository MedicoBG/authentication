package repo

import (
	"authentication/repo/db"
	"authentication/utils"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"
)

type ModeratorRepoQuerier interface {
}

type ModeratorRepoMutator interface {
}

type ModeratorRepo struct {
	connection *pgx.Conn
	queries    *db.Queries
}

var (
	_ BaseRepo             = (*ModeratorRepo)(nil)
	_ ModeratorRepoQuerier = (*ModeratorRepo)(nil)
	_ ModeratorRepoMutator = (*ModeratorRepo)(nil)
)

func NewModeratorRepo(ctxOptional ...context.Context) *ModeratorRepo {
	ctx := utils.FirstContextOrBackground(ctxOptional)
	connection := newConnection(utils.GetDatabaseConfig(), ctx)
	return &ModeratorRepo{
		connection: connection,
		queries:    db.New(),
	}
}

func (m *ModeratorRepo) Close(ctxOptional ...context.Context) error {
	ctx := utils.FirstContextOrBackground(ctxOptional)
	return m.connection.Close(ctx)
}
