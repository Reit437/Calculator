package main

import (
	"fmt"
	"net/http"

	server "github.com/Reit437/Calculator/internal/server"
)

func main() {
	http.HandleFunc("/api/v1/calculate", server.CalculateHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

//(Invoke-WebRequest -Method Post -Uri "http://localhost:8080/api/v1/calculate" -Headers @{"Content-Type"="application/json"} -Body '{"expression":"80*9/f75/2"}').Content
//curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"1*(1+1+)\"}" http://localhost:8080/api/v1/calculate
