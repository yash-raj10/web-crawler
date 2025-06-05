# 🕷️ Go Web Crawler

A simple, concurrent web crawler written in Go that extracts all anchor links from a given domain. It recursively crawls only URLs that belong to the same domain, ensuring focused and efficient exploration.


## 🚀 Features

- 🌐 Extracts all anchor (`<a href="">`) links
- 🔄 Recursively crawls internal links
- 🔒 Skips SSL verification to support crawling HTTPS pages without valid certificates
- ⚡ Uses goroutines for concurrency
- 🧠 Avoids duplicate crawls with a visited URL map
- 🧪 Built with `net/http`, `url`, and [steelx/extractlinks](https://github.com/steelx/extractlinks)

---

## 📦 Installation

```bash
git clone https://github.com/your-username/go-web-crawler.git
cd go-web-crawler
go mod tidy
```

---



## 🧪 Usage

```bash
go run main.go https://example.com

