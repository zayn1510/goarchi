package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goarchi",
	Short: color.HiCyanString("Goarchi CLI for generating Golang boilerplate code"),
	Long: color.HiWhiteString(`
%s

%s

%s

%s

%s

%s

%s

%s

%s

%s

%s
`,
		color.New(color.FgHiBlue, color.Bold).Sprint("📦 Goarchi - Simple Layered Architecture Generator for Golang"),

		color.HiGreenString("🔧 Controller:\n  goarchi archi controller [name]")+
			"\n    → Generate a controller (e.g. UserController)",

		color.HiGreenString("🛠️  Service:\n  goarchi archi service [name]")+
			"\n    → Generate a service layer (e.g. UserService)",

		color.HiGreenString("📝 Request:\n  goarchi archi request [name] [fields...]")+
			"\n    → Generate a request struct with validation (e.g. name:string age:int)",

		color.HiGreenString("📦 Resource:\n  goarchi archi resource [name]")+
			"\n    → Generate a response formatter (DTO/transformer)",

		color.HiGreenString("🧩 Model:\n  goarchi archi model [name] [fields...]")+
			"\n    → Generate a GORM model with tags\n    → Example: goarchi archi model users \"id:int;primaryKey\" \"name:string;not null\"",

		color.HiGreenString("🛠️  Migration:\n  goarchi archi migration [name]")+
			"\n    → Generate a migration file in 'database/migrations'",

		color.HiGreenString("🧬 Migrate:\n  goarchi archi migrate [up|down]")+
			"\n    → 'up' applies migrations, 'down' rolls them back\n    → Looks for .sql files in 'database/migrations'",

		color.HiYellowString("📌 Installation via Go (Linux/macOS/Windows):")+
			"\n  goarchi install"+
			"\n  → Will build and (optionally) move the binary to your PATH",

		color.HiYellowString("📁 After install:")+
			"\n  You can use 'goarchi' globally from any folder.",

		color.HiYellowString("📌 Upgrade via Go (Linux/macOS/Windows):")+
			"\n  goarchi upgrade"+
			"\n  → Rebuilds the binary and (optionally) replaces the existing one in your PATH",
	),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
