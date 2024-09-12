package main

import (
  "net/http"
  "io/ioutil"
  "os"
  "encoding/csv"
  "log"
)

func main() {
  //Initialize error table
  err_t := [3]error{}

  //Loads all content from files when sever initializes
  static_html_content, tmp := ioutil.ReadFile("index.html")
  err_t[0] = tmp
  todo_list_file, tmp = os.Open("todo_list.csv")
  err_t[1] = tmp

  //All files opened are closed when program terminates
  defer todo_list_file.Close()

  tdl_reader := csv.NewReader(todo_list_file)
  tdl_reader.FieldsPerRecord = -1
  todo_list_data, tmp := tdl_reader.ReadAll()
  err_t[2] = tmp
  
  //If any errors in loading critical data, program crashes
  if(func(arr){for _, v := range arr {if v != nil {return true}}; return false(err_t)}{
    log.Panic("Error: Server failed to initialize", err_table)
    return
  }

  //Generic response for http
  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Write(static_html_content)
    return
  })
 
  //Routes 
  http.HandleFunc("/fetch/todo", func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/json")
    w.Write(todo_list_data)
    return
  })

  //http server listens on port 80 and program crashes if any errors
  log.Panic(http.ListenAndServe(":80", nil))
  return
}
