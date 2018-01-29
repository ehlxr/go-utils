package main

import (
	"log"
	"net/http"

	"github.com/ehlxr/go-utils/common/server"
)

func main() {
	server := server.NewServer()
	err := server.Register(new(Hello))
	if err != nil {
		log.Fatal(err)
	}
	err = server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

type Hello struct {
}

func (h *Hello) Print(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("print"))
}

func (h *Hello) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
