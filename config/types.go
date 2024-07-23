package config

type config struct {
	Db db
}

type db struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
