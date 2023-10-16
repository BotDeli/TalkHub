package config

import "fmt"

type Config struct {
	Http     *HttpConfig     `yaml:"http" env-required:"true"`
	Grpc     *GRPCConfig     `yaml:"grpc" env-required:"true"`
	Postgres *PostgresConfig `yaml:"postgres" env-required:"true"`
}

type HttpConfig struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-required:"true"`
}

func (c *HttpConfig) GetAddress() string {
	return c.Host + ":" + c.Port
}

type GRPCConfig struct {
	Address string `yaml:"address" env-required:"true"`
}

type PostgresConfig struct {
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password"`
	Address  string `yaml:"address" env-required:"true"`
	Dbname   string `yaml:"dbname" env-required:"true"`
	Sslmode  string `yaml:"sslmode" env-default:"false"`
}

func (cfg *PostgresConfig) GetSourceName() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Address,
		cfg.Dbname,
		cfg.Sslmode,
	)
}
