package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type CountResp struct {
	Count int `json:count`
}

var lock sync.Mutex
var count = 0

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Request")
	if r.Method == "POST" {
		var data map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(data)

		lock.Lock()
		count = count + 1
		resp := &CountResp{Count: count}
		lock.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Listening")
	err := http.ListenAndServe("0.0.0.0:9000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
