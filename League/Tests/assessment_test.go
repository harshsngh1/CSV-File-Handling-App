package tests

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
)

func loadTestFile(t *testing.T) [][]string {
	file, err := os.Open("../resources/matrix.csv")
	if err != nil {
		t.Fatalf("Failed to open file : %v", err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV file : %v", err)
	}
	return records
}

func TestEcho(t *testing.T) {
	records := loadTestFile(t)
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if response != expected {
		t.Errorf("Echo function returned unexpected result: got %v want %v", response, expected)
	}
}

func TestInvert(t *testing.T) {
	records := loadTestFile(t)
	n := len(records)
	inverted := make([][]string, n)
	for i := range inverted {
		inverted[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			inverted[j][i] = records[i][j]
		}
	}

	var response string
	for _, row := range inverted {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	expected := "1,4,7\n2,5,8\n3,6,9\n"
	if response != expected {
		t.Errorf("Invert function returned unexpected result: got %v want %v", response, expected)
	}
}

func TestFlatten(t *testing.T) {
	records := loadTestFile(t)
	var flat []string
	for _, row := range records {
		flat = append(flat, row...)
	}

	response := strings.Join(flat, ",")
	expected := "1,2,3,4,5,6,7,8,9"
	if response != expected {
		t.Errorf("Flatten function returned unexpected result: got %v want %v", response, expected)
	}
}

func TestSum(t *testing.T) {
	records := loadTestFile(t)
	sum := 0
	for _, row := range records {
		for _, val := range row {
			var num int
			fmt.Sscanf(val, "%d", &num)
			sum += num
		}
	}

	expected := 45
	if sum != expected {
		t.Errorf("Sum function returned unexpected result: got %v want %v", sum, expected)
	}
}

func TestMultiply(t *testing.T) {
	records := loadTestFile(t)
	product := 1
	for _, row := range records {
		for _, val := range row {
			var num int
			fmt.Sscanf(val, "%d", &num)
			product *= num
		}
	}

	expected := 362880
	if product != expected {
		t.Errorf("Multiply function returned unexpected result: got %v want %v", product, expected)
	}
}
