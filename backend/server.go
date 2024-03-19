package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type PasswordResponse struct {
	Password string `json:"password"`
}

func generatePassword(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>/?`~"
	var password strings.Builder

	for i := 0; i < length; i++ {
		password.WriteByte(characters[rand.Intn(len(characters))])
	}

	return password.String()
}

func handleGeneratePassword(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
        // Pre-flight request. Reply with CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "http://jwt2706.ca")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.WriteHeader(http.StatusNoContent)
        return
    }

    w.Header().Set("Access-Control-Allow-Origin", "http://jwt2706.ca")

    length := 10
    if len(r.URL.Query()["length"]) > 0 {
        length, _ = strconv.Atoi(r.URL.Query()["length"][0])
    }

    password := generatePassword(length)
    response := PasswordResponse{Password: password}
    json.NewEncoder(w).Encode(response)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/generate-password", handleGeneratePassword)
	http.ListenAndServe(":8080", nil)
}
