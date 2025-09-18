# ZapLog - 灵活且高性能的 Go 日志工具

ZapLog 是一个轻量级、灵活的 Go 日志工具，基于快速且结构化的日志库 [zap](https://github.com/uber-go/zap) 构建。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

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

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
