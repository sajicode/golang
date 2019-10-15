package main

import "log/syslog"

func main() {
	//* creates a new syslog client
	logger, err := syslog.New(syslog.LOG_LOCAL3, "narwhal")
	if err != nil {
		panic("Cannot attach to syslog")
	}
	defer logger.Close()

	//* sends the logger a variety of messages
	logger.Debug("Debug message")
	logger.Notice("Notice message")
	logger.Warning("Warning message")
	logger.Alert("Alert message")
}
