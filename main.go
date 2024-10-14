package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const NOOFDAYSINYEAR = 365

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/hello-world") {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}

	if (r.Method != "GET") {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World! You are at %s path \n", r.URL.Path)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/form-age-calc" || r.Method != "POST") {
		fmt.Fprintln(w, "404 not fount", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, err)
		log.Fatal(err)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")
	ageInt, err := strconv.Atoi(age); 
	if err != nil {
		log.Fatal(err)
		return
	}

	ageInDays := NOOFDAYSINYEAR * ageInt

	fmt.Fprintf(w, "Hello %s. You are %s years old. You are %d days old (Note: we didn't include leap years).", name, age, ageInDays)
}

func main () {
	fileServer := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/form-age-calc", formHandler)

	// Listener
	if err := http.ListenAndServe(":3010", nil); err != nil {
		log.Fatal(err)
	}
}