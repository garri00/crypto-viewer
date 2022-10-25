package main

import (
	"crypto-viewer/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server start")

	r := api.NewRouter()

	http.ListenAndServe(":8080", r)
	fmt.Println("server STOP")

}
