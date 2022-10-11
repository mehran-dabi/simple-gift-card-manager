package helper

import "github.com/gin-gonic/gin"

func GinResponse(ctx *gin.Context, code int, body interface{}) {
	type response struct {
		Status int         `json:"status"`
		Body   interface{} `json:"body"`
	}

	resp := response{
		Status: code,
		Body:   body,
	}
	ctx.JSON(code, resp)
}
