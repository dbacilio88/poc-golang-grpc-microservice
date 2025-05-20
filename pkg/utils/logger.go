package utils

import (
	"fmt"
	"github.com/dbacilio88/poc-golang-grpc-microservice/pkg/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
*
* logger
* <p>
* logger file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 8/12/2024
*
 */

var Core zapcore.Core
var instance *zap.Logger

func LoggerConfiguration(level string) (*zap.Logger, error) {

	mapLevel := map[string]zapcore.Level{
		"debug":      zapcore.DebugLevel,
		"info":       zapcore.InfoLevel,
		"warn":       zapcore.WarnLevel,
		"error":      zapcore.ErrorLevel,
		"fatal":      zapcore.FatalLevel,
		"panic":      zapcore.PanicLevel,
		"disabled":   zapcore.DPanicLevel,
		"deprecated": zapcore.DPanicLevel,
	}
	levelValue, exists := mapLevel[level]

	if !exists {
		levelValue = zapcore.InfoLevel
	}

	log := zap.NewAtomicLevelAt(levelValue)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		CallerKey:     "caller",

		EncodeName:     zapcore.FullNameEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Usar colores
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // Formato ISO8601 para el tiempo
		EncodeDuration: zapcore.SecondsDurationEncoder,   // Duración en segundos
		EncodeCaller:   zapcore.ShortCallerEncoder,       // Información completa de la llamada
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	filename := fmt.Sprintf("/%s/%s-%s.log", env.YAML.Server.Logs, env.YAML.Server.Name, env.YAML.Server.Environment)
	fmt.Println(filename)
	logLumberjack := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   true,
		MaxAge:     30,
	}

	fileCore := zapcore.NewCore(encoder, zapcore.AddSync(logLumberjack), log)
	consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), log)

	Core = zapcore.NewTee(
		fileCore.With(
			[]zap.Field{
				zap.String("app", env.YAML.Server.Name),
				zap.String("env", env.YAML.Server.Environment),
			},
		),

		consoleCore.With(
			[]zap.Field{
				zap.String("app", env.YAML.Server.Name),
				zap.String("env", env.YAML.Server.Environment),
			}),
	)
	instance = zap.New(Core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.Development())
	return instance, nil
}
