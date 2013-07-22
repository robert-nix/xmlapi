package main

import (
  "github.com/Mischanix/applog"
  "github.com/Mischanix/evconf"
  "github.com/Mischanix/wait"
  "os"
)

var ready = wait.NewFlag(false)

var logStdout = false

var config struct {
  HttpPort      int                 `json:"http_port"`
  RemoteBaseUrl string              `json:"remote_base_url"`
  RemoteSuffix  string              `json:"remote_suffix"`
  Endpoints     map[string][]string `json:"endpoints"`
}

func defaultConfig() {
  config.HttpPort = 9003
}

func start(name string, ready *wait.Flag) {
  // Log setup
  applog.Level = applog.DebugLevel
  if logStdout {
    applog.SetOutput(os.Stdout)
  } else {
    if logFile, err := os.OpenFile(
      name+".log",
      os.O_WRONLY|os.O_CREATE|os.O_APPEND,
      os.ModeAppend|0666,
    ); err != nil {
      applog.SetOutput(os.Stdout)
      applog.Error("Unable to open log file: %v", err)
    } else {
      applog.SetOutput(logFile)
    }
  }
  applog.Info("starting...")

  // Config setup
  conf := evconf.New(name+".json", &config)
  conf.OnLoad(func() {
    ready.Set(true)
  })
  defaultConfig()
  go func() {
    conf.Ready()
  }()
}

func main() {
  start("xmlapi", ready)

  ready.WaitFor(true)

  httpServer()
}
