# Quick Start Guide

Get started with Watchface Builder in 5 minutes!

## Installation

### Option 1: Download Pre-built Binary (Recommended)

1. Go to [Releases](https://github.com/Toooony/WatchfaceBuilder/releases)
2. Download the binary for your platform:
   - Windows: `watchface-builder-windows-amd64.exe`
   - macOS: `watchface-builder-darwin-amd64`
   - Linux: `watchface-builder-linux-amd64`
3. Make it executable (macOS/Linux):
   ```bash
   chmod +x watchface-builder-*
   ```

### Option 2: Build from Source

```bash
git clone https://github.com/Toooony/WatchfaceBuilder.git
cd WatchfaceBuilder
go build -o watchface-builder ./cmd/cli
```

## Your First Watchface

### Method 1: Interactive Mode (Easiest)

```bash
./watchface-builder -i
```

Follow the prompts:
1. Enter watchface name: `My First Watchface`
2. Version: Press Enter for default `1.0.0`
3. Author: Enter your name or press Enter for `Anonymous`
4. Description: Optional, press Enter to skip
5. Template: Enter `1` for Simple template
6. Tags: Optional, press Enter to skip
7. Preview: Press Enter for `Y` (yes)

Done! You'll see a ZIP file created in the current directory.

### Method 2: Command Line (Quick)

```bash
./watchface-builder --name "My First Watchface" --template simple
```

### Method 3: With All Options

```bash
./watchface-builder \
  --name "My First Watchface" \
  --template analog \
  --version 1.0.0 \
  --author "Your Name" \
  --description "A beautiful analog clock" \
  --tags "classic,elegant" \
  --output ./output
```

## Test Your Watchface

### 1. Extract the ZIP

```bash
unzip My_First_Watchface_v1.0.0_*.zip -d my-watchface
cd my-watchface
```

### 2. Open in Browser

```bash
# macOS
open index.html

# Linux
xdg-open index.html

# Windows
start index.html
```

You should see your watchface running in the browser!

## Next Steps

### Try Different Templates

**Simple (Digital Clock)**:
```bash
./watchface-builder --name "Simple Clock" --template simple
```

**Analog (Clock with Hands)**:
```bash
./watchface-builder --name "Analog Clock" --template analog
```

**Digital (Tech Style)**:
```bash
./watchface-builder --name "Digital Clock" --template digital
```

### Customize a Template

Create your own HTML/CSS/JS files:

**mytemplate.html**:
```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>My Custom Watchface</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div id="time">00:00</div>
    <script src="script.js"></script>
</body>
</html>
```

**mystyle.css**:
```css
body {
    background: black;
    color: lime;
    font-family: monospace;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}
#time {
    font-size: 4rem;
}
```

**myscript.js**:
```javascript
function updateTime() {
    const now = new Date();
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    document.getElementById('time').textContent = hours + ':' + minutes;
}
updateTime();
setInterval(updateTime, 1000);
```

Build with custom files:
```bash
./watchface-builder \
  --name "Custom Clock" \
  --template custom \
  --custom-html-file mytemplate.html \
  --custom-css-file mystyle.css \
  --custom-js-file myscript.js
```

### Batch Create Multiple Watchfaces

Create a script `build-all.sh`:

```bash
#!/bin/bash

for template in simple analog digital; do
  ./watchface-builder \
    --name "Watchface ${template^}" \
    --template $template \
    --author "Your Name" \
    --output ./output
done

echo "Built 3 watchfaces!"
```

Run it:
```bash
chmod +x build-all.sh
./build-all.sh
```

## Common Issues

### "Permission denied" error

On macOS/Linux, make the binary executable:
```bash
chmod +x watchface-builder
```

### "command not found" error

Use `./watchface-builder` instead of just `watchface-builder`, or add it to your PATH:

```bash
# macOS/Linux
export PATH=$PATH:$(pwd)

# Or move to a directory in PATH
sudo mv watchface-builder /usr/local/bin/
```

### Preview image not generated

The preview image requires the `gg` library. If it fails, the watchface will still build successfully without the preview. You can use `--no-preview` to skip preview generation:

```bash
./watchface-builder --name "My Watchface" --template simple --no-preview
```

## Tips & Tricks

### 1. Use Aliases

Add to your `~/.bashrc` or `~/.zshrc`:

```bash
alias wfb='~/path/to/watchface-builder'
```

Then use:
```bash
wfb --name "Quick Watchface" --template simple
```

### 2. Default Template

Always use the same template? Set it as default in a wrapper script:

```bash
#!/bin/bash
~/path/to/watchface-builder --template analog "$@"
```

### 3. Version Incrementing

Automatically increment version:

```bash
#!/bin/bash
VERSION=$(cat VERSION)
./watchface-builder --name "My Watchface" --version $VERSION --template simple
echo $((VERSION + 1)) > VERSION
```

### 4. Watch Mode (Auto-rebuild on changes)

Use with file watcher:

```bash
# Install fswatch (macOS)
brew install fswatch

# Watch and rebuild
fswatch -o mytemplate.html | xargs -n1 -I{} ./watchface-builder \
  --name "Auto" --template custom --custom-html-file mytemplate.html
```

## Getting Help

- 📖 [Full Documentation](../README.md)
- 🐛 [Report Issues](https://github.com/Toooony/WatchfaceBuilder/issues)
- 💬 [Discussions](https://github.com/Toooony/WatchfaceBuilder/discussions)

## What's Next?

- Read the [Template Guide](TEMPLATES.md) to understand template structure
- Check out [Examples](../examples/) for inspiration
- Read the [API Documentation](API.md) if you want to integrate programmatically
- Learn about [Contributing](../CONTRIBUTING.md) to add your own templates

Happy building! 🎉
