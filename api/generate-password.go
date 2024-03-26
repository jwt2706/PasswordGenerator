package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "math/rand"
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

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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