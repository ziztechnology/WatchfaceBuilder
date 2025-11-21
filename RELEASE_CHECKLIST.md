# Release Checklist

## Pre-release 检查清单

在发布到 GitHub 之前，请确保完成以下步骤：

### 1. 代码质量
- [x] 代码编译通过 (`go build ./cmd/cli`)
- [x] 所有测试通过 (`go test ./...`)
- [ ] 代码格式化 (`go fmt ./...`)
- [ ] 代码检查 (`golangci-lint run`)

### 2. 文档
- [x] README.md (英文) 完整且准确
- [x] README_zh.md (中文) 完整且准确
- [x] CONTRIBUTING.md 包含贡献指南
- [x] LICENSE 文件存在 (MIT)
- [x] docs/QUICKSTART.md 快速开始指南
- [x] docs/API.md API 文档
- [x] PROJECT_SUMMARY.md 项目总结

### 3. 示例和测试
- [x] examples/ 目录包含示例
- [ ] 测试所有三个模板 (simple, analog, digital)
- [ ] 测试自定义模板功能
- [ ] 测试交互式模式
- [ ] 测试命令行模式

### 4. 构建和发布
- [x] Makefile 包含所有必要的构建命令
- [x] .gitignore 配置正确
- [x] go.mod 和 go.sum 是最新的
- [x] GitHub Actions workflows 配置正确
- [x] Git tags 创建 (v0.1.0)

### 5. GitHub 准备
- [ ] 创建 GitHub 仓库: https://github.com/Toooony/WatchfaceBuilder
- [ ] 推送代码到 GitHub
- [ ] 创建第一个 Release
- [ ] 上传预构建的二进制文件
- [ ] 启用 GitHub Discussions
- [ ] 启用 GitHub Issues
- [ ] 添加仓库描述和标签

## 发布步骤

### 步骤 1: 创建 GitHub 仓库

1. 访问 https://github.com/new
2. 仓库名称: `WatchfaceBuilder`
3. 描述: `🎨 A lazy-friendly tool to quickly generate H5 watchface packages`
4. 设置为 Public
5. 不要初始化 README (我们已经有了)
6. 创建仓库

### 步骤 2: 推送代码

```bash
cd D:\47.107.248.143\Toooony\WatchfaceBuilder

# 添加远程仓库
git remote add origin https://github.com/Toooony/WatchfaceBuilder.git

# 推送代码和 tags
git push -u origin main
git push --tags
```

### 步骤 3: 创建 Release

#### 方式 1: 通过 GitHub Actions (推荐)

推送 tag 后，GitHub Actions 会自动构建并创建 Release：

```bash
# Tag 已经创建，只需推送
git push --tags
```

#### 方式 2: 手动创建 Release

如果自动发布失败，可以手动创建：

1. 访问 https://github.com/Toooony/WatchfaceBuilder/releases/new
2. 选择 tag: `v0.1.0`
3. Release 标题: `v0.1.0 - Initial Release`
4. 描述：

```markdown
## WatchfaceBuilder v0.1.0 - Initial Release 🎉

🎨 A lazy-friendly tool to quickly generate H5 watchface packages for WebView-based smartwatches.

### Features

- ✅ **CLI Tool**: Easy-to-use command-line interface with interactive mode
- ✅ **Multiple Templates**: Built-in templates (Simple, Analog, Digital)
- ✅ **Custom Template Support**: Use your own HTML/CSS/JS
- ✅ **Auto Preview**: Automatically generate preview images
- ✅ **Cross-Platform**: Windows, macOS, Linux support

### Templates

1. **Simple** - Minimalist digital clock with gradient background
2. **Analog** - Classic clock with Canvas rendering
3. **Digital** - Tech-style digital clock with neon effects
4. **Custom** - Fully customizable with your own code

### Installation

Download the binary for your platform:

- **Linux**: `watchface-builder-v0.1.0-linux-amd64.tar.gz`
- **macOS (Intel)**: `watchface-builder-v0.1.0-darwin-amd64.tar.gz`
- **macOS (Apple Silicon)**: `watchface-builder-v0.1.0-darwin-arm64.tar.gz`
- **Windows**: `watchface-builder-v0.1.0-windows-amd64.zip`

### Quick Start

```bash
# Extract (Linux/macOS)
tar -xzf watchface-builder-v0.1.0-*.tar.gz

# Make executable (Linux/macOS)
chmod +x watchface-builder-*

# Run interactive mode
./watchface-builder -i
```

Windows:
```cmd
unzip watchface-builder-v0.1.0-windows-amd64.zip
watchface-builder-windows-amd64.exe -i
```

### Documentation

- 📖 [Quick Start Guide](docs/QUICKSTART.md)
- 📖 [README (English)](README.md)
- 📖 [中文文档](README_zh.md)
- 🤝 [Contributing Guide](CONTRIBUTING.md)

### What's Next?

- More built-in templates
- Web UI interface
- Template marketplace
- Live preview in browser

---

🙏 Thank you for using WatchfaceBuilder!

Report issues: https://github.com/Toooony/WatchfaceBuilder/issues
```

5. 上传构建的二进制文件
6. 发布

### 步骤 4: 配置仓库

1. **启用 Discussions**:
   - Settings → Features → Discussions → 勾选

2. **配置 Issues**:
   - Settings → Features → Issues → 勾选

3. **添加 Topics (标签)**:
   - 点击仓库页面的 "Add topics"
   - 添加: `watchface`, `h5`, `webview`, `smartwatch`, `cli-tool`, `golang`, `template-engine`

4. **更新仓库描述**:
   - Description: `🎨 A lazy-friendly tool to quickly generate H5 watchface packages`
   - Website: 留空或添加文档链接

### 步骤 5: 推广

1. **在 README 中添加 badges**:
   - GitHub Stars
   - Go Report Card
   - License
   - Release version

2. **社交媒体分享** (可选):
   - Twitter
   - Reddit (r/golang, r/opensource)
   - Hacker News
   - 掘金
   - V2EX

3. **提交到项目列表** (可选):
   - awesome-go
   - awesome-cli
   - awesome-tools

## 发布后检查

- [ ] README badges 显示正常
- [ ] GitHub Actions 运行成功
- [ ] Release 页面显示所有二进制文件
- [ ] 下载链接可用
- [ ] Issues 和 Discussions 已启用
- [ ] 仓库描述和标签正确

## 维护计划

### 短期 (1-2 周)
- 修复用户报告的 bugs
- 改进文档
- 添加更多示例

### 中期 (1-2 月)
- 添加新模板
- 实现 Web UI
- 改进预览图生成

### 长期 (3-6 月)
- Web API 服务
- 模板市场
- 桌面 GUI 应用

## 注意事项

1. **版本号规范**: 遵循语义化版本 (Semantic Versioning)
   - Major: 重大变更
   - Minor: 新功能
   - Patch: Bug 修复

2. **Changelog**: 维护 CHANGELOG.md 记录所有变更

3. **Breaking Changes**: 在 Release Notes 中明确标注

4. **安全**: 定期更新依赖，关注安全漏洞

5. **社区**: 及时回复 Issues 和 Pull Requests

---

**准备开源**: ✅ 所有必要文件已准备就绪

**下一步**: 创建 GitHub 仓库并推送代码
