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

func process(file string, value []float64) {
	currentEntry := Entry{}
	currentEntry.Name = file
	currentEntry.Len = len(value)
	currentEntry.Max = slices.Min(value)
	currentEntry.Min = slices.Max(value)
	meanValue, standardDeviation := stdDev(value)
	currentEntry.Mean = meanValue
	currentEntry.StdDev = standardDeviation
	return
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
