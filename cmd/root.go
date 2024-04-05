package cmd

import (
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-curl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		u, err := url.Parse(args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("unable to parse URL")
		}

		host := u.Hostname()
		port := u.Port()
		path := u.Path

		if port == "" {
			port = "80"
		}

		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))

		if err != nil {
			log.Fatal().Err(err).Msg("error connecting to tcp")
		}

		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Fatal().Err(err).Msg("error closing tcp connection")
			}
		}(conn)

		_, err = fmt.Fprintf(conn, "GET %s HTTP/1.0\r\nHost: %s\r\n\r\n", path, host)
		if err != nil {
			log.Info().Err(err).Msg("error writing to tcp connection")
			return
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			log.Fatal().Err(err).Msg("error reading from tcp connection")
		}

		fmt.Println(string(buf[:n]))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-curl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
