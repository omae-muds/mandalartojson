package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type IMandalartAtom interface {
}

type MandalartAtom struct {
	Name    string
	Props   map[string]any
	Arounds [8]*MandalartAtom
}

func main() {
	m := MandalartAtom{
		Name:    "core",
		Props:   map[string]any{"achieved": false},
		Arounds: [8]*MandalartAtom{},
	}
	m2 := MandalartAtom{
		Name:    "child",
		Props:   map[string]any{"achieved": true},
		Arounds: [8]*MandalartAtom{},
	}
	m.Arounds[0] = &m2

	err := json.NewEncoder(os.Stdout).Encode(m)
	if err != nil {
		log.Fatal(err)
	}

	// Read CSV
	f, err := os.Open("example.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	var noblankrows [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Skip empty line
		if strings.Join(record, "") == "" {
			continue
		}
		noblankrows = append(noblankrows, record)
		fmt.Println(record)
	}

	h := len(noblankrows)
	w := len(noblankrows[0])
	// expectmaxh, expectmaxw := h, w
	// if h > w {
	// 	expectmaxh = w
	// } else if w > h {
	// 	expectmaxw = h
	// }

	// Find empty column & Trim
	candidate := make([]bool, w)
	for i := 0; i < w; i++ {
		candidate[i] = true
	}
	for _, record := range noblankrows {
		for col, str := range record {
			record[col] = strings.Trim(str, " 　")
			str = record[col]
			if str != "" {
				candidate[col] = false
			}
		}
	}
	// Remove empty column
	// var noblankrowscols [][]string
	noblankrowscols := make([][]string, h*w)
	for icol, b := range candidate {
		if b {
			for irow, str := range noblankrows {
				noblankrowscols[irow] = append(str[:icol], str[icol+1:]...)
			}
		}
	}
	fmt.Println(noblankrowscols)
}
