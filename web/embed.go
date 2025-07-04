package web

import (
	"context"
	"embed"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

//go:embed dist
var webFS embed.FS

func InitWebRouter(h *server.Hertz, contextPath string) {
	path := contextPath
	if path == "" {
		path = "/"
	}
	processVueRouter(h, path, webFS, "dist")
}

func processVueRouter(h *server.Hertz, contextPath string, vfs embed.FS, dir string) {
	h.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		path := string(c.Path())
		if contextPath != "" && contextPath != "/" {
			path = strings.Replace(path, contextPath, "", -1)
		}
		hlog.Infof("Request path: %s", path)

		if path == "" || strings.HasSuffix(path, "/") {
			homePage := dir + "/index.html"
			data, err := vfs.ReadFile(homePage)
			if err != nil {
				hlog.Errorf("Error reading %s: %v", homePage, err)
				return
			}
			c.Data(200, "text/html; charset=utf-8", data)
		} else {
			f := dir + path
			data, err := vfs.ReadFile(f)
			if err != nil {
				hlog.Errorf("File not found: %s: %+v", f, err)
				return
			}

			ext := filepath.Ext(path)
			if err != nil {
				hlog.Errorf("Error reading %s: %v", f, err)
				return
			}

			contentType := GetContentType(ext)
			if len(contentType) == 0 {
				c.Data(200, "application/octet-stream", data)
			} else {
				c.Data(200, contentType, data)
			}
		}
	})
}

// GetContentType returns the MIME type for the given file extension
func GetContentType(ext string) string {
	switch strings.ToLower(ext) {
	case ".html", ".htm":
		return "text/html; charset=utf-8"
	case ".css":
		return "text/css; charset=utf-8"
	case ".js":
		return "application/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	case ".txt":
		return "text/plain; charset=utf-8"
	case ".pdf":
		return "application/pdf"
	default:
		return ""
	}
}
