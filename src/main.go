package main

import (
"fmt"
"net/http"
)

func myhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "はろー、%s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", myhandler)
	http.ListenAndServe(":8080", nil)
}
