// Code generated by goctl. DO NOT EDIT.

package mapper

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	saveFieldNames          = builder.RawFieldNames(&Save{})
	saveRows                = strings.Join(saveFieldNames, ",")
	saveRowsExpectAutoSet   = strings.Join(stringx.Remove(saveFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	saveRowsWithPlaceHolder = strings.Join(stringx.Remove(saveFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	saveModel interface {
		Insert(ctx context.Context, data *Save) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Save, error)
		Update(ctx context.Context, data *Save) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSaveModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Save struct {
		Id  int64          `db:"id"`
		Db1 sql.NullString `db:"db1"`
		Db2 sql.NullString `db:"db2"`
		Db3 sql.NullString `db:"db3"`
		Db4 sql.NullString `db:"db4"`
		Db5 sql.NullString `db:"db5"`
	}
)

func newSaveModel(conn sqlx.SqlConn) *defaultSaveModel {
	return &defaultSaveModel{
		conn:  conn,
		table: "`save`",
	}
}

func (m *defaultSaveModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSaveModel) FindOne(ctx context.Context, id int64) (*Save, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", saveRows, m.table)
	var resp Save
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSaveModel) Insert(ctx context.Context, data *Save) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, saveRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Db1, data.Db2, data.Db3, data.Db4, data.Db5)
	return ret, err
}

func (m *defaultSaveModel) Update(ctx context.Context, data *Save) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, saveRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Db1, data.Db2, data.Db3, data.Db4, data.Db5, data.Id)
	return err
}

func (m *defaultSaveModel) tableName() string {
	return m.table
}
