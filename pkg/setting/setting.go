package setting

import (
	"fmt"
	"github.com/gilbertlim/member-service-go/config"
	"github.com/spf13/viper"
	"os"
)

func Setup() {
	profile := initProfile()
	viper.SetConfigName(profile)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config.RuntimeConf)
	if err != nil {
		panic(err)
	}
}

func initProfile() string {
	profile := os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("GO_PROFILE: " + profile)
	return profile
}
