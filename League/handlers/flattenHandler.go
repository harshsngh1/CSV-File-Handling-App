package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func Flatten(w http.ResponseWriter, r *http.Request) {
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
	var flattened []string
	for _, row := range records {
		flattened = append(flattened, row...)
	}
	response := strings.Join(flattened, ",")
	fmt.Fprint(w, response)
}
