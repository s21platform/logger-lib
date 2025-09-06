//go:generate mockgen -destination=mock_contract.go -package=${GOPACKAGE} -source=contract.go

package logger_lib

import "context"

type LoggerInterface interface {
	WithField(ctx context.Context, key string, value any) context.Context
	Info(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
}
