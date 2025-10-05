package config

import (
	"subs-service/pkg/database/postgres"
	"subs-service/pkg/http/server"
)

type ServiceConfig struct {
	DebugMode bool `yaml:"debug_mode" env-default:"true"`
}

type DataConfig struct {
	MaxPrice             int64 `yaml:"max_price" env-default:"100000"`
	MaxServiceNameLength int   `yaml:"max_service_name_length" env-default:"50"`
	DefaultPageSize      int   `yaml:"default_page_size" env-default:"20"`
	MaxPageSize          int   `yaml:"max_page_size" env-default:"100"`
}

type PathConfig struct {
	Api        string `yaml:"api" env-required:"true"`
	PostSub    string `yaml:"post_sub" env-required:"true"`
	GetSub     string `yaml:"get_sub" env-required:"true"`
	PutSub     string `yaml:"put_sub" env-required:"true"`
	DeleteSub  string `yaml:"delete_sub" env-required:"true"`
	ListSubs   string `yaml:"list_subs" env-required:"true"`
	GetSummary string `yaml:"get_summary" env-required:"true"`
}

type Config struct {
	HttpCfg     server.HttpConfig       `yaml:"http"`
	PostgresCfg postgres.PostgresConfig `yaml:"postgres"`
	SvcCfg      ServiceConfig           `yaml:"service"`
	DataCfg     DataConfig              `yaml:"data"`
	PathCfg     PathConfig              `yaml:"paths"`
}
