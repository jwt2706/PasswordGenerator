package handler

import (
    "encoding/json"
    "net/http"
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
        case len(req.Password) < 6:
            strength = "Weak"
        case len(req.Password) < 10:
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