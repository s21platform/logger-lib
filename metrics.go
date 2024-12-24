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

type Logger struct {
	url      string
	service  string
	env      string
	funcName string
}

func New(host, port, service, env string) *Logger {
	url := fmt.Sprintf("http://%s:%s/loki/api/v1/push", host, port)
	return &Logger{
		url:     url,
		service: service,
		env:     env,
	}
}

func (l *Logger) AddFuncName(name string) {
	l.funcName = l.service + "_" + name
}

func (l *Logger) Info(msg string) {
	l.sendToLoki("info", msg)
}

func (l *Logger) Error(msg string) {
	l.sendToLoki("error", msg)
}

func (l *Logger) Warn(msg string) {
	l.sendToLoki("warn", msg)
}

func (l *Logger) sendToLoki(level string, message string) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())

	entry := Entry{
		Streams: []StreamEntry{
			{
				Stream: Stream{
					Service:      l.service,
					Level:        level,
					Environment:  l.env,
					FunctionName: l.funcName,
				},
				Values: [][]string{
					{
						timestamp,
						message,
					},
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

func FromContext(ctx context.Context, key interface{}) *Logger {
	value := ctx.Value(key)
	if value == nil {
		// Обрабатываем ситуацию, когда значение отсутствует в контексте
		return nil
	}

	logger, ok := value.(*Logger)
	if !ok {
		// Обрабатываем ситуацию, когда значение есть, но неверного типа
		return nil
	}

	return logger
}
