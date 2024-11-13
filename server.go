package main

import (
  "net/http"
  "io"
  "io/ioutil"
  "os"
  "bufio"
  "log"
  "errors"
  "bytes"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
    w.Header().Set("Content-Type", "text/html")
    u, err := url.Parse(req.URL.path)
    if err == nil {
      html, read_err := ioutil.ReadFile("index.html")
      if read_err != nil {
        w.Write("File failed to open.")
      } else
    } else {
      w.Write("Server could not read url.")
    }
    
    return
  })

  go log.Panic(http.ListenAndServeTLS(":443", "crt.crt", "key.key", nil)) 
  return
}