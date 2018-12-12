package models

type DaoRequest struct {
	ID      string              `db:"ID" json:"id"`
	Method  string              `db:"Method" json:"method"`
	Headers map[string][]string `db:"Headers" json:"headers,omitempty"`
	Address string              `db:"Address" json:"address"`
	Body    string              `db:"Body" json:"body,omitempty"`
}

type DaoResponse struct {
	ID            string              `db:"ID" json:"id"`
	HttpStatus    string              `db:"HttpStatus" json:"httpStatus"`
	Headers       map[string][]string `db:"Headers" json:"headers"`
	ContentLength int64               `db:"ContentLength" json:"contentLength"`
}
