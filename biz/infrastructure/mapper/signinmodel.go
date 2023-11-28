package mapper

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"SEproject/biz/infrastructure/config"
)

var _ SignInModel = (*customSignInModel)(nil)

type (
	// SignInModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSignInModel.
	SignInModel interface {
		signInModel
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
