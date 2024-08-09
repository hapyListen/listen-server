package songlist

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"listen-server/app/user/cmd/api/user/internal/logic/songlist"
	"listen-server/app/user/cmd/api/user/internal/svc"
	"listen-server/app/user/cmd/api/user/internal/types"
)

// 获取歌单
func GetSongListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSongListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := songlist.NewGetSongListLogic(r.Context(), svcCtx)
		resp, err := l.GetSongList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
