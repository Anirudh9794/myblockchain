package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteJsonResponse writes a response ResponseWriter
func WriteJsonResponse(writer http.ResponseWriter, resp interface{}) {
	writer.Header().Set("Content-Type", "application/json")

	respBytes, err := json.Marshal(resp)
	if err != nil {
		WriteError(writer, err)
		return
	}

	fmt.Fprint(writer, string(respBytes))
}

// WriteError writes a json error to the ResponseWriter
func WriteError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusInternalServerError)

	resp := fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
	fmt.Fprint(writer, resp)
}
