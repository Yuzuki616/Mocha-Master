package main

import (
	"flag"
	"fmt"
	"github.com/Yuzuki616/Mocha-Master/conf"
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/Yuzuki616/Mocha-Master/handle"
	"github.com/Yuzuki616/Mocha-Master/log"
	"github.com/Yuzuki616/Mocha-Master/middleware"
	"github.com/Yuzuki616/Mocha-Master/router"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var confPath = flag.String("conf", "config.json", "config file")
var version = "0.0.1"

func main() {
	log.Info("Mocha-Master starting...")
	log.Info(fmt.Sprintf("Version: %s", version))
	log.Info("Dev: github.com/Yuzuki616")
	flag.Parse()
	c := conf.New(*confPath)
	err := c.Load()
	if err != nil {
		log.Fatal("Load config failed.", zap.Error(err))
	}
	log.SetLevel(c.LogLevel)
	d, err := data.New(c.DbPath)
	if err != nil {
		log.Fatal("Open db failed.", zap.Error(err))
	}
	m := middleware.New(c)
	h := handle.NewHandle(d)
	r := router.NewRouter(h, m)
	go func() {
		err = r.Start(c.Addr)
		if err != nil {
			log.Fatal("Start router failed.", zap.Error(err))
		}
	}()
	log.Info("Mocha-Master started.")

	runtime.GC()
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-osc
}
