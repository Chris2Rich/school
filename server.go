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

func read_bytes(reader *bufio.Reader, start, count int)([]byte, error){
  for i := 1; i < start; i++ {
    _, err := reader.ReadBytes('\n')
    if err != nil {
      return nil, errors.New("Error: line outside of file")
    }
  }

  var res [][]byte
  for i:= 0; i < count; i++ {
    buffer, err := reader.ReadBytes('\n')
    if err != nil {
      if err == io.EOF {
        if len(buffer) > 0 {
          res = append(res, buffer)
        }
        break
      }
      return nil, err
    }
    res = append(res, buffer)
  }

  return bytes.Join(res, nil), nil
}

func main() {
  //Initialize error table
  err_t := [3]error{}

  //Loads all content from files when sever initializes
  static_html_content, tmp := ioutil.ReadFile("index.html")
  err_t[0] = tmp
  todo_list_file, tmp := os.Open("todo_list.txt")
  err_t[1] = tmp

  //All files opened are closed when program terminates
  defer todo_list_file.Close()

  tdl_reader := bufio.NewReader(todo_list_file)
  todo_list_data, tmp := read_bytes(tdl_reader, 0, 0xFFFFFF)
  err_t[2] = tmp
  
  //If any errors in loading critical data, program crashes
  if(func(arr []error)(bool){for _, v := range arr {if v != nil {return true}}; return false}(err_t[:])){
    log.Panic("Error: Server failed to initialize", err_t)
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
