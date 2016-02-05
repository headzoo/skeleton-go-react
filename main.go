package main

import (
    "fmt"
    "flag"
    "net/http"
    "time"
    "github.com/golang/glog"
    "github.com/headzoo/website/app"
    "github.com/headzoo/website/app/common"
    "log"
)

func main() {
    flag.Parse()
    defer glog.Flush()
    
    http.Handle("/", httpLogger(app.GetIndex))
    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)
    
    log.Println("Listening...")
    err := http.ListenAndServe(":8099", nil)
    if err != nil {
        fmt.Println(err)
    }
}

func httpLogger(h http.HandlerFunc) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        startTime := time.Now()
        h(w, req)
        finishTime := time.Now()
        elapsedTime := finishTime.Sub(startTime)
        common.LogAccess(w, req, elapsedTime)
    })
}
