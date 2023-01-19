package config

var DefaultConfig = map[string]interface{}{
	"name":         "INVITATION-RPC-BOILERPLATE",
	"port":         ":3001",
	"postgres_dsn": "host=localhost user=root password=password dbname=invitation_auth port=5432 sslmode=disable TimeZone=Asia/Jakarta",
	"log_level":    "DEBUG",
	"log_format":   "json",
	"secret": "INVITATION-RPC-AUTH",
  	"exp_auth": 168,
	"redis": map[string]interface{}{
		"server":   "localhost:6379",
		"timeout":  10,
		"authPass": "",
		"poolSize": 10,
	},
}
