package config

type DBConfigStruct struct {
	Host     string
	Port     int
	User     string
	DBname   string
	Password string
}

var DBConfig DBConfigStruct
