package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
)

const name = "App"

type App struct {
	r io.Reader
	l *log.Logger
}

// NewApp returns a new App.
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r: r, l: l}
}

func (a *App) Run(ctx context.Context) {
	newCtx, span := otel.Tracer(name).Start(ctx, runFuncName())
	defer span.End()

	a.Run1(newCtx)
	a.Run2(newCtx)
}

func (a *App) Run1(ctx context.Context) {
	newCtx, span := otel.Tracer(name).Start(ctx, runFuncName())
	defer span.End()

	a.Run1_1(newCtx)
}

func (a *App) Run1_1(ctx context.Context) {
	_, span := otel.Tracer(name).Start(ctx, runFuncName())
	defer span.End()
}

func (a *App) Run2(ctx context.Context) {
	_, span := otel.Tracer(name).Start(ctx, runFuncName())
	defer span.End()
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	name := f.Name()
	fmt.Println("name", name)
	names := strings.Split(name, ".")
	return names[len(names)-1]
}
