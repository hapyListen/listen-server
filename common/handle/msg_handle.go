package handle

import (
	"go/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
)

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusUnauthorized, xhttp.BaseResponse[types.Nil]{
		Code: http.StatusUnauthorized,
		Msg:  "Token authentication failed",
	})
}

func ErrorHandler(err error) (int, any) {
	switch e := err.(type) {
	case *errors.CodeMsg:
		return e.Code, xhttp.BaseResponse[types.Nil]{
			Code: e.Code,
			Msg:  e.Msg,
		}
	default:
		return http.StatusInternalServerError, err
	}
}
