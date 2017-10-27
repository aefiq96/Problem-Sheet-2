//G00324934 - Nuzhafiq Iqbal Noor Azman
//https://gobyexample.com/structs
package main

import (
//	"fmt"	
	"net/http"
	"html/template"
	"math/rand"
	"time"
	"strconv"
)
type message struct{
	Message string
	Guess int
	WinOrNot bool
	MessageTwo string

}
func Guess(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	http.ServeFile(w,r,"index.html")
}
//this generates random number between given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func GuessRoute(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	//cookies
	rand := xrand(1,20)

	cookie, err := r.Cookie("target");
	if err != nil {
	// Create a cookie instance and set the cookie.
	// You can delete the Expires line (and the time import) to make a session cookie.
		cookie = &http.Cookie{
		Name:    "target",
		Value:   strconv.Itoa(rand),
		Expires: time.Now().Add(72 * time.Hour),
	}
		http.SetCookie(w, cookie)
	}

	// If we could read it, try to convert its value to an int.
	guessTarget, _ := strconv.Atoi(r.FormValue("guess"))

	//create a string 
	str := "Guess a number between 1 and 20"
	//pass it to a struct 
	m := &message{Message: str, Guess: guessTarget, WinOrNot: false}


	CookieVal, _ := strconv.Atoi(cookie.Value)


	if CookieVal == guessTarget{

     m.MessageTwo = "You won!!"
     m.WinOrNot = true

	}else if guessTarget < CookieVal{
 
      m.MessageTwo = "Guess Too Low"
	}else if guessTarget > CookieVal{
		 m.MessageTwo = "Guess Too High"
	}
	
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

