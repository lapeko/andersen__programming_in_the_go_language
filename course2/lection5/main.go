package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

type Payload struct {
	A int
	B int
	C int
}

type Response struct {
	A      int
	B      int
	C      int
	NRoots int
}

var lastPayload *Payload

func main() {
	http.HandleFunc("/grab", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		payload := &Payload{}
		err := json.NewDecoder(r.Body).Decode(payload)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		lastPayload = payload
		w.WriteHeader(http.StatusAccepted)
	})

	http.HandleFunc("/solve", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if lastPayload == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if lastPayload.A == 0 || lastPayload.B == 0 || lastPayload.C == 0 {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Response{
				A:      lastPayload.A,
				B:      lastPayload.B,
				C:      lastPayload.C,
				NRoots: 0,
			})
			return
		}

		descr := lastPayload.B*lastPayload.B - 4*lastPayload.A*lastPayload.C
		nRoots := 0

		if descr > 0 {
			nRoots = 2
		} else if descr == 0 {
			nRoots = 1
		} else {
			nRoots = 0
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Response{
			A:      lastPayload.A,
			B:      lastPayload.B,
			C:      lastPayload.C,
			NRoots: nRoots,
		})
	})

	log.Printf("Server is running on: http://localhost:%d", PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
