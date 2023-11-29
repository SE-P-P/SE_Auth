package mapper

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"SEproject/biz/infrastructure/config"
)

var _ SignInModel = (*customSignInModel)(nil)

type (
	// SignInModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSignInModel.
	SignInModel interface {
		signInModel
		FindByEmail(ctx context.Context, email string) (*SignIn, error)
	}

	customSignInModel struct {
		*defaultSignInModel
	}
)

// NewSignInModel returns a model for the database table.
func NewSignInModel(c *config.Config) SignInModel {
	conn := sqlx.NewMysql(c.MySql.URL)
	return &customSignInModel{
		defaultSignInModel: newSignInModel(conn),
	}
}

func (m *defaultSignInModel) FindByEmail(ctx context.Context, email string) (*SignIn, error) {
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", signInRows, m.table)
	var resp SignIn
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
