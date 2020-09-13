package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hsimao/go_gateway_demo/golang_common/lib"
	"github.com/hsimao/go_gateway_demo/router"
)

func main() {
	lib.InitModule("./conf/dev/")
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
