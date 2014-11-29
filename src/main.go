package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./low.m3u8")))
	http.ListenAndServe(":8080", nil)
}
