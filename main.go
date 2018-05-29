package main

import (
  "fmt"

  "github.com/kelseyhightower/envconfig"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
  "github.com/hiro-kun/AwsBillingNotifyGo/aws"
)

const (
	dimensionValue = "USD"
)

func main() {
  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    fmt.Println(err)
  }

	billing := aws.GetBilling()
	fmt.Printf("%v %v\n", dimensionValue, billing)
}
