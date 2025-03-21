package ai

// Provider is an interface for AI providers
type Provider interface {
	// Generate generates a response from the AI model
	Generate(prompt string) (string, error)
}

// ImprovePrompt improves a prompt using the given AI provider
func ImprovePrompt(prompt string, provider Provider) (string, error) {
	improvementPrompt := `You are an AI prompt engineer with extensive expertise in crafting effective prompts.
Your task is to improve the following prompt to make it more clear, specific, and effective.
Consider the following aspects:
1. Clarity: Is the prompt clear about what it's asking for?
2. Specificity: Does it include enough details and constraints?
3. Context: Does it provide necessary context?
4. Structure: Is it well-organized and easy to understand?
5. Language: Is the language precise and unambiguous?

Here's the prompt to improve:

"""
` + prompt + `
"""

Provide an improved version of the prompt with explanations of what you changed and why.
Format your response as follows:

IMPROVED PROMPT:
[Your improved prompt goes here]

EXPLANATIONS:
[Your explanations of the changes made]
`

	return provider.Generate(improvementPrompt)
}
