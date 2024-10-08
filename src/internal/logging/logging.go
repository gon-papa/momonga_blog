package logging

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

var (
    AppLogger    *slog.Logger
    AccessLogger *slog.Logger
    ErrorLogger  *slog.Logger

    appLogFile    *os.File
    accessLogFile *os.File
    errorLogFile  *os.File
)

func Init() error {
	opt := &slog.HandlerOptions{
		AddSource: true,
	}

    // AppLoggerの初期化
    appLogFile, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return fmt.Errorf("アプリケーションログが開けませんでした: %w", err)
    }
    appHandler := slog.NewJSONHandler(appLogFile, opt)
    AppLogger = slog.New(appHandler)

    // ErrorLoggerの初期化
    errorLogFile, err := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return fmt.Errorf("エラーログが開けませんでした: %w", err)
    }
    errorHandler := slog.NewJSONHandler(errorLogFile, opt)
    ErrorLogger = slog.New(errorHandler)

    // AccessLoggerの初期化
    accessLogFile, err := os.OpenFile("./logs/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return fmt.Errorf("アクセスログが開けませんでした: %w", err)
    }
    accessHandler := slog.NewJSONHandler(accessLogFile, nil)
    AccessLogger = slog.New(accessHandler)

    return nil
}

func LogWithStackTrace() {
    buf := make([]byte, 1024)
    n := runtime.Stack(buf, false)
    ErrorLogger.Error("Stack trace", "stack", string(buf[:n]))
}

func Close() {
    if appLogFile != nil {
        appLogFile.Close()
    }
    if errorLogFile != nil {
        errorLogFile.Close()
    }
    if accessLogFile != nil {
        accessLogFile.Close()
    }
}

