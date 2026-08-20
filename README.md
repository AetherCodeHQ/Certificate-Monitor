# Certificate Monitor

![CI](https://github.com/Qyroxen/Certificate-Monitor/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Certificate-Monitor/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Certificate-Monitor?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Certificate-Monitor)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Certificate-Monitor)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Certificate-Monitor?style=social)](https://github.com/Qyroxen/Certificate-Monitor/stargazers)

## What is it?

Certificate Monitor is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Certificate-Monitor.git
cd Certificate-Monitor
go build -o certificatemonitor .

# Run
./certificatemonitor --help
```

## CLI Usage

```bash
# Basic usage
./certificatemonitor

# With flags
./certificatemonitor --verbose --output json

# Get help
./certificatemonitor --help
```

## Examples

```bash
# Example 1
./certificatemonitor example1

# Example 2
./certificatemonitor example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o certificatemonitor .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Certificate-Monitor/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Certificate-Monitor?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Certificate-Monitor/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Certificate-Monitor?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Certificate-Monitor/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Certificate-Monitor" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Certificate-Monitor/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Certificate-Monitor" alt="Pull Requests">
  </a>
</p>
