package logger

import (
	"go.uber.org/zap"
)

type CommonLogitAdapter struct {
	commonLogger *zap.SugaredLogger
	productID    string
	environment  string
	application  string
}

// NewLogitAdapter ... which uses general log interface
func NewCommonLogitAdapter(moduleName, logLevel, productID, environment, application string, initialFields map[string]interface{}) *CommonLogitAdapter {
	zapLogger := NewCommonLogger(moduleName, logLevel, productID, environment, application, initialFields)
	return &CommonLogitAdapter{
		commonLogger: zapLogger.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)),
		productID:    productID,
		environment:  environment,
		application:  application,
	}
}

func addKeyVals(productID, environment, application, activity, event string, otherKeyVals ...any) []any {
	var keyVals []any
	keyVals = append(keyVals,
		"activity", activity,
		"event", event,
		"product", productID,
		"environment", environment,
		"application", application,
	)
	keyVals = append(keyVals, otherKeyVals...)

	return keyVals
}

func (log *CommonLogitAdapter) Debug(msg string, keyvals ...any) {
	keyvals = addKeyVals(log.productID, log.environment, log.application, "activity", "event", keyvals...)
	log.commonLogger.Debugf(msg, keyvals...)
}

func (log *CommonLogitAdapter) Info(msg string, keyvals ...any) {
	log.commonLogger.Infof(msg, keyvals...)
}

func (log *CommonLogitAdapter) Warn(msg string, keyvals ...any) {
	log.commonLogger.Warnf(msg, keyvals...)
}

func (log *CommonLogitAdapter) Error(msg string, keyvals ...any) {
	log.commonLogger.Errorf(msg, keyvals...)
}

func (log *CommonLogitAdapter) Debugw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(log.productID, log.environment, log.application, activity, event, otherKeyVals...)
	log.commonLogger.Debugw(message, keyVals...)
}

func (log *CommonLogitAdapter) Infow(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(log.productID, log.environment, log.application, activity, event, otherKeyVals...)
	log.commonLogger.Infow(message, keyVals...)
}

func (log *CommonLogitAdapter) Warnw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(log.productID, log.environment, log.application, activity, event, otherKeyVals...)
	log.commonLogger.Warnw(message, keyVals...)
}

func (log *CommonLogitAdapter) Errorw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(log.productID, log.environment, log.application, activity, event, otherKeyVals...)
	log.commonLogger.Errorw(message, keyVals...)
}

func (log *CommonLogitAdapter) Fatalw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(log.productID, log.environment, log.application, activity, event, otherKeyVals...)
	log.commonLogger.Fatalw(message, keyVals...)
}
