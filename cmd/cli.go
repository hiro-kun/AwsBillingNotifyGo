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

	billingInfo, err := aws.GetBilling()
  if err != nil {
    return conf.ExitCodeError, err
  }
	msg := fmt.Sprintf(" \n 想定金額: %v %v\n 想定金額確定日: %v ", conf.DimensionValue, billingInfo["estimatePrice"], billingInfo["timestamp"])

  lineApi := line.NewLineApi(msg, &config, config.LINE_NOTIFY_API_TOKEN)
  err = lineApi.MessageApiCall()
  if err != nil {
    return conf.ExitCodeError, err
  }

  return conf.ExitCodeOk, nil
}
