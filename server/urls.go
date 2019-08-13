package server

import "net/http"

func static() {
	http.Handle("/public", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
}

func index() {}
