package main

import (
  "fmt"
  "html/template"
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
  t := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
  err := t.Execute(w, "Hello")
  if err != nil {
    fmt.Fprint(w, err)
  }
}

func WildcardHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, r.PathValue("v"))
}
