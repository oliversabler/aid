# Aid

An LLM-powered terminal assistant that helps you understand shell commands and learn how to perform tasks.

## Prerequisites

- [Go](https://go.dev/dl/)
- [Ollama](https://ollama.com/) with a model installed, e.g. [mistral:7b](https://ollama.com/library/mistral:7b)

```sh
ollama pull mistral:7b
```

## Installation

```sh
./install.sh
source ~/.zprofile
```

## Configuration

Create a config file:

```sh
aid --init
```

Or manually create `~/.config/aid/config.toml`:

```toml
[api]
endpoint = "http://localhost:11434/v1/chat/completions"
model = "mistral:7b"
```

## Usage

```sh
# Ask how to do something
aid -q "find all .go files larger than 1MB"
aid --question "list files by size"
```
