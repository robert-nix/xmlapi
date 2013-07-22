package main

import (
  "encoding/json"
  "github.com/Mischanix/applog"
  "net/http"
)

func responseHeaders(w http.ResponseWriter) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Server", "go/xmlapi")
}

func writeJson(w http.ResponseWriter, result interface{}) {
  responseHeaders(w)
  w.WriteHeader(200)
  if err := json.NewEncoder(w).Encode(result); err != nil {
    applog.Error("writeJson: json Encode failed: %v", err)
  }
}

func errJson(w http.ResponseWriter, error string, code int) {
  responseHeaders(w)
  w.WriteHeader(code)
  if err := json.NewEncoder(w).Encode(struct {
    Status string `json:"status"`
    Reason string `json:"reason"`
  }{"nok", error}); err != nil {
    applog.Error("writeJson: json Encode failed: %v", err)
  }
}
