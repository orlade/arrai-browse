package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func dir(w http.ResponseWriter, r *http.Request) {
	files, err := listArrai(".")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(toHtml(files)))
}

// listArrai returns a list of paths to arr.ai files that can be evaluated.
func listArrai(root string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".arrai") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}

func toHtml(files []string) string {
	wd, _ := os.Getwd()
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("<h1>Arr.ai files under %s</h1><ul>", wd))
	for _, path := range files {
		b.WriteString(fmt.Sprintf(`<li><a href="%s">%s</a></li>`, path, path))
	}
	b.WriteString("</ul>")
	return b.String()
}
