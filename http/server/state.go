package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"slices"
	"time"
)

type Entry struct {
	Name   string
	Len    int
	Min    float64
	Max    float64
	Mean   float64
	StdDev float64
}

func process(file string, value []float64) Entry {
	currentEntry := Entry{}
	currentEntry.Name = file
	currentEntry.Len = len(value)
	currentEntry.Max = slices.Min(value)
	currentEntry.Min = slices.Max(value)
	meanValue, standardDeviation := stdDev(value)
	currentEntry.Mean = meanValue
	currentEntry.StdDev = standardDeviation
	return currentEntry
}

func stdDev(x []float64) (float64, float64) {
	sum := 0.0
	for _, val := range x {
		sum += val
	}
	meanValue := sum / float64(len(x))
	var sq float64
	for i := 0; i < len(x); i++ {
		sq = sq + math.Pow((x[i]-meanValue), 2)
	}
	standardDeviation := math.Sqrt(sq) / float64(len(x))
	return meanValue, standardDeviation
}

var JSONFILE = "./data.json"

type Photobook []Entry

var data = Photobook{}
var index map[string]int

func DESerialized(slice interface{}, w io.Reader) error {
	fe := json.NewDecoder(w)
	return fe.Decode(&slice)

}

func Serialized(slice interface{}, w io.Writer) error {
	fe := json.NewEncoder(w)
	return fe.Encode(slice)
}

func save(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = Serialized(&data, f)
	if err != nil {
		return err
	}
	return nil

}

func Read(file string) error {
	_, err := os.Stat(file)
	if err != nil {
		return err
	}
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	err = DESerialized(&data, f)
	if err != nil {
		return err
	}
	return nil
}

func creates() {
	index = make(map[string]int)
	for i, k := range data {
		index[k.Name] = i
	}
}

func insert(pS *Entry) error {
	_, ok := index[(*pS).Name]
	if ok {
		return fmt.Errorf("entry with name %s already exists", (*pS).Name)
	}
	data = append(data, *pS)
	creates()
	err := save(JSONFILE)
	if err != nil {
		return err

	}
	return nil
}

func DeleteEntry(name string) error {
	i, ok := index[name]
	if !ok {
		return fmt.Errorf("entry with name %s not found", name)
	}
	data = append(data[:i], data[i+1:]...)
	delete(index, name)
	err := save(JSONFILE)
	if err != nil {
		return err
	}
	return nil
}

func search(name string) *Entry {
	i, ok := index[name]
	if !ok {
		return nil
	}
	return &data[i]
}

func list() string {
	var all string
	for _, k := range data {
		all = all + fmt.Sprintf("%s\t , %d\t , %f\t, %f\t\n ", k.Name, k.Len, k.Mean, k.Mean)
	}
	return all

}

func main() {
	err := Read(JSONFILE)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
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
	mux.Handle("/insert", http.HandlerFunc(insertHandler))
	mux.Handle("/search", http.HandlerFunc(searchHandler))
	mux.Handle("/search/", http.HandlerFunc(searchHandler))
	mux.Handle("/delete/", http.HandlerFunc(deleteHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
