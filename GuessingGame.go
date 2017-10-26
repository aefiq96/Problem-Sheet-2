//G00324934 - Nuzhafiq Iqbal Noor Azman
//https://gobyexample.com/structs
package main

import (
//	"fmt"	
	"net/http"
	"html/template"
)
type message struct{
	Message string
}
func Guess(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	http.ServeFile(w,r,"index.html")
}

func GuessRoute(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	//create a string 
	str := "Guess a number between 1 and 20"
	//pass it to a struct 
	m := &message{Message: str}
	//parse a file
	tem, _ := template.ParseFiles("guess.tmpl")
	tem.Execute(w,m)

}

func main() {
	//localhost:8080
	http.HandleFunc("/", Guess)	
	http.HandleFunc("/guess", GuessRoute)
	http.ListenAndServe(":8080", nil)
}
