package main

import (
	"fmt"
	"net/http"
)


func main() {
	
	fmt.Printf("Starting server at port 3000\n")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}