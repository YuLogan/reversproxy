package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var tarurl = &url.URL{
	Scheme: "http",
	Path:   "/",
	Host:   "127.0.0.1:9001",
}

func main() {
	reproxy := httputil.NewSingleHostReverseProxy(tarurl)
	fmt.Println("recover proxy start at port 80")
	http.ListenAndServe(":80", reproxy)
}
