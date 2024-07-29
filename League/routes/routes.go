package routes

import (
	"League/handlers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/flatten", handlers.Flatten)
	http.HandleFunc("/invert", handlers.Invert)
	http.HandleFunc("/multiply", handlers.Multiply)
	http.HandleFunc("/sum", handlers.Sum)
}
