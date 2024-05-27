package web

import (
	"fmt"
	"html/template"
	"net/http"
)

func Print(w http.ResponseWriter, r *http.Request) { // w for send data from server to user and r for take data from user
	tmp, err := template.ParseFiles("./templet/indix.html") // pionter in file html
	if r.URL.Path != "/" {                                  // handel if url was not valide
		http.NotFound(w, r) // enter to func for print not found
		return
	}
	if err != nil {
		http.Error(w, "server down", http.StatusInternalServerError) // hundul if was file html not
		return
	}
	tmp = template.Must(tmp, err)                              // some one
	font := r.FormValue("select")                              // take user's select font
	user_input := r.FormValue("user_input")                    // take user's input
	tmp.Execute(w, SplitAndPrint(user_input, Font(font)))      // send data to showing in page in form respons and creat ascii art
	if (font == "" || user_input == "") && r.Method != "GET" { // if user doesn't  input any thing
		fmt.Fprintln(w, `<marquee behavior="" direction="right"> You did't enter any thing , Please try again : </marquee>`)
		return
	}
}
