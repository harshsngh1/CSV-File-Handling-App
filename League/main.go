package main

import (
	"League/routes"
	"net/http"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8080", nil)
}
