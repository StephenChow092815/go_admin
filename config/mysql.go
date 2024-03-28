package config

type MysqlConfig struct {
	Database   string `json:"database"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	UseDefault bool   `json:"use_default"` // 使用 default 的账号密码
}

