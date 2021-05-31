package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		// Fprint  = formats to i/o ?? thats about all i am seeing as a 'bonus'
		// Printf  = formats and writes to standard output
		// Sprintf = formats without printing ex value := Sprintf("Hey %s", field)
		fmt.Fprint(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}

	JSON(w, http.StatusBadRequest, nil)
}
