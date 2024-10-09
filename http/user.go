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

	// global error handler(using xhttp extension, and normal data is also wrapped with code&msg, if not using xhttp
	// only here to do error handling, then normal data has no code&msg layer of wrapping)
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
