# Aid

## What is this?
This is a minimal cheat sheet CLI tool, containing commands I want to learn or commands I tend to forget.

## Usage
```sh
aid [command] [flags]
```

## Installation
The install.sh script is tailored to my development environment.

```sh
# Run install script
. ./install.sh
```

What it does:
1. Builds application from source
2. Creates the directory "~/.go_cli" if it does not already exist
3. Moves the `aid` binary to directory ~./go_cli/
4. Adds an `aid` alias to the .zprofile

Don't forget to source your .zprofile after alias has been added.
