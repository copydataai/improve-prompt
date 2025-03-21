# Improve Prompt (impv) 🔄✨

A CLI tool that enhances AI prompts using OpenAI or local Ollama models.

![CLI Demo](https://github.com/copydataai/improve-prompt/blob/95d8e1a1d0615c4a5aa6ca1bb15047a81d2240ed/impv.gif)

## Features 🚀

- **AI-Powered Improvements** - Optimize prompts using state-of-the-art models
- **Multi-Provider Support** - Choose between OpenAI or local Ollama models
- **Simple Interface** - Streamlined CLI workflow with smart defaults
- **Configurable** - Customize model parameters and output formats

## Installation 📦

```bash
go install github.com/copydataai/improve-prompt/cmd/impv@latest
```

## Usage 📖

```bash
impv "Your initial prompt here" \
  --provider -p [openai|ollama] \
  --model -m [gpt-4o-mini|gemma3:1b] \
  --verbose -v \
  --input -i prompts.txt \
  --output -o improved-prompts.txt
```

## Configuration ⚙️

### OpenAI

```bash
export OPENAI_API_KEY="your-api-key"
```

### Ollama

```bash
# If you want to specify the Ollama url use --server or -s
impv --provider ollama --server http://localhost:11434 --model gemma3:1b
```

## Examples 🌟

```bash
# Basic usage with OpenAI
impv "help to write a poem like DaVinci" --provider openai

# Batch process prompts from file
impv --input prompts.txt --output improved.txt

# Use local Ollama model
impv "design a database for a hotel" --provider ollama --model gemma3:1b

# Use verbose mode
impv "Write a poem about AI" --verbose
```

## Development 🛠️

Contributions welcome! Please see our [contribution guidelines](CONTRIBUTING.md).

---

📄 **License**: [MIT](LICENSE)  
💬 **Support**: Open an issue
