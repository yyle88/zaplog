[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/zaplog/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/zaplog/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/zaplog)](https://pkg.go.dev/github.com/yyle88/zaplog)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/zaplog/main.svg)](https://coveralls.io/github/yyle88/zaplog?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/zaplog.svg)](https://github.com/yyle88/zaplog/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/zaplog)](https://goreportcard.com/report/github.com/yyle88/zaplog)

# ZapLog - Flexible and High-Performance Logging for Go

ZapLog is a lightweight, flexible logging utility for Go applications, built on top of the fast and structured logging pkg [zap](https://github.com/uber-go/zap).

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Installation

```bash
go get github.com/yyle88/zaplog
```

## Core Features

### 1. **Basic Logging**

You can log messages with key-value pairs using `zaplog.LOG`:

```go
zaplog.LOG.Debug("Debug message", zap.String("key", "value"))
zaplog.LOG.Error("Error message", zap.Error(errors.New("error")))
```

### 2. **Logging Multiple Key-Value Pairs**

Log multiple fields in a single log entry by passing multiple key-value pairs:

```go
zaplog.LOG.Debug("Debug message", zap.String("key1", "value1"), zap.Int("key2", 2))
zaplog.LOG.Error("Error message", zap.Int("key1", 1), zap.Error(errors.New("error")))
```

### 3. **Using `SugaredLogger`**

For simpler logging, you can use `zaplog.SUG`, which supports variadic arguments for logging:

```go
SUG.Debug("Simplified log", "key1", "value1", "key2", 2)
SUG.Error("Simplified error", errors.New("error"))
```

### 4. **Creating Sub-Loggers (SubZap)**

You can create sub-loggers with default fields for additional context, making your logs more informative. Use `SubZap`, `NewZap` for creating sub-loggers:

#### SubLogger Creation with `SubZap`:

```go
zp := zaplog.LOGGER.SubZap(zap.String("module", "abc"), zap.String("key", "value"))
zp.LOG.Debug("Sub-log message", zap.Int("a", 1))
zp.SUG.Error("Simplified sub-log error", 1, 2, 3)
```

#### SubLogger Creation with `NewZap`:

```go
zp := zaplog.LOGGER.NewZap("module", "abc", zap.String("key", "value"))
zp.LOG.Debug("Sub-log message 2", zap.Int("a", 2))
```

### 5. **Handling Multiple Arguments in Sugared Logger**

With `SugaredLogger`, you can pass various argument types, including arrays and slices:

```go
zaplog.SUG.Debug("Debug message", 1, 2, 3)
zaplog.SUG.Debug([]int{0, 1, 2})
```

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

## GitHub Stars

[![starring](https://starchart.cc/yyle88/zaplog.svg?variant=adaptive)](https://starchart.cc/yyle88/zaplog)
