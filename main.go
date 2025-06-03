// Package main Mocha-Master API
//
//	@title			Mocha-Master API
//	@version		0.0.1
//	@description	A relay and tunnel panel API for managing traffic forwarding rules and server configurations
//	@termsOfService	http://swagger.io/terms/
//
//	@contact.name	Yuzuki616
//	@contact.url	https://github.com/Yuzuki616
//
//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT
//
//	@host		localhost:8080
//	@BasePath	/
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Access token for admin endpoints
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
	"github.com/dgraph-io/ristretto"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime"

	"github.com/eko/gocache/lib/v4/cache"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
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
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     100,
		BufferItems: 64,
	})
	if err != nil {
		log.Fatal("Init ristretto cache failed.", zap.Error(err))
	}
	ristrettoStore := ristretto_store.NewRistretto(ristrettoCache)
	cache := cache.New[any](ristrettoStore)
	d, err := data.New(c.DbPath)
	if err != nil {
		log.Fatal("Open db failed.", zap.Error(err))
	}
	m := middleware.New(c)
	h := handle.NewHandle(d, cache, c)
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
