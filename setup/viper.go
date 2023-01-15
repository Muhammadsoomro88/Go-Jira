package setup

import (
	"log"

	"github.com/spf13/viper"
)

func GetViper() *viper.Viper {
	vi := viper.New()
	vi.SetConfigFile("config/local.yaml")
	err := vi.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return vi
}
