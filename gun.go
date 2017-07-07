package aliMail

import (
	"net/smtp"
	"strings"
)

type GunConfig struct {
	User     string
	Password string
	Host     string
}

type Gun struct {
	GunConfig
	Auth smtp.Auth
}

func NewGun(c *GunConfig) *Gun {
	g := &Gun{}
	g.GunConfig = *c
	hp := strings.Split(c.Host, ":")
	g.Auth = smtp.PlainAuth("", c.User, c.Password, hp[0])
	return g
}
