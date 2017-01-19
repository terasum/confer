package cmd

import (
"github.com/spf13/cobra"
"fmt"
"github.com/terasum/confer/conf"
//"strings"
	"encoding/json"
)
//var vartype string
func init() {
	RootCmd.AddCommand(readCmd)
	//genprivCmd.Flags().StringVarP(&vartype, "vartype", "t", "string", "specific the value type: string, int or bool")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read the config key's value",
	Long: `read the config key's value by key such as key1.key2`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<2{
			fmt.Println("the param error, please specific the config path and the key")
			fmt.Println("Example:")
			fmt.Println("confer read config.yaml global.peerconfig")
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
		result := conf.Get(args[1])
		//result = strings.TrimSpace(result)
		//if result == ""{
		//	fmt.Println("null")
		//	return
		//}
		b,err := json.Marshal(result)
		if err != nil{
			fmt.Println("read err")
			return
		}

		fmt.Printf("%v\n",string(b))
	},
}