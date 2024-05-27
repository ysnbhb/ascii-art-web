package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Notfound(w http.ResponseWriter) {
	file, err := os.ReadFile("notfound.txt") // read file
	if err != nil {
		http.Error(w, "ERRER 404 PAGE NOT FOUND ...", 404) // if file not found and send sign to surver
	}
	http.Error(w, string(file), 404) //  send sign to surver
}

func Print(w http.ResponseWriter, r *http.Request) { // w for send data from server to user and r for take data from user
	tmp, err := template.ParseFiles("./templet/indix.html") // pionter in file html
	if r.URL.Path != "/" {                                  // handel if url was not valide
		Notfound(w) // enter to func for print not found
		return
	}
	if err != nil {
		http.Error(w, "server down", 500) // hundul if was file html not
		return
	}
	tmp = template.Must(tmp, err) // some one
	font := r.FormValue("select") // take user's select font
	user_input := r.FormValue("user_input")
	fmt.Println(font)                                     // take user's input
	tmp.Execute(w, SplitAndPrint(user_input, Font(font))) // send data to showing in page in form respons and creat ascii art
	if font == "" {                                       // if user doesn't  input any thing
		fmt.Fprintln(w, `<marquee behavior="" direction="right"> You did't enter any thing , Please try again : </marquee>`)
		return
	}
}
