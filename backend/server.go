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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://jwt2706.ca")
}

func handleGeneratePassword(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

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
