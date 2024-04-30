package logger

import (
	"go.uber.org/zap"
)

type CommonLogAdapter struct {
	commonLogger *zap.SugaredLogger
}

var commonLogger CommonLogAdapter

func intializeLogAdapter(logger *zap.SugaredLogger) {
	commonLogger = CommonLogAdapter{
		commonLogger: logger.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)),
	}
}

func GetCommonLogger() *CommonLogAdapter {
	return &commonLogger
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

func (log *CommonLogAdapter) Debug(msg string, keyvals ...any) {
	log.commonLogger.Debugf(msg, keyvals...)
}

func (log *CommonLogAdapter) Info(msg string, keyvals ...any) {
	log.commonLogger.Infof(msg, keyvals...)
}

func (log *CommonLogAdapter) Warn(msg string, keyvals ...any) {
	log.commonLogger.Warnf(msg, keyvals...)
}

func (log *CommonLogAdapter) Error(msg string, keyvals ...any) {
	log.commonLogger.Errorf(msg, keyvals...)
}

func (log *CommonLogAdapter) Debugw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	log.commonLogger.Debugw(message, keyVals...)
}

func (log *CommonLogAdapter) Infow(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	log.commonLogger.Infow(message, keyVals...)
}

func (log *CommonLogAdapter) Warnw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	log.commonLogger.Warnw(message, keyVals...)
}

func (log *CommonLogAdapter) Errorw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	log.commonLogger.Errorw(message, keyVals...)
}

func (log *CommonLogAdapter) Fatalw(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	log.commonLogger.Fatalw(message, keyVals...)
}

func Debug(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogger.commonLogger.Debugw(message, keyVals...)
}

func Info(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogger.commonLogger.Infow(message, keyVals...)
}

func Warn(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogger.commonLogger.Warnw(message, keyVals...)
}

func Error(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogger.commonLogger.Errorw(message, keyVals...)
}

func Fatal(message string, activity string, event string, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogger.commonLogger.Fatalw(message, keyVals...)
}
