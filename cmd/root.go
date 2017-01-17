package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/terasum/confer/conf"
	"strings"
)

var RootCmd = &cobra.Command{
	Use:   "confer",
	Short: "Confer is a simple config file reader tool",
	Long: `A simple config file reader tool built with love by terasum and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<2{
			fmt.Println("the param error, please specific the config path and the key")
			fmt.Println("Example:")
			fmt.Println("confer config.yaml global.peerconfig")
			fmt.Println()
			fmt.Println("============================")
			fmt.Println()
			cmd.Help()
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
		result := conf.GetString(args[1])
		result = strings.TrimSpace(result)
		if result == ""{
			fmt.Println("null")
			return
		}
		fmt.Println(result)
	},
}