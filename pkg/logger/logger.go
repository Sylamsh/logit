package logger

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewCommonLogger(moduleName, logLevel, productID, environment, application string, initialFields map[string]interface{}) *zap.SugaredLogger {
	var rawJSON = []byte(`{
		"level": "` + logLevel + `",
		"encoding": "json",
		"encoderConfig": {
			"levelKey": "level",
			"messageKey": "message",
			"levelEncoder": "lowercase",
			"nameKey":"name",
			"stacktraceKey":"stack"
		}
	}`)

	config := setUpConfig(moduleName, rawJSON, initialFields)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	logit, _ := config.Build()

	return logit.Sugar()
}

func NewFastLogger(moduleName, logLevel string, initialFields map[string]interface{}) *zap.Logger {
	var rawJSON = []byte(`{
		"level": "` + logLevel + `",
		"encoding": "json",
		"encoderConfig": {
			"levelKey": "level",
			"messageKey": "message",
			"levelEncoder": "lowercase",
			"nameKey":"name",
			"stacktraceKey":"stack"
		}
	}`)

	config := setUpConfig(moduleName, rawJSON, initialFields)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	logit, _ := config.Build()

	return logit
}

func setUpConfig(moduleName string, rawJSON []byte, initialFields map[string]interface{}) (config zap.Config) {
	_ = json.Unmarshal(rawJSON, &config)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05Z")
	if config.InitialFields == nil {
		config.InitialFields = make(map[string]interface{})
	}
	config.InitialFields["module"] = moduleName
	for k, v := range initialFields {
		config.InitialFields[k] = v
	}
	return config
}
