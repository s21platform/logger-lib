package logger_lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type loggerKeyType string

const loggerKey = loggerKeyType("logger")

type Logger struct {
	url    string
	labels map[string]string
	fields map[string]any
}

func New(host, port, service, env string) *Logger {
	url := fmt.Sprintf("http://%s:%s/loki/api/v1/push", host, port)
	logger := &Logger{
		url:    url,
		labels: make(map[string]string),
		fields: make(map[string]any),
	}

	// Set default labels
	logger.withLabel("service", service)
	logger.withLabel("env", env)
	return logger
}

func (l *Logger) info(msg string) {
	l.sendToLoki("info", msg)
}

func (l *Logger) error(msg string) {
	l.sendToLoki("error", msg)
}

func (l *Logger) warn(msg string) {
	l.sendToLoki("warn", msg)
}

func (l *Logger) sendToLoki(level string, message string) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())

	l.withField("message", message)

	messageJson, err := json.Marshal(l.fields)
	if err != nil {
		log.Println("failed to marshall fields")
		return
	}

	entry := Entry{
		Streams: []StreamEntry{
			{
				Stream: Stream{
					Level:       level,
					Environment: l.labels["env"],
					Service:     l.labels["service"],
				},
				Values: [][]string{
					{timestamp, string(messageJson)},
				},
			},
		},
	}

	body, err := json.Marshal(entry)
	if err != nil {
		log.Println("failed to marshall log")
		return
	}

	resp, err := http.Post(l.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("failed to send loki")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		log.Println("failed to send loki. StatusCode:", resp.StatusCode)
	}
}

func (l *Logger) withLabel(key, value string) {
	l.labels[key] = value
}

func (l *Logger) withField(key string, value any) {
	l.fields[key] = value
}

// NewContext returns a new context with the logger.
func NewContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext returns the logger from the context.
// If the logger is not found, it returns nil!! Need to check for nil before using it.
func FromContext(ctx context.Context) *Logger {
	value, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		return nil
	}
	return value
}

// WithField returns a new context with the logger and the field.
func WithField(ctx context.Context, key string, value any) context.Context {
	logger := FromContext(ctx)
	if logger == nil {
		return ctx
	}
	logger.withField(key, value)
	return context.WithValue(ctx, loggerKey, logger)
}

func WithError(ctx context.Context, err error) context.Context {
	return WithField(ctx, "error", err.Error())
}

func WithUserUuid(ctx context.Context, userUuid string) context.Context {
	return WithField(ctx, "user_uuid", userUuid)
}

// Info logs an info message.
func Info(ctx context.Context, msg string) {
	logger := FromContext(ctx)
	if logger == nil {
		return
	}
	logger.info(msg)
}

// Error logs an error message.
func Error(ctx context.Context, msg string) {
	logger := FromContext(ctx)
	if logger == nil {
		return
	}
	logger.error(msg)
}

// Warn logs a warning message.
func Warn(ctx context.Context, msg string) {
	logger := FromContext(ctx)
	if logger == nil {
		return
	}
	logger.warn(msg)
}
