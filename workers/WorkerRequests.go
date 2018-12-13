package workers

import (
	. "HttpRequestExecutionService/models"
	"bytes"
	"log"
	"net/http"
	"time"
)

func WorkerRequests(in <-chan DaoRequest, outReq chan<- DaoRequest, outResp chan<- DaoResponse) {
	for jobRequest := range in {
		outReq <- jobRequest
		client := &http.Client{Timeout: time.Second * 2}
		request, err := http.NewRequest(
			jobRequest.Method,
			jobRequest.Address,
			bytes.NewBuffer([]byte(jobRequest.Body)),
		)
		if err != nil {
			log.Fatal(err)
		}

		if jobRequest.Headers != nil {
			for k, v := range jobRequest.Headers {
				for _, val := range v {
					request.Header.Add(k, val)
				}
			}
		}

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}

		outResp <- DaoResponse{
			ID:            jobRequest.ID,
			HttpStatus:    response.Status,
			Headers:       response.Header,
			ContentLength: response.ContentLength,
		}
	}
}
