package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var factorialTestCases = []struct {
	input    int
	expected int
}{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{5, 120},
	{6, 720},
}

func TestFactorialTests(t *testing.T) {
	for idx, testCase := range factorialTestCases {
		if res := factorial(testCase.input); res != testCase.expected {
			t.Fatalf("case %d. factorial(%d) returns %d, when %d expected\n", idx, testCase.input, res, testCase.expected)
		}
	}
}

func TestHttpFactorial(t *testing.T) {
	for idx, testCase := range factorialTestCases {
		t.Run(fmt.Sprintf("Http factorial test number: %d", idx), func(t *testing.T) {
			jsonBody, _ := json.Marshal(RequestBody{Number: testCase.input})
			request := httptest.NewRequest(http.MethodPost, "/factorial", bytes.NewBuffer(jsonBody))
			recorder := httptest.NewRecorder()
			GetHttpFactorialFunc(recorder, request)
			res := recorder.Result()

			var responseBody ResponseBody
			responseBodyString, err := io.ReadAll(res.Body)

			if err != nil {
				t.Errorf("expexted error to be nil. got %v", err)
			}

			defer res.Body.Close()

			err = json.Unmarshal(responseBodyString, &responseBody)
			if err != nil {
				t.Errorf("expexted error to be nil. got %v", err)
			}

			if responseBody.Factorial != testCase.expected {
				t.Errorf("expexted factorial should be %v. got %v", testCase.expected, responseBody.Factorial)
			}
		})
	}
}
