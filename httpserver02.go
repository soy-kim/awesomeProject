package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":80", mux)
	// can connect to web browser http://127.0.0.1/
}
