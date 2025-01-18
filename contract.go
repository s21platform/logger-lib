//go:generate mockgen -destination=mock_contract.go -package=${GOPACKAGE} -source=contract.go

package logger_lib

type LoggerInterface interface {
	AddFuncName(name string)
	Info(msg string)
	Error(msg string)
	Warn(msg string)
}
