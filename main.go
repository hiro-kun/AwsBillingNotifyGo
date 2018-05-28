package main

import (
  "fmt"

  // "github.com/aws/aws-sdk-go"
  "github.com/kelseyhightower/envconfig"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
)

func main() {
  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    fmt.Println(err)
  }

  // fmt.Println(config.AWS_ACCESS_KEY)
}
