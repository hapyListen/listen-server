package songlist

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"listen-server/app/user/cmd/user/internal/logic/songlist"
	"listen-server/app/user/cmd/user/internal/svc"
	"listen-server/app/user/cmd/user/internal/types"
)

// 创建歌单
func CreateSonglistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSonglistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := songlist.NewCreateSonglistLogic(r.Context(), svcCtx)
		resp, err := l.CreateSonglist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
