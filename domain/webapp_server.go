package domain

import (
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

// StaticServe serves the data
func StaticServe(urlPrefix string, fs *packr.Box) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if exists(fs, urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func exists(fs *packr.Box, prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		name := path.Join("/", p)
		if fs.HasDir(name) {
			index := path.Join(name, "index.html")
			if !fs.Has(index) {
				return false
			}
		} else if !fs.Has(name) {
			return false
		}

		return true
	}
	return false
}
