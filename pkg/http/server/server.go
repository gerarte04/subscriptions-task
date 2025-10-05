package server

import (
	"net/http"
	"time"
)

type HttpConfig struct {
	Address      string        `yaml:"address" env:"HTTP_ADDRESS" env-required:"true"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-required:"true"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-required:"true"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-required:"true"`
}

func CreateServer(handler http.Handler, cfg HttpConfig) error {
    s := &http.Server{
        Addr: cfg.Address,
        ReadTimeout: cfg.ReadTimeout,
        WriteTimeout: cfg.WriteTimeout,
        IdleTimeout: cfg.IdleTimeout,

        Handler: handler,
    }

    return s.ListenAndServe()
}
