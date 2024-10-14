package main

// Imports
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Constants
const NOOFDAYSINYEAR = 365

// helloWorldHandler is an HTTP handler for the /hello-world path
// It returns "Hello World! You are at /hello-world path \n" as the response
// It only supports GET requests
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path is /hello-world
	if r.URL.Path != "/hello-world" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the method is GET
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// Print the response
	fmt.Fprintf(w, "Hello World! You are at %s path \n", r.URL.Path)
}

// formHandler is an HTTP handler for the /form-age-calc path
// It returns the age in days (ignoring leap years) given the age in years
// It only supports POST requests
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the path is /form-age-calc
	if r.URL.Path != "/form-age-calc" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Check if the method is POST
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// Parse the form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, err)
		log.Fatal(err)
		return
	}

	// Get the name and age from the form
	name := r.FormValue("name")
	age := r.FormValue("age")

	// Convert the age from string to int
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Calculate the age in days (ignoring leap years)
	ageInDays := NOOFDAYSINYEAR * ageInt

	// Print the response
	fmt.Fprintf(w, "Hello %s. You are %s years old. You are %d days old (Note: we didn't include leap years).", name, age, ageInDays)
}

// main is the entry point for the program
func main() {
	// Set up file server. This is the handler for requests to the root path
	// The handler is for the directory ./frontend/static
	fileServer := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/", fileServer)

	// Set up the hello-world handler. The handler is for the /hello-world path
	http.HandleFunc("/hello-world", helloWorldHandler)

	// Set up the form-age-calc handler. The handler is for the /form-age-calc path
	http.HandleFunc("/form-age-calc", formHandler)

	// Set up the listener. The listener is for port 3010. The handler is nil
	// because the handler for each path is set up above
	if err := http.ListenAndServe(":3010", nil); err != nil {
		// If there is an error, log it
		log.Fatal(err)
	}
}
