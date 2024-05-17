package main

import "github.com/nvg14/logit/pkg/logger"

func main() {
	logger.Info("This is a message", ActivityExample, EventExample, "key1", "value1")

	loggerVariable := logger.GetCommonLogger()
	loggerVariable.Info("Logger Adapter variable is meant to be passed to other third party frameworks like Temporal", "key1", "value1")
	loggerVariable.Infow("Recommended to use the 'w' versions", ActivityExample, EventExample, "key1", "value1")

	CustomConfigurationLogger := logger.NewSugaredLogger("info", "jammer", "jammer-bugbash", "api-server", map[string]interface{}{"KEY": "VALUE"})
	CustomConfigurationLogger.Warnf("NOT RECOMMENDED TO USE", "key1", "value1")
}
