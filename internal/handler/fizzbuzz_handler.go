package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/fizz-buzz-adevinta/internal/service"
)

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    int1, _ := strconv.Atoi(params.Get("int1"))
    int2, _ := strconv.Atoi(params.Get("int2"))
    limit, _ := strconv.Atoi(params.Get("limit"))
    str1 := params.Get("str1")
    str2 := params.Get("str2")

    if int1 <= 0 || int2 <= 0 {
		http.Error(w, "Invalid Integer parameters, must be positive or a number", http.StatusBadRequest)
        return
	}
	if limit <= 0 || limit > 1000 {
		http.Error(w, "Invalid limit parameters, must be between 1 and 1000", http.StatusBadRequest)
        return
	}
	if str1 == "" || str2 == "" {
		http.Error(w, "Invalid string parameters, must be not empty", http.StatusBadRequest)
        return
	}

    result := service.GenerateFizzBuzz(int1, int2, limit, str1, str2)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}