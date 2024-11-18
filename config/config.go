package config

import "time"

type Configurations struct {
	Postgres  PostgresConfig `mapstructure:"postgres"`
	AppConfig AppConfig      `mapstructure:"app"`
	SMTP      SMTPConfig     `mapstructure:"smtp"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}

type AppConfig struct {
	JWTSecret            string        `mapstructure:"jwtSecret"`
	JWTDuration          time.Duration `mapstructure:"jwtDuration"`
	RefreshTokenDuration time.Duration `mapstructure:"refreshTokenDuration"`
}

type SMTPConfig struct {
	SMTPHost     string `mapstructure:"host"`
	SMTPPort     string `mapstructure:"port"`
	SMTPUsername string `mapstructure:"username"`
	SMTPPassword string `mapstructure:"password"`
	SenderEmail  string `mapstructure:"from"`
}
