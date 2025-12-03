package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	err := insert(&entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		Body := "Failed to add record\n"
		fmt.Fprintf(w, "%s", Body)
	} else {
		Body := "New record added successfully\n"
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", Body)
	}
	log.Println("Serving:", r.URL.Path, "from", r.Host)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	//Get Search value from URL
	paramstr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramstr)
	if len(paramstr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not enough arquments:"+r.URL.Path)
		return
	}
	var body string
	datset := paramstr[2]
	t := search(datset)
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		body = "Could not find record\n" + datset + "\n"
	} else {
		w.WriteHeader(http.StatusOK)
		body = fmt.Sprintf("%s,%d,%f,%f\n", t.Name, t.Len, t.Mean, t.StdDev)
	}
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	fmt.Fprintf(w, "%s", body)
}

func main() {

	err := save(JSONFILE)
	if err != nil && err != io.EOF {
		fmt.Println("Errors", err)
		return
	}
	creates()

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/list", http.HandlerFunc(listHandler))

	mux.Handle("/insert/", http.HandlerFunc(insertHandler))

	mux.Handle("/insert", http.HandlerFunc(insertHandler))

	mux.Handle("/search", http.HandlerFunc(searchHandler))

	mux.Handle("/search/", http.HandlerFunc(searchHandler))

	mux.Handle("/delete/", http.HandlerFunc(deleteHandler))

	mux.Handle("/status", http.HandlerFunc(statusHandler))

	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at ", PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
