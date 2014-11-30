package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"bufio"
	"./utils"
)

func movieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-mpegurl")
	w.Header().Set("Accept-Ranges", "bytes")

	index := utils.GetRequestIndex(r.URL.RawQuery)
	if index < 0 {
		index = 0
	}

	f, err := os.Open("src/low.m3u8")
	check(err)

	delete_flag := false

	br := bufio.NewReaderSize(f, 4096)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		check(err)

		str := string(line)

		if index == 0 {
			fmt.Fprintln(w, str)
			continue
		}

		if utils.IsExtinf(str) {
			delete_flag = true
			continue
		}

		if utils.IsTsFile(str) && delete_flag {
			index--
			delete_flag = false
			continue
		}

		fmt.Fprintln(w, str)
	}
}

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path：%s\n", r.URL.Path)
	fmt.Fprintf(w, "Host:%s\n", r.URL.Host)
	fmt.Fprintf(w, "RawQuery:%s\n", r.URL.RawQuery)
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.HandleFunc("/movie.m3u8", movieHandler)
	http.HandleFunc("/hoge/", hogeHandler)
	http.ListenAndServe(":8080", nil)
	readFile("src/low.m3u8")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readFile(path string) {
	f, err := os.Open(path)
	check(err)

	br := bufio.NewReaderSize(f, 4096)
	for {
		line, _, err := br.ReadLine()
		fmt.Println(string(line))
		if err == io.EOF {
			break
		}
		check(err)
	}
}

