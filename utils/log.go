/*
 * @Author: rui.li
 * @Date: 2024-04-24 17:40:30
 * @LastEditors: rui.li
 * @LastEditTime: 2024-04-24 17:43:34
 * @FilePath: /GoOpenAI/utils/log.go
 */
package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.SugaredLogger
)

func _getLogger() *zap.SugaredLogger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapConfig.Encoding = "json"
	zapConfig.DisableStacktrace = true

	l, err := zapConfig.Build()

	if err != nil {
		panic(err)
	}

	return l.Sugar()
}

func init() {
	Logger = _getLogger()
}
