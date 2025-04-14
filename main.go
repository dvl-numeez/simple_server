package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Server up and runnning",
		}
		data, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(200)
		_, err = w.Write(data)
		if err != nil {
			log.Println(err)
		}

	})
	log.Println("Server up .............")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
