package main


import (
	"net/http"
)


func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./")))

	var server *http.Server
	server = &http.Server{}
	server.Handler = mux
	server.Addr = "localhost:8080"

	server.ListenAndServe()

}
