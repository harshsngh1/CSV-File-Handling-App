package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

func Multiply(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	product := 1
	for _, row := range records {
		for _, val := range row {
			num, err := strconv.Atoi(val)
			if err != nil {
				w.Write([]byte(fmt.Sprintf("invalid integer value: %v", err)))
				return
			}
			product *= num
		}
	}
	fmt.Fprint(w, product)
}
