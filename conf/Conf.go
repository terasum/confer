package conf

import (
	"github.com/terasum/viper"
	"fmt"
)

type Confer struct{
	configPath string
	viper *viper.Viper
}

func NewConfer(configPath string)(*Confer,error){
	var confer Confer
	confer.viper = viper.New()
	confer.configPath = configPath
	confer.viper.SetConfigFile(configPath)
	err := confer.viper.ReadInConfig()
	if err != nil{
		fmt.Println("Cannot read the config file!")
		fmt.Println(err)
		return nil,err
	}
	return &confer,nil

}

func (confer *Confer)GetString(key string)string{
	return confer.viper.GetString(key)
}

func (confer *Confer)GetBool(key string) bool{
	return confer.viper.GetBool(key)
}