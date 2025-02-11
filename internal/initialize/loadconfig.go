package initialize

import (
	"fmt"

	"github.com/anonydev/e-commerce-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to look for the config file in
	viper.SetConfigName("local")      // file name
	viper.SetConfigType("yaml")       // file type

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err) // Terminate the program
	}

	// configuration structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}
}
