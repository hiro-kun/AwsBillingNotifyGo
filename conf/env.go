package conf

type Config struct {
    AWS_ACCESS_KEY        string `required:"true"`
    AWS_SECRET_ACCESS_KEY string `required:"true"`
}
