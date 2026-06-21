package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	addr := ":" + port
	fmt.Printf("父亲节快乐！服务已启动: http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(addr, mux))
}

var indexHTML = template.Must(template.New("index").Parse(htmlContent))

const htmlContent = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>父亲节快乐</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            background: linear-gradient(145deg, #1a1a2e 0%, #16213e 40%, #0f3460 100%);
            font-family: -apple-system, "Noto Sans SC", "PingFang SC", system-ui, sans-serif;
            padding: 20px;
            position: relative;
            overflow-x: hidden;
        }
        body::before {
            content: "";
            position: fixed;
            top: 0; left: 0; right: 0; bottom: 0;
            background-image:
                radial-gradient(1.5px 1.5px at 10% 20%, rgba(255,255,255,0.5), transparent),
                radial-gradient(1px 1px at 30% 65%, rgba(255,255,255,0.4), transparent),
                radial-gradient(2px 2px at 50% 10%, rgba(255,255,255,0.6), transparent),
                radial-gradient(1px 1px at 70% 40%, rgba(255,255,255,0.3), transparent),
                radial-gradient(1.5px 1.5px at 90% 80%, rgba(255,255,255,0.5), transparent),
                radial-gradient(2px 2px at 20% 90%, rgba(255,255,255,0.4), transparent),
                radial-gradient(1px 1px at 80% 15%, rgba(255,255,255,0.3), transparent),
                radial-gradient(1.5px 1.5px at 45% 50%, rgba(255,255,255,0.2), transparent);
            pointer-events: none;
            z-index: 0;
        }
        .card {
            position: relative;
            z-index: 1;
            max-width: 640px;
            width: 100%;
            background: rgba(255,255,255,0.06);
            backdrop-filter: blur(16px);
            -webkit-backdrop-filter: blur(16px);
            border: 1px solid rgba(255,255,255,0.12);
            border-radius: 24px;
            padding: 48px 36px;
            text-align: center;
            box-shadow: 0 24px 64px rgba(0,0,0,0.5);
        }
        .heart { font-size: 56px; margin-bottom: 12px; display: block; }
        h1 {
            font-size: 42px;
            font-weight: 700;
            color: #fff;
            margin-bottom: 16px;
            letter-spacing: 4px;
        }
        .subtitle {
            font-size: 20px;
            color: rgba(255,255,255,0.75);
            line-height: 1.7;
            margin-bottom: 32px;
        }
        .footer {
            font-size: 14px;
            color: rgba(255,255,255,0.35);
        }
    </style>
</head>
<body>
    <div class="card">
        <span class="heart">&#10084;</span>
        <h1>父亲节快乐</h1>
        <p class="subtitle">您是我心中永远的超级英雄<br>感谢您一路的陪伴与守护</p>
        <p class="subtitle" style="font-size:17px; margin-bottom:0;">爸爸，我爱您 &#10084;</p>
        <p class="footer">Happy Father&prime;s Day</p>
    </div>
</body>
</html>`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = indexHTML.Execute(w, nil)
}