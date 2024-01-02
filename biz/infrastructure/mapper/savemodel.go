package mapper

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"SEproject/biz/infrastructure/config"
)

var _ SaveModel = (*customSaveModel)(nil)

type (
	// SaveModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSaveModel.
	SaveModel interface {
		saveModel
	}

	customSaveModel struct {
		*defaultSaveModel
	}
)

// NewSaveModel returns a model for the database table.
func NewSaveModel(c *config.Config) SaveModel {
	conn := sqlx.NewMysql(c.MySql.URL)
	return &customSaveModel{
		defaultSaveModel: newSaveModel(conn),
	}
}
