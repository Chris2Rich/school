package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
    w.Header().Set("Content-Type", "text/html")
    u, err := url.Parse(req.URL.Path)
    if err == nil {
      if strings.HasPrefix(u.String(), "key") {
        w.Write([]byte("Unauthorized access."))
        return
      }
      if u.String() == "/" {
        html, read_err := ioutil.ReadFile("index.html")
        if read_err != nil {
          w.Write([]byte("File failed to open."))
        } else {
          w.Write(html)
        }
        return
      }
      html, read_err := ioutil.ReadFile(u.String() + "index.html")
      if read_err != nil {
        w.Write([]byte("File failed to open."))
      } else {
        w.Write(html)
      }
    } else {
      w.Write([]byte("Server could not read url."))
    }
    
    return
  })

  go log.Panic(http.ListenAndServeTLS(":443", "secret/crt.crt", "secret/key.key", nil)) 
  return
}