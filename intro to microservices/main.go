package main

import (
	"log"
	"net/http"
)

func main() {
	http.ListenAndServe(":9090", nil)
	log.Panic("sdsd")
}
