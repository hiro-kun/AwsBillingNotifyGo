package main

// 設置場所
// $GOPATH/src/github.com/hiro-kun/AwsBillingNotifyGo

import (
  "fmt"

  "github.com/kelseyhightower/envconfig"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
  "github.com/hiro-kun/AwsBillingNotifyGo/aws"
  "github.com/hiro-kun/AwsBillingNotifyGo/line"
)

func main() {

  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    fmt.Println(err)
  }

	billing := aws.GetBilling()
	msg := fmt.Sprintf("%v %v\n", conf.DimensionValue, billing)

  line.ApiCall(&line.LineApi{
    Msg:    msg,
    Config: &config,
  })
}
