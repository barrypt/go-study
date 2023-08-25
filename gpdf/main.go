package main

import (
   _ "fmt"
    "html/template"
    "net/http"
	"github.com/jung-kurt/gofpdf"
)

func main() {
    http.HandleFunc("/", PrintPreviewHandler)
	http.HandleFunc("/pdf", PDFPreviewHandler)
    http.ListenAndServe(":8080", nil)
}

func PrintPreviewHandler(w http.ResponseWriter, r *http.Request) {
    // 模拟从数据库或其他数据源获取需要打印的数据
    data := struct {
        Title   string
        Content string
    }{
        Title:   "Print Preview Example",
        Content: "This is the content to be printed...",
    }

    // 使用Go的html/template包生成HTML页面
    tmpl := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>{{.Title}}</title>
        </head>
        <body>
            <h1>{{.Title}}</h1>
            <p>{{.Content}}</p>
            <button onclick="window.print()">Print</button>
        </body>
        </html>
    `

    t, err := template.New("printPreview").Parse(tmpl)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // 将生成的HTML页面写入ResponseWriter
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func PDFPreviewHandler(w http.ResponseWriter, r *http.Request) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Hello, PDF Generation!")

    // 设置响应头，指定内容为PDF
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "inline; filename=preview.pdf")

    // 将PDF内容写入ResponseWriter
    err := pdf.Output(w)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
