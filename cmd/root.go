package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	initFlag     bool
	previousFlag bool
	questionStr  string
)

var rootCmd = &cobra.Command{
	Use:   "aid",
	Short: "Aid is an LLM-powered terminal assistant",
	Long: `Aid helps you understand shell commands and learn how to perform tasks.

Examples:
  aid --init                  Create config file at ~/.config/aid/config.toml
  aid -p                      Explain the previous command and what went wrong
  aid -q "list files"         Get help on how to accomplish a task`,
	Run: func(cmd *cobra.Command, args []string) {
		if initFlag {
			runInit()
			return
		}

		if !previousFlag && questionStr == "" {
			cmd.Help()
			return
		}

		cfg, err := LoadConfig()
		if err != nil {
			PrintError(err.Error())
			os.Exit(1)
		}

		client := NewLLMClient(cfg)

		if previousFlag {
			runPrevious(client)
		} else if questionStr != "" {
			runQuestion(client, questionStr)
		}
	},
}

func runPrevious(client *LLMClient) {
	prevCmd, err := GetPreviousCommand()
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	PrintInfo("Analyzing: " + prevCmd)

	response, err := client.AnalyzePreviousCommand(prevCmd)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	PrintResponse(response)
}

func runQuestion(client *LLMClient, q string) {
	response, err := client.Question(q)
	if err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	PrintResponse(response)
}

func runInit() {
	if err := InitConfig(); err != nil {
		PrintError(err.Error())
		os.Exit(1)
	}

	PrintInfo("Config file created at " + DefaultConfigPath())
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&initFlag, "init", false, "Create config file")
	rootCmd.Flags().BoolVarP(&previousFlag, "previous", "p", false, "Analyze the previous shell command")
	rootCmd.Flags().StringVarP(&questionStr, "question", "q", "", "Ask how to do something")
}
