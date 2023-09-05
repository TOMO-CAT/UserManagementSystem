package config

type umsConfig struct {
	Env string `toml:"env"`

	Redis struct {
		IP                 string `toml:"ip"`
		Port               int    `toml:"port"`
		MaxIdle            int    `toml:"max-idle"`
		MaxActive          int    `toml:"max-active"`
		IdleTimeoutSeconds int    `toml:"idle-timeout-seconds"`
	} `toml:"redis"`

	Mysql struct {
		IP           string `toml:"ip"`
		Port         int    `toml:"port"`
		User         string `toml:"user"`
		Password     string `toml:"password"`
		Database     string `toml:"database"`
		MaxOpenConns int    `toml:"max-open-conns"`
		MaxIdleConns int    `toml:"max-idle-conns"`
	} `toml:"mysql"`

	GrpcServer struct {
		Port int `toml:"port"`
	} `toml:"grpc-server"`

	HttpServer struct {
		Port int `toml:"port"`
	} `toml:"http-server"`
}
