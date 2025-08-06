package config

import (
	"github.com/caarlos0/env/v11"
)

func MustLoad[T any]() *T {
	t := env.Must(env.ParseAs[T]())
	return &t
}
