package repo

import (
	"authentication/repo/db"
	"authentication/utils"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type AdminRepo interface {
	GetAdminAuthByEmail(email string, ctxOptional ...context.Context) (db.AdminAuth, error)
	CreateModeratorAuth(email, password string, moderatorType db.ModeratorType, ctxOptional ...context.Context) error
}

type adminRepo struct {
	connection *Connection
	queries    *db.Queries
}

func NewAdminRepository() AdminRepo {
	connection := newConnection(utils.GetDatabaseConfig(), context.Background())
	return &adminRepo{
		connection: connection,
		queries:    db.New(connection),
	}
}

func (repo *adminRepo) Close(ctx context.Context) error {
	return repo.connection.Close(ctx)
}

func (repo *adminRepo) GetAdminAuthByEmail(email string, ctxOptional ...context.Context) (db.AdminAuth, error) {
	ctx := utils.FirstContextOrBackground(ctxOptional)
	return repo.queries.GetAdminAuthByEmail(ctx, email)
}

func (repo *adminRepo) CreateModeratorAuth(email, password string, moderatorType db.ModeratorType, ctxOptional ...context.Context) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), utils.GetHashingCost())
	if err != nil {
		return err
	}

	ctx := utils.FirstContextOrBackground(ctxOptional)
	return repo.queries.CreateModeratorAuth(ctx, db.CreateModeratorAuthParams{
		Email:    email,
		Password: hashedPassword,
		Type:     moderatorType,
	})
}
