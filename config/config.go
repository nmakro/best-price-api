package config

type DBConfig struct {
	DbDriver string
	DbName   string
	DbUser   string
	Password string
	Options  string
	URI      string
}

type AuthConfig struct {
	Username string
	Password string
}
