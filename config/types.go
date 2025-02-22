package config

type config struct {
	Log logger `json:"log"`
	Db  db     `json:"db"`
}

type logger struct {
	Debug bool `json:"debug"`
	File  bool `json:"file"`
}

type db struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
