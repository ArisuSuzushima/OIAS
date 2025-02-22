package config

var logConfig = logger{
	Debug: true,
	File:  false,
}

var dbConfig = db{
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "demodb",
	Database: "oias",
}
