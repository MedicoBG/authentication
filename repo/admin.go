package repo

import (
	"authentication/repo/db"
	"authentication/utils"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type AdminRepoQuerier interface {
	GetAdminAuthByEmail(email string, ctxOptional ...context.Context) (db.AdminAuth, error)
}

type AdminRepoMutator interface {
	CreateModeratorAuth(email, password string, moderatorType db.ModeratorType, ctxOptional ...context.Context) error
}

type AdminRepo struct {
	connection *Connection
	queries    *db.Queries
}

var (
	_ BaseRepo         = (*AdminRepo)(nil)
	_ AdminRepoQuerier = (*AdminRepo)(nil)
	_ AdminRepoMutator = (*AdminRepo)(nil)
)

func NewAdminRepository(ctxOptional ...context.Context) *AdminRepo {
	ctx := utils.FirstContextOrBackground(ctxOptional)
	connection := newConnection(utils.GetDatabaseConfig(), ctx)
	return &AdminRepo{
		connection: connection,
		queries:    db.New(connection),
	}
}

func (a *AdminRepo) Close(ctxOptional ...context.Context) error {
	ctx := utils.FirstContextOrBackground(ctxOptional)
	return a.connection.Close(ctx)
}

func (a *AdminRepo) GetAdminAuthByEmail(email string, ctxOptional ...context.Context) (db.AdminAuth, error) {
	ctx := utils.FirstContextOrBackground(ctxOptional)
	return a.queries.GetAdminAuthByEmail(ctx, email)
}

func (a *AdminRepo) CreateModeratorAuth(email, password string, moderatorType db.ModeratorType, ctxOptional ...context.Context) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), utils.GetHashingCost())
	if err != nil {
		return err
	}

	ctx := utils.FirstContextOrBackground(ctxOptional)
	return a.queries.CreateModeratorAuth(ctx, db.CreateModeratorAuthParams{
		Email:    email,
		Password: hashedPassword,
		Type:     moderatorType,
	})
}
