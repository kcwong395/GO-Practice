package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// documentation of viper: https://github.com/spf13/viper
func main() {
	// init config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	// read config and output read error
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Not found")
		} else {
			// Config file was found but another error was produced
		}
	}

	// work with string
	ver := viper.Get("app-version")
	fmt.Println(ver) // v1.2.5

	// work with map
	dbCred := viper.Get("database-credential").(map[string]interface{})
	fmt.Println(dbCred["username"], dbCred["password"]) // admin pa55w0rd

}
