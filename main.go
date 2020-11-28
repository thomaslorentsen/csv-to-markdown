package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <csv files...>", os.Args[0])
		os.Exit(0)
	}

	var heading []string
	var records [][]string

	filenames := os.Args[1:]
	for _, filename := range filenames {
		csvfile, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error reading file: %s", err.Error())
			os.Exit(1)
		}
		r := csv.NewReader(csvfile)
		line := 0
		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			if len(heading) == 0 {
				heading = record
			} else if line > 0 {
				records = append(records, record)
			}
			line++

		}
	}

	for i := 0; i != len(heading); i++ {
		if i == 0 {
			fmt.Printf("|")
		}
		fmt.Printf(" %s |", heading[i])
	}
	fmt.Printf("\n")
	for i := 0; i != len(heading); i++ {
		if i == 0 {
			fmt.Printf("|")
		}
		fmt.Printf(" --- |")
	}
	fmt.Printf("\n")
	for _, record := range records {
		for i := 0; i != len(record); i++ {
			if i == 0 {
				fmt.Printf("|")
			}
			fmt.Printf(" %s |", record[i])
		}
		fmt.Printf("\n")
	}
}
