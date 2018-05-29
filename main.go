package main

import (
  "fmt"

  "log"
  "net/http"
  "net/url"
  "strings"

  "github.com/kelseyhightower/envconfig"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
  "github.com/hiro-kun/AwsBillingNotifyGo/aws"
)

func main() {
  var config conf.Config
  err := envconfig.Process("", &config)
  if err != nil {
    fmt.Println(err)
  }

	billing := aws.GetBilling()
	msg := fmt.Sprintf("%v %v\n", conf.DimensionValue, billing)

  accessToken := config.LINE_NOTIFY_API_TOKEN

  URL := "https://notify-api.line.me/api/notify"

  u, err := url.ParseRequestURI(URL)
  if err != nil {
      log.Fatal(err)
  }

  c := &http.Client{}

  form := url.Values{}
  form.Add("message", msg)

  body := strings.NewReader(form.Encode())

  req, err := http.NewRequest("POST", u.String(), body)
  if err != nil {
      log.Fatal(err)
  }

  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Set("Authorization", "Bearer "+accessToken)

  _, err = c.Do(req)
  if err != nil {
      log.Fatal(err)
  }
}
