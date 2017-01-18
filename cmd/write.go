package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/terasum/viper"
	"strconv"
	"bufio"
	"os"
)
var(
	isConfirm bool
	vartype string

)
func init() {
	RootCmd.AddCommand(genprivCmd)
	genprivCmd.Flags().BoolVarP(&isConfirm, "confirm", "y", false, "confirm replace or not")
	genprivCmd.Flags().StringVarP(&vartype, "vartype", "t", "string", "specific the value type: string, int or bool")
}

var genprivCmd = &cobra.Command{
	Use:   "write",
	Short: "write a config value into target config file",
	Long:  `specific oigin config file and target config file path, then specific the key and value`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4{
			fmt.Println("Please specific the origin config file path and target file path")
			fmt.Println("Example:")
			fmt.Println("confer write ./origin.json ./target.json config.num 1 -t int")
			return
		}
		fmt.Println("Read the origin config file")
		originConfigFilePath := Abs(args[0])
		if !Exist(originConfigFilePath){
			fmt.Println("Origin config file is not exist!")
			return
		}
		targetConfigFilePath := Abs(args[1])
		originViper := viper.New()
		originViper.SetConfigFile(originConfigFilePath)
		err := originViper.ReadInConfig()
		if err != nil{
			fmt.Println("Read the config file failed!")
			return
		}
		targetKey := args[2]
		targetValue := args[3]

		fmt.Printf("Will replace the origin config file:\n=========================\n%s\n=========================\n",originConfigFilePath)
		fmt.Printf("Target config file is:\n=========================\n%s\n=========================\n",targetConfigFilePath)
		fmt.Printf("Key: %s\n",targetKey)
		fmt.Printf("Origin Value: %v\n",originViper.Get(targetKey))
		fmt.Printf("Replace as: %s\n",targetValue)
		switch vartype {
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
		default:{
			fmt.Printf("the type: %s you specific is wrong!\n",vartype)
			fmt.Println("Abort")
			return
			}
		}

		if isConfirm{
			err = originViper.WriteConfigAs(targetConfigFilePath)
		}else{
			fmt.Printf("Are you sure replace the key-value from: \n %s \n into\n %s ? \n (yes/YES/Yes) [default no]\n",originConfigFilePath,targetConfigFilePath)
			yesflag := "no"
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			yesflag = scan.Text()

			if yesflag == "yes" || yesflag == "Yes" || yesflag == "YES"{
				err = originViper.WriteConfigAs(targetConfigFilePath)
			}else{
				fmt.Println("Abort")
				return
			}
		}

		if err != nil{
			fmt.Println("Write new config file failed!")
			fmt.Errorf("Error info is %v\n",err)
			return
		}
		fmt.Println("replace the config file success!")


	},
}