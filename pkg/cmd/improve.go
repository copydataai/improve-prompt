package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/copydataai/improve-prompt/ai"
	"github.com/spf13/cobra"
)

func improveFunc(cmd *cobra.Command, args []string) error {
	provider, _ := cmd.Flags().GetString("provider")
	model, _ := cmd.Flags().GetString("model")
	server, _ := cmd.Flags().GetString("server")
	input, _ := cmd.Flags().GetString("input")
	output, _ := cmd.Flags().GetString("output")
	verbose, _ := cmd.Flags().GetBool("verbose")

	// Get the prompt from args, file, or stdin
	var prompt string
	var err error

	if len(args) > 0 {
		prompt = strings.Join(args, " ")
	} else if input != "" {
		content, err := os.ReadFile(input)
		if err != nil {
			return fmt.Errorf("failed to read input file: %w", err)
		}
		prompt = string(content)
	} else {
		// Read from stdin
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read from stdin: %w", err)
		}
		prompt = string(data)
	}

	if prompt == "" {
		return errors.New("empty prompt. Please provide a prompt to improve")
	}

	// Initialize AI provider
	var aiProvider ai.Provider
	switch strings.ToLower(provider) {
	case "openai":
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			return errors.New("OPENAI_API_KEY environment variable not set")
		}
		aiProvider = ai.NewOpenAIProvider(apiKey, model)
	case "ollama":
		aiProvider = ai.NewOllamaProvider(server, model)
	default:
		return fmt.Errorf("unsupported provider: %s", provider)
	}

	// Improve the prompt
	var improved string
	if verbose {
		improved, err = ai.ImprovePrompt(prompt, aiProvider)
		if err != nil {
			return fmt.Errorf("failed to improve prompt: %w", err)
		}
	} else {
		improved, err = ai.ImprovePromptSimple(prompt, aiProvider)
		if err != nil {
			return fmt.Errorf("failed to improve prompt: %w", err)
		}
	}

	// Output the improved prompt
	if output != "" {
		err = os.WriteFile(output, []byte(improved), 0644)
		if err != nil {
			return fmt.Errorf("failed to write to output file: %w", err)
		}
		fmt.Printf("Improved prompt written to %s\n", output)
	} else {
		fmt.Println(improved)
	}

	return nil
}
