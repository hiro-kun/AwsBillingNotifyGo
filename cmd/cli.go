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

// float64, int, error
func run() (int, error) {
  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    return conf.ExitCodeError, err
  }

	billing, _, err := aws.GetBilling()
  if err != nil {
    return conf.ExitCodeError, err
  }

	msg := fmt.Sprintf("%v %v\n", conf.DimensionValue, billing)

  line.MessageApiCall(&line.LineApi{
    Msg:    msg,
    Config: &config,
  })

  return conf.ExitCodeOk, nil
}
