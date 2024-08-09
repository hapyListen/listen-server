package songlist

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"listen-server/app/user/cmd/user/internal/logic/songlist"
	"listen-server/app/user/cmd/user/internal/svc"
	"listen-server/app/user/cmd/user/internal/types"
)

// 删除歌单
func DeleteSonglistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteSonglistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := songlist.NewDeleteSonglistLogic(r.Context(), svcCtx)
		resp, err := l.DeleteSonglist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
