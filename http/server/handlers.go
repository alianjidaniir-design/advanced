package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const PORT = ":1234"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := "Hello World!"
	fmt.Fprintln(w, body)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	paramstr := strings.Split(r.URL.Path, "/")
	fmt.Println(paramstr)
	if len(paramstr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 not found", r.URL.Path)
		return
	}
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	dataset := paramstr[2]
	err := DeleteEntry(dataset)
	if err != nil {
		fmt.Println(err)
		Body := err.Error() + "\n"
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, Body)
		return
	}
	body := dataset + "deleted!\n"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, body)
}
