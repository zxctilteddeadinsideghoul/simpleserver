package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"development"`
	HTTPServer `yaml:"HTTPServer"`
}

type HTTPServer struct {
	Address     string `yaml:"address" env:"HTTP_ADDRESS" env-default:":8080"`
	Mode        string `yaml:"mode" env:"HTTP_MODE" env-default:"release"`
	Timeout     int    `yaml:"timeout" env:"HTTP_TIMEOUT" env-default:"4"`
	IdleTimeout int    `yaml:"idleTimeout" env:"HTTP_IDLE_TIMEOUT" env-default:"60"`
}

func MustLoad() Config {
	/*configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}
	*/
	var config Config
	cleanenv.ReadConfig("", &config)
	//err := cleanenv.ReadConfig("", &config)

	//if err != nil {
	//	log.Fatal(err)
	//}

	return config
}
