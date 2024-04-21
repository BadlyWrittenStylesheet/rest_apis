package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type ResponseBody struct {
	Result int `json:"result"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	var requestBody RequestBody
	// jsonData, _ := json.Marshal(r.Body)
	// fmt.Println(jsonData)
	// fmt.Println()
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return
	}
	result := requestBody.Num1 + requestBody.Num2
	req, _ := json.Marshal(requestBody)
	fmt.Println("Processing request:", string(req))
	responseBody := ResponseBody{
		Result: result,
	}

	jsonResponse, _ := json.Marshal(responseBody)
	fmt.Println("Responding with:", string(jsonResponse))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	port := 8080
	http.HandleFunc("/api", handleRequest)
	fmt.Printf("Listening on port %d...", port)
	fmt.Println(fmt.Sprintf(":%d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Fail.")
	}
}
