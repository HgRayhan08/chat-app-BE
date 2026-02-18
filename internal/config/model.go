package config

type Config struct {
	Server   Server
	Database Database
	JWT      JWT
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
	TZ       string
}
type JWT struct {
	Key string
	Exp int
}
