package log

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/SyarifKA/crowdfunding-api/pkg/env"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger        *logrus.Logger
	JSONFormatter logrus.JSONFormatter
	TextFormatter logrus.TextFormatter = logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	}

	mu          sync.Mutex
	currentTime string
	currentFile *os.File
)

type (
	Level = logrus.Level
)

const (
	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	TraceLevel = logrus.TraceLevel
)

type Config struct {
	logrus.Formatter
	logrus.Level
	LogName string
}

func InitLogger(cfg *Config) error {
	if cfg.LogName == "" {
		return errors.New("log name is empty")
	}

	Logger = logrus.New()
	Logger.SetFormatter(cfg.Formatter)
	Logger.SetLevel(cfg.Level)

	if !env.IsDevelopment() {
		_ = os.MkdirAll("logs", os.ModePerm)
		logFile := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02_15-04"))

		Logger.SetOutput(&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    1, // megabytes
			MaxBackups: 7,
			MaxAge:     30, // days
			Compress:   true,
		})
	}

	return nil
}

// Rotasi log berdasarkan menit saat ada request
func RotateLogIfNeeded() {
	now := time.Now().Format("2006-01-02_15-04")

	mu.Lock()
	defer mu.Unlock()

	if currentTime == now {
		return
	}

	logFile := fmt.Sprintf("logs/%s.log", now)
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Gagal membuka file log: %v\n", err)
		return
	}

	if currentFile != nil {
		currentFile.Close()
	}

	currentFile = file
	currentTime = now

	Logger.SetOutput(file)
}

func getServiceFields() logrus.Fields {
	return logrus.Fields{
		env.EnvironmentName: env.ServiceEnv(),
		env.GoVersionName:   env.GetVersion(),
	}
}

// Helper log
func Debug(args ...interface{}) {
	Logger.WithFields(getServiceFields()).Debug(args...)
}

func Info(args ...interface{}) {
	Logger.WithFields(getServiceFields()).Info(args...)
}

func Warn(args ...interface{}) {
	Logger.WithFields(getServiceFields()).Warn(args...)
}

func Error(args ...interface{}) {
	Logger.WithFields(getServiceFields()).Error(args...)
}

func Fatal(args ...interface{}) {
	Logger.WithFields(getServiceFields()).Fatal(args...)
}
