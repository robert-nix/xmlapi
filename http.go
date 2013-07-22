package main

import (
  "fmt"
  "github.com/Mischanix/applog"
  "net/http"
  "net/url"
  "strings"
  "time"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  applog.Info("http: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

  uri, err := url.ParseRequestURI(r.RequestURI)
  if err != nil {
    applog.Info("http: ParseRequestURI failed: %v", err)
    errJson(w, "Bad request", 400)
    return
  }

  last := strings.LastIndex(uri.Path, "/")
  if last < 0 {
    last = len(uri.Path)
  }
  dir := uri.Path[:last]

  fields, ok := config.Endpoints[dir]
  if !ok {
    errJson(w, "Not found", 404)
    return
  }

  remoteUri := config.RemoteBaseUrl + uri.Path + config.RemoteSuffix
  applog.Info("grabbing xml from %s", remoteUri)
  resp, err := http.Get(remoteUri)
  if err != nil {
    applog.Info("http: Get failed: %v", err)
    errJson(w, "Bad gateway", 502)
    return
  }

  writeJson(w, readXml(resp.Body, fields))

  applog.Info("http: %s %s took %v to process",
    r.Method, r.RequestURI,
    time.Now().Sub(start),
  )
}

func httpServer() {
  dialString := fmt.Sprintf(":%d", config.HttpPort)
  if err := http.ListenAndServe(dialString, &handler{}); err != nil {
    applog.Error("httpServer: ListenAndServe error: %v", err)
  }
}
