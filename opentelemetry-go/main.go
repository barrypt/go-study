package main

import (
	"context"
	"io"
	"log"
	"os"
	"otlego/app"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

// 创建 Exporter
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// 创建表示当前进程的 Resource 对象
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("gostudy"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "dev"), // 添加 attribute
		),
	)
	return r
}

func main() {
	l := log.New(os.Stdout, "", 0)

	// trace 导出到文件
	f, err := os.Create("tracesimple.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	// 创建一个 Exporter
	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	// 创建 TracerProvider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	app := app.NewApp(os.Stdin, l)
	app.Run(context.Background())
}
