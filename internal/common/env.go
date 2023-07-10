package common

import (

	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	// check if prod
	// prod := os.Getenv("PROD")
	prod:= viper.GetString("PROD")

	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
