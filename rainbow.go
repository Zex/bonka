package main

import (
  "flag"
  "github.com/zex/bonka/rainbow"
//  "bytes"
)

var (
  start_point = flag.String("start", "157820", "start point of the chain")
)

func main() {
  flag.Parse()
  app := rainbow.NewRainbow()
  app.Start(*start_point)
  defer app.Stop()
}
