package main

import "fmt"

import (
    "github.com/splitio/go-client/v6/splitio/client"
    "github.com/splitio/go-client/v6/splitio/conf"
)

func main() {

  fmt.Println("Hello World!")

  cfg := conf.Default()
  factory, err := client.NewSplitFactory("1s46c6r6tfl9m7usijqsmicfckm04mn8b321", cfg)
  if err != nil {
    fmt.Printf("SDK init error: %s\n", err.Error())
    return
  } else {
    fmt.Println("obtained factory...")
  }

  client := factory.Client()
  fmt.Println("obtained client...")

  err2 := client.BlockUntilReady(10)
  if err2 != nil {
    fmt.Printf("SDK timeout: %s\n", err2.Error())
    return
  }
  fmt.Println("finished block until ready...")


  treatment := client.Treatment("user_id", "jsga_demo", nil) 

  if treatment == "on" {
    fmt.Println("on")
  } else if treatment == "off" {
    fmt.Println("THIS IS REALLY off")
  } else {
    fmt.Println(treatment)
  }

}
