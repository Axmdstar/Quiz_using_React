package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func responsewithJSON(w http.ResponseWriter, statuscode int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Encode json payload : %v \n", err)
	}
	w.Write(jsonData)
}

func responseWithERROR(w http.ResponseWriter, statuscode int, err string) {
	responsewithJSON(w, statuscode, map[string]string{"error": err})
}
