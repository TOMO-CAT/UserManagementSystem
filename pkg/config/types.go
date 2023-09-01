package config

type umsConfig struct {
	Env string `toml:"env"`

	MySql struct {
		Port int `toml:"port"`
	} `toml:"mysql"`

	GrpcServer struct {
		Port int `toml:"port"`
	} `toml:"grpc-server"`

	HttpServer struct {
		Port int `toml:"port"`
	} `toml:"http-server"`
}
