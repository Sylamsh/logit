package logger

import (
	"go.uber.org/zap"
)

type CommonLogAdapter struct {
	commonLogger *zap.SugaredLogger
}

var commonLogAdapter CommonLogAdapter

func intializeLogAdapter(logger *zap.SugaredLogger) {
	commonLogAdapter = CommonLogAdapter{
		commonLogger: logger.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)),
	}
}

func GetCommonLogger() *CommonLogAdapter {
	return &commonLogAdapter
}

func addKeyVals(activity, event string, otherKeyVals ...any) []any {
	var keyVals []any
	keyVals = append(keyVals,
		"activity", activity,
		"event", event,
	)
	keyVals = append(keyVals, otherKeyVals...)

	return keyVals
}

func (adapter *CommonLogAdapter) Debug(msg string, keyvals ...any) {
	adapter.commonLogger.Debugf(msg, keyvals...)
}

func (adapter *CommonLogAdapter) Info(msg string, keyvals ...any) {
	adapter.commonLogger.Infof(msg, keyvals...)
}

func (adapter *CommonLogAdapter) Warn(msg string, keyvals ...any) {
	adapter.commonLogger.Warnf(msg, keyvals...)
}

func (adapter *CommonLogAdapter) Error(msg string, keyvals ...any) {
	adapter.commonLogger.Errorf(msg, keyvals...)
}

func (adapter *CommonLogAdapter) Debugw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Debugw(message, keyVals...)
}

func (adapter *CommonLogAdapter) Infow(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Infow(message, keyVals...)
}

func (adapter *CommonLogAdapter) Warnw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Warnw(message, keyVals...)
}

func (adapter *CommonLogAdapter) Errorw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Errorw(message, keyVals...)
}

func (adapter *CommonLogAdapter) Fatalw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Fatalw(message, keyVals...)
}

func Debug(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Debugw(message, keyVals...)
}

func Info(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Infow(message, keyVals...)
}

func Warn(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Warnw(message, keyVals...)
}

func Error(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Errorw(message, keyVals...)
}

func Fatal(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Fatalw(message, keyVals...)
}
