package main

import (
  "net/http"
  "io/ioutil"
  "log"
)

func main() {
  content, err_initio := ioutil.ReadFile("index.html")
  if(err_initio != nil){
    log.Fatal("Error: Server failed to initialize - html was not read")
    return
  }

  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Write(content)
    return
  })
 
  http.HandleFunc("/fetch/todo", func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/json")
    w.Write(content)
    return
  })

  log.Fatal(http.ListenAndServe(":80", nil))
  return
}
