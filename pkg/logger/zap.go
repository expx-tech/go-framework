package logger

import (
	"go.uber.org/zap"
)

type ZapLog struct {
	client *zap.Logger
}

func NewZapLog(level string) (*ZapLog, error) {
	logLevel := zap.ErrorLevel
	var err error
	if err = logLevel.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}
	var cl *zap.Logger
	if logLevel == zap.DebugLevel {
		cl, err = zap.NewDevelopment()
	} else {
		cl, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}
	defer cl.Sync()

	return &ZapLog{
		client: cl,
	}, nil
}

func (zl *ZapLog) Info(msg string) {
	zl.client.Info(
		msg,
	)
}

func (zl *ZapLog) Debug(msg string) {
	zl.client.Debug(
		msg,
	)
}

func (zl *ZapLog) DebugHTTP(url, method string) {
	zl.client.Debug(
		"incoming request",
		zap.String("url", url),
		zap.String("method", method),
	)
}

func (zl *ZapLog) Error(err error) {
	zl.client.Error(
		err.Error(),
	)
}

func (zl *ZapLog) Warn(msg string) {
	zl.client.Warn(
		msg,
	)
}
