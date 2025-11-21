# Watchface Builder

[English](README.md) | [中文](README_zh.md)

🎨 A lazy-friendly tool to quickly generate H5 watchface packages for WebView-based smartwatches.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

## ✨ Features

- 🚀 **Quick Start**: Generate a complete watchface package in seconds
- 🎨 **Multiple Templates**: Built-in templates (Simple, Analog, Digital)
- 🛠️ **CLI Tool**: Easy-to-use command-line interface with interactive mode
- 📦 **Standards Compliant**: Generated packages follow H5 watchface specifications
- 🔧 **Customizable**: Support for custom HTML/CSS/JS
- 🖼️ **Auto Preview**: Automatically generate preview images
- 🌐 **Web API**: Can be integrated as a web service

## 📦 Installation

### Pre-built Binaries

Download the latest release from [Releases](https://github.com/ziztechnology/WatchfaceBuilder/releases).

### Build from Source

```bash
git clone https://github.com/ziztechnology/WatchfaceBuilder.git
cd WatchfaceBuilder
go build -o watchface-builder ./cmd/cli
```

## 🚀 Quick Start

### Interactive Mode (Recommended for Beginners)

```bash
./watchface-builder -i
```

The tool will prompt you for:
- Watchface name (required)
- Version (default: 1.0.0)
- Author (default: Anonymous)
- Description (optional)
- Template selection (1=simple, 2=analog, 3=digital)
- Tags (optional, comma-separated)

### Command-Line Mode

```bash
./watchface-builder -name "My Watchface" -template analog -version 1.0.0
```

### List Available Templates

```bash
./watchface-builder -list
```

## 📋 Templates

### 1. Simple (Digital Clock)
A minimalist digital clock with gradient background.
- Time display (HH:MM:SS)
- Date display
- Responsive design

### 2. Analog (Clock with Hands)
Classic analog clock with Canvas rendering.
- Hour, minute, second hands
- Clock face with markings
- Smooth animations

### 3. Digital (Tech-Style Clock)
Modern digital clock with neon effects.
- Large time display
- Date and day of week
- Futuristic styling

### 4. Custom
Fully customizable with your own HTML/CSS/JS.

## 📖 Usage

### CLI Parameters

```
  -name string
        Watchface name (required)
  -version string
        Version number (default "1.0.0")
  -author string
        Author name (default "Anonymous")
  -description string
        Watchface description
  -template string
        Template type: simple, analog, digital (default "simple")
  -tags string
        Tags, comma-separated
  -output string
        Output directory (default ".")
  -no-preview
        Do not generate preview image
  -i    Interactive mode
  -list
        List available templates
```

### Generated Package Structure

```
watchface.zip
├── manifest.json       # Watchface metadata
├── index.html          # Entry point
├── style.css           # Styles
├── script.js           # JavaScript logic
└── preview.png         # Preview image (optional)
```

### Example manifest.json

```json
{
  "name": "My Watchface",
  "version": "1.0.0",
  "author": "Your Name",
  "description": "A beautiful watchface",
  "entrypoint": "index.html",
  "tags": ["minimal", "digital"],
  "created_at": "2025-01-21T10:30:00Z"
}
```

## 🌐 Web API Mode

You can run Watchface Builder as a web service:

```bash
./watchface-builder serve --port 8080
```

### API Endpoints

**Build Watchface**
```http
POST /api/build
Content-Type: application/json

{
  "name": "My Watchface",
  "version": "1.0.0",
  "template": "analog",
  "author": "Your Name",
  "tags": ["classic", "analog"]
}
```

**List Templates**
```http
GET /api/templates
```

## 🔧 Advanced Usage

### Custom Template

Create your own watchface with custom HTML/CSS/JS:

```bash
./watchface-builder -name "Custom" -template custom \
  -custom-html "$(cat my-template.html)" \
  -custom-css "$(cat my-style.css)" \
  -custom-js "$(cat my-script.js)"
```

### Batch Generation

```bash
for template in simple analog digital; do
  ./watchface-builder -name "Watchface-$template" -template $template
done
```

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) first.

### Adding New Templates

1. Fork this repository
2. Create your template in `pkg/templates/`
3. Add template metadata to `pkg/templates/registry.go`
4. Submit a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [fogleman/gg](https://github.com/fogleman/gg) - For image generation
- [spf13/cobra](https://github.com/spf13/cobra) - For CLI framework

## 📞 Support

- 🐛 [Report Issues](https://github.com/ziztechnology/WatchfaceBuilder/issues)
- 💬 [Discussions](https://github.com/ziztechnology/WatchfaceBuilder/discussions)
- 📖 [Documentation](docs/)

## 🗺️ Roadmap

- [ ] More built-in templates
- [ ] Template marketplace
- [ ] Live preview in browser
- [ ] GUI application
- [ ] Template editor
- [ ] Cloud-based builder

---

Made with ❤️ by [Toooony](https://github.com/Toooony)
