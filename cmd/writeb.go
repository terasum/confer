package cmd


import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/terasum/viper"
	"strconv"
)
var(
	b_vartype string
)
func init() {
	RootCmd.AddCommand(writebCmd)
	writebCmd.Flags().StringVarP(&b_vartype, "vartype", "t", "string", "specific the value type: string, int or bool")
}

var writebCmd = &cobra.Command{
	Use:   "writeb",
	Short: "write a config value and output the new config into stdout",
	Long:  `specific oigin config file path, then specific the key and value`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3{
			fmt.Println("Please specific the origin config file path ")
			fmt.Println("Example:")
			fmt.Println("confer writeb ./origin.json config.num 1 -t int")
			return
		}
		originConfigFilePath := Abs(args[0])
		if !Exist(originConfigFilePath){
			fmt.Println("Origin config file is not exist!")
			return
		}
		originViper := viper.New()
		originViper.SetConfigFile(originConfigFilePath)
		err := originViper.ReadInConfig()
		if err != nil{
			fmt.Println("Read the config file failed!")
			return
		}
		targetKey := args[1]
		targetValue := args[2]
		switch b_vartype {
		case "int": {
			intvar,err:=strconv.Atoi(targetValue)
			if err != nil{
				fmt.Errorf("connot convert the value %s into integer\n",targetValue)
				return
			}
			originViper.Set(targetKey,intvar)
		}
		case "bool":{
			if targetValue == "true"{
				originViper.Set(targetKey,true)
			}else{
				originViper.Set(targetKey,false)
			}
		}
		case "string":{
			originViper.Set(targetKey,targetValue)
		}
		case "map":{
			fmt.Println("unsupport value type map")
			return
		}
		default:{
			fmt.Printf("the type: %s you specific is wrong!\n",vartype)
			fmt.Println("Abort")
			return
		}
		}
		buf,err := originViper.WriteConfigBuffer()
		if err != nil{
			fmt.Errorf("Something Wrong: %v\n",err)
			return
		}
		fmt.Println((string)(buf))
	},
}