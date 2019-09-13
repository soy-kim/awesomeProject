package main

import "net/http"

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":80", nil)
	// can connect to web browser http://127.0.0.1/
}
