package gotils

import (
	"encoding/json"
	"io"
	"net/http"
)

type httpResponse struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

//Check will log output to console on error
func Check(err error) {
	if(err != nil){
		panic(err)
	}
} 

// SendError - writes error message to io.Writer
func SendError(err error, w io.Writer) {
	w.Write([]byte(err.Error()))
}

// SendHTTPError - sends 500 error with http writer
func SendHTTPError(err error, w http.ResponseWriter) {
	d := httpResponse{
		StatusCode: 500,
		Message: err.Error(),
	}

	json.NewEncoder(w).Encode(d)
}