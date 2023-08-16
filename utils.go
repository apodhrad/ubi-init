package main

import (
	"io/fs"
	"os"
	"strings"
)

func exists(filename string) (fs.FileInfo, bool) {
	info, err := os.Stat(filename)
	return info, !os.IsNotExist(err)
}

func templatePath(root, path string, d fs.DirEntry) string {
	templatePath, _ := strings.CutPrefix(path, root)
	if templatePath == "" {
		return ""
	}
	if strings.HasPrefix(templatePath, "/") {
		templatePath = templatePath[1:]
	}
	if d.IsDir() && !strings.HasSuffix(templatePath, "/") {
		templatePath += "/"
	}
	return templatePath
}
