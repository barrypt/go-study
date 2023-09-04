package main

import (
	"github.com/natefinch/lumberjack"
	"log/slog"
)

func main() {
	r := &lumberjack.Logger{
		Filename:   "./foo.log",
		LocalTime:  true,
		MaxSize:    1,
		MaxAge:     3,
		MaxBackups: 5,
		Compress:   true,
	}
	logs := slog.New(slog.NewJSONHandler(r, nil))
	slog.SetDefault(logs)
	slog.Info("`1223333")
	slog.Info("`1223333", "1", 2, "3", 4, "5", 6)

}
