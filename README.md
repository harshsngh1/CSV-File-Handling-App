# CSV File Handling App

This is a simple web service written in Golang for processing CSV files. It provides various operations such as echoing the content, flattening rows, inverting columns, multiplying values, and summing values.

## Introduction
This project implements a RESTful service for managing and processing CSV files. It allows users to:
- Echo the content of a CSV file.
- Flatten the rows of a CSV file into a single line.
- Invert the rows and columns of a CSV file.
- Multiply all integer values in a CSV file.
- Sum all integer values in a CSV file.

## Project Structure
```
League/
├── handlers/
│   ├── echoHandler.go
│   ├── flattenHandler.go
│   ├── invertHandler.go
│   ├── multiplyHandler.go
│   └── sumHandler.go
├── resources/
│   └── matrix.csv
├── routes/
│   └── routes.go
├── tests/
│   └── assessment_test.go
├── main.go
├── go.mod
├── dockerfile
└── README.md
```

## Installation

- Clone the repository:
```
git clone https://github.com/harshsngh1/CSV-File-Handling-App
```
- Navigate to project directory
```
cd League
```
- Run the program
```
go run main.go
```
### Same can be run using Docker as well
- Build the Docker image:
```
docker build -t league-app .
```
- Run the built image
```
docker run -p 8080:8080 league-app
```


## Testing
- To run tests go to /tests dir and run
```
go test
```
- To run Tests via API, run these curl requests
```
curl -F 'file=@resources/matrix.csv' "localhost:8080/echo"
curl -F 'file=@resources/matrix.csv' "localhost:8080/flatten"
curl -F 'file=@resources/matrix.csv' "localhost:8080/sum"
curl -F 'file=@resources/matrix.csv' "localhost:8080/multiply"
curl -F 'file=@resources/matrix.csv' "localhost:8080/invert"
```
