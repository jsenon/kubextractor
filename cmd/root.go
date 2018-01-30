package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config Struct of JSON Kubernetes Config File
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
var output string

var rootCmd = &cobra.Command{
	Use:   "kubextractor",
	Short: "Extract k8s context from global config file",
	Long: `Extract kubernetes context ie. configuration user and endpoint.
				  Complete documentation is available at https://github.com/jsenon/kubextractor
				  After export Use kubectl config use-context YOURCONTEXT --kubeconfig output.json to use it`,
	// Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		defaultfilejson := "/.kube/config.json"
		defaultfile := "/.kube/config"

		tempfile := ".convert.json"

		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		cmdName := "kubectl"
		cmdArgs := []string{"config", "view", "-o", "json", "--raw", "--kubeconfig", usr.HomeDir + defaultfile}

		// If no value for config k8s file, use default config but we need to convert to json
		if cfgFile == "" && jsonfile == "" {

			cfgFile = usr.HomeDir + defaultfile
			out, erro := exec.Command(cmdName, cmdArgs...).Output() // #nosec
			if erro != nil {
				log.Fatal(erro)
			}

			err = ioutil.WriteFile(tempfile, out, 0644)
			if err != nil {
				log.Fatal(err)
			}

			jsonfile = tempfile
		}

		// If value for config k8s but no json we need to generate a json output
		if cfgFile != "" && jsonfile == "" {

			cmdArgs := []string{"config", "view", "-o", "json", "--raw", "--kubeconfig", cfgFile}

			cfgFile = usr.HomeDir + defaultfile
			out, erro := exec.Command(cmdName, cmdArgs...).Output() // #nosec
			if erro != nil {
				log.Fatal(erro)
			}

			err = ioutil.WriteFile(tempfile, out, 0644)
			if err != nil {
				log.Fatal(err)
			}

			jsonfile = tempfile

		}

		// Used default value for json config file
		// Exit if doesn't exist
		if jsonfile == "" {
			jsonfile = usr.HomeDir + defaultfilejson

		}

		file, err := os.Open(jsonfile)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close() // nolint: errcheck

		b, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		res := &Config{}
		var configoutput Config

		err = json.Unmarshal([]byte(string(b)), &res)
		if err != nil {
			log.Fatal(err)
		}

		configoutput.APIVersion = res.APIVersion
		configoutput.Kind = res.Kind

		// Loop over Clusters matching with context asked
		for _, coutput := range res.Clusters {

			if coutput.Name == context {
				configoutput.Clusters = append(configoutput.Clusters, coutput)

			}

		}

		// Loop over Users matching with contexy asked
		for _, coutput := range res.Users {

			if coutput.Name == context {
				configoutput.Users = append(configoutput.Users, coutput)
			}

		}

		// Loop over Contexts matching with contexy asked
		for _, coutput := range res.Contexts {

			if coutput.Name == context {
				configoutput.Contexts = append(configoutput.Contexts, coutput)
			}

		}

		// Output to console
		if output == "" {
			body, erro := json.MarshalIndent(configoutput, "", "   ")
			if erro != nil {
				log.Fatal(erro)
			}
			fmt.Println(string(body))
		} else {

			//Write to output file specified in args
			body, erro := json.MarshalIndent(configoutput, "", "   ")
			if erro != nil {
				log.Fatal(erro)
			}

			err = ioutil.WriteFile(output, body, 0644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Exported to:", output)

		}

		// Delete temporary file
		err = os.Remove(tempfile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute Viper Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Func to init Cobra Flag and bind flag
func init() {

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "k8s config file default ($HOME/.kube/config)")
	rootCmd.PersistentFlags().StringVarP(&jsonfile, "configjson", "j", "", "k8s config file JSON default ($HOME/.kube/config.json)")

	rootCmd.PersistentFlags().StringVarP(&context, "context", "e", "", "MANDATORY: Name of  context to extract")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Name of output file")

	err := rootCmd.MarkPersistentFlagRequired("context")
	if err != nil {
		log.Fatal(err)
	}

	err = viper.BindPFlag("jsonfile", rootCmd.PersistentFlags().Lookup("jsonfile"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("context", rootCmd.PersistentFlags().Lookup("context"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	if err != nil {
		log.Fatal(err)
	}
	err = viper.BindPFlag("cfgFile", rootCmd.PersistentFlags().Lookup("cfgFile"))
	if err != nil {
		log.Fatal(err)
	}

}
