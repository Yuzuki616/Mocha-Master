package conf

import (
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"os"
)

type Conf struct {
	path        string
	LogLevel    string `json:"LogLevel"`
	DbPath      string `json:"DbPath"`
	Addr        string `json:"Addr"`
	AccessToken string `json:"AccessToken" validate:"required"`
	PullToken   string `json:"PullToken"   validate:"required"`
}

func New(path string) *Conf {
	c := &Conf{
		path:        path,
		LogLevel:    "info",
		DbPath:      "data.db",
		Addr:        ":8080",
		AccessToken: "",
		PullToken:   "",
	}
	return c
}

func (c *Conf) Load() error {
	v := validator.New()
	f, err := os.Open(c.path)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return err
	}
	err = v.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
