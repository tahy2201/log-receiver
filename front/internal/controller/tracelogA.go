package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"log-receiver/internal/common"
	"log-receiver/internal/config"

	"go.uber.org/zap"
)

func TracelogA(w http.ResponseWriter, r *http.Request) {
	logger := config.GenerateLogger()
	logger.Debug("call tracelogA")

	// parse and validate json
	bs, _ := ioutil.ReadAll(r.Body)
	bsJson, jsErr := common.GetParamMap(common.TRACE_LOG_A, bs)
	if jsErr != nil {
		eMsg := "Error parse json"
		logger.Error(eMsg, zap.Error(jsErr))
		writeErrorResponse(w, eMsg)
		return
	}
	logger.Debug("Publish:", zap.Any("json", bsJson))

	msgId, eMsg, err := publishMessage(r, common.TRACE_LOG_A, bsJson)
	if err != nil {
		logger.Error(eMsg, zap.Error(err))
		writeErrorResponse(w, eMsg)
		return
	}
	logger.Info("Published message with custom attributes.", zap.String("msg ID", msgId))

	res, _ := json.Marshal(StdResponse{0, "", msgId})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
