package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Data struct {
	Elment, Value string
}

var output Data

func Print(w http.ResponseWriter, r *http.Request) { // w for send data from server to user and r for take data from user
	dir, _ := os.ReadDir("draw")
	fmt.Println(dir)
	fmt.Println(r.PostForm)
	tmp, err := template.ParseFiles("./templet/index.html") // pionter in file html
	if r.URL.Path != "/" {                                  // handel if url was not valide
		http.NotFound(w, r) // enter to func for print not found
		return
	}
	if err != nil {
		http.Error(w, "server down ", http.StatusInternalServerError) // hundul if was file html not
		return
	}
	tmp = template.Must(tmp, err)
	err = (tmp.Execute(w, output)) //f send data to showing in page in form respons and creat ascii art
	if err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
	}
}

func Handel_input(w http.ResponseWriter, r *http.Request) {
	font := r.FormValue("select")
	user_input := r.FormValue("user_input")
	fmt.Println(r.)
	if len(font) == 0 || len(user_input) == 0 {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	mapDraw := Font(font)
	if mapDraw == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	Stock(user_input, mapDraw, &output)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Stock(s string, mapDraw map[int][]string, r *Data) {
	r.Value = s
	r.Elment = SplitAndPrint(s, mapDraw)
}
