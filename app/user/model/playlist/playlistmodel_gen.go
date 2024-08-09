// Code generated by goctl. DO NOT EDIT.

package playlist

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	playlistFieldNames          = builder.RawFieldNames(&Playlist{})
	playlistRows                = strings.Join(playlistFieldNames, ",")
	playlistRowsExpectAutoSet   = strings.Join(stringx.Remove(playlistFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	playlistRowsWithPlaceHolder = strings.Join(stringx.Remove(playlistFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	playlistModel interface {
		Insert(ctx context.Context, data *Playlist) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Playlist, error)
		Update(ctx context.Context, data *Playlist) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPlaylistModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Playlist struct {
		Id           int64          `db:"id"`             // 歌单ID,主键
		Name         string         `db:"name"`           // 歌单名称
		LastPlayTime sql.NullTime   `db:"last_play_time"` // 最后一次播放时间
		PlayCount    sql.NullInt64  `db:"play_count"`     // 播放次数
		IsPublic     sql.NullInt64  `db:"is_public"`      // 是否公开
		Avatar       sql.NullString `db:"avatar"`         // 歌曲头像，默认用最后一次收藏的歌曲图片
		CreateTime   sql.NullTime   `db:"create_time"`    // 创建时间
		UpdateTime   sql.NullTime   `db:"update_time"`    // 更新时间
	}
)

func newPlaylistModel(conn sqlx.SqlConn) *defaultPlaylistModel {
	return &defaultPlaylistModel{
		conn:  conn,
		table: "`playlist`",
	}
}

func (m *defaultPlaylistModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPlaylistModel) FindOne(ctx context.Context, id int64) (*Playlist, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", playlistRows, m.table)
	var resp Playlist
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPlaylistModel) Insert(ctx context.Context, data *Playlist) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, playlistRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.LastPlayTime, data.PlayCount, data.IsPublic, data.Avatar)
	return ret, err
}

func (m *defaultPlaylistModel) Update(ctx context.Context, data *Playlist) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, playlistRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.LastPlayTime, data.PlayCount, data.IsPublic, data.Avatar, data.Id)
	return err
}

func (m *defaultPlaylistModel) tableName() string {
	return m.table
}
