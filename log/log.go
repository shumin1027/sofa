package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

//var RedisLogger *zap.Logger
//var LmdbLogger *zap.Logger
//var HttpLogger *zap.Logger

func InitLogger(role string) {
	switch role {
	case "server":
		Logger = NewLogger("./logs/sofa-server.log", zapcore.InfoLevel, 128, 30, 7, true, true, "sofa-server")

	case "agent":
		Logger = NewLogger("./logs/sofa.log", zapcore.InfoLevel, 128, 30, 7, true, false, "sofa")
	}
	//MainLogger = NewLogger("./logs/main.log", zapcore.InfoLevel, 128, 30, 7, true, "Main")
	//GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
	//RedisLogger = NewLogger("./logs/redis.log", zapcore.InfoLevel, 128, 30, 7, true, "redis")
	//LmdbLogger = NewLogger("./logs/lmdb.log", zapcore.InfoLevel, 128, 30, 7, true, "lmdb")
	//HttpLogger = NewLogger("./logs/http.log", zapcore.InfoLevel, 128, 30, 7, true, "http")
	Logger.Info("zap logger init success")
}
