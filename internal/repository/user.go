package repository

import (
	"context"
	"database/sql"
	"go-fiber-postgre/domain"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

// FindByEmail implements [domain.UserRepository].
func (u *userRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &user)
	return 
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}
