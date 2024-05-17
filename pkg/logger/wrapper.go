package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
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

func addKeyVals(activity LogActivity, event LogEvent, otherKeyVals ...any) []any {
	var keyVals []any

	keyVals = append(keyVals,
		LGR_KEY_ACTIVITY, string(activity),
		LGR_KEY_EVENT, string(event),
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

func (adapter *CommonLogAdapter) Debugw(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Debugw(message, keyVals...)
}

func (adapter *CommonLogAdapter) Infow(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Infow(message, keyVals...)
}

func (adapter *CommonLogAdapter) Warnw(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Warnw(message, keyVals...)
}

func (adapter *CommonLogAdapter) Errorw(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Errorw(message, keyVals...)
}

func (adapter *CommonLogAdapter) Fatalw(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	adapter.commonLogger.Fatalw(message, keyVals...)
}

func Debug(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Debugw(message, keyVals...)
}

func Info(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Infow(message, keyVals...)
}

func Warn(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Warnw(message, keyVals...)
}

func Error(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Errorw(message, keyVals...)
}

func Fatal(message string, activity LogActivity, event LogEvent, otherKeyVals ...any) {
	keyVals := addKeyVals(activity, event, otherKeyVals...)
	commonLogAdapter.commonLogger.Fatalw(message, keyVals...)
}

// context containing
//    span created by otel
//    activity set with key CTX_KEY_ACTIVITY(required)
//    amg_id set with key CTX_KEY_AMGID
//    channel_id set with key CTX_KEY_CHANNELID

func DebugCtx(ctx context.Context, message string, event LogEvent, otherKeyVals ...any) {
	otherKeyVals = append(otherKeyVals, getValuesFromContext(ctx)...)
	keyVals := addKeyVals(getActivity(ctx), event, otherKeyVals...)
	commonLogAdapter.commonLogger.Debugw(message, keyVals...)
}

func InfoCtx(ctx context.Context, message string, event LogEvent, otherKeyVals ...any) {
	otherKeyVals = append(otherKeyVals, getValuesFromContext(ctx)...)
	keyVals := addKeyVals(getActivity(ctx), event, otherKeyVals...)
	commonLogAdapter.commonLogger.Infow(message, keyVals...)
}

func WarnCtx(ctx context.Context, message string, event LogEvent, otherKeyVals ...any) {
	otherKeyVals = append(otherKeyVals, getValuesFromContext(ctx)...)
	keyVals := addKeyVals(getActivity(ctx), event, otherKeyVals...)
	commonLogAdapter.commonLogger.Warnw(message, keyVals...)
}

func ErrorCtx(ctx context.Context, message string, event LogEvent, otherKeyVals ...any) {
	otherKeyVals = append(otherKeyVals, getValuesFromContext(ctx)...)
	keyVals := addKeyVals(getActivity(ctx), event, otherKeyVals...)
	commonLogAdapter.commonLogger.Errorw(message, keyVals...)
}

func FatalCtx(ctx context.Context, message string, event LogEvent, otherKeyVals ...any) {
	otherKeyVals = append(otherKeyVals, getValuesFromContext(ctx)...)
	keyVals := addKeyVals(getActivity(ctx), event, otherKeyVals...)
	commonLogAdapter.commonLogger.Fatalw(message, keyVals...)
}

func getActivity(ctx context.Context) LogActivity {
	activity, ok := ctx.Value(CTX_KEY_ACTIVITY).(string)
	if !ok {
		panic("Activity is not set in ctx for logging")
	}
	return LogActivity(activity)
}

func getValuesFromContext(ctx context.Context) []any {
	var keyVals []any
	amgid, ok := ctx.Value(CTX_KEY_AMGID).(string)
	if ok {
		keyVals = append(keyVals, LGR_KEY_AMGID, amgid)
	}
	channelid, ok := ctx.Value(CTX_KEY_CHANNELID).(string)
	if ok {
		keyVals = append(keyVals, LGR_KEY_CHANNELID, channelid)
	}
	traceContext := getTraceContext(ctx)
	if traceContext != nil {
		keyVals = append(keyVals, LGR_KEY_TRACECONTEXT, traceContext)
	}
	return keyVals
}

func getTraceContext(ctx context.Context) map[string]string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.HasTraceID() || !spanCtx.HasSpanID() {
		return nil
	}
	return map[string]string{
		LGR_KEY_TRACEPARENT: "00-" + spanCtx.TraceID().String() + "-" + spanCtx.SpanID().String() + "-" + spanCtx.TraceFlags().String(),
		LGR_KEY_TRACESTATE:  spanCtx.TraceState().String(),
	}
}
