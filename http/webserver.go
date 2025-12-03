package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served %s\n", r.Host)

}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("Mon Jan 02 15:04:05 -0700 2006")
	Body := "My time is: "
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)

	fmt.Fprintf(w, "Serving: %s\\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arg := os.Args
	if len(arg) != 1 {
		PORT = ":" + arg[1]
	}
	fmt.Println("Using port number: " + PORT)

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
