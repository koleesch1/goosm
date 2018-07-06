package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("settings")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.goosm")
	if err := viper.ReadInConfig(); err != nil {
		//    if err !=
		//panic(
		fmt.Printf("Fatal error reading config file: %s \n", err)
	}

	parseCommandLineFlags()

	if viper.GetBool("init") { //&& !viper.IsSet("MySensebox") {
		// create config for my opensensebox
		viper.Set("MySensebox", "5a436432ec97390020d9c2ee")
		viper.WriteConfig()
	}

}

func parseCommandLineFlags() {
	//using standard library "flag" package
	flag.Bool("init", false, "initiliazes the app with default values")
	flag.Bool("debug", false, "print debug messages")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// get init Flag
	viper.Set("init", viper.GetBool("init"))
	// debug flag
	viper.Set("debug", viper.GetBool("debug"))
	// debug flag

	if viper.GetBool("debug") {

		fmt.Println("debug mode is set")
	}

}
