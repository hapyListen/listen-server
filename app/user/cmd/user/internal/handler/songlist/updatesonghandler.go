package songlist

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"listen-server/app/user/cmd/user/internal/logic/songlist"
	"listen-server/app/user/cmd/user/internal/svc"
	"listen-server/app/user/cmd/user/internal/types"
)

// 更新歌曲
func UpdateSongHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSongReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := songlist.NewUpdateSongLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSong(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
