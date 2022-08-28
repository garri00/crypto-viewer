package main

import (
	"cryptoViewer/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/hello", handlers.HomeHandler)
	http.HandleFunc("/random", handlers.RandomHandler)
	http.ListenAndServe(":8080", nil)

	//go func() {
	//	http.HandleFunc("/hello", hello)
	//	http.HandleFunc("/random", random)
	//
	//	http.ListenAndServe(":8080", nil)
	//
	//}()
	//time.Sleep(100 * time.Millisecond)

	requestURL := fmt.Sprintf("http://localhost:8080/random")
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

}
