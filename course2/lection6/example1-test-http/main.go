package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RequestBody struct {
	Number int `json:"number"`
}

type ResponseBody struct {
	Factorial int `json:"factorial"`
}

func main() {
	http.HandleFunc("/factorial", GetHttpFactorialFunc)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

var hash = make(map[int]int)

func GetHttpFactorialFunc(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	bodyString, err := io.ReadAll(request.Body)
	defer request.Body.Close()

	if err != nil {
		log.Println("body format error")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var body RequestBody
	err = json.Unmarshal(bodyString, &body)

	if err != nil {
		log.Println("body format error")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	hashNum, ok := hash[body.Number]

	if ok {
		_ = json.NewEncoder(response).Encode(ResponseBody{Factorial: hashNum})
		return
	}

	result := factorial(body.Number)
	hash[body.Number] = result
	_ = json.NewEncoder(response).Encode(ResponseBody{Factorial: result})
}

func factorial(num int) int {
	res := 1
	if num < 2 {
		return res
	}
	for i := 2; i <= num; i++ {
		res *= i
	}
	return res
}
