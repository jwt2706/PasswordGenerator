package handler

import (
    "encoding/json"
    "net/http"
    "regexp"
)

type PasswordRequest struct {
    Password string `json:"password"`
}

type StrengthResponse struct {
    Strength string `json:"strength"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    } else if r.Method == http.MethodPost {
        var req PasswordRequest
        err := json.NewDecoder(r.Body).Decode(&req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        var strength string
        switch {
        case isWeak(req.Password):
            strength = "Weak"
        case isMedium(req.Password):
            strength = "Medium"
        default:
            strength = "Strong"
        }

        res := StrengthResponse{Strength: strength}
        json.NewEncoder(w).Encode(res)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func isWeak(password string) bool {
    length := len(password)
    hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
    hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9\s]`).MatchString(password)

    return length < 8 || !hasLower || !hasUpper || !hasNumber || !hasSpecial
}

func isMedium(password string) bool {
    length := len(password)
    hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
    hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
    hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
    hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9\s]`).MatchString(password)

    return length >= 8 && ((hasLower && hasUpper && hasNumber) || (hasLower && hasUpper && hasSpecial) || (hasLower && hasNumber && hasSpecial) || (hasUpper && hasNumber && hasSpecial))
}