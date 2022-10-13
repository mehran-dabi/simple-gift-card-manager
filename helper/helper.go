package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Status bool        `json:"status"`
	Body   interface{} `json:"body"`
}

func GinResponse(ctx *gin.Context, code int, body interface{}) {
	resp := response{
		Status: true,
		Body:   body,
	}
	ctx.JSON(code, resp)
}

func GinErrResponse(ctx *gin.Context, code int, err error) {
	resp := response{
		Status: false,
		Body:   err.Error(),
	}

	ctx.JSON(code, resp)
}

func ReplaceEscapeCharacter(s string) string {
	re := strings.NewReplacer("?", "\\?", "(", "\\(", ")", "\\)", ".", "\\.")
	return re.Replace(s)
}
