package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

// InitLogger مقداردهی اولیه logger را انجام می‌دهد و لاگ‌ها را در فایل می‌ریزد
func InitLogger() {
	once.Do(func() {
		// تنظیم فایل لاگ
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./logs/logs.log", // مسیر فایل لاگ
			MaxSize:    10,               // مگابایت
			MaxBackups: 5,
			MaxAge:     7,                // روز
			Compress:   true,             // فشرده‌سازی gzip
		})

		// کانفیگ ساده برای خروجی JSON
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			w,
			zapcore.DebugLevel,
		)

		l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		logger = l.Sugar()
	})
}

// GetLogger همیشه همان logger را برمی‌گرداند
func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		InitLogger()
	}
	return logger
}
