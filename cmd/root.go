package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// type Config struct {
// }

var jsonfile string
var context string
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "kubextractor",
	Short: "Extract k8s context from global config file",
	Long: `Extract kubernetes context ie. configuration user and endpoint.
				  Complete documentation is available at https://github.com/jsenon/kubextractor`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		content, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			fmt.Print("Error:", err)
		}
		// var conf Config
		// err = json.Unmarshal(content, &conf)
		// if err != nil {
		// 	fmt.Print("Error:", err)
		// }
		fmt.Println(content)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mytest.yaml)")

	rootCmd.PersistentFlags().StringVarP(&jsonfile, "jsonfile", "c", "", "config file (default is $HOME/.kube/config")
	rootCmd.PersistentFlags().StringVarP(&context, "context", "e", "", "Name of  context to extract")

	viper.BindPFlag("jsonfile", rootCmd.PersistentFlags().Lookup("jsonfile"))
	viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mytest" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mytest")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
