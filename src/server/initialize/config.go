package initialize

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"fmt"
)

func SetupConfig() {
	var (
		conf = flag.String("conf", "", "config file path")
	)

	flag.Parse()

	fmt.Println(*conf)
	viper.SetConfigName(*conf)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config fail:", err.Error())
	}
}
