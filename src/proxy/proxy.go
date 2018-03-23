package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const remoteAddr = "http://ryzen2.utv.lokal:9200"
const addr = "localhost:3000"
const timeout = 30 * time.Second

func main() {
	fmt.Printf("Starting server, listening at %s\n", addr)

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(catchAllHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}
	log.Fatal(srv.ListenAndServe())
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from : %s\n", r.RemoteAddr)

	path := r.URL.Path
	uri := remoteAddr + path

	request, err := http.NewRequest("GET", uri, r.Body)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	buffer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Write(buffer)
}
