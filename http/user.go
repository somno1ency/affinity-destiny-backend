package main

import (
	"flag"
	"fmt"

	"ad.com/http/internal/config"
	"ad.com/http/internal/handler"
	"ad.com/http/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 全局错误处理(在hander中使用了xhttp扩展实现了类似效果,并且正常数据也一并带了code&msg,不使用xhttp只在这里做错误处理,那么正常数据是没有code&msg这层包装的,只有错误有)
	// httpx.SetErrorHandler(func(err error) (int, any) {
	// 	switch e := err.(type) {
	// 	case *errors.CodeMsg:
	// 		return http.StatusOK, xhttp.BaseResponse[string]{
	// 			Code: e.Code,
	// 			Msg:  e.Msg,
	// 		}
	// 	default:
	// 		return http.StatusInternalServerError, nil
	// 	}
	// })

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
