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
  Token   string
}

func NewLineApi(msg string, config *conf.Config, token string) *LineApi {
    return &LineApi{
        Msg:    msg,
        Config: config,
        Token:  token,
    }
}

func (l *LineApi) MessageApiCall() (error) {

  accessToken := l.Token

  URL := conf.LineEndPointURL

  u, err := url.ParseRequestURI(URL)
  if err != nil {
    return fmt.Errorf("url parse error : %s", err)
  }

  c := &http.Client{}

  form := url.Values{}
  form.Add("message", l.Msg)

  body := strings.NewReader(form.Encode())

  req, err := http.NewRequest("POST", u.String(), body)
  if err != nil {
    return fmt.Errorf("api request error : %s", err)
  }

  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Set("Authorization", "Bearer "+accessToken)

  _, err = c.Do(req)
  if err != nil {
    return fmt.Errorf("api call error : %s", err)
  }

  return nil
}
