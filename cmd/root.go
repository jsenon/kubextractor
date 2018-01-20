package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// type Config struct {
// }

var cfgFile string
var context string

var rootCmd = &cobra.Command{
	Use:   "kubextractor",
	Short: "Extract k8s context from global config file",
	Long: `Extract kubernetes context ie. configuration
				  user and endpoint.
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

func init() {

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.kube/config")
	rootCmd.PersistentFlags().StringP("author", "a", "SENON Julien", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringP("userLicense", "l", "Apache", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().StringVarP(&context, "c", "", "Default", "Name of  context to extract")

	viper.BindPFlag("cfgFile", rootCmd.PersistentFlags().Lookup("cfgFile"))
	viper.BindPFlag("userLicense", rootCmd.PersistentFlags().Lookup("userLicense"))
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))

	viper.SetDefault("author", "SENON Julien julien.senon@gmail.com")
	viper.SetDefault("userLicense", "apache")

}
