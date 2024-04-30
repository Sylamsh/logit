package logger

import (
	"encoding/json"
	"sync"

	"github.com/nvg14/logit/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

const (
	KEY_PRODUCT_ID       = "product"
	KEY_ENVIRONMENT      = "environment"
	KEY_APPLICATION_NAME = "application"
)

func NewSugaredLogger(moduleName, logLevel, productID, environment, application string, initialFields map[string]interface{}) *zap.SugaredLogger {
	initialFields[KEY_PRODUCT_ID] = productID
	initialFields[KEY_ENVIRONMENT] = environment
	initialFields[KEY_APPLICATION_NAME] = application
	return NewZapLogger(moduleName, logLevel, initialFields).Sugar()
}

func NewZapLogger(moduleName, logLevel string, initialFields map[string]interface{}) *zap.Logger {
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
	zapLogger, _ := config.Build()

	return zapLogger
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

func init() {
	once.Do(func() {
		zapLogger := NewZapLogger(
			utils.GetEnvWithDefault("MODULE_NAME", "fabric"),
			utils.GetEnvWithDefault("LOG_LEVEL", "info"),
			map[string]interface{}{
				KEY_PRODUCT_ID:       utils.GetEnvWithDefault("PRODUCT_ID", "fabric"),
				KEY_ENVIRONMENT:      utils.GetEnvWithDefault("Stack", "jammer-dev"),
				KEY_APPLICATION_NAME: utils.GetEnvWithDefault("APPLICATION_NAME", "master"),
			},
		)
		defer zapLogger.Sync()
		intializeLogAdapter(zapLogger.Sugar())
	})
}
