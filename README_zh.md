# Watchface Builder

[English](README.md) | [中文](README_zh.md)

🎨 为 WebView 套壳智能手表快速生成 H5 表盘包的懒人工具

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

## ✨ 特性

- 🚀 **快速开始**: 几秒钟生成完整的表盘包
- 🎨 **多种模板**: 内置模板（简约、指针、数字）
- 🛠️ **CLI 工具**: 易用的命令行界面，支持交互式模式
- 📦 **符合规范**: 生成的包符合 H5 表盘规范
- 🔧 **可定制**: 支持自定义 HTML/CSS/JS
- 🖼️ **自动预览**: 自动生成预览图
- 🌐 **Web API**: 可作为 Web 服务集成

## 📦 安装

### 预编译二进制文件

从 [Releases](https://github.com/Toooony/WatchfaceBuilder/releases) 下载最新版本。

### 从源码构建

```bash
git clone https://github.com/Toooony/WatchfaceBuilder.git
cd WatchfaceBuilder
go build -o watchface-builder ./cmd/cli
```

## 🚀 快速开始

### 交互式模式（推荐新手）

```bash
./watchface-builder -i
```

工具会提示你输入：
- 表盘名称（必需）
- 版本号（默认: 1.0.0）
- 作者（默认: Anonymous）
- 描述（可选）
- 模板选择（1=简约, 2=指针, 3=数字）
- 标签（可选，逗号分隔）

### 命令行模式

```bash
./watchface-builder -name "我的表盘" -template analog -version 1.0.0
```

### 列出可用模板

```bash
./watchface-builder -list
```

## 📋 模板

### 1. Simple（简约数字时钟）
简洁的数字时钟，带渐变背景。
- 时间显示（HH:MM:SS）
- 日期显示
- 响应式设计

### 2. Analog（指针时钟）
经典指针时钟，Canvas 绘制。
- 时针、分针、秒针
- 表盘刻度
- 流畅动画

### 3. Digital（数字显示时钟）
现代数字时钟，霓虹灯效果。
- 大号时间显示
- 日期和星期
- 科技感设计

### 4. Custom（自定义）
使用你自己的 HTML/CSS/JS 完全自定义。

## 📖 使用方法

### CLI 参数

```
  -name string
        表盘名称（必需）
  -version string
        版本号（默认 "1.0.0"）
  -author string
        作者名称（默认 "Anonymous"）
  -description string
        表盘描述
  -template string
        模板类型: simple, analog, digital（默认 "simple"）
  -tags string
        标签，逗号分隔
  -output string
        输出目录（默认 "."）
  -no-preview
        不生成预览图
  -i    交互式模式
  -list
        列出可用模板
```

### 生成的包结构

```
watchface.zip
├── manifest.json       # 表盘元数据
├── index.html          # 入口文件
├── style.css           # 样式
├── script.js           # JavaScript 逻辑
└── preview.png         # 预览图（可选）
```

### manifest.json 示例

```json
{
  "name": "我的表盘",
  "version": "1.0.0",
  "author": "你的名字",
  "description": "一个漂亮的表盘",
  "entrypoint": "index.html",
  "tags": ["简约", "数字"],
  "created_at": "2025-01-21T10:30:00Z"
}
```

## 🌐 Web API 模式

你可以将 Watchface Builder 作为 Web 服务运行：

```bash
./watchface-builder serve --port 8080
```

### API 端点

**构建表盘**
```http
POST /api/build
Content-Type: application/json

{
  "name": "我的表盘",
  "version": "1.0.0",
  "template": "analog",
  "author": "你的名字",
  "tags": ["经典", "指针"]
}
```

**列出模板**
```http
GET /api/templates
```

## 🔧 高级用法

### 自定义模板

使用自定义 HTML/CSS/JS 创建表盘：

```bash
./watchface-builder -name "自定义" -template custom \
  -custom-html "$(cat my-template.html)" \
  -custom-css "$(cat my-style.css)" \
  -custom-js "$(cat my-script.js)"
```

### 批量生成

```bash
for template in simple analog digital; do
  ./watchface-builder -name "表盘-$template" -template $template
done
```

## 🤝 贡献

欢迎贡献！请先阅读我们的[贡献指南](CONTRIBUTING.md)。

### 添加新模板

1. Fork 这个仓库
2. 在 `pkg/templates/` 中创建你的模板
3. 将模板元数据添加到 `pkg/templates/registry.go`
4. 提交 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 🙏 致谢

- [fogleman/gg](https://github.com/fogleman/gg) - 图像生成
- [spf13/cobra](https://github.com/spf13/cobra) - CLI 框架

## 📞 支持

- 🐛 [报告问题](https://github.com/Toooony/WatchfaceBuilder/issues)
- 💬 [讨论](https://github.com/Toooony/WatchfaceBuilder/discussions)
- 📖 [文档](docs/)

## 🗺️ 路线图

- [ ] 更多内置模板
- [ ] 模板市场
- [ ] 浏览器实时预览
- [ ] GUI 应用程序
- [ ] 模板编辑器
- [ ] 云端构建器

---

用 ❤️ 制作 by [Toooony](https://github.com/Toooony)
