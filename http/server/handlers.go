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
	}
}
