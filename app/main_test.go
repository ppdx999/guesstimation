package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)


func TestRootHandler(t *testing.T) {
  router := newRouter()
  r := httptest.NewRequest("GET", "/", nil)
  w := httptest.NewRecorder()
  router.ServeHTTP(w, r)

  if w.Code != http.StatusOK {
    t.Errorf("Home page must return status code %v, but got %v", http.StatusOK, w.Code)
  }
}


func TestWildcardHandler(t *testing.T) {
  router := newRouter()
  r := httptest.NewRequest("GET", "/wildcard/anything", nil)
  w := httptest.NewRecorder()
  router.ServeHTTP(w, r)
  if w.Code != http.StatusOK {
    t.Errorf("Wildcard page must return status code %v, but got %v", http.StatusOK, w.Code)
  }

  if w.Body.String() != "anything" {
    t.Errorf("Wildcard page must return 'anything', but got %v", w.Body.String())
  }
}
