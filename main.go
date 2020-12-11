package main

import (
	"fmt"
	"net/http"
)

const VERSION = "0.0.1"

func main() {
	fmt.Printf("KabWeb v%s\n", VERSION)
	fmt.Println("Waiting for connections...")

	http.Handle("/", http.FileServer(http.Dir("./web")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
