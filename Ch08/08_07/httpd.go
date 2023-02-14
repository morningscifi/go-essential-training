/*
	Key/Value database

$ curl -d'hello' http://localhost:8080/k1
$ curl http://localhost:8080/k1
hello
$ curl -i http://localhost:8080/k2
404 not found

Limit value size to 1k
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	db DB
}

type dbResponse struct {
	Error string `json:"error"`
	Data  []byte `json:"data"`
}

type dbRequest struct {
	Key string `json:"key"`
}

// POST /key Store request body as value
// GET /<key> Send back value, or 404 if key not found
func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var dbr dbRequest
	dec := json.NewDecoder(r.Body)

	fmt.Printf("Received:\n%v#+", r.Body)

	if err := dec.Decode(&dbr); err != nil {
		http.Error(w, "decode failed", http.StatusBadRequest)
		log.Printf("error: %s", err)
		return
	}

	resp := dbResponse{Error: "", Data: s.db.Get(dbr.Key)}
	// Encode result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(&resp); err != nil {
		log.Printf("can't encode %v - %s", resp, err)
	}
}

func main() {
	// TODO: Routing & start server
	var s Server
	s.db.Set("yo", []byte("dog"))
	http.HandleFunc("/get", s.getHandler)

	addr := ":8080"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
