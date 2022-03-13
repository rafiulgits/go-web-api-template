package logger

import (
	"net/http"
	"time"
	"webapi/configs"

	"github.com/go-chi/chi/v5/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func NewLogger(cfg *configs.LogConfig) {
	rotator, err := rotatelogs.New(
		cfg.FilePath,
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAgeInHour)*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(cfg.RotationTimeInHour)),
	)
	if err != nil {
		panic(err)
	}
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Name",
		CallerKey:      "Caller",
		MessageKey:     "Msg",
		StacktraceKey:  "St",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// add the encoder config and rotator to create a new zap logger
	writer := zapcore.AddSync(rotator)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encCfg),
		writer,
		zap.InfoLevel)

	zapLogger := zap.New(core)
	Log = zapLogger
}

func ZapFileLogging(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.Info("Served",
					zap.String("proto", r.Proto),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.Duration("duration", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()))
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
