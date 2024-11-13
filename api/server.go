package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getPort() string {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	return port
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		u, err := url.Parse(req.URL.Path)
		if err == nil {
			if strings.HasPrefix(u.String(), "key") {
				w.Write([]byte("Unauthorized access."))
				return
			}
			if u.String() == "/" {
				html, read_err := os.ReadFile("index.html")
				if read_err != nil {
					w.Write([]byte("File failed to open."))
				} else {
					w.Write(html)
				}
				return
			}
			html, read_err := os.ReadFile("." + u.String() + "index.html")
			if read_err != nil {
				w.Write([]byte(read_err.Error()))
			} else {
				w.Write(html)
			}
		} else {
			w.Write([]byte([]byte(err.Error())))
		}
	})

	log.Panic(http.ListenAndServe(":"+getPort(), nil))
}
