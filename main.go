package main

import (
	"net/http"

	web "web/func" // package creat by team
)

func main() {
	http.HandleFunc("/", web.Print)   // if i was in url / use func Print
	http.ListenAndServe(":8080", nil) // this func for run server
}
