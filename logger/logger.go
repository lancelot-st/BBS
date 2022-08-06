package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	fileWriter := getLogWriter()
	enCoder := getEncoder()
	core := zapcore.NewCore(enCoder, fileWriter, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//使用os打开当前目录下的BBS.log 文件并返回一个file句柄
	//file, _ := os.Create("./logger/BBS.log")
	//传入要打开的文件到AddSync函数中并返回WriteSyncer 句柄
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./logger/logger.log",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(lumberjackLogger)
}

func SimpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		SugarLogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		SugarLogger.Info(
			"Success",
			zap.String("Status", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}

}
