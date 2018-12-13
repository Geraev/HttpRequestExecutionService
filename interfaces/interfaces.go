package interfaces

import (
	. "HttpRequestExecutionService/models"
)

type RepositoryReq interface {
	Put(string, DaoRequest)
	Get(string) (DaoRequest, bool)
	Delete(string)
}

type RepositoryResp interface {
	Put(string, DaoResponse)
	Get(string) (DaoResponse, bool)
	Delete(string)
}
