package http_utils

import (
	"encoding/json"
	"github.com/comfysweet/bookstore_utils-go/errors"
	"github.com/comfysweet/bookstore_utils-go/logger"
	"go.uber.org/zap"
	"net/http"
)

func RespondJson(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err errors.RestErr) {
	RespondJson(w, err.Status(), err)
}

const (
	// Успешный статус события
	successEvent = "SUCCESS"
	// Ошибочный статус события
	errorEvent = "ERROR"
)

func LogSuccessEvent(msg, info string, event string) {
	logger.Info(
		msg,
		zap.String("information", info),
		zap.String("uid", event),
		zap.String("name", event),
		zap.String("namespace", event),
		zap.String("operation", event),
		zap.String("status", successEvent),
	)
}

// LogErrorEvent - логирование эвента с ошибкой
func LogErrorEvent(err error, info string, event string) {
	LogSuccessEvent(err.Error(), info, event)
	logger.Info(
		err.Error(),
		zap.String("information", info),
		zap.String("uid", event),
		zap.String("name", event),
		zap.String("namespace", event),
		zap.String("operation", event),
		zap.String("status", errorEvent),
	)
}
