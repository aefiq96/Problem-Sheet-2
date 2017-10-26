//G00324934 - Nuzhafiq Iqbal Noor Azman
package main

import (
	"fmt"	
	"net/http"
)

func Guess(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Guessing game! asuhh")
}

func main() {
	//open localhost on port 8080
	http.HandleFunc("/", Guess)	
	http.ListenAndServe(":8080", nil)
}
