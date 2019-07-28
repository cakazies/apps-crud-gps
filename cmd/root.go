package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	conf "github.com/local/app-gps/application/models"
	"github.com/local/app-gps/routes"
	"github.com/local/app-gps/utils"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "Management GPS",
	Short: "Apps Management GPS",
	Long:  `Apps Management GPS with Login`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("apps GPS is avaible running")
		routes.Route()
	},
}

func init() {
	cobra.OnInitialize(splash, InitViper, conf.Connect)
}

// Opened
func splash() {
	fmt.Println(`
	_____ ____________________            __________________  _________
	/  _  \\______   \______   \          /  _____/\______   \/   _____/
   /  /_\  \|     ___/|     ___/  ______ /   \  ___ |     ___/\_____  \ 
  /    |    \    |    |    |     /_____/ \    \_\  \|    |    /        \
  \____|__  /____|    |____|              \______  /|____|   /_______  /
		  \/                                     \/                  \/ 
	`)
	// http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20
}

// InitViper from file toml
func InitViper() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
	log.Println("Using Config File: ", viper.ConfigFileUsed())
}

// Execute from Cobra Firsttime
func Execute() {
	err := rootCmd.Execute()
	utils.FailError(err, "Error Execute RootCMD")
}
