# Grist

Grist stands for "grain that is made into flour" - a CLI tool to generate and transform data.

## Features

- **UUID Generation**: Generate random UUIDs
- **Time Conversion**: Convert between different time formats (RFC, PostgreSQL, seconds, milliseconds, microseconds, nanoseconds, protobuf)
- **Password Hashing**: Generate bcrypt hashes for passwords
- **Color Conversion**: Convert between color formats (hex, RGB, HSL, Tailwind)

## Installation

```bash
go install github.com/Aenluu/grist@latest
```

## Shell Setup

After installation, you need to ensure the Go binary directory is in your PATH.

### For Zsh 

Add this line to your `~/.zshrc` file:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then reload your shell configuration:

```bash
source ~/.zshrc
```

### For Bash

Add this line to your `~/.bashrc` or `~/.bash_profile` file:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then reload your shell configuration:

```bash
source ~/.bashrc  # or source ~/.bash_profile
```
