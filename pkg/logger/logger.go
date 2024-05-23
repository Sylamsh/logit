package logger

import (
	"encoding/json"
	"sync"

	"github.com/nvg14/logit/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var once sync.Once

func NewSugaredLogger(logLevel, productID, environment, application string, initialFields map[string]interface{}) *zap.SugaredLogger {
	initialFields[LGR_KEY_PRODUCT_ID] = productID
	initialFields[LGR_KEY_ENVIRONMENT] = environment
	initialFields[LGR_KEY_APPLICATION_ID] = application
	return NewZapLogger(logLevel, initialFields).Sugar()
}

func NewZapLogger(logLevel string, initialFields map[string]interface{}) *zap.Logger {
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

	config := setUpConfig(rawJSON, initialFields)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	zapLogger, _ := config.Build()

	return zapLogger
}

func setUpConfig(rawJSON []byte, initialFields map[string]interface{}) (config zap.Config) {
	_ = json.Unmarshal(rawJSON, &config)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05Z")
	if config.InitialFields == nil {
		config.InitialFields = make(map[string]interface{})
	}
	for k, v := range initialFields {
		config.InitialFields[k] = v
	}
	return config
}

func init() {
	once.Do(func() {
		zapLogger := NewZapLogger(
			utils.GetEnvWithDefault("LOG_LEVEL", "info"),
			map[string]interface{}{
				LGR_KEY_PRODUCT_ID:     utils.GetEnvWithDefault("PRODUCT_ID", "amaginow"),
				LGR_KEY_ENVIRONMENT:    utils.GetEnvWithDefault("ENVIRONMENT", "development"),
				LGR_KEY_APPLICATION_ID: utils.GetEnvWithDefault("APPLICATION_ID", "master"),
			},
		)
		defer zapLogger.Sync()
		intializeLogAdapter(zapLogger.Sugar())
	})
}
