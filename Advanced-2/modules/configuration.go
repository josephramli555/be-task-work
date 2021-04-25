package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"github.com/spf13/viper"
)

type Configuration struct{}

var config *viper.Viper


func (c Configuration) Init() {
	var err error

	files, err := ioutil.ReadDir("./conf/")
	if err != nil {
		log.Fatal(err)
	}

	var names []string

	for _, f := range files {
		fmt.Println(f.Name())
		names = append(names, strings.Split(f.Name(), ".")[0])
	}

	config = viper.New()


	config.SetConfigType("json")
	config.AddConfigPath("./conf/")
	for i, n := range names {
		config.SetConfigName(n)
		if i == 0 {
			err = config.ReadInConfig()
		} else {
			err = config.MergeInConfig()
		}
		if err != nil {
			log.Fatal(err)
		}
	}


	config.AutomaticEnv()
}

func (c Configuration) GetConfig() *viper.Viper {
	return config
}