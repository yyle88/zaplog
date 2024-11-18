# ZapLog - Flexible and High-Performance Logging for Go

ZapLog is a lightweight, flexible logging utility for Go applications, built on top of the fast and structured logging library [zap](https://github.com/uber-go/zap). It’s perfect for developers who need fine-grained control over their log outputs and want to handle concurrent logging safely and efficiently.

Whether you're working on a small project or a complex, multi-threaded application, ZapLog lets you configure logging levels and handle log outputs in an easy-to-use interface, allowing you to focus on what really matters—building your app.

## README
[中文说明](README.zh.md)

## Key Features

- **Control Logging Levels**: Skip logs at different verbosity levels, making it easy to adjust what information gets logged based on the environment or execution stage.
- **Concurrency-Safe**: Perfect for multi-threaded applications. ZapLog safely handles concurrent log access with `mutexmap.Map`.
- **Optimized for Performance**: Built on `zap`, so it’s fast and efficient even in high-load applications.
- **Easy to Integrate**: Simple setup with a familiar interface, using `zap` as the backend.

## Why Use ZapLog?

### 1. **Simple and Flexible Logging Levels**
With ZapLog, you don’t have to worry about excessive log outputs cluttering your console or log files. You can easily control what’s logged, based on configurable verbosity levels (from `P0` to `P4`). Whether you’re debugging, in production, or need detailed traces, ZapLog has you covered.

### 2. **Concurrency-Safe Logging**
If you're working with a multi-threaded or multi-goroutine application, you need to ensure that logging doesn’t cause race conditions. ZapLog uses `mutexmap.Map` to provide thread-safe access to loggers, ensuring your logs are handled efficiently and safely, even in highly concurrent environments.

### 3. **Seamless Integration**
Integrating ZapLog into your project is a breeze. It works with your existing `zap` logger setup, and it’s simple to configure different log levels for different parts of your application.

## Installation

```bash
go get github.com/yyle88/zaplog
```

## Getting Started

### Basic Example

```go
package main

import (
	"go.uber.org/zap"
	"github.com/yyle88/zaplog"
)

func main() {
	// Create a new zap logger instance
	logger, _ := zap.NewProduction()

	// Initialize ZapLog with the logger
	logs := zaplog.NewSkipZaps(logger)

	// Use different log levels
	logs.Pn(0).Info("This is a P0 log")
	logs.Pn(1).Warn("This is a P1 log")
	logs.Pn(2).Error("This is a P2 log")
}
```

### Advanced Usage: Concurrency

ZapLog handles concurrent access to logs using a mutex-based map, making it ideal for multi-threaded applications. Here’s how you can use it in a concurrent setting:

```go
logs := zaplog.NewSkipZaps(logger)
logs.Pn(3).Info("Thread-safe logging example")
```

## Contributing

We welcome contributions! Whether you’ve fixed a bug, improved documentation, or added a new feature, your contributions are greatly appreciated. Please follow the standard fork-and-pull request process.

## License

ZapLog is open-source and released under the [MIT License](LICENSE).

## Thank You

If you find this package valuable, give it a star on GitHub! Thank you!!!
