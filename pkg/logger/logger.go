package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() (*Logger, error) {
	logger, err := zap.NewProduction() // или zap.NewDevelopment() для разработки
	if err != nil {
		return nil, fmt.Errorf("ошибка инициализации логгера: %w", err)
	}
	return &Logger{z: logger}, nil
}

type Logger struct {
	z *zap.Logger
}

// Sync is flushing any buffered entries.
func (l *Logger) Sync() error {
	return l.z.Sync()
}
func (l *Logger) Debug(msg string, fields ...any) { l.z.Debug(msg, toZapFields(fields)...) }
func (l *Logger) Info(msg string, fields ...any)  { l.z.Info(msg, toZapFields(fields)...) }
func (l *Logger) Warn(msg string, fields ...any)  { l.z.Debug(msg, toZapFields(fields)...) }
func (l *Logger) Error(msg string, fields ...any) { l.z.Debug(msg, toZapFields(fields)...) }
func (l *Logger) Fatal(msg string, fields ...any) { l.z.Debug(msg, toZapFields(fields)...) }

func toZapFields(fields []interface{}) []zap.Field {
	var zf []zap.Field
	for i := 0; i < len(fields); i += 2 {
		// If odd fields count
		if i == len(fields)-1 {
			zf = appendField(zf, "_", fields[i])
			break
		}
		zf = appendField(zf, fields[i].(string), fields[i+1])
	}
	return zf
}

// Custom fields append.
func appendField(fields []zap.Field, key string, value interface{}) []zap.Field {
	switch v := value.(type) {
	case error:
		fields = append(fields, zap.Field{Key: key, Type: zapcore.StringType, String: v.Error()})
		return fields
	}
	fields = append(fields, zap.Any(key, value))
	return fields
}
