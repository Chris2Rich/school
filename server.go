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

func read_bytes(filename string, start, count int)([]byte, error){
  file, err := os.Open(filename)
  if(err != nil){
    return nil, err
  }

  reader := bufio.NewReader(file)
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

  file.Close()

  return bytes.Join(res, nil), nil
}

func write_bytes(filename string, data string)(error){ 
  file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
  if(err != nil){
    return err
  }

  writer := bufio.NewWriter(file)
  _, err = writer.WriteString(data)
  if err != nil {
    return err
  }

  return writer.Flush()
}

func main() {
  err_t := [2]error{}

  static_html_content, tmp := ioutil.ReadFile("index.html")
  err_t[0] = tmp

  /*
  todo_list_data, tmp := read_bytes("todo_list.txt", 0, 0xFFFFFF)
  err_t[1] = tmp
  */

  if(func(arr []error)(bool){for _, v := range arr {if v != nil {return true}}; return false}(err_t[:])){
    log.Panic("Error: Server failed to initialize", err_t)
    return
  }

  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
    w.Header().Set("Content-Type", "text/html")
    w.Write(static_html_content)
    return
  })
 
  /*
  http.HandleFunc("/fetch/todo", func(w http.ResponseWriter, req *http.Request){
    w.Header().Set("Content-Type", "text/plain")
    w.Write(todo_list_data)
    return
  })

  http.HandleFunc("/write/todo", func(w http.ResponseWriter, req *http.Request){
    w.Header().Set("Content-Type", "text/plain")
    err := write_bytes("todo_list.txt", "This is a test\nwords have meaning\nschema34444\nthisis line 4")
    if(err != nil){
      w.Write([]byte("Null"))
      log.Panic("Error: writing to todo_list.txt failed", err)
    } else {
      todo_list_data, err = read_bytes("todo_list.txt", 0, 0xFFFFFF)
      if(err != nil){
        log.Panic("Error: todo_list_data out of sync", err)
      } else {  
        w.Write(todo_list_data)
      }
    }
    return
  })
  */
  go log.Panic(http.ListenAndServeTLS(":443", "crt.crt", "key.key", nil)) 
  return
}
