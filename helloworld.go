package main

import "fmt"
import "time"

import (
    "github.com/splitio/go-client/v6/splitio/client"
    "github.com/splitio/go-client/v6/splitio/conf"
    "github.com/splitio/go-client/v6/splitio/impressionListener"
)

type CustomImpressionListener struct {
}

func (i *CustomImpressionListener) LogImpression(data impressionlistener.ILObject) {
  fmt.Println("got an impression...")
}

func main() {
  customImpressionListener := &CustomImpressionListener{}

  fmt.Println("Hello World!")

  cfg := conf.Default()
  cfg.OperationMode = "redis-consumer"
  cfg.Redis.Host = "localhost"
  cfg.Redis.Port = 6379
  cfg.Advanced.ImpressionListener = customImpressionListener

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

  attributes := make(map[string]interface{})
  attributes["region"] = "CA";

  start := time.Now()
  treatment := client.Treatment("user_id", "multivariant_demo", attributes) 
  t := time.Now()
  fmt.Println(t.Sub(start))

  fmt.Println("With 'CA' as region...");
  if treatment == "on" {
    fmt.Println("on")
  } else if treatment == "off" {
    fmt.Println("THIS IS REALLY off")
  } else {
    fmt.Println(treatment)
  } // could also handle "control" specifically

  att := make(map[string]interface{})
  att["region"] = "MA";
  tment := client.Treatment("user_id", "multivariant_demo", att) 
  fmt.Println("With 'MA' as region...");
  fmt.Println(tment)
 
}
