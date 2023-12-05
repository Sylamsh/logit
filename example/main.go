package main

import "github.com/nvg14/logit/pkg/logger"

func main() {
	lg := logger.NewCommonLogitAdapter("fabric", "debug", "jammer", "development", "master", map[string]interface{}{"project": "jammer"})
	lg.Infow("This is a message", "activity_name", "event_name")
}
