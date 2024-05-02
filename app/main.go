package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", hander)
  http.ListenAndServe(":8080", nil)
}

func hander(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}
