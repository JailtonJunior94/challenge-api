package responses

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.Errorf(err.Error())
	}
}

func Error(w http.ResponseWriter, statusCode int, errorMessage string) {
	JSON(w, statusCode, struct {
		ErrorMessage string `json:"errorMessage"`
	}{
		ErrorMessage: errorMessage,
	})
}
