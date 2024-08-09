package playlist

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PlaylistModel = (*customPlaylistModel)(nil)

type (
	// PlaylistModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlaylistModel.
	PlaylistModel interface {
		playlistModel
		withSession(session sqlx.Session) PlaylistModel
	}

	customPlaylistModel struct {
		*defaultPlaylistModel
	}
)

// NewPlaylistModel returns a model for the database table.
func NewPlaylistModel(conn sqlx.SqlConn) PlaylistModel {
	return &customPlaylistModel{
		defaultPlaylistModel: newPlaylistModel(conn),
	}
}

func (m *customPlaylistModel) withSession(session sqlx.Session) PlaylistModel {
	return NewPlaylistModel(sqlx.NewSqlConnFromSession(session))
}
