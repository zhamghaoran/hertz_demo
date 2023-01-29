package Controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func V1controller(c context.Context, ctx *app.RequestContext) {
	val, _ := ctx.Get("test")
	ctx.JSON(http.StatusOK, utils.H{
		"response": val,
	})

}
