package main

import (
	"net/http"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	for {
		simpleHttpGet("www.baidu.com")
		simpleHttpGet("http://www.baidu.com")
	}
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 打印函数行数
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 格式化时间 可自定义
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 保存console 文件日志切割
func getLogWriter() zapcore.WriteSyncer {

	std := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // 日志名称
		MaxSize:    1,            // 文件内容大小, MB
		MaxBackups: 5,            // 保留旧文件最大个数
		MaxAge:     30,           // 保留旧文件最大天数
		Compress:   false,        // 文件是否压缩
		LocalTime:  true,
	}
	file := zapcore.AddSync(lumberJackLogger)
	return zapcore.NewMultiWriteSyncer(std, file)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error 这是一个错误日志 fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
