package line

import(
  "fmt"

  "net/http"
  "net/url"
  "strings"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
)

func ApiCall(msg string, config *conf.Config) {

  accessToken := config.LINE_NOTIFY_API_TOKEN

  URL := conf.LineEndPointURL

  u, err := url.ParseRequestURI(URL)
  if err != nil {
    fmt.Println(err)
  }

  c := &http.Client{}

  form := url.Values{}
  form.Add("message", msg)

  body := strings.NewReader(form.Encode())

  req, err := http.NewRequest("POST", u.String(), body)
  if err != nil {
    fmt.Println(err)
  }

  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Set("Authorization", "Bearer "+accessToken)

  _, err = c.Do(req)
  if err != nil {
    fmt.Println(err)
  }
}
