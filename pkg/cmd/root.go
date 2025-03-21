package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "impv",
	Short: "A CLI tool to improve prompts using AI models",
	Long: `impv "improve-prompt" is a CLI tool that helps you improve your prompts
for AI models. It uses AI to analyze and enhance your prompts,
making them more effective for your specific use case.

It supports various AI providers including local models through Ollama.`,
	Run: improveCmd.Run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you will define your flags and configuration settings.
	rootCmd.PersistentFlags().StringP("model", "m", "gpt-4o-mini", "AI model to use for improving prompts")
	rootCmd.PersistentFlags().StringP("provider", "p", "openai", "AI provider to use (openai, ollama)")
	rootCmd.PersistentFlags().StringP("server", "s", "http://localhost:11434", "Server address for Ollama")
}
