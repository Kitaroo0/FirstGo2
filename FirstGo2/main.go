package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var reqBody JsonRequest
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if reqBody.Message == "" {
		response := JsonResponse{
			Status:  "400",
			Message: "Некорректное JSON-сообщение",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Printf("Received message: %s\n", reqBody.Message)

	response := JsonResponse{
		Status:  "success",
		Message: "Данные успешно приняты",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
