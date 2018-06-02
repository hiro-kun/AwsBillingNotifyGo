package line

import(
  "fmt"

  "net/http"
  "net/url"
  "strings"

  "github.com/hiro-kun/AwsBillingNotifyGo/conf"
)

type LineApi struct {
  Msg     string
  Config  *conf.Config
}

func MessageApiCall(l *LineApi) (int, error) {

  accessToken := l.Config.LINE_NOTIFY_API_TOKEN

  URL := conf.LineEndPointURL

  u, err := url.ParseRequestURI(URL)
  if err != nil {
    return conf.ExitCodeError, fmt.Errorf("url parse error : %s", err)
  }

  c := &http.Client{}

  form := url.Values{}
  form.Add("message", l.Msg)

  body := strings.NewReader(form.Encode())

  req, err := http.NewRequest("POST", u.String(), body)
  if err != nil {
    return conf.ExitCodeError, fmt.Errorf("api request error : %s", err)
  }

  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Set("Authorization", "Bearer "+accessToken)

  _, err = c.Do(req)
  if err != nil {
    return conf.ExitCodeError, fmt.Errorf("api call error : %s", err)
  }

  return conf.ExitCodeOk, nil
}
