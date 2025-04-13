package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zayn1510/goarchi/cmd/install"
)

var makeCmd = &cobra.Command{
	Use:   "archi",
	Short: "Generate code components like controller, model, etc.",
}
var airCmd = &cobra.Command{
	Use:   "air",
	Short: "Jalankan Goarchi dalam development mode dengan Air (hot reload)",
	Run: func(cmd *cobra.Command, args []string) {
		install.RunDev()
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Goarchi CLI globally",
	Run: func(cmd *cobra.Command, args []string) {
		install.RunInstall()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(makeCmd)
	rootCmd.AddCommand(airCmd)
}
