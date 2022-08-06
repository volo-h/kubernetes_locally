package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/google/uuid"
  "github.com/kelseyhightower/envconfig"
)

type Config struct {
  Port int `envconfig:"PORT" default:"8080"`
}

func main() {
  var conf Config
  if err := envconfig.Process("KLAUS", &conf); err != nil {
    log.Fatalf("%v", err)
  }
  log.Println("[INFO] Starting klaus server ...")

  http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    id, err := uuid.NewUUID()
    if err != nil {
      fmt.Fprintf(w, err.Error())
      return
    }
    log.Printf("[INFO] ReqID: %v", id.String())
    fmt.Fprintf(w, "Hello Klaus!!")
  })

  log.Printf("[INFO] Listening on port :%d", conf.Port)
  http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
}
