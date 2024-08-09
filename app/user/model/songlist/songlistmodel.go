package songlist

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SonglistModel = (*customSonglistModel)(nil)

type (
	// SonglistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSonglistModel.
	SonglistModel interface {
		songlistModel
		withSession(session sqlx.Session) SonglistModel
	}

	customSonglistModel struct {
		*defaultSonglistModel
	}
)

// NewSonglistModel returns a model for the database table.
func NewSonglistModel(conn sqlx.SqlConn) SonglistModel {
	return &customSonglistModel{
		defaultSonglistModel: newSonglistModel(conn),
	}
}

func (m *customSonglistModel) withSession(session sqlx.Session) SonglistModel {
	return NewSonglistModel(sqlx.NewSqlConnFromSession(session))
}
