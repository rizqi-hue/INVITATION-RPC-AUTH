package logger

import (
	"context"
	"strings"

	"github.com/INVITATION-RPC-AUTH/pkg/config"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func Configure() {
	format := config.GetString("log_format")
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	lvl := strings.ToLower(config.GetString("log_level"))
	switch lvl {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	log = logrus.WithField("app", config.GetString("name"))
}

func setDefault() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	log = logrus.WithField("app", config.GetString("name"))
}

func GetLogger(pkg, funcName string) *logrus.Entry {
	if log == nil {
		setDefault()
	}
	return log.WithFields(logrus.Fields{
		"function": funcName,
		"package":  pkg,
	})
}

func GetLoggerContext(ctx context.Context, pkg, funcName string) *logrus.Entry {
	if log == nil {
		setDefault()
	}
	return log.WithContext(ctx).WithFields(logrus.Fields{
		"function": funcName,
		"package":  pkg,
	})
}
