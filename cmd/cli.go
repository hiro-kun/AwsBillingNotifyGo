package main

// 設置場所
// $GOPATH/src/github.com/hiro-kun/AwsBillingNotifyGo

import (
  "fmt"
  "os"

  "github.com/kelseyhightower/envconfig"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
  "github.com/hiro-kun/AwsBillingNotifyGo/aws"
  "github.com/hiro-kun/AwsBillingNotifyGo/line"
)

func main() {

  exitCode, err := run()
  if err != nil {
    fmt.Println(err)
  }
  os.Exit(exitCode)
}

func run() (int, error) {
  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    return conf.ExitCodeError, err
  }

	billingInfo, _, err := aws.GetBilling()
  if err != nil {
    return conf.ExitCodeError, err
  }
	msg := fmt.Sprintf(" 想定金額: %v %v\n 請求金額作成日 %v ", conf.DimensionValue, billingInfo["estimatePrice"], billingInfo["timestamp"])

  line.MessageApiCall(&line.LineApi{
    Msg:    msg,
    Config: &config,
  })

  return conf.ExitCodeOk, nil
}
