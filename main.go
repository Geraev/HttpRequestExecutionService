package main

import (
	"HttpRequestExecutionService/controllers"
	. "HttpRequestExecutionService/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type requestRepo interface {
	Put()
	Get()
	Delete()
}

func Worker(in <-chan DaoRequest) {
	for request := range in {
		fmt.Printf("Hey, new request %+v \n", request)
	}
	fmt.Printf("Worker stoped")
}

func main() {
	router := mux.NewRouter()
	jobs := make(chan DaoRequest, 500)
	go Worker(jobs)

	controller := controllers.Controller{}

	router.HandleFunc("/", controller.QueueHandler(jobs)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", router))

}




