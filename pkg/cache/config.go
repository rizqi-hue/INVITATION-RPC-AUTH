package cache

// RedisConfig define configuration model for redis
type RedisConfig struct {
	Server   string `json:"server"`
	Timeout  int64  `json:"timeout"`
	AuthPass string `json:"authPass"`
	PoolSize int    `json:"poolSize"`
}
