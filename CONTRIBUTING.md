# Contributing to Watchface Builder

Thank you for your interest in contributing to Watchface Builder! 🎉

## Code of Conduct

Please be respectful and constructive in all interactions with the project.

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported in [Issues](https://github.com/Toooony/WatchfaceBuilder/issues)
2. If not, create a new issue with:
   - Clear title and description
   - Steps to reproduce
   - Expected vs actual behavior
   - Your environment (OS, Go version, etc.)

### Suggesting Features

1. Check if the feature has been suggested in [Issues](https://github.com/Toooony/WatchfaceBuilder/issues)
2. Create a new issue with:
   - Clear description of the feature
   - Use cases and benefits
   - Possible implementation ideas

### Submitting Pull Requests

1. **Fork** the repository
2. **Create a branch** for your feature:
   ```bash
   git checkout -b feature/my-awesome-feature
   ```
3. **Make your changes** with clear, descriptive commits
4. **Test** your changes thoroughly
5. **Update documentation** if needed
6. **Submit a pull request** to the `main` branch

### Pull Request Guidelines

- One feature/fix per PR
- Follow existing code style
- Add tests for new features
- Update README if adding new functionality
- Keep commits atomic and well-described

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/WatchfaceBuilder.git
cd WatchfaceBuilder

# Add upstream remote
git remote add upstream https://github.com/Toooony/WatchfaceBuilder.git

# Install dependencies
go mod download

# Build
go build -o watchface-builder ./cmd/cli

# Test
go test ./...
```

## Project Structure

```
WatchfaceBuilder/
├── cmd/
│   ├── cli/           # CLI application
│   └── server/        # Web server application
├── pkg/
│   ├── builder/       # Core builder logic
│   ├── templates/     # Template definitions
│   └── util/          # Utility functions
├── docs/              # Documentation
├── examples/          # Example watchfaces
└── tests/             # Integration tests
```

## Adding New Templates

To add a new template:

1. Create template files in `pkg/templates/your_template/`
2. Register the template in `pkg/templates/registry.go`
3. Add template documentation
4. Add test cases
5. Update README with template description

Example template structure:

```go
// pkg/templates/your_template/template.go
package your_template

type YourTemplate struct{}

func (t *YourTemplate) Generate(options Options) (map[string]string, error) {
    return map[string]string{
        "index.html": generateHTML(options),
        "style.css":  generateCSS(options),
        "script.js":  generateJS(options),
    }, nil
}

func (t *YourTemplate) Metadata() TemplateMetadata {
    return TemplateMetadata{
        ID:          "your_template",
        Name:        "Your Template Name",
        Description: "Template description",
        Category:    "digital", // or "analog", "minimal", etc.
        Difficulty:  "easy",    // or "medium", "hard"
        Features:    []string{"feature1", "feature2"},
    }
}
```

## Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions small and focused

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test ./pkg/builder -run TestBuildSimple
```

## Documentation

- Update README.md for user-facing changes
- Update inline code comments for API changes
- Add examples for new features
- Keep documentation clear and concise

## Commit Message Guidelines

Use conventional commit format:

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat(templates): add minimalist template

Add a new minimalist template with clean design.

Closes #123

fix(builder): handle empty author name

Previously crashed when author was empty.
Now defaults to "Anonymous".

docs(readme): update installation instructions
```

## Review Process

1. Maintainers will review your PR
2. Address any feedback or requested changes
3. Once approved, maintainers will merge your PR
4. Your contribution will be included in the next release!

## Getting Help

- 💬 [Discussions](https://github.com/Toooony/WatchfaceBuilder/discussions) - Ask questions
- 📖 [Documentation](docs/) - Read the docs
- 🐛 [Issues](https://github.com/Toooony/WatchfaceBuilder/issues) - Report bugs

## Recognition

Contributors will be recognized in:
- README.md acknowledgments
- Release notes
- CONTRIBUTORS.md file

Thank you for contributing! 🙏
