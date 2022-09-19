package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Data          interface{} `json:"data"`
	Message       string      `json:"message"`
	contentType   string
	respondeWrite http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		respondeWrite: rw,
		contentType:   "application/json",
	}
}

func (res *Response) Send() {
	res.respondeWrite.Header().Set("Content-Type", res.contentType)
	res.respondeWrite.WriteHeader(res.Status)

	output, _ := json.Marshal(&res.Data)
	fmt.Fprintln(res.respondeWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource not found"
}

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity not found"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()
}
