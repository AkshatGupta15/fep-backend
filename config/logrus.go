package config

import "github.com/sirupsen/logrus"

func logrusConfig() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
}
