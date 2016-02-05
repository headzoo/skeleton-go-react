package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/headzoo/skeleton-go-react/app"
	"github.com/headzoo/skeleton-go-react/app/common"
	"github.com/headzoo/skeleton-go-react/app/controller"
	"github.com/headzoo/skeleton-go-react/app/model"
	"log"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := app.InitSettings("./conf/site.yml"); err != nil {
		log.Fatal(err)
	}
	if err := app.InitRedis(); err != nil {
		log.Fatal(err)
	}
	if err := model.InitDatabase(); err != nil {
		log.Fatal(err)
	}

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", static)
	http.Handle("/", logHandler(controller.GetIndex))

	log.Println("Listening to", app.Settings.Server.Address)
	log.Fatal(http.ListenAndServe(app.Settings.Server.Address, nil))
}

func logHandler(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		h(w, req)
		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)
		common.LogAccess(w, req, elapsedTime)
	})
}
