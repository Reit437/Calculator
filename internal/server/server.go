package server

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/Reit437/Calculator/pkg/calc"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func isValidExpression(expression string) bool {
	matched, err := regexp.MatchString(`^[\d+\-*/().]+$`, expression)
	if err != nil {
		return false
	}
	return matched
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var req CalculateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !isValidExpression(req.Expression) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		errResp := ErrorResponse{Error: "Expression is not valid", Code: http.StatusUnprocessableEntity}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	result, err := calc.Calc(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errResp := ErrorResponse{Error: err.Error(), Code: http.StatusInternalServerError}
		json.NewEncoder(w).Encode(errResp)
		return
	}

	resp := CalculateResponse{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
