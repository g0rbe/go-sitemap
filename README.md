# go-sitemap

[![Go Report Card](https://goreportcard.com/badge/github.com/g0rbe/go-sitemap)](https://goreportcard.com/report/github.com/g0rbe/go-sitemap)
[![Go Reference](https://pkg.go.dev/badge/github.com/g0rbe/go-sitemap.svg)](https://pkg.go.dev/github.com/g0rbe/go-sitemap)

Golang module to work with Sitemaps 

Get:
```bash
go get github.com/g0rbe/go-sitemap@latest
```

Get the latest tag (if Go module proxy is not updated):
```bash
go get "github.com/g0rbe/go-sitemap@$(curl -s 'https://api.github.com/repos/g0rbe/go-sitemap/tags' | jq -r '.[0].name')"
```

Get the latest commit (if Go module proxy is not updated):
```bash
go get "github.com/g0rbe/go-sitemap@$(curl -s 'https://api.github.com/repos/g0rbe/go-sitemap/commits' | jq -r '.[0].sha')"
```