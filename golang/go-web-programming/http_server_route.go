package main

import (
    "fmt"
    "net/http"
    "log"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        sayhelloName(w, r)
        return
    }
    http.NotFound(w, r)
    return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello myroute!")
}

func main() {
    mux := &MyMux{}
    err := http.ListenAndServe(":9090", mux)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}