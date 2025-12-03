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
	fmt.Fprintf(w, "%s", body)
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
		fmt.Fprintf(w, "%s", Body)
		return
	}
	body := dataset + "deleted!\n"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", body)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := list()
	fmt.Fprintf(w, "%s", body)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	body := fmt.Sprintf("Total entries: %d\n", len(list()))
	fmt.Fprintf(w, "%s", body)

}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	paramstr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramstr)
	if len(paramstr) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Not enough arquments:"+r.URL.Path)
		return
	}
	dataset := paramstr[2]
	dataStr := paramstr[3:]
	data := make([]float64, 0)
	for _, v := range dataStr {
		val, err := strconv.ParseFloat(v, 64)
		if err == nil {
			data = append(data, val)
		}
	}
	entry := process(dataset, data)
}
