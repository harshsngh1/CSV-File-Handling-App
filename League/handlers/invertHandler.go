package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func Invert(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading file: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading CSV: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if len(records) == 0 {
		http.Error(w, "error: CSV file is empty", http.StatusBadRequest)
		return
	}

	numCols := len(records[0])
	for _, row := range records {
		if len(row) != numCols {
			http.Error(w, "error: CSV rows have inconsistent number of columns", http.StatusBadRequest)
			return
		}
	}

	n := len(records)
	inverted := make([][]string, numCols)
	for i := range inverted {
		inverted[i] = make([]string, n)
		for j := range records {
			inverted[i][j] = records[j][i]
		}
	}

	var response string
	for _, row := range inverted {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
