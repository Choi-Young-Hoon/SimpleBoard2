package webserver

import (
	"net/http"
	"simpleboard2/webserver/handler"
)

func HttpServerStart() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/css/", fs)
	http.Handle("/js/", fs)

	http.HandleFunc("/ws", handler.WebSocketHandler)
	http.HandleFunc("/", handler.IndexHandler)

	err := http.ListenAndServe(":50000", nil)
	if err != nil {
		panic(err)
	}
}
