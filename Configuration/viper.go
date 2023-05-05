package Configuration

import "github.com/spf13/viper"

type AppConfig struct {
	DBHost       string `mapstructure:"DB_HOST"`
	MODE         string `mapstructure:"MODE"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	Port         string `mapstructure:"PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func NewConfig(path string) (config *AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, err
}
