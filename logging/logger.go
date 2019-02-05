/**
 * Copyright 2017 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ------------------------------------------------------------------------------
 */

package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	CRITICAL = 50
	ERROR    = 40
	WARN     = 30
	INFO     = 20
	DEBUG    = 10
)

// Set the calldepth so we get the right file when logging
const (
	CALLDEPTH = 3
)

type Logger struct {
	logger *zap.Logger
	sugar *zap.SugaredLogger
	level  int
}

var _LOGGER *Logger= nil
var list []interface{}

func Get() *Logger {
	if _LOGGER == nil {
		cfg := zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message",

				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey:    "timestamp",
				EncodeTime: zapcore.ISO8601TimeEncoder,

				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}
		var logger, _ = cfg.Build()
		var sugared = logger.Sugar()
		_LOGGER = &Logger{
			logger: logger,
			sugar: sugared,
			level: DEBUG,
		}
	}
	return _LOGGER
}

func (self *Logger) Output(format string, v ...interface{}) {
	if self.level <= DEBUG {
		self.logf("DEBUG", format, v...)
	}
}

func (self *Logger) SetLevel(level int) {
	self.level = level
}

func (self *Logger) Debugf(format string, v ...interface{}) {
	if self.level <= DEBUG {
		self.logf("DEBUG", format, v...)
	}
}

func (self *Logger) Debug(v ...interface{}) {
	if self.level <= DEBUG {
		self.log("DEBUG", v...)
	}
}

func (self *Logger) Infof(format string, v ...interface{}) {
	if self.level <= INFO {
		self.logf("INFO", format, v...)
	}
}

func (self *Logger) Info(v ...interface{}) {
	if self.level <= INFO {
		self.log("INFO", v...)
	}
}

func (self *Logger) Warnf(format string, v ...interface{}) {
	if self.level <= WARN {
		self.logf("WARN", format, v...)
	}
}

func (self *Logger) Warn(v ...interface{}) {
	if self.level <= WARN {
		self.log("WARN", v...)
	}
}

func (self *Logger) Errorf(format string, v ...interface{}) {
	if self.level <= ERROR {
		self.logf("ERROR", format, v...)
	}
}

func (self *Logger) Error(v ...interface{}) {
	if self.level <= ERROR {
		self.log("ERROR", v...)
	}
}

func (self *Logger) Criticalf(format string, v ...interface{}) {
	if self.level <= CRITICAL {
		self.logf("CRITICAL", format, v...)
	}
}

func (self *Logger) Critical(v ...interface{}) {
	if self.level <= CRITICAL {
		self.log("CRITICAL", v...)
	}
}

func (self *Logger) logf(prefix string, format string, v ...interface{}) {
	msg := "["+prefix+"] "+ fmt.Sprintf(format, v...)

	switch prefix {
	case "DEBUG":
		self.logger.Debug(format)
	case "INFO":
		self.logger.Info(msg)
	case "ERROR":
		self.logger.Error(msg)
	case "CRITICAL":
		self.logger.Info(msg)
	case "WARN":
		self.logger.Warn(msg)
	default:
		self.logger.Info(msg)

	}
}

func (self *Logger) SetDefaultKeyValues(keysAndValues ...interface{}) {
	list = keysAndValues
}

func (self *Logger) log(prefix string, v ...interface{}) {
	msg := "["+prefix+"] "+fmt.Sprint(v...)

	switch prefix {
	case "DEBUG":
		self.sugar.Debug(msg)
	case "INFO":
		self.sugar.Infow(msg, list...)
	case "ERROR":
		self.sugar.Error(msg)
	case "CRITICAL":
		self.sugar.Info(msg)
	case "WARN":
		self.sugar.Warn(msg)
	default:
		self.sugar.Info(msg)

	}
}
