package config

import (
	"log"

	"github.com/spf13/viper"
)

// Settings struct    stores configuration variables,
// each with a mapstructure for Viper.
// The values stored here, will be used across the app.
type Settings struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPass     string `mapstructure:"DB_PASS"`
	Env        string `mapstructure:"ENV"`
	JwtExpires string `mapstructure:"JWT_EXPIRES"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
}

// New function    Loads the configuration from the env file
// and unmarshal the data into the Settings struct
func New() *Settings {
	var cfg Settings

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No env file, using environment variables.", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Error trying to Unmarshal configuration", err)
	}

	return &cfg

}
