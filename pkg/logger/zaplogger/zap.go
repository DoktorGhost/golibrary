package zaplogger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger - обертка над Zap для реализации интерфейса Logger
type ZapLogger struct {
	zap *zap.Logger
}

// Создание нового экземпляра логгера
func NewZapLogger() (*ZapLogger, error) {
	config := zap.NewProductionConfig() // Или zap.NewDevelopmentConfig()
	config.Encoding = "console"         // Используем текстовый формат
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "msg",   // Ключ для сообщения
		LevelKey:   "level", // Ключ для уровня логирования
		TimeKey:    "time",  // Ключ для времени
		CallerKey:  "",      // Отключаем отображение caller
		EncodeTime: zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("02/01/2006 15:04")) // Устанавливаем формат времени
		}),
		EncodeLevel: zapcore.CapitalLevelEncoder,
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{zap: logger}, nil
}

// Метод для записи информационного сообщения
func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		l.zap.Sugar().Infow(msg, fields...)
	} else {
		l.zap.Sugar().Info(msg)
	}
}

// Метод для записи сообщения об ошибке
func (l *ZapLogger) Error(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		l.zap.Sugar().Errorw(msg, fields...)
	} else {
		l.zap.Sugar().Error(msg)
	}
}

// Метод для записи фатального сообщения
func (l *ZapLogger) Fatal(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		l.zap.Sugar().Fatalw(msg, fields...)
	} else {
		l.zap.Sugar().Fatal(msg)
	}
}

func (l *ZapLogger) Sync() {
	_ = l.zap.Sync()
}