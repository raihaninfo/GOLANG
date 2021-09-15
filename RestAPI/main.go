package main

import (
	"GOLANG/RestAPI/structs"
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "ping",
			}

			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":8000", mux)
}
