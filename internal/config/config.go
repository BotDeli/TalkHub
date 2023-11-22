package config

import "fmt"

type Config struct {
	Http     *HttpConfig     `yaml:"http" env-required:"true"`
	Grpc     *GRPCConfig     `yaml:"grpc" env-required:"true"`
	Postgres *PostgresConfig `yaml:"postgres" env-required:"true"`
	Meeting  *MeetingConfig  `yaml:"meeting" env-required:"true"`
	Webrtc   *WebrtcConfig   `yaml:"webrtc" event-required:"true"`
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

type MeetingConfig struct {
	MaxCountConnections int `yaml:"max_count_connections" env-required:"true"`
}

type WebrtcConfig struct {
	StunUrl    string `yaml:"stun" json:"stun" env-required:"true"`
	TurnUrl    string `yaml:"turn" json:"turn" env-required:"true"`
	Username   string `yaml:"username" json:"username" env-required:"true"`
	Credential string `yaml:"credential" json:"credential" env-required:"true"`
	AudioCodec string `yaml:"audio_codec" json:"audio" env-required:"true"`
	VideoCodec string `yaml:"video_codec" json:"video" env-required:"true"`
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
