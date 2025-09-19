package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	stdin := os.Stdin
	csvReader := csv.NewReader(stdin)

	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	if len(records) == 0 {
		fmt.Println("[]")
		return
	}

	headers := records[0]
	data := []map[string]string{} // slice of maps

	for rowIndex, row := range records[1:] {
		if len(row) != len(headers) {
			fmt.Fprintf(os.Stderr, "ERROR: row %d has %d columns, expected %d\n", rowIndex+1, len(row), len(headers))
			os.Exit(1)
		}

		entry := make(map[string]string, len(headers))
		for i, header := range headers {
			entry[header] = row[i]
		}
		data = append(data, entry)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
