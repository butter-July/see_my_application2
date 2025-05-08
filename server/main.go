package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	currentWindowTitle := ""

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			_, _ = responseWriter.Write([]byte(currentWindowTitle))
		case http.MethodPost:
			body, _ := io.ReadAll(request.Body)
			defer request.Body.Close()
			currentWindowTitle = string(body)
		}
	})

	log.Println("HTTP server starting")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
