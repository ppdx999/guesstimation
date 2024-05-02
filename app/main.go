package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.ListenAndServe(":8080", newRouter())
}

func newRouter() *http.ServeMux {
  mux := http.NewServeMux()

  mux.HandleFunc("/", hander)
  mux.HandleFunc("GET /wildcard/{v}", WildcardHandler)

  return mux
}

func hander(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World")
}

func WildcardHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, r.PathValue("v"))
}
