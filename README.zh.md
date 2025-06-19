# ZapLog - 灵活且高性能的 Go 日志工具

ZapLog 是一个轻量级、灵活的 Go 日志工具，基于快速且结构化的日志库 [zap](https://github.com/uber-go/zap) 构建。

## README
[English Documentation](README.md)

## 安装

```bash
go get github.com/yyle88/zaplog
```

## 核心功能

### 1. **基础日志打印**

您可以使用 `zaplog.LOG` 记录带有键值对的日志：

```go
zaplog.LOG.Debug("调试信息", zap.String("key", "value"))
zaplog.LOG.Error("错误信息", zap.Error(errors.New("错误")))
```

### 2. **打印多个键值对**

通过传递多个键值对，您可以在单条日志中打印多个字段：

```go
zaplog.LOG.Debug("调试信息", zap.String("key1", "value1"), zap.Int("key2", 2))
zaplog.LOG.Error("错误信息", zap.Int("key1", 1), zap.Error(errors.New("错误")))
```

### 3. **使用 `SugaredLogger`**

对于更简化的日志记录，您可以使用 `zaplog.SUG`，它支持变参形式的日志记录：

```go
SUG.Debug("简化日志", "key1", "value1", "key2", 2)
SUG.Error("简化错误", errors.New("错误"))
```

### 4. **创建子日志（SubZap）**

您可以创建带有默认字段的子日志，以提供额外的上下文信息，使日志更具可读性。使用 `SubZap`、`NewZap` 创建子日志：

#### 使用 `SubZap` 创建子日志：

```go
zp := zaplog.LOGGER.SubZap(zap.String("module", "abc"), zap.String("key", "value"))
zp.LOG.Debug("子日志信息", zap.Int("a", 1))
zp.SUG.Error("简化子日志错误", 1, 2, 3)
```

#### 使用 `NewZap` 创建子日志：

```go
zp := zaplog.LOGGER.NewZap("module", "abc", zap.String("key", "value"))
zp.LOG.Debug("子日志信息2", zap.Int("a", 2))
```

### 5. **在 SugaredLogger 中处理多参数**

在 `SugaredLogger` 中，您可以传递各种类型的参数，包括数组和切片：

```go
zaplog.SUG.Debug("调试信息", 1, 2, 3)
zaplog.SUG.Debug([]int{0, 1, 2})
```

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
