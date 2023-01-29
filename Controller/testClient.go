package Controller

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
)

func main() {
	// send get request
	c, err := client.NewClient()
	if err != nil {
		return
	}
	//status, body, _ := c.Get(context.Background(), nil, "http://www.baidu.com")
	//fmt.Println(status)
	//fmt.Println(string(body))
	// send post request
	var postArgs protocol.Args
	postArgs.Add("test", "test")
	post, i, _ := c.Post(context.Background(), nil, "http://localhost/v1/v2", &postArgs)
	fmt.Println(post)
	fmt.Println(string(i))
}
