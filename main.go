package main

import (
	"HttpRequestExecutionService/controllers"
	. "HttpRequestExecutionService/models"
	"HttpRequestExecutionService/repository"
	. "HttpRequestExecutionService/workers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	jobs := make(chan DaoRequest, 10)
	queueReq := make(chan DaoRequest, 10)
	queueResp := make(chan DaoResponse, 10)

	mapReqStorage := repository.NewMapRequest()
	mapRespStorage := repository.NewMapResponse()

	go WorkerRequests(jobs, queueReq, queueResp)
	go WorkerReqSaver(queueReq, mapReqStorage)
	go WorkerRespSaver(queueResp, mapRespStorage)

/*	go func() {
		time.Sleep(10 * time.Second)
		for ch1 := range queueReq {
			fmt.Printf("queueReq: %+v \n", ch1)

		}
	}()
	go func() {
		time.Sleep(15 * time.Second)
		for ch2 := range queueResp {
			fmt.Printf("queueResp: %+v \n", ch2)
		}
	}()
*/
	controller := controllers.Controller{}
	router.HandleFunc("/", controller.QueueHandler(jobs)).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))

}
