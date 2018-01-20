package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Kind        string `json:"kind"`
	APIVersion  string `json:"apiVersion"`
	Preferences struct {
	} `json:"preferences"`
	Clusters []struct {
		Name    string `json:"name"`
		Cluster struct {
			Server                string `json:"server"`
			InsecureSkipTLSVerify bool   `json:"insecure-skip-tls-verify"`
		} `json:"cluster"`
	} `json:"clusters"`
	Users []struct {
		Name string `json:"name"`
		User struct {
			ClientCertificateData string `json:"client-certificate-data"`
			ClientKeyData         string `json:"client-key-data"`
		} `json:"user"`
	} `json:"users"`
	Contexts []struct {
		Name    string `json:"name"`
		Context struct {
			Cluster string `json:"cluster"`
			User    string `json:"user"`
		} `json:"context"`
	} `json:"contexts"`
	CurrentContext string `json:"current-context"`
}

var jsonfile string
var context string
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "kubextractor",
	Short: "Extract k8s context from global config file",
	Long: `Extract kubernetes context ie. configuration user and endpoint.
				  Complete documentation is available at https://github.com/jsenon/kubextractor`,
	Run: func(cmd *cobra.Command, args []string) {
		// jsonfile = "/Users/julien/.kube/config.json"
		// Do Stuff Here
		fmt.Println("JSON:", jsonfile)
		file, err := os.Open(jsonfile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)

		// str := string(b)
		// fmt.Println(str)

		res := &Config{}
		json.Unmarshal([]byte(string(b)), &res)
		// fmt.Println(res)
		fmt.Println("Context Asked:", context)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&jsonfile, "config", "c", "/Users/julien/.kube/config.json", "k8s config file ")
	rootCmd.PersistentFlags().StringVarP(&context, "context", "e", "", "Name of  context to extract")

	viper.BindPFlag("jsonfile", rootCmd.PersistentFlags().Lookup("jsonfile"))
	viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))

}
