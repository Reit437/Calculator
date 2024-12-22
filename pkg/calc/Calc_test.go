package calc_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Reit437/Calculator/pkg/calc"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

func TestCalculatorServer(t *testing.T) {
	testCases := []struct {
		name        string
		expression  string
		expected    float64
		expectedErr bool
	}{
		{"Addition", "1+2", 3, false},
		{"Subtraction", "10-5", 5, false},
		{"Multiplication", "3*4", 12, false},
		{"Division", "10/2", 5, false},
		{"ComplexExpression", "(1+2)*3+4", 13, false},
		{"InvalidExpression", "1++2", 0, true},
		{"DivisionByZero", "10/0", 0, true},
		{"Parentheses", "(5+5)/2", 5, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqBody, err := json.Marshal(CalculateRequest{Expression: tc.expression})
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			resp, err := http.Post(
				"http://localhost:8080/api/v1/calculate",
				"application/json",
				bytes.NewBuffer(reqBody),
			)
			if err != nil {
				if tc.expectedErr {
					t.Logf("Expected error for %s: %v", tc.name, err)
					return
				}
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK && !tc.expectedErr {
				t.Errorf("Expected status code 200, got %d", resp.StatusCode)
				return
			}
			if resp.StatusCode == http.StatusOK {

				var respBody CalculateResponse
				if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
					t.Fatalf("Failed to unmarshal response body: %v", err)
				}

				if respBody.Result != tc.expected {
					t.Errorf("Expected result %f, got %f", tc.expected, respBody.Result)
				}
			}
		})
	}
}

func TestCalcFunction(t *testing.T) {
	testCases := []struct {
		name        string
		expression  string
		expected    float64
		expectedErr bool
	}{
		{"Addition", "1+2", 3, false},
		{"Subtraction", "10-5", 5, false},
		{"Multiplication", "3*4", 12, false},
		{"Division", "10/2", 5, false},
		{"ComplexExpression", "(1+2)*3+4", 13, false},
		{"InvalidExpression", "1++2", 0, true},
		{"DivisionByZero", "10/0", 0, true},
		{"Parentheses", "(5+5)/2", 5, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Calc(tc.expression)
			if err != nil {
				if !tc.expectedErr {
					t.Errorf("Unexpected error: %v", err)
				}
				return
			}
			if tc.expectedErr {
				t.Errorf("Expected error, but got result: %f", result)
				return
			}
			if result != tc.expected {
				t.Errorf("Expected %f, got %f", tc.expected, result)
			}
		})
	}
}
