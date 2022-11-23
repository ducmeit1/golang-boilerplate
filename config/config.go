package config

type Configurations struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}
