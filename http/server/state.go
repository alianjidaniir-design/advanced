package main

import (
	"net/http"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"time"
)

type Entry struct {
	Name string
}
