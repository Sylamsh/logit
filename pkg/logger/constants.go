package logger

type LogEvent string
type LogActivity string
type ContextKey string

const (
	CTX_KEY_ACTIVITY  ContextKey = "ACTIVITY"
	CTX_KEY_AMGID     ContextKey = "AMG_ID"
	CTX_KEY_CHANNELID ContextKey = "CHANNEL_ID"
)

const (
	LGR_KEY_ENVIRONMENT    = "environment"
	LGR_KEY_APPLICATION_ID = "application_id"
	LGR_KEY_PRODUCT_ID     = "product_id"

	LGR_KEY_AMGID     = "amg_id"
	LGR_KEY_CHANNELID = "channel_id"

	LGR_KEY_TRACECONTEXT = "tracecontext"
	LGR_KEY_TRACEPARENT  = "traceparent"
	LGR_KEY_TRACESTATE   = "tracestate"

	LGR_KEY_ACTIVITY = "activity"
	LGR_KEY_EVENT    = "event"
)
