package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(""))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /")
	w.Write([]byte("res"))
	// fmt.Fprint(w, value)
}

func main() {
	//http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("server start on port 80")
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err := listen.Accept()
		if err != nil {
			continue
		}
	}
}
