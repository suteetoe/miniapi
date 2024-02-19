package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap/zapcore"
)

type LoggerFileWriterSyncer struct {
	logDir string
}

func NewLoggerFileWriterSyncer(logDir string) *LoggerFileWriterSyncer {
	logFileWriter := &LoggerFileWriterSyncer{
		logDir: logDir,
	}
	logFileWriter.createLogDirectory()
	return logFileWriter
}

func (f *LoggerFileWriterSyncer) createLogDirectory() error {
	// Define a custom log directory
	//logDir := "logs"
	// Create the log directory if it doesn't exist
	err := os.MkdirAll(f.logDir, os.ModePerm)
	if err != nil {
		//fmt.Printf("Failed to create log directory: %v\n", err)
		return err
	}
	return nil
}

func (f *LoggerFileWriterSyncer) GetFileWriter(t time.Time) zapcore.WriteSyncer {

	// Define the log file path with date
	logPath := fmt.Sprintf("%s/app-%s.log", f.logDir, t.Format("2006-01-02"))

	// Create a file output writer
	fileOutput, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return os.Stderr
	}

	return zapcore.AddSync(fileOutput)

}
