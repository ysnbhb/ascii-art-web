package main

import (
	"fmt"
	"net/http"

	web "web/func" // package creat by team
)

func main() {
	http.HandleFunc("/", web.Print) // if i was in url / use func Print
	fmt.Println("Server run in http://localhost:8080/")
	http.ListenAndServe(":8080", nil) // this func for run server
}
