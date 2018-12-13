package workers

import (
	. "HttpRequestExecutionService/interfaces"
	. "HttpRequestExecutionService/models"
)

func WorkerReqSaver(in <-chan DaoRequest, storage RepositoryReq) {
	for item := range in {
		storage.Put(item.ID, item)
	}
}

func WorkerRespSaver(in <-chan DaoResponse, storage RepositoryResp) {
	for item := range in {
		storage.Put(item.ID, item)
	}
}
