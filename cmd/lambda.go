package main

// 設置場所
// $GOPATH/src/github.com/hiro-kun/AwsBillingNotifyGo

import (
  "fmt"

  "github.com/kelseyhightower/envconfig"
  "github.com/aws/aws-lambda-go/lambda"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
  "github.com/hiro-kun/AwsBillingNotifyGo/aws"
  "github.com/hiro-kun/AwsBillingNotifyGo/line"
)

type Event struct {}

type Response struct {
  Message string
}

func run() error {
  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    return err
  }

	billingInfo, err := aws.GetBilling()
  if err != nil {
    return err
  }
	msg := fmt.Sprintf(" \n 想定金額: %v %v\n 想定金額確定日: %v ", conf.DimensionValue, billingInfo["estimatePrice"], billingInfo["timestamp"])

  err = line.MessageApiCall(&line.LineApi{
    Msg:    msg,
    Config: &config,
  })
  if err != nil {
    return err
  }

  return nil
}

func call(event Event) (Response, error) {

  err := run()
  if err != nil {
    return Response{Message: err.Error()}, nil
  }

  return Response{Message: "lambda end success."}, nil
}

func main() {
  lambda.Start(call)
}
