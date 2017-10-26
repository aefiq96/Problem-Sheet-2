//G00324934 - Nuzhafiq Iqbal Noor Azman
package main

import (
//	"fmt"	
	"net/http"
)

func Guess(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint(w, "<h1>Guessing game</h1>")
	http.ServeFile(w,r,"index.html")
}

func main() {
	//localhost:8080
	http.HandleFunc("/", Guess)	
	http.ListenAndServe(":8080", nil)
}
