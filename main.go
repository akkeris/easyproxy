package main

import (
	"net/http/httputil"
        "net/http"
        "os"
	"net/url"
)


func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}

func getListenAddress() string {
	port := os.Getenv("PORT")
	return ":" + port
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {

	serveReverseProxy(os.Getenv("PROXY_URL"), res, req)
}


func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

