package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to look for the config file in
	viper.SetConfigName("local")      // file name
	viper.SetConfigType("yaml")       // file type

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err) // Terminate the program
	}

	fmt.Println("port:", viper.GetInt("server.port"))
	fmt.Println("jwt key:", viper.GetInt("security.jwt.key"))

	// configuration structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}

	fmt.Println("port:", config.Server.Port)

	// Syntax _, db -> ignore the index
	for _, db := range config.Databases {
		fmt.Println("user:", db.User)
		fmt.Println("password:", db.Password)
		fmt.Println("host:", db.Host)
	}
}
