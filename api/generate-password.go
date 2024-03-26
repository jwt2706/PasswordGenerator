package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "math/rand"
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
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    } else if r.Method == http.MethodPost {
        length := 10
        if len(r.URL.Query()["length"]) > 0 {
            length, _ = strconv.Atoi(r.URL.Query()["length"][0])
        }

        password := generatePassword(length)
        response := PasswordResponse{Password: password}
        json.NewEncoder(w).Encode(response)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}