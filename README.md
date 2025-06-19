[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/zaplog/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/zaplog/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/zaplog)](https://pkg.go.dev/github.com/yyle88/zaplog)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/zaplog/master.svg)](https://coveralls.io/github/yyle88/zaplog?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/zaplog.svg)](https://github.com/yyle88/zaplog/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/zaplog)](https://goreportcard.com/report/github.com/yyle88/zaplog)

# ZapLog - Flexible and High-Performance Logging for Go

ZapLog is a lightweight, flexible logging utility for Go applications, built on top of the fast and structured logging pkg [zap](https://github.com/uber-go/zap). 

## README
[中文说明](README.zh.md)

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

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

## GitHub Stars

[![starring](https://starchart.cc/yyle88/zaplog.svg?variant=adaptive)](https://starchart.cc/yyle88/zaplog)
