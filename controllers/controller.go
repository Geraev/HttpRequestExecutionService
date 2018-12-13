package controllers

import (
	. "HttpRequestExecutionService/models"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func GenerateID() string {
	return fmt.Sprint(uuid.Must(uuid.NewV4()))
}

type Controller struct{}

func (c Controller) QueueHandler(out chan<- DaoRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newRequest DaoRequest
		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			log.Fatal(err)
		}
		newRequest.ID = GenerateID()
		out <- newRequest
		w.Header().Set("Content-Type", "text/plain")
		//noinspection GoUnhandledErrorResult
		fmt.Fprintf(w, "In progress. ID = %s", newRequest.ID)
	}
}

