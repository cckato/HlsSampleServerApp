package main

import (
	"fmt"
	"net/http"
)

func movieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-mpegurl")
	w.Header().Set("Accept-Ranges", "bytes")
	fmt.Fprintln(w, "#EXTM3U")
	fmt.Fprintln(w, "#EXT-X-VERSION:3")
	fmt.Fprintln(w, "#EXT-X-MEDIA-SEQUENCE:0")
	fmt.Fprintln(w, "#EXT-X-ALLOW-CACHE:YES")
	fmt.Fprintln(w, "#EXT-X-TARGETDURATION:14")
	fmt.Fprintln(w, "#EXTINF:14.000000,")
	fmt.Fprintln(w, "http://ec2-54-65-87-57.ap-northeast-1.compute.amazonaws.com:8080/kinmoza/kinmoza_000.ts")
	fmt.Fprintln(w, "#EXTINF:2.300000,")
	fmt.Fprintln(w, "http://ec2-54-65-87-57.ap-northeast-1.compute.amazonaws.com:8080/kinmoza/kinmoza_001.ts")
	fmt.Fprintln(w, "#EXTINF:8.633333,")
	fmt.Fprintln(w, "http://ec2-54-65-87-57.ap-northeast-1.compute.amazonaws.com:8080/kinmoza/kinmoza_002.ts")
	fmt.Fprintln(w, "#EXT-X-ENDLIST")
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.HandleFunc("/movie.m3u8", movieHandler)
	http.ListenAndServe(":8080", nil)
}
