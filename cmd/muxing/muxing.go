package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func getParamName(w http.ResponseWriter, r *http.Request) {
	lParam := mux.Vars(r)["PARAM"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", lParam)
}

func getBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func postData(w http.ResponseWriter, r *http.Request) {
	lBody, lError := io.ReadAll(r.Body)
	if lError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%s", string(lBody))
}

func postHeaders(w http.ResponseWriter, r *http.Request) {
	lValue1Text := r.Header.Get("a")
	lValue2Text := r.Header.Get("b")

	lValue1, lError1 := strconv.Atoi(lValue1Text)
	lValue2, lError2 := strconv.Atoi(lValue2Text)

	if lError1 != nil || lError2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("a+b", strconv.Itoa(lValue1+lValue2))
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", getParamName).Methods("GET")
	router.HandleFunc("/bad", getBad).Methods("GET")
	router.HandleFunc("/data", postData).Methods("POST")
	router.HandleFunc("/headers", postHeaders).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
