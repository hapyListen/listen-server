package song

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SongModel = (*customSongModel)(nil)

type (
	// SongModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSongModel.
	SongModel interface {
		songModel
		withSession(session sqlx.Session) SongModel
	}

	customSongModel struct {
		*defaultSongModel
	}
)

// NewSongModel returns a model for the database table.
func NewSongModel(conn sqlx.SqlConn) SongModel {
	return &customSongModel{
		defaultSongModel: newSongModel(conn),
	}
}

func (m *customSongModel) withSession(session sqlx.Session) SongModel {
	return NewSongModel(sqlx.NewSqlConnFromSession(session))
}
