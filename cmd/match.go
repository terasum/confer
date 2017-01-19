package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/terasum/confer/conf"
	"strings"
	"encoding/json"
	"reflect"
	"strconv"
)
func init() {
	RootCmd.AddCommand(matchCmd)
}

var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "match the config key's value",
	Long: `match the config key's value by key such as key1.key2`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<2{
			fmt.Println("the param error, please specific the config path and the key")
			fmt.Println("Example:")
			fmt.Println("confer match config.yaml global.peerconfigs mapkey 1")
			return
		}
		if !Exist(Abs(args[0])){
			fmt.Println("the config file is not exist",Abs(args[0]))
			return
		}

		conf,err := conf.NewConfer(Abs(args[0]))
		if err != nil{
			fmt.Println("read config file err")
			fmt.Println(err)
			return
		}
		slices := conf.Get(args[1])
		mapkey := args[2]
		value  := args[3]
		switch slices.(type){
		case []interface {}:{
			matchFlag := false
			slice := slices.([](interface{}))
			for i:=0;i<len(slice);i++{
				innermap := slice[i].(map[string]interface{})
				toMatchValue,ok := innermap[mapkey]
				if !ok {
					fmt.Printf("cannot find the key's item, key: %s \n ",mapkey)
					return
				}
				toMatchValueStr := ""
				switch reflect.TypeOf(toMatchValue).Name(){
				case "string":{
					toMatchValueStr = toMatchValue.(string)
				}
				case "int":{
					toMatchValueStr = strconv.Itoa(toMatchValue.(int))
				}
				case "bool":{
					if toMatchValue.(bool) {
						toMatchValueStr = "true"
					}else{
						toMatchValueStr = "false"
					}
				}
				case "float64":{
					toMatchValueStr = strconv.Itoa((int)(toMatchValue.(float64)))
				}
				default:
					fmt.Printf("cannot match thoes type %v\n",toMatchValue)
					return
				}


				if strings.EqualFold(toMatchValueStr,value) {
					matchFlag = true
					innermapByte,err := json.Marshal(innermap)
					if err != nil{
						fmt.Printf("json marshal failed, %v",err)
						return
					}
					fmt.Println(string(innermapByte))
				}
			}
			if !matchFlag{
				fmt.Println("Match Noting!")
			}
			return
			}
		default:{
			fmt.Println("this patten not match a array, please use the `confer read` to get the item value!")
		}
		}
	},
}