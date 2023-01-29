package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
	"hertz_demo/Controller"
	"net/http"
)

func MyMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		fmt.Println("pre-handle")
		c.Next(ctx)
		fmt.Println("post-handle")
	}
}
func testClient() {
	//send get request
	c, err := client.NewClient()
	if err != nil {
		return
	}
	status, body, _ := c.Get(context.Background(), nil, "http://www.baidu.com")
	fmt.Println(status)
	fmt.Println(string(body))
	//send post request
	var postArgs protocol.Args
	postArgs.Add("test", "test")
	post, i, _ := c.Post(context.Background(), nil, "http://localhost/v1/v2", &postArgs)
	fmt.Println(post)
	fmt.Println(string(i))

}
func main() {
	h := server.Default()
	g1 := h.Group("/v1")
	{
		g1.POST("/v2", Controller.V1controller)
	}
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		testClient()
		ctx.JSON(http.StatusOK, utils.H{
			"response": "pong",
		})

	})
	h.Spin()

}
