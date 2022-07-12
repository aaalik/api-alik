package alog

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

var Logger *logrus.Logger

type Logging interface {
	InitLogging()
	GetTracer()
}

func Init() {
	Logger = logrus.New()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   os.Getenv(LogLocation),
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     90, //days
		Level:      logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
		},
	})

	if err != nil {
		Logger.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	Logger.SetReportCaller(true)

	if os.Getenv(AppEnv) == "production" {
		Logger.AddHook(rotateFileHook)
	}
}

func GetTracer(err error) string {
	pc, fn, line, _ := runtime.Caller(1)
	tracer := fmt.Sprintf("[error] in %s[%s:%d] %v \n", runtime.FuncForPC(pc).Name(), fn, line, err)

	return tracer
}
