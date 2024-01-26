package config

type config struct {
	Env          string `yaml:"env" env-default:"local"`
	Storage_path string `yaml:"storage_path"`
	Grpc
}
