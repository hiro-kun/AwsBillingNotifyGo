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

func call(event Event) (Response, error) {

  fmt.Println("=== lambda start. ===")

  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    fmt.Println(err)
  }

  billing := aws.GetBilling()
  msg := fmt.Sprintf("%v %v\n", conf.DimensionValue, billing)

  line.MessageApiCall(&line.LineApi{
    Msg:    msg,
    Config: &config,
  })

  return Response{Message: "=== lambda end. ==="}, nil
}

func main() {
  lambda.Start(call)
}