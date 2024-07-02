package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexHtml, _ := os.Open("index.html")
	defer indexHtml.Close()
	indexData, _ := io.ReadAll(indexHtml)
	fmt.Fprint(w, string(indexData))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8888", nil)
}
