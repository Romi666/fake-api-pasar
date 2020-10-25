package models

type ResponseValidation struct {
	ResponseCode string `json:"RC"`
	Deskripsi    string `json:"Deskripsi"`
	Data         struct {
		DeskripsiError interface{} `json:"error"`
	} `json:"Data"`
}


type DataResponse struct {
	DeskripsiError string `json:"error"`
	DateFromat     string `json:"issue_at"`
}

type Response struct {
	ResponseCode string       `json:"RC"`
	Deskripsi    string       `json:"Deskripsi"`
	Data         DataResponse `json:"Data"`
}