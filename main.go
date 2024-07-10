package main

import (
	"fmt"
	"net/http"

	web "web/func" // package creat by team
)

func main() {
	http.HandleFunc("/", web.Print) // if i was in url / use func Print
	http.HandleFunc("/ascii-art", web.Handel_input)
	imgeServe := http.FileServer(http.Dir("image"))
	styeServe := http.FileServer(http.Dir("style"))
	http.Handle("/pictures/", http.StripPrefix("/pictures", imgeServe))
	http.Handle("/css/", http.StripPrefix("/css", styeServe))
	fmt.Print(http.ListenAndServe(":8081", nil)) // this func for run server
}
