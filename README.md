# Aid

An LLM-powered terminal assistant that helps you understand shell commands and learn how to perform tasks.

## Prerequisites

- [Go](https://go.dev/dl/)
- [Ollama](https://ollama.com/) with a model installed, e.g. [mistral:7b](https://ollama.com/library/mistral:7b)

```sh
ollama pull mistral:7b
```

### Shell configuration

By default, zsh only writes to history when the session closes. For `aid -p` to work with recent commands, add to your `~/.zshrc`:

```sh
setopt INC_APPEND_HISTORY
```

For bash, add to your `~/.bashrc`:

```sh
shopt -s histappend
PROMPT_COMMAND="history -a;$PROMPT_COMMAND"
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
# Analyze the previous shell command (what went wrong, how to fix it)
aid -p
aid --previous

# Ask how to do something
aid -q "find all .go files larger than 1MB"
aid --question "list files by size"
```
